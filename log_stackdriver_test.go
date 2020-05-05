package aelog_test

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"testing"

	"cloud.google.com/go/logging"
	"github.com/1919yuan/aelog"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"gotest.tools/assert"
)

type FakeServer struct {
	Addr string
	l    net.Listener
	Gsrv *grpc.Server
}

func NewFakeServer(opts ...grpc.ServerOption) (*FakeServer, error) {
	l, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 0))
	if err != nil {
		return nil, err
	}
	s := &FakeServer{
		Addr: l.Addr().String(),
		l:    l,
		Gsrv: grpc.NewServer(opts...),
	}
	return s, nil
}

func (s *FakeServer) Start() {
	go func() {
		if err := s.Gsrv.Serve(s.l); err != nil {
			log.Printf("FakeServer.Start: %v", err)
		}
	}()
}

func (s *FakeServer) Close() {
	s.Gsrv.Stop()
	s.l.Close()
}

func TestStackDriverLogger(t *testing.T) {
	var client *logging.Client = nil
	server, err := NewFakeServer()
	if err != nil {
		log.Fatalf("Failed to start fake server: %v", err)
	}
	server.Start()
	defer server.Close()
	if os.Getenv("GOOGLE_CLOUD_PROJECT") == "" {
		conn, err := grpc.Dial(server.Addr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Dialing %q: %v", server.Addr, err)
		}
		client, err = logging.NewClient(context.Background(), "",
			option.WithGRPCConn(conn))
		if err != nil {
			log.Fatalf("Creating client for fake at %q: %v", server.Addr, err)
		}
	}

	logger := aelog.CreateStackDriverLogger(context.Background(), client)
	defer logger.Close()

	logger.Debug("Testing StackDriverLogger: Debug")
	logger.Info("Testing StackDriverLogger: Info")
	logger.Warning("Testing StackDriverLogger: Warning")
	logger.Error("Testing StackDriverLogger: Error")
	// Fatal will call log.Fatalln, so test will fail.
	// logger.Fatal("Testing StackDriverLogger: Fatal")
	assert.Assert(t, true)
}
