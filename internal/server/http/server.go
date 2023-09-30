package http

import (
	"fmt"
	"github.com/anvlad11/testapp-20230927/docs/api"
	"github.com/anvlad11/testapp-20230927/internal/interfaces/services"
	"github.com/anvlad11/testapp-20230927/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	_ "net/http/pprof"
)

type Server struct {
	cfg     config.HTTPServer
	handler services.HandlerService
	echo    *echo.Echo
}

func NewServer(
	cfg config.HTTPServer,
	handler services.HandlerService,
	e *echo.Echo,
) *Server {
	server := &Server{
		cfg:     cfg,
		handler: handler,
		echo:    e,
	}

	v1 := e.Group("/v1")

	v1.POST("/order", handler.CreateOrder)

	e.GET(
		"/_/docs/*",
		echo.WrapHandler(http.FileServer(http.FS(api.UI))),
		middleware.Rewrite(map[string]string{"/_/docs/*": "/openapi_ui/$1"}),
	)
	e.GET(
		"/_/docs/openapi.yaml",
		echo.WrapHandler(http.FileServer(http.FS(api.OpenAPISpec))),
		middleware.Rewrite(map[string]string{"/_/docs/*": "/$1"}),
	)

	return server
}

func (s *Server) Start() error {
	return s.echo.Start(fmt.Sprintf("%s:%s", s.cfg.Address, s.cfg.Port))
}

func (s *Server) Close() error {
	return s.echo.Close()
}

func NewEchoServer() *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodOptions, http.MethodPost},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	return e
}
