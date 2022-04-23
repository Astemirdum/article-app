package article

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Serv struct {
	serv *http.Server
}

type AddrHttp struct {
	Addr string
	Port int
}

func NewServer(handler http.Handler, addrHttp AddrHttp) *Serv {
	return &Serv{
		serv: &http.Server{
			Addr:         fmt.Sprintf("%s:%d", addrHttp.Addr, addrHttp.Port),
			Handler:      handler,
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
