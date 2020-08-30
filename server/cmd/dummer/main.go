package main

import (
	"context"
	"net/http"
	"log"
	"net"
	"strings"

	dummy "github.com/veganafro/mono/pkg/dummy_gen"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/veganafro/mono/pkg/dummer_gen"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"google.golang.org/grpc"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const (
	host = "0.0.0.0:3002"
)

type dummerServer struct {
	pb.UnimplementedDummerServiceServer
}

func (server *dummerServer) GetWorld(ctx context.Context, request *dummy.DummyRequest) (*dummy.DummyResponse, error) {
	conn, error := grpc.Dial("0.0.0.0:3001", grpc.WithInsecure())
	if error != nil {
		log.Println("Failed to dial dummy service | ", error)
		return nil, error
	}
	defer conn.Close()

	client := dummy.NewDummyServiceClient(conn)
	response, error := client.GetHello(context.Background(), &empty.Empty{})
	if error != nil {
		log.Println("Failed to request dummy.GetHello | ", error)
		return nil, error
	}
	
	return response, nil
}

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterDummerServiceServer(grpcServer, &dummerServer{})

	dialOpts := []grpc.DialOption{ grpc.WithInsecure() }

	ctx := context.Background()
	gatewayMux := runtime.NewServeMux()
	error := pb.RegisterDummerServiceHandlerFromEndpoint(ctx, gatewayMux, host, dialOpts)
	if error != nil {
		log.Fatal("Failed to register service handler from endpoint | ", error)
	}

	httpMux := http.NewServeMux()
	httpMux.Handle("/world", gatewayMux)

	httpHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			httpMux.ServeHTTP(w, r)
		}
	})

	http2Server := &http2.Server{}
	http1Server := &http.Server{
		Addr: host,
		Handler: h2c.NewHandler(httpHandler, http2Server),
	}

	conn, error := net.Listen("tcp", host)
	if error != nil {
		log.Fatal("Failed to start tcp connection | ", error)
	}

	error = http1Server.Serve(conn)
	if error != nil {
		log.Fatal("Failed to start http1 server | ", error)
	}
}

