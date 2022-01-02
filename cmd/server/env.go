package main

import (
	"fmt"
	"os"
)

type envConfig struct {
	mongoConnString     string
	mongoDBName         string
	mongoCollectionName string
	grpcAddr            string
}

func loadEnvConfig() (*envConfig, error) {
	const (
		provideEnvErrorMsg = `please provide "%s" environment variable`

		mongoConnStringEnv     = "MONGO_URI"
		mongoDBNameEnv         = "MONGO_DB_NAME"
		mongoCollectionNameEnv = "MONGO_COLLECTION_NAME"
		grpcAddrEnv            = "GRPC_SERVER_ADDR"
	)

	var ok bool

	cfg := &envConfig{}

	cfg.mongoConnString, ok = os.LookupEnv(mongoConnStringEnv)
	if !ok {
		return nil, fmt.Errorf(provideEnvErrorMsg, mongoConnStringEnv)
	}

	cfg.mongoDBName, ok = os.LookupEnv(mongoDBNameEnv)
	if !ok {
		return nil, fmt.Errorf(provideEnvErrorMsg, mongoDBNameEnv)
	}

	cfg.mongoCollectionName, ok = os.LookupEnv(mongoCollectionNameEnv)
	if !ok {
		return nil, fmt.Errorf(provideEnvErrorMsg, mongoCollectionNameEnv)
	}

	cfg.grpcAddr, ok = os.LookupEnv(grpcAddrEnv)
	if !ok {
		return nil, fmt.Errorf(provideEnvErrorMsg, grpcAddrEnv)
	}

	return cfg, nil
}
