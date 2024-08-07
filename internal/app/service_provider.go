package app

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/saver89/microservices_auth/internal/api/user"
	"github.com/saver89/microservices_auth/internal/client/db"
	"github.com/saver89/microservices_auth/internal/client/db/pg"
	"github.com/saver89/microservices_auth/internal/client/db/transaction"
	"github.com/saver89/microservices_auth/internal/closer"
	"github.com/saver89/microservices_auth/internal/config"
	"github.com/saver89/microservices_auth/internal/repository"
	userRepository "github.com/saver89/microservices_auth/internal/repository/user"
	userLogRepository "github.com/saver89/microservices_auth/internal/repository/user_log"
	"github.com/saver89/microservices_auth/internal/service"
	userService "github.com/saver89/microservices_auth/internal/service/user"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient  db.Client
	txManager db.TxManager

	userService        service.UserService
	userRepository     repository.UserRepository
	userLogRepository  repository.UserLogRepository
	userImplementation *user.Implementation

	log *slog.Logger
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN(), s.Log())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewUserService(s.UserRepository(ctx), s.UserLogRepository(ctx), s.TxManager(ctx))
	}

	return s.userService
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewUserRepository(s.DBClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) UserLogRepository(ctx context.Context) repository.UserLogRepository {
	if s.userLogRepository == nil {
		s.userLogRepository = userLogRepository.NewUserLogRepository(s.DBClient(ctx))
	}

	return s.userLogRepository
}

func (s *serviceProvider) UserImplementation(ctx context.Context) *user.Implementation {
	if s.userImplementation == nil {
		s.userImplementation = user.NewImplementation(s.Log(), s.UserService(ctx))
	}

	return s.userImplementation
}

func (s *serviceProvider) Log() *slog.Logger {
	if s.log == nil {
		s.log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}

	return s.log
}
