package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/veganafro/mono/server/pkg/dummy/v1"

	"google.golang.org/grpc"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const (
	host = "0.0.0.0"
)

type dummyServer struct {
	pb.UnimplementedDummyServiceServer
}

func (server *dummyServer) GetHello(ctx context.Context, request *empty.Empty) (*pb.GetHelloResponse, error) {
	return &pb.GetHelloResponse{Rsp: "Hello world"}, nil
}

func main() {
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		log.Fatal("Failed to get port from environment")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDummyServiceServer(grpcServer, &dummyServer{})

	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	ctx := context.Background()
	gatewayMux := runtime.NewServeMux()
	error := pb.RegisterDummyServiceHandlerFromEndpoint(ctx, gatewayMux, host+":"+PORT, dialOpts)
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

	http2Server := &http2.Server{}
	http1Server := &http.Server{
		Addr:    host + ":" + PORT,
		Handler: h2c.NewHandler(httpHandler, http2Server),
	}

	conn, error := net.Listen("tcp", host+":"+PORT)
	if error != nil {
		log.Fatal("Failed to start tcp connection | ", error)
	}

	error = http1Server.Serve(conn)
	if error != nil {
		log.Fatal("Failed to start http server | ", error)
	}
}
