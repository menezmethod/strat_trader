// Code generated by sqlc-grpc (https://github.com/walterwanderley/sqlc-grpc). DO NOT EDIT.

package main

import (
	"database/sql"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	pb_trades "strat_trader/api/trades/v1"
	"strat_trader/internal/server"
	app_trades "strat_trader/internal/trades"
)

func registerServer(logger *zap.Logger, db *sql.DB) server.RegisterServer {
	return func(grpcServer *grpc.Server) {
		pb_trades.RegisterTradesServiceServer(grpcServer, app_trades.NewService(logger, app_trades.New(db)))

	}
}

func registerHandlers() []server.RegisterHandler {
	var handlers []server.RegisterHandler

	handlers = append(handlers, pb_trades.RegisterTradesServiceHandler)

	return handlers
}