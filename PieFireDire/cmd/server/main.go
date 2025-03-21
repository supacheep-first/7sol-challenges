// cmd/server/main.go
package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"piefiredire/internal/services"
	"piefiredire/proto/summary"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	go startGRPCServer()

	r := gin.Default()
	r.POST("/beef/summary", func(c *gin.Context) {
		var requestBody struct {
			Text string `json:"text"`
		}

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		if requestBody.Text == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Text is required"})
			return
		}

		conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to gRPC server"})
			return
		}
		defer conn.Close()

		client := summary.NewSummaryServiceClient(conn)
		resp, err := client.GetSummary(context.Background(), &summary.SummaryRequest{Text: requestBody.Text})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get summary"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"beef": resp.WordCount})
	})

	log.Println("HTTP server is running on port 8080...")
	r.Run(":8080")
}

func startGRPCServer() {
	listener, err := net.Listen("tcp", ":50051")
	if (err != nil) {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	summary.RegisterSummaryServiceServer(grpcServer, &services.SummaryServiceServer{})
	reflection.Register(grpcServer)

	log.Println("gRPC server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
