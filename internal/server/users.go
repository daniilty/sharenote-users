package server

import (
	"net/http"
	"strings"
)

type usersResponse struct {
	Users []user `json:"users"`
}

func (u *usersResponse) writeJSON(w http.ResponseWriter) error {
	return writeJSONResponse(w, http.StatusOK, u)
}

func (h *HTTP) getUsersByID(w http.ResponseWriter, r *http.Request) {
	resp := h.getUsersByIDResponse(r)

	resp.writeJSON(w)
}

func (h *HTTP) getUsersByIDResponse(r *http.Request) response {
	const idsParamName = "ids"

	ids := r.FormValue(idsParamName)
	if ids == "" {
		msg := idsParamName + " cannot be empty"

		return getBadRequestWithMsgResponse(msg)
	}

	splittedIDs := strings.Split(ids, ",")

	users, err := h.service.GetUsers(r.Context(), splittedIDs)
	if err != nil {
		h.logger.Errorw("Get users by ids.", "ids", ids, "err", err)

		return getInternalServerErrorResponse()
	}

	return &usersResponse{
		Users: convertCoreUsersToResponse(users),
	}
}
