package server

import (
	grpc2 "github.com/ce-final-project/backend_rest_api/account_service/internal/account/delivery/grpc"
	accountService "github.com/ce-final-project/backend_rest_api/account_service/proto/account"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

const (
	maxConnectionIdle = 5
	gRPCTimeout       = 15
	maxConnectionAge  = 5
	gRPCTime          = 10
)

func (s *server) newAccountGrpcServer() (func() error, *grpc.Server, error) {
	l, err := net.Listen("tcp", s.cfg.GRPC.Port)
	if err != nil {
		return nil, nil, errors.Wrap(err, "net.Listen")
	}

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: maxConnectionIdle * time.Minute,
			Timeout:           gRPCTimeout * time.Second,
			MaxConnectionAge:  maxConnectionAge * time.Minute,
			Time:              gRPCTime * time.Minute,
		}),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_recovery.UnaryServerInterceptor(),
			s.im.Logger,
		),
		),
	)

	accountGrpcAccount := grpc2.NewAccountGrpcService(s.log, s.cfg, s.v, s.as, s.metrics)
	accountService.RegisterAccountServiceServer(grpcServer, accountGrpcAccount)
	grpc_prometheus.Register(grpcServer)

	if s.cfg.GRPC.Development {
		reflection.Register(grpcServer)
	}

	go func() {
		s.log.Infof("Account gRPC server is listening on port: %s", s.cfg.GRPC.Port)
		s.log.Fatal(grpcServer.Serve(l))
	}()

	return l.Close, grpcServer, nil
}
