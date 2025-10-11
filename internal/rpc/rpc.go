package rpc

import (
	"News-portal/internal/newsportal"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/vmkteam/zenrpc/v2"
)

type NewsService struct {
	zenrpc.Service
	m *newsportal.Manager
}
type ShortNewsService struct {
	zenrpc.Service
	m *newsportal.Manager
}
type TagsService struct {
	zenrpc.Service
	m *newsportal.Manager
}
type CategoriesService struct {
	zenrpc.Service
	m *newsportal.Manager
}

//go:generate zenrpc
func Init() {
	addr := flag.String("addr", "localhost:9999", "listen address")
	flag.Parse()

	rpc := zenrpc.NewServer(zenrpc.Options{
		ExposeSMD:              true,
		AllowCORS:              true,
		DisableTransportChecks: true,
	})

	rpc.Register("news", NewsService{})
	rpc.Register("shortNews", ShortNewsService{})
	rpc.Register("tags", TagsService{})
	rpc.Register("categories", CategoriesService{})
	rpc.Use(zenrpc.Logger(log.New(os.Stderr, "", log.LstdFlags)))

	http.Handle("/", rpc)

	log.Printf("starting arithsrv on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))

}
