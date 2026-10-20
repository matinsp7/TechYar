package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/romina-kh/NetEng-Project/backend/internal/config"
	"github.com/romina-kh/NetEng-Project/backend/internal/db"
	"github.com/romina-kh/NetEng-Project/backend/internal/handler"
	"github.com/romina-kh/NetEng-Project/backend/internal/repository"
	"github.com/romina-kh/NetEng-Project/backend/internal/server"
	"github.com/romina-kh/NetEng-Project/backend/internal/service"
	utility "github.com/romina-kh/NetEng-Project/backend/internal/utils"
)

func main() {
	if err := run(); err != nil {
		log.Printf("application failed: %v", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("load congif: %w", err)
	}

	jwtManager, err := utility.NewJWTManager(cfg.JWT.Secret)
	if err != nil {
		return fmt.Errorf("initialize JWT manager: %w", err)
	}

	postgresDB, err := db.NewPostgresDB(cfg.Postgres)
	if err != nil {
		return fmt.Errorf("initialize postgres: %w", err)
	}

	sqlDB, err := postgresDB.DB()
	if err != nil {
		return fmt.Errorf("get underlying sql database: %w", err)
	}

	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Printf("close postgres database: %v", err)
		}
	}()

	userRepo := repository.NewUserRepo(postgresDB)
	productRepo := repository.NewProductRepository(postgresDB)

	authService := service.NewAuthService(userRepo)
	userService := service.NewUserService(userRepo)
	productService := service.NewProductService(productRepo)

	authHandler := handler.NewAuthHandler(jwtManager, authService)
	userHandler := handler.NewUserHandler(userService)
	productHandler := handler.NewProductHandler(productService)

	httpServer := server.NewHttpServer(
		cfg.Server.Address,
		authHandler,
		userHandler,
		productHandler,
		jwtManager,
	)

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("starting http server on %s", httpServer.Addr)
		serverErrors <- httpServer.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("run HTTP server: %w", err)
		}

		return nil

	case <-ctx.Done():
		log.Println("shutdown signal received")

		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			10*time.Second,
		)
		defer cancel()

		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			if closeErr := httpServer.Close(); closeErr != nil {
				return fmt.Errorf(
					"graceful shutdown failed: %w; force close failed: %v",
					err,
					closeErr,
				)
			}

			return fmt.Errorf("graceful shutdown HTTP server: %w", err)
		}

		if err := <-serverErrors; err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("HTTP server stopped with an error: %w", err)
		}

		log.Println("HTTP server stopped")
		return nil
	}
}
