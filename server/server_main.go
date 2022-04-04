package main

import (
	"fmt"
	"net/http"
	"os"
	authclient "pinterest/clients/auth"
	authfacade "pinterest/interfaces/auth"
	"pinterest/interfaces/routing"
	authproto "pinterest/services/auth/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func runServer(addr string) {
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

	// err = godotenv.Load("s3.env")
	// if err != nil {
	// 	sugarLogger.Fatal("Could not load s3.env file", zap.String("error", err.Error()))
	// }

	err = godotenv.Load("docker_vars.env")
	if err != nil {
		sugarLogger.Fatal("Could not load docker_vars.env file", zap.String("error", err.Error()))
	}

	err = godotenv.Load("vk_info.env")
	if err != nil {
		sugarLogger.Fatal("Could not load vk_info.env file", zap.String("error", err.Error()))
	}

	dockerStatus := os.Getenv("CONTAINER_PREFIX")
	if dockerStatus != "DOCKER" && dockerStatus != "LOCALHOST" {
		sugarLogger.Fatalf("Wrong prefix: %s , should be DOCKER or LOCALHOST", dockerStatus)
	}

	// sess := entity.ConnectAws()
	// TODO divide file

	sessionAuth, err := grpc.Dial(os.Getenv(dockerStatus+"_AUTH_PREFIX")+":8081", grpc.WithInsecure())
	if err != nil {
		sugarLogger.Fatal("Can not create session for Auth service")
	}
	defer sessionAuth.Close()

	authClient := authclient.NewAuthClient(authproto.NewAuthClient(sessionAuth))

	authFacade := authfacade.NewAuthFacade(authClient, logger)
	// TODO divide file

	r := routing.CreateRouter(authClient, authFacade, os.Getenv("CSRF_ON") == "true", os.Getenv("HTTPS_ON") == "true")

	allowedOrigins := make([]string, 0)
	switch os.Getenv("HTTPS_ON") {
	case "true":
		allowedOrigins = append(allowedOrigins, "https://pinterbest.ru:8081", "https://pinterbest.ru",
			"https://127.0.0.1:8081", "https://164.90.222.152") // TODO: replace with actual
	case "false":
		allowedOrigins = append(allowedOrigins, "http://pinterbest.ru:8081", "http://pinterbest.ru",
			"http://127.0.0.1:8081", "http://164.90.222.152")
	default:
		sugarLogger.Fatal("HTTPS_ON variable is not set")
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	})

	handler := c.Handler(r)
	fmt.Printf("Starting server at localhost%s\n", addr)

	switch os.Getenv("HTTPS_ON") {
	case "true":
		sugarLogger.Fatal(http.ListenAndServeTLS(addr, "cert.pem", "key.pem", handler))
	case "false":
		sugarLogger.Fatal(http.ListenAndServe(addr, handler))
	}
}

func main() {
	runServer(":8080")
}
