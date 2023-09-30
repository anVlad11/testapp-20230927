package main

import (
	"flag"
	"github.com/anvlad11/testapp-20230927/internal/config"
	"github.com/anvlad11/testapp-20230927/internal/server/http"
	"github.com/anvlad11/testapp-20230927/internal/services/handler"
	"github.com/anvlad11/testapp-20230927/internal/services/order"
	"github.com/labstack/gommon/log"
)

var configPath = flag.String("config-path", "./config.yaml", "Path to the application config")

func main() {
	flag.Parse()

	cfg, err := config.NewConfig(*configPath)
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	orderService := order.NewService(cfg.Orders)
	handlerService := handler.NewService(orderService)

	e := http.NewEchoServer()
	server := http.NewServer(
		cfg.HTTPServer,
		handlerService,
		e,
	)

	err = server.Start()
	if err != nil {
		log.Fatalf("server error: %v", err)
	}

}
