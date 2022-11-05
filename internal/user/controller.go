package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type (
	controller struct {
		s userService
	}
)

func newController(s *service) *controller {
	return &controller{s: s}
}

func (c *controller) Create(w http.ResponseWriter, r *http.Request) {
	req := createUserRequest{}
	err := unmarshalRequest(r, &req)
	if err != nil {
		logrus.WithError(err).Error("unmarshal request")
		writeResponse(w, response{
			Success: false,
			Errors: []errResponse{
				{
					Code:    "1",
					Context: err.Error(),
				},
			},
		}, http.StatusBadRequest)
		return
	}

	u, err := req.toUser()
	if err != nil {
		logrus.
			WithField("request", req).
			WithError(err).
			Error("convert request to user")
		writeErrorResponse(w, http.StatusBadRequest, err, "1")
		return
	}

	userID, err := c.s.Create(r.Context(), u)
	if err != nil {
		logrus.
			WithFields(logrus.Fields{
				"user":    u,
				"request": req,
			}).
			WithError(err).
			Error("failed to create user")
		writeErrorResponse(w, http.StatusBadRequest, err, "2")
		return
	}

	resp := createUserResponse{userID}
	rawResp, err := json.Marshal(resp)
	if err != nil {
		logrus.WithError(err).Error("marshal response")
		writeErrorResponse(w, http.StatusInternalServerError, err, "3")
		return
	}

	writeResponse(w, response{
		Success: true,
		Body:    rawResp,
	}, http.StatusCreated)
}

func (c *controller) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		logrus.WithField("vars", vars).Error("id is not set")
		writeErrorResponse(w, http.StatusBadRequest, errEmptyUserID, "j6")
		return
	}

	id, _ := strconv.ParseInt(idStr, 10, 64)

	user, err := c.s.Get(r.Context(), id)
	if errors.Is(err, ErrUserNotFound) {
		writeResponse(w, response{
			Success: false,
			Errors: []errResponse{
				{
					Code:    "89302",
					Context: fmt.Sprintf("err: %s, id: %d", ErrUserNotFound.Error(), id),
				},
			},
		}, http.StatusNotFound)
	}
	if err != nil {
		logrus.WithField("id", id).WithError(err).Error("failed to get user")
		writeErrorResponse(w, http.StatusInternalServerError, err, "j6asd")
		return
	}

	resp := getUserResponse{
		ID:          user.ID,
		Name:        user.Name,
		DateOfBirth: user.DateOfBirth.Format(dobFormat),
	}

	rawResp, err := json.Marshal(resp)
	if err != nil {
		logrus.WithError(err).Error("marshal response")
		writeErrorResponse(w, http.StatusInternalServerError, err, "3")
		return
	}

	writeResponse(w, response{
		Success: true,
		Body:    rawResp,
	}, http.StatusOK)
}

func (c *controller) Update(w http.ResponseWriter, r *http.Request) {
	req := updateUserRequest{}
	err := unmarshalRequest(r, &req)
	if err != nil {
		logrus.WithError(err).Error("unmarshal request")
		writeResponse(w, response{
			Success: false,
			Errors: []errResponse{
				{
					Code:    "1",
					Context: err.Error(),
				},
			},
		}, http.StatusBadRequest)
		return
	}

	u, err := req.toUser()
	if err != nil {
		logrus.
			WithField("request", req).
			WithError(err).
			Error("convert request to user")
		writeErrorResponse(w, http.StatusBadRequest, err, "1")
		return
	}

	err = c.s.Update(r.Context(), u)
	if err != nil {
		logrus.
			WithFields(logrus.Fields{
				"user":    u,
				"request": req,
			}).
			WithError(err).
			Error("failed to update user")
		writeErrorResponse(w, http.StatusBadRequest, err, "2")
		return
	}

	writeResponse(w, response{
		Success: true,
	}, http.StatusAccepted)
}

func (c *controller) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		logrus.WithField("vars", vars).Error("id is not set")
		writeErrorResponse(w, http.StatusBadRequest, errEmptyUserID, "j6")
		return
	}

	id, _ := strconv.ParseInt(idStr, 10, 64)

	err := c.s.Delete(r.Context(), id)
	if err != nil {
		logrus.WithField("id", id).WithError(err).Error("failed to delete user")
		writeErrorResponse(w, http.StatusInternalServerError, err, "j6asd")
		return
	}

	writeResponse(w, response{
		Success: true,
	}, http.StatusAccepted)

}

func (c *controller) Register(r *mux.Router) {
	r.PathPrefix("/user/")

	r.HandleFunc("", c.Create).Methods(http.MethodPost)
	r.HandleFunc("", c.Update).Methods(http.MethodPut)
	r.HandleFunc("{id:[0-9]+}", c.Get).Methods(http.MethodGet)
	r.HandleFunc("{id:[0-9]+}", c.Delete).Methods(http.MethodGet)
}

func unmarshalRequest(r *http.Request, req json.Unmarshaler) error {
	body, err := r.GetBody()
	if err != nil {
		return fmt.Errorf("get body: %w", err)
	}
	b, err := io.ReadAll(body)
	if err != nil {
		return fmt.Errorf("read all: %w", err)
	}
	err = json.Unmarshal(b, req)
	if err != nil {
		logrus.
			WithField("raw_body", string(b)).
			WithError(err).
			Error("failed to unmarshal json body to req")

		return fmt.Errorf("unmarshal json: %w", err)
	}

	return nil
}

func writeResponse(w http.ResponseWriter, resp response, statusCode int) {
	b, err := json.Marshal(resp)
	if err != nil {
		f := logrus.Fields{
			"response":      resp,
			"marshal_error": err.Error(),
		}

		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte(hardFailError))
		if err != nil {
			f["write_error"] = err.Error()
		}

		logrus.
			WithError(errFailedToWriteResponse).
			WithFields(f).
			Error("failed to marshal response")
	}

	w.WriteHeader(statusCode)
	_, err = w.Write(b)
	if err != nil {
		logrus.
			WithField("response", resp).
			WithError(err).
			Errorf("failed to write response")
	}
}

func writeErrorResponse(w http.ResponseWriter, statusCode int, err error, code string) {
	writeResponse(w, response{
		Success: false,
		Errors: []errResponse{
			{
				Code:    code,
				Context: err.Error(),
			},
		},
	}, statusCode)
}
