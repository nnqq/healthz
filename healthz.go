package healthz

import (
	"log"
	"net/http"
)

func NewHealthz(opts ...Option) *healthz {
	defaults := &options{
		mux:      http.NewServeMux(),
		addr:     "0.0.0.0:80",
		endpoint: "/healthz",
		response: nil,
		logger:   log.Default(),
	}
	for _, opt := range opts {
		opt(defaults)
	}

	defaults.mux.Handle(defaults.endpoint, http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, err := w.Write(defaults.response)
		if err != nil {
			defaults.logger.Printf("response write error: %w", err)
		}
	}))

	return &healthz{
		mux:      defaults.mux,
		addr:     defaults.addr,
		endpoint: defaults.endpoint,
		response: defaults.response,
		logger:   defaults.logger,
	}
}

func (h *healthz) Mux() *http.ServeMux {
	return h.mux
}

func (h *healthz) Serve() error {
	return http.ListenAndServe(h.addr, h.mux)
}

type Option func(*options)

func Mux(mux *http.ServeMux) Option {
	return func(o *options) {
		o.mux = mux
	}
}

func Addr(addr string) Option {
	return func(o *options) {
		o.addr = addr
	}
}

func Endpoint(endpoint string) Option {
	return func(o *options) {
		o.endpoint = endpoint
	}
}

func Response(response []byte) Option {
	return func(o *options) {
		o.response = response
	}
}

func Logger(logger Printer) Option {
	return func(o *options) {
		o.logger = logger
	}
}

type Printer interface {
	Printf(format string, v ...interface{})
}

type healthz options

type options struct {
	mux      *http.ServeMux
	addr     string
	endpoint string
	response []byte
	logger   Printer
}
