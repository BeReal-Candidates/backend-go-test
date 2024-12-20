package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pbdiscovery "github.com/BeReal-Candidates/backend-go-test/protob/discovery/v1"

	"github.com/BeReal-Candidates/backend-go-test/discovery"
)

var (
	DefaultPageSize = uint(3)
)

func main() {
	ctx, cf := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)    // interrupt signal sent from terminal
	signal.Notify(sig, syscall.SIGTERM) // sigterm signal sent from kubernetes

	lis, err := CreateListener("localhost", 8080)
	if err != nil {
		panic(err)
	}
	srv := CreateGrpcServer(false)
	discoveryServer, err := discovery.NewServer(discovery.NewService(getPageSize()))
	if err != nil {
		panic(err)
	}
	pbdiscovery.RegisterDiscoveryServiceServer(srv, discoveryServer)

	chanAllClose := make(chan bool, 1)
	go func() {
		<-ctx.Done()
		log.Println("closing resources...")
		srv.GracefulStop()
		lis.Close()
		discoveryServer.Close()
		close(chanAllClose)
	}()
	go func() {
		srvErr := srv.Serve(lis)
		if srvErr != nil {
			panic(srvErr)
		}
		sig <- os.Kill
	}()

	<-sig
	log.Println("received a close signal")
	cf()
	<-chanAllClose
	log.Println("everything is closed, bye")
}

func CreateListener(host string, port int) (net.Listener, error) {
	addr := fmt.Sprintf("%s:%d", host, port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	log.Printf("successfully listening on tcp %s\n", addr)
	return lis, nil
}

func CreateGrpcServer(enableReflection bool) *grpc.Server {
	grpcServer := grpc.NewServer()
	if enableReflection {
		reflection.Register(grpcServer)
	}
	return grpcServer
}

func getPageSize() uint {
	pageSizeStr, exists := os.LookupEnv("PAGE_SIZE")
	if !exists {
		return DefaultPageSize
	}

	i, err := strconv.ParseUint(pageSizeStr, 10, 0)
	if err != nil {
		return DefaultPageSize
	}
	return uint(i)
}
