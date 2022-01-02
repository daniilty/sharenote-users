package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	schema "github.com/daniilty/sharenote-grpc-schema"
	"github.com/daniilty/sharenote-users/internal/core"
	"github.com/daniilty/sharenote-users/internal/mongo"
	"github.com/daniilty/sharenote-users/internal/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	exitCodeInitError = 2
)

func run() error {
	cfg, err := loadEnvConfig()
	if err != nil {
		return err
	}

	client, err := mongo.Connect(context.Background(), cfg.mongoConnString)
	if err != nil {
		return err
	}

	db := client.Database(cfg.mongoDBName)
	collection := db.Collection(cfg.mongoCollectionName)

	d := mongo.NewDBImpl(db, collection)

	service := core.NewServiceImpl(d)

	loggerCfg := zap.NewProductionConfig()

	logger, err := loggerCfg.Build()
	if err != nil {
		return err
	}

	sugaredLogger := logger.Sugar()

	wg := &sync.WaitGroup{}
	listener, err := net.Listen("tcp", cfg.grpcAddr)
	if err != nil {
		return fmt.Errorf("net.Listen: %w", err)
	}

	grpcServer := grpc.NewServer()
	grpcService := server.NewGRPC(service)
	schema.RegisterUsersServer(grpcServer, grpcService)

	sugaredLogger.Infow("GRPC server is starting.", "addr", listener.Addr())

	wg.Add(1)
	go func() {
		err = grpcServer.Serve(listener)
		if err != nil {
			sugaredLogger.Errorw("Server failed to start.", "err", err)
		}
		wg.Done()
	}()

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-termChan

	sugaredLogger.Info("Gracefully stopping GRPC server.")
	grpcServer.GracefulStop()

	wg.Wait()

	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(exitCodeInitError)
	}
}
