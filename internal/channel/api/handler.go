package api

import (
	"fmt"
	"net/http"

	svc "github.com/vvlsouza/c3pcid/internal/channel/service"
	"github.com/vvlsouza/c3pcid/internal/rest"
)

type channelService interface {
	GetChannelMessagesContent(string) ([]svc.ChannelMessage, error)
	UpdateData(string) error
	GetRandomChannelMessage(string) (*svc.ChannelMessage, error)
}

// Handler is used to aggregate all endpoints related
type Handler struct {
	svc channelService
}

// NewHandler Create a new API handler
func NewHandler(svc channelService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) RandomMessage(w http.ResponseWriter, r *http.Request) {
	channelID, err := rest.GetString(r, "channel-id")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	resp, err := h.svc.GetRandomChannelMessage(channelID)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	rest.SendJSON(w, resp)
}

func (h *Handler) ListChannelMessages(w http.ResponseWriter, r *http.Request) {
	channelID, err := rest.GetString(r, "channel-id")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	msgs, err := h.svc.GetChannelMessagesContent(channelID)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	rest.SendJSON(w, msgs)
}

func (h *Handler) UpdateData(w http.ResponseWriter, r *http.Request) {
	channelID, err := rest.GetString(r, "channel-id")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	err = h.svc.UpdateData(channelID)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	rest.SendJSON(w, http.StatusOK)
}
