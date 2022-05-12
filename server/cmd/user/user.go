package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"pinterest/domain"
	userapp "pinterest/services/user/application"
	users3 "pinterest/services/user/clients/s3"
	userrepo "pinterest/services/user/infrastructure"
	userfacade "pinterest/services/user/interfaces"
	userproto "pinterest/services/user/proto"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func runService(addr string) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	sugarLogger := logger.Sugar()

	err := godotenv.Load(".env")
	if err != nil {
		sugarLogger.Fatal("Could not load .env file", zap.String("error", err.Error()))
	}

	err = godotenv.Load("passwords.env")
	if err != nil {
		sugarLogger.Fatal("Could not load passwords.env file", zap.String("error", err.Error()))
	}

	err = godotenv.Load("s3.env")
	if err != nil {
		sugarLogger.Fatal("Could not load s3.env file", zap.String("error", err.Error()))
	}

	err = godotenv.Load("docker_vars.env")
	if err != nil {
		sugarLogger.Fatal("Could not load docker_vars.env file", zap.String("error", err.Error()))
	}

	dbPrefix := os.Getenv("DB_PREFIX")
	if dbPrefix != "AMAZON" && dbPrefix != "LOCAL" {
		sugarLogger.Fatalf("Wrong prefix: %s , should be AMAZON or LOCAL", dbPrefix)
	}

	postgresConnectionString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv(dbPrefix+"_DB_USER"), os.Getenv(dbPrefix+"_DB_PASSWORD"), os.Getenv(dbPrefix+"_DB_HOST"),
		os.Getenv(dbPrefix+"_DB_PORT"), os.Getenv(dbPrefix+"_DB_NAME"))
	postgresConn, err := pgxpool.Connect(context.Background(), postgresConnectionString)
	if err != nil {
		sugarLogger.Fatal("Could not connect to postgres database", zap.String("error", err.Error()))
		return
	}

	fmt.Println("Successfully connected to postgres database")
	defer postgresConn.Close()

	dockerStatus := os.Getenv("CONTAINER_PREFIX")
	if dockerStatus != "DOCKER" && dockerStatus != "LOCALHOST" {
		sugarLogger.Fatalf("Wrong prefix: %s , should be DOCKER or LOCALHOST", dockerStatus)
	}

	server := grpc.NewServer()

	s3session, err := domain.ConnectAws()
	if err != nil {
		sugarLogger.Fatal("Could not connect to s3 bucket", zap.String("error", err.Error()))
	}

	service := userfacade.NewUserFacade(userapp.NewUserApp(
		userrepo.NewUserRepo(postgresConn), users3.NewS3Client(s3session, os.Getenv("AWS_BUCKET_NAME"))))
	userproto.RegisterUserServer(server, service)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Listen user error", err)
	}

	fmt.Printf("Starting server at localhost%s\n", addr)
	err = server.Serve(lis)
	if err != nil {
		log.Fatalln("Serve user error", err)
	}
}

func main() {
	runService(":8082")
}
