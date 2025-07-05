package main

import (
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	pb "progress-tracker/api/proto/service"
	interceptors "progress-tracker/cmd/GRPC/interceprtors"
	"progress-tracker/internal/config"
	"progress-tracker/internal/handlers"
	"progress-tracker/internal/services"
)

func main() {
	config.Configurate()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("server start error: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		interceptors.AuthInterceptor,
	))

	dbURL := viper.GetString("database.url")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {

		log.Fatal(err)
	}

	jobService := services.NewJobService(db)
	progressService := services.NewProgressService()
	progressService.StartQueueWorker()

	pb.RegisterJobServiceServer(grpcServer, handlers.NewJobRpcServer(jobService))

	log.Println("grpc server started on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("server error: %v", err)
	}

}
