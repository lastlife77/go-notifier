// Package handler provides HTTP endpoints.
// It handles managing notifies.
package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/lastlife77/go-notifier/internal/broker"
	"github.com/lastlife77/go-notifier/internal/domain"
)

// Broker defines the contract for message broker.
// It provides methods for managing notifies.
type Broker interface {
	SendMsg(id string, msg string, time time.Time) error
	GetStatus(id string) (string, error)
	DeleteMsg(id string) error
}

// Handler defines HTTP handlers.
// It uses a Broker to perform notifies operations.
type Handler struct {
	brok Broker
}

// New creates a new Handler with the given Broker.
func New(b Broker) *Handler {
	return &Handler{
		brok: b,
	}
}

// CreateNotify handles the creation of a new notify.

// @Summary Create notify
// @Tags notify
// @Accept json
// @Produce json
// @Param notify body domain.Notify true "Notify"
// @Router /notify [post]
func (h *Handler) CreateNotify(w http.ResponseWriter, r *http.Request) {
	var req domain.Notify

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t, err := convertTime(req.Time)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.brok.SendMsg(req.Id, req.Msg, t)
	if err != nil {
		var brokerErr *broker.Error
		if errors.As(err, &brokerErr) {
			http.Error(w, brokerErr.Reason, brokerErr.Code)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}

// GetNotifyStatus handles retrieving status for a specific notify.

// @Summary Get status
// @Tags notify
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Router /notify/{id} [get]
func (h *Handler) GetNotifyStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	status, err := h.brok.GetStatus(id)
	if err != nil {
		var brokerErr *broker.Error
		if errors.As(err, &brokerErr) {
			http.Error(w, brokerErr.Reason, brokerErr.Code)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(status))
}

// DeleteEvent handles the deletion of an notify.

// @Summary Delete notify
// @Tags notify
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Router /notify/{id} [delete]
func (h *Handler) DeleteNotify(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.brok.DeleteMsg(id)
	if err != nil {
		var brokerErr *broker.Error
		if errors.As(err, &brokerErr) {
			http.Error(w, brokerErr.Reason, brokerErr.Code)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}

func convertTime(t string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04", t, time.Local)
}
