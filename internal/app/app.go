package app

import (
	"banners_oto/config"
	"banners_oto/internal/delivery/metrics"
	"banners_oto/internal/delivery/route"
	"banners_oto/internal/utils/functions"
	"banners_oto/microservices"
	"banners_oto/services"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

func Run(logger *zap.Logger) {
	functions.InitDtoValidator(logger)
	cfg := config.Config

	// i: mertrics

	reg := prometheus.NewRegistry()
	m := metrics.NewMetrics("gateway", reg)

	// i: microservices Clients

	grpcClients := microservices.Init(logger, m)
	// i: service Cluster
	serviceCluster := services.Init(logger)

	defer func(clients *microservices.Clients) {
		err := clients.SessionConn.Close()
		if err != nil {
			(logger.Error(fmt.Sprintf("Error while closign connection ~SessionConn %v", err)))
		}

		err = clients.UserConn.Close()
		if err != nil {
			(logger.Error(fmt.Sprintf("Error while closign connection ~UserConn %v", err)))
		}

		err = clients.CommentConn.Close()
		if err != nil {
			(logger.Error(fmt.Sprintf("Error while closign connection ~CommentConn %v", err)))
		}

		err = clients.AuthConn.Close()
		if err != nil {
			(logger.Error(fmt.Sprintf("Error while closign connection ~AuthConn %v", err)))
		}

		err = clients.RestConn.Close()
		if err != nil {
			(logger.Error(fmt.Sprintf("Error while closign connection ~RestConn %v", err)))
		}

	}(grpcClients)
	r := mux.NewRouter()
	handler := route.Setup(r, serviceCluster, grpcClients, m, reg, logger)

	serverCfg := cfg.Server
	serverAddr := fmt.Sprintf(":%d", serverCfg.Port)
	httpServer := http.Server{
		Handler:      handler,
		Addr:         serverAddr,
		ReadTimeout:  serverCfg.ReadTimeout,
		WriteTimeout: serverCfg.WriteTimeout,
		IdleTimeout:  serverCfg.IdleTimeout,
	}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			logger.Error(fmt.Sprintf("Error while running server %v", err))
		}

	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), serverCfg.ShutdownDuration)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Error(fmt.Sprintf("Error while shutdown server %v", err))

	}
	logger.Info("Succe on shutdown server ::..")

}
