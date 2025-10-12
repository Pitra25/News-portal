package rpc

import (
	"News-portal/internal/newsportal"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/vmkteam/zenrpc/v2"
)

type Service struct {
	zenrpc.Service
	m *newsportal.Manager
}

//go:generate zenrpc
func Run(rpc zenrpc.Server) {
	addr := flag.String("addr", "localhost:9999", "listen address")
	flag.Parse()

	rpc.Register("", Service{})
	rpc.Use(zenrpc.Logger(log.New(os.Stderr, "rpc", log.LstdFlags)))

	http.Handle("/", rpc)

	log.Printf("starting arithsrv on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))

}
