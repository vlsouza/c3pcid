package api

import (
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/mux"
	"github.com/vvlsouza/c3pcid/internal/channel/service"
)

// Config is used to setup the API.
type Config struct {
	DiscordGo *discordgo.Session
	Router    *mux.Router
}

// New is used to initialize the API.
func New(c Config) {
	handler := NewHandler(service.New(c.DiscordGo))

	SetRoutes(handler, c.Router)
}

// SetRoutes is used to declare all endpoints managed by this API.
func SetRoutes(handler *Handler, router *mux.Router) {
	router.HandleFunc("/channels/{channel-id}/messages/random", handler.RandomMessage).Methods(http.MethodGet)
	router.HandleFunc("/channels/{channel-id}/messages", handler.ListChannelMessages).Methods(http.MethodGet)
	router.HandleFunc("/channels/{channel-id}/messages", handler.UpdateData).Methods(http.MethodPut)
}
