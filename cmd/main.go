package main

import (
	"context"
	"os"
	"runtime"
	"time"

	"go.uber.org/fx"

	"github.com/efectn/fiber-boilerplate/app/middleware"
	"github.com/efectn/fiber-boilerplate/app/module/article"
	"github.com/efectn/fiber-boilerplate/app/router"
	"github.com/efectn/fiber-boilerplate/internal/database"
	"github.com/efectn/fiber-boilerplate/internal/logger"
	"github.com/efectn/fiber-boilerplate/utils/config"
	"github.com/efectn/fiber-boilerplate/utils/response"
	fxzerolog "github.com/efectn/fx-zerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	_ "go.uber.org/automaxprocs"
)

func NewFiber(cfg *config.Config) *fiber.App {
	// Setup Webserver
	app := fiber.New(fiber.Config{
		ServerHeader:          cfg.App.Name,
		AppName:               cfg.App.Name,
		Prefork:               cfg.App.Prefork,
		ErrorHandler:          response.ErrorHandler,
		IdleTimeout:           cfg.App.IdleTimeout * time.Second,
		EnablePrintRoutes:     cfg.App.PrintRoutes,
		DisableStartupMessage: true,
	})

	// Pass production config to check it
	response.IsProduction = cfg.App.Production

	return app
}

func Start(lifecycle fx.Lifecycle, cfg *config.Config, fiber *fiber.App, router *router.Router, middlewares *middleware.Middleware, database *database.Database, log zerolog.Logger) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				// Register middlewares & routes
				middlewares.Register()
				router.Register()

				// Custom Startup Messages
				host, port := config.ParseAddr(cfg.App.Port)
				if host == "" {
					if fiber.Config().Network == "tcp6" {
						host = "[::1]"
					} else {
						host = "0.0.0.0"
					}
				}

				// Information message
				log.Info().Msg(fiber.Config().AppName + " is running at the moment!")

				// Debug informations
				if !cfg.App.Production {
					prefork := "Enabled"
					procs := runtime.GOMAXPROCS(0)
					if !cfg.App.Prefork {
						procs = 1
						prefork = "Disabled"
					}

					log.Debug().Msgf("Version: %s", "-")
					log.Debug().Msgf("Host: %s", host)
					log.Debug().Msgf("Port: %s", port)
					log.Debug().Msgf("Prefork: %s", prefork)
					log.Debug().Msgf("Handlers: %d", fiber.HandlersCount())
					log.Debug().Msgf("Processes: %d", procs)
					log.Debug().Msgf("PID: %d", os.Getpid())
				}

				// Listen the app (with TLS Support)
				if cfg.App.TLS.Enable {
					log.Debug().Msg("TLS support was enabled.")

					if err := fiber.ListenTLS(cfg.App.Port, cfg.App.TLS.CertFile, cfg.App.TLS.KeyFile); err != nil {
						log.Error().Err(err).Msg("An unknown error occurred when to run server!")
					}
				}

				go func() {
					if err := fiber.Listen(cfg.App.Port); err != nil {
						log.Error().Err(err).Msg("An unknown error occurred when to run server!")
					}
				}()

				database.ConnectDatabase()
				database.Migrate()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				log.Info().Msg("Shutting down the app...")
				if err := fiber.Shutdown(); err != nil {
					log.Panic().Err(err).Msg("")
				}

				log.Info().Msg("Running cleanup tasks...")
				log.Info().Msg("1- Shutdown the database")
				database.ShutdownDatabase()
				log.Info().Msgf("%s was successful shutdown.", cfg.App.Name)
				log.Info().Msg("\u001b[96msee you again👋\u001b[0m")

				return nil
			},
		},
	)
}

func main() {
	fx.New(
		// Provide patterns
		fx.Provide(NewFiber),
		fx.Provide(config.NewConfig),
		fx.Provide(logger.NewLogger),
		fx.Provide(database.NewDatabase),
		fx.Provide(middleware.NewMiddleware),
		fx.Provide(router.NewRouter),

		// Provide modules
		article.NewArticleModule,

		// Start Application
		fx.Invoke(Start),

		// Define logger
		fx.WithLogger(fxzerolog.Init(log.Logger)),
	).Run()
}
