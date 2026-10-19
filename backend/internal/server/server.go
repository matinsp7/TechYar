package server

import (
	"net/http"
	"time"

	"github.com/romina-kh/NetEng-Project/backend/internal/handler"
	"github.com/romina-kh/NetEng-Project/backend/internal/router"
)

func NewHttpServer(
	address string,
	authHandler *handler.Handler,
	userHandler *handler.UserHandler,
	productHandler *handler.ProductHandler,
) *http.Server {
	r := router.SetupRouter(authHandler, userHandler, productHandler)

	return &http.Server{
		Addr:        address,
		Handler:     r,
		ReadTimeout: 5 * time.Second,
	}
}
