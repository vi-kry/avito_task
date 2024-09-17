package app

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	changeStatusBid "avito_task/internal/handler/http/bid/changeStatus"
	createBid "avito_task/internal/handler/http/bid/create"
	editBid "avito_task/internal/handler/http/bid/edit"
	getBids "avito_task/internal/handler/http/bid/get"
	getBidsByTenderId "avito_task/internal/handler/http/bid/getByTenderId"
	submitBid "avito_task/internal/handler/http/bid/submit"
	changeStatusTender "avito_task/internal/handler/http/tender/changeStatus"
	"avito_task/internal/handler/http/tender/create"
	editTender "avito_task/internal/handler/http/tender/edit"
	"avito_task/internal/handler/http/tender/getAll"
	"avito_task/internal/handler/http/tender/getAllByUserId"
	"avito_task/internal/server"

	"avito_task/internal/config"
	"avito_task/internal/configure/database/postgres"
	repositoryBid "avito_task/internal/repository/bid"
	repositoryEmployee "avito_task/internal/repository/employee"
	repositoryTender "avito_task/internal/repository/tender"
	usecaseCreateBid "avito_task/internal/usecase/bid/create"
	usecaseFetchBids "avito_task/internal/usecase/bid/fetch"
	usecaseSubmitBids "avito_task/internal/usecase/bid/submit"
	usecaseTender "avito_task/internal/usecase/tender"
	"avito_task/pkg/logger"
)

func Run() {
	cfg := config.InitConfig()

	log := logger.SetupLogger(cfg.Env)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db := setupDatabase(ctx, log, &cfg)

	// repo
	repoTender := repositoryTender.NewTenderRepo(db.GetDB())
	repoEmployee := repositoryEmployee.NewEmployeeRepo(db.GetDB())

	repoBid := repositoryBid.NewBidRepo(db.GetDB())

	// use case
	createTenderUseCase := usecaseTender.NewTenderUseCase(repoTender, repoEmployee)
	getTendersByUserIdUseCase := usecaseTender.NewGetTendersByUserIdUseCase(repoTender, repoEmployee)

	createBidUseCase := usecaseCreateBid.NewBidUseCase(repoBid, repoEmployee)
	getBidsByUsername := usecaseFetchBids.NewBidUseCase(repoBid, repoEmployee)
	submitBidUseCase := usecaseSubmitBids.NewSubmitUseCase(repoBid, repoTender)

	// handler
	createTenderHandler := create.NewHandler(createTenderUseCase, log)
	getAllTendersHandler := getAll.NewHandler(repoTender, log)
	changeStatusTenderHandler := changeStatusTender.NewHandler(repoTender, log)
	getTendersByUserIdHandler := getAllByUserId.NewHandler(getTendersByUserIdUseCase, log)
	editTenderHandler := editTender.NewHandler(repoTender, log)

	createBidHandler := createBid.NewHandler(createBidUseCase, log)
	getBidsByUsernameHandler := getBids.NewHandler(getBidsByUsername, log)
	getBidsByTenderIdHandler := getBidsByTenderId.NewHandler(repoBid, log)
	editBidHandler := editBid.NewHandler(repoBid, log)
	changeStatusBidHandler := changeStatusBid.NewHandler(repoBid, log)
	submitBidHandler := submitBid.NewHandler(submitBidUseCase, log)

	// router
	r := chi.NewRouter()
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	r.Route("/api", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Route("/tenders", func(r chi.Router) {
				r.Get("/", getAllTendersHandler.Handle)
				r.Post("/new", createTenderHandler.Handle)
				r.Post("/status", changeStatusTenderHandler.Handle)
				r.Get("/my", getTendersByUserIdHandler.Handle)

				r.Patch("/{tenderId}/edit", editTenderHandler.Handle)
			})

			r.Route("/bids", func(r chi.Router) {
				r.Post("/new", createBidHandler.Handle)
				r.Post("/status", changeStatusBidHandler.Handle)
				r.Get("/my", getBidsByUsernameHandler.Handle)
				r.Get("/{tenderId}/list", getBidsByTenderIdHandler.Handle)

				r.Post("/submit_decision", submitBidHandler.Handle)

				r.Patch("/{bidId}/edit", editBidHandler.Handle)
			})
		})
	})

	// gc
	srv := new(server.Server)
	go func() {
		log.Info("start http server")
		err := srv.Run(cfg.HTTP.Address, r)
		if err != http.ErrServerClosed {
			log.Error("error occurred while running http server: %s", err.Error())
			os.Exit(1)
		}
	}()

	signalListner := make(chan os.Signal, 1)
	signal.Notify(signalListner,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	stop := <-signalListner
	fmt.Println(cfg)
	log.Info("Shutting Down app: %s", stop)

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Error("error occurred on server shutting down: %s", err.Error())
	}
}

func setupDatabase(ctx context.Context, log *slog.Logger, cfg *config.Config) *postgres.PgDB {
	postgresDB, err := postgres.NewPostgresDB(ctx, log, postgresDSN(&cfg.Postgres))
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	return postgresDB
}

func postgresDSN(psqlCfg *config.PostgresConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		psqlCfg.Host, psqlCfg.Port, psqlCfg.Username, psqlCfg.Password, psqlCfg.DbName, psqlCfg.SslMode)
}
