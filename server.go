package article

import (
	"context"
	"net/http"
	"time"
)

type Serv struct {
	serv *http.Server
}

func NewServer(port string, hand http.Handler) *Serv {
	return &Serv{
		serv: &http.Server{
			Addr:         ":" + port,
			Handler:      hand,
			ReadTimeout:  time.Second * 5,
			WriteTimeout: time.Second * 5,
		},
	}
}

func (s *Serv) Run() error {
	return s.serv.ListenAndServe()
}

func (s *Serv) Shutdown(ctx context.Context) error {
	return s.serv.Shutdown(ctx)
}
