package server

import (
	"github.com/2413Yang/lotterysvr/internal/conf"
	"github.com/2413Yang/lotterysvr/internal/interfaces"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, h *interfaces.Handler) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			MiddlewareTraceID(),
			MiddlewareLog(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	// v1.RegisterLotteryHTTPServer(srv, greeter)
	srv.HandlePrefix("/", interfaces.NewRouter(h))
	return srv
}
