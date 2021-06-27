package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	pb "github.com/rafaelturon/seed-converter/proto"
	"github.com/rafaelturon/seed-converter/seedgen"
	"github.com/tyler-smith/go-bip39"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement ApiServer.
type server struct {
	pb.UnimplementedApiServer
}

func (s *server) GetSeed(ctx context.Context, in *pb.SeedRequest) (*pb.SeedReply, error) {
	log.Printf("Received %v", in.GetCoin())
	seed, _ := seedgen.GenerateRandomSeed(seedgen.RecommendedSeedLen)
	seedStr, _ := bip39.NewMnemonic(seed)
	return &pb.SeedReply{Seed: seedStr}, nil
}

func main() {

	fmt.Println("gRPC Server")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	cert, _ := ioutil.ReadFile("../../conf/certs/client.crt")
	privateKey, _ := ioutil.ReadFile("../../conf/certs/client.key")
	clientCert, _ := tls.X509KeyPair(cert, privateKey)

	c := credentials.NewTLS(&tls.Config{Certificates: []tls.Certificate{clientCert}})
	s := grpc.NewServer(grpc.Creds(c))
	reflection.Register(s)

	pb.RegisterApiServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
