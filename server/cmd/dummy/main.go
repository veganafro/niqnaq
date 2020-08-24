package main

import (
	"context"
	"net/http"
	"log"
	"net"
	"strings"
	"crypto/tls"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/veganafro/mono/pkg/insecure"
	pb "github.com/veganafro/mono/pkg/dummy_gen"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc"
)

const (
	host = "0.0.0.0:3001"
)

type dummyServer struct {
	pb.UnimplementedDummyServiceServer
}

func (server *dummyServer) GetHello(ctx context.Context, request *empty.Empty) (*pb.DummyResponse, error) {
	return &pb.DummyResponse{ Rsp: "Hello world" }, nil
}

func main() {
	serverOpts := []grpc.ServerOption{
		grpc.Creds(credentials.NewClientTLSFromCert(insecure.GetCertPool(), host)),
	}
	grpcServer := grpc.NewServer(serverOpts...)
	pb.RegisterDummyServiceServer(grpcServer, &dummyServer{})

	dialCreds := credentials.NewTLS(&tls.Config{
		ServerName: host,
		RootCAs: insecure.GetCertPool(),
	})
	dialOpts := []grpc.DialOption{ grpc.WithTransportCredentials(dialCreds) }

	ctx := context.Background()
	gatewayMux := runtime.NewServeMux()
	error := pb.RegisterDummyServiceHandlerFromEndpoint(ctx, gatewayMux, host, dialOpts)
	if error != nil {
		log.Fatal("Failed to register service handler from endpoint | ", error)
	}

	httpMux := http.NewServeMux()
	httpMux.Handle("/hello", gatewayMux)

	httpHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			httpMux.ServeHTTP(w, r)
		}
	})

	httpServer := &http.Server{
		Addr: host,
		Handler: httpHandler,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{*insecure.GetKeyPair()},
			NextProtos: []string{"h2"},
		},
	}

	conn, error := net.Listen("tcp", host)
	if error != nil {
		log.Fatal("Failed to start tcp connection | ", error)
	}

	error = httpServer.Serve(tls.NewListener(conn, httpServer.TLSConfig))
	if error != nil {
		log.Fatal("Failed to start http server | ", error)
	}
}

