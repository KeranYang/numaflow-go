package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	sourcepb "github.com/KeranYang/numaflow-go/pkg/apis/proto/source/v1"
	functionsdk "github.com/KeranYang/numaflow-go/pkg/function"
	"github.com/KeranYang/numaflow-go/pkg/info"
	sourcesdk "github.com/KeranYang/numaflow-go/pkg/source"
)

type server struct {
	svc *sourcesdk.Service
}

// New creates a new server object.
func New() *server {
	s := new(server)
	s.svc = new(sourcesdk.Service)
	return s
}

func (s *server) RegisterPendingHandler(m sourcesdk.PendingHandler) *server {
	s.svc.PendingHandler = m
	return s
}

// Start starts the gRPC server via unix domain socket at configs.Addr and return error.
func (s *server) Start(ctx context.Context, inputOptions ...Option) error {
	var opts = &options{
		sockAddr:           functionsdk.UDS_ADDR,
		maxMessageSize:     functionsdk.DefaultMaxMessageSize,
		serverInfoFilePath: info.ServerInfoFilePath,
	}

	for _, inputOption := range inputOptions {
		inputOption(opts)
	}

	// Write server info to the file
	serverInfo := &info.ServerInfo{Protocol: info.UDS, Language: info.Go, Version: info.GetSDKVersion()}
	if err := info.Write(serverInfo, info.WithServerInfoFilePath(opts.serverInfoFilePath)); err != nil {
		return err
	}

	cleanup := func() error {
		// err if no opts.sockAddr should be ignored
		if _, err := os.Stat(opts.sockAddr); err == nil {
			return os.RemoveAll(opts.sockAddr)
		}
		return nil
	}

	if err := cleanup(); err != nil {
		return err
	}

	ctxWithSignal, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	lis, err := net.Listen(functionsdk.UDS, opts.sockAddr)
	if err != nil {
		return fmt.Errorf("failed to execute net.Listen(%q, %q): %v", functionsdk.UDS, functionsdk.UDS_ADDR, err)
	}
	defer func() { _ = lis.Close() }()
	grpcServer := grpc.NewServer(
		grpc.MaxRecvMsgSize(opts.maxMessageSize),
		grpc.MaxSendMsgSize(opts.maxMessageSize),
	)
	defer log.Println("Successfully stopped the gRPC server")
	defer grpcServer.GracefulStop()
	sourcepb.RegisterUserDefinedSourceServer(grpcServer, s.svc)

	errCh := make(chan error, 1)
	defer close(errCh)
	// start the grpc server
	go func(ch chan<- error) {
		log.Println("starting the gRPC server with unix domain socket...", lis.Addr())
		err = grpcServer.Serve(lis)
		if err != nil {
			ch <- fmt.Errorf("failed to start the gRPC server: %v", err)
		}
	}(errCh)

	select {
	case err := <-errCh:
		return err
	case <-ctxWithSignal.Done():
		log.Println("Got a signal: terminating gRPC server...")
	}

	return nil
}
