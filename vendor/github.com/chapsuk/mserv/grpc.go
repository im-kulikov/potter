package mserv

import (
	"net"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	addr   string
	server *grpc.Server
}

func NewGRPCServer(addr string, server *grpc.Server) Server {
	if addr == "" {
		return nil
	}

	return &GRPCServer{
		addr:   addr,
		server: server,
	}
}

func (g *GRPCServer) Start() {
	lis, err := net.Listen("tcp", g.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		if err := g.server.Serve(lis); err != nil {
			if err != grpc.ErrServerStopped {
				log.Fatalf("start grpc server %s error %s", g.addr, err)
			}
		}
	}()
}

func (g *GRPCServer) Stop() {
	if g.server == nil {
		return
	}
	g.server.GracefulStop()
}
