package rpc

import (
	"News-portal/internal/newsportal"
	"net/http"

	"github.com/vmkteam/zenrpc/v2"
)

type NewsService struct {
	zenrpc.Service
	m *newsportal.Manager
}

var noContentError = zenrpc.NewStringError(http.StatusNotFound, "not found")

//go:generate zenrpc
func New(m *newsportal.Manager) zenrpc.Server {
	srv := zenrpc.NewServer(zenrpc.Options{
		ExposeSMD:              true,
		AllowCORS:              true,
		DisableTransportChecks: true,
	})
	srv.Register("news", NewsService{m: m})

	return srv
}

func newInternalError(err error) error {
	return zenrpc.NewStringError(http.StatusInternalServerError, err.Error())
}
