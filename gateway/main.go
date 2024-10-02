package main

import (
	"log"
	"net/http"

	common "github.com/fantom0052/market/common"
	pb "github.com/fantom0052/market/common/api"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpAddr         = common.EnvString("HTTP_ADDR", ":3000")
	ordersServiceAddr = "localhost:3000"
)

func main() {
	conn, err := grpc.Dial(ordersServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
		log.Fatalf("%v", err)
	}
	defer conn.Close()
	c := pb.NewOrderServiceClient(conn)
	

	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.registerRoutes(mux)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatalf("Failed to star http server %s", err.Error())
	}
}
