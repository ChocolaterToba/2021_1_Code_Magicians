package main

import (
	"fmt"
	"net/http"
	"os"

	authclient "pinterest/clients/auth"
	productclient "pinterest/clients/product"
	userclient "pinterest/clients/user"
	authfacade "pinterest/interfaces/auth"
	productfacade "pinterest/interfaces/product"
	profilefacade "pinterest/interfaces/profile"
	"pinterest/interfaces/routing"
	authproto "pinterest/services/auth/proto"
	productproto "pinterest/services/product/proto"
	userproto "pinterest/services/user/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/rs/cors"
)

func runServer(addr string) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	sugarLogger := logger.Sugar()

	err := loadSettings()
	if err != nil {
		sugarLogger.Fatal(err.Error())
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

	sessionUser, err := grpc.Dial(os.Getenv(dockerStatus+"_USER_PREFIX")+":8082", grpc.WithInsecure())
	if err != nil {
		sugarLogger.Fatal("Can not create session for User service")
	}
	defer sessionUser.Close()

	sessionProduct, err := grpc.Dial(os.Getenv(dockerStatus+"_PRODUCT_PREFIX")+":8083", grpc.WithInsecure())
	if err != nil {
		sugarLogger.Fatal("Can not create session for Product service")
	}
	defer sessionProduct.Close()

	authClient := authclient.NewAuthClient(authproto.NewAuthClient(sessionAuth), os.Getenv("HTTPS_ON") == "true")
	userClient := userclient.NewUserClient(userproto.NewUserClient(sessionUser))
	productClient := productclient.NewProductClient(productproto.NewProductServiceClient(sessionProduct))

	authFacade := authfacade.NewAuthFacade(authClient, logger)
	profileFacade := profilefacade.NewProfileFacade(userClient, authClient, logger)
	productFacade := productfacade.NewProductFacade(productClient, logger)

	r := routing.CreateRouter(authClient, authFacade, profileFacade, productFacade, os.Getenv("CSRF_ON") == "true")

	allowedOrigins := make([]string, 0)
	switch os.Getenv("HTTPS_ON") {
	case "true":
		allowedOrigins = append(allowedOrigins, "https://gears4us.ru:8081", "https://gears4us.ru",
			"https://127.0.0.1:8081", "https://51.250.76.99", "https://localhost:3001") // TODO: replace with actual
	case "false":
		allowedOrigins = append(allowedOrigins, "http://gears4us.ru:8081", "http://gears4us.ru",
			"http://127.0.0.1:8081", "http://51.250.76.99", "http://localhost:3001")
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

	switch os.Getenv("SERVE_HTTPS_ON") {
	case "true":
		sugarLogger.Fatal(http.ListenAndServeTLS(addr, "cert.pem", "key.pem", handler))
	case "false":
		sugarLogger.Fatal(http.ListenAndServe(addr, handler))
	default:
		sugarLogger.Fatal("SERVE_HTTPS_ON variable is not set")
	}
}

func main() {
	runServer(":8080")
}

func loadSettings() (err error) {
	err = godotenv.Load(".env")
	if err != nil {
		return errors.Wrap(err, "Could not load .env file")
	}

	err = godotenv.Load("passwords.env")
	if err != nil {
		return errors.Wrap(err, "Could not load passwords.env file")
	}

	// err = godotenv.Load("s3.env")
	// if err != nil {
	// 	sugarLogger.Fatal("Could not load s3.env file", zap.String("error", err.Error()))
	// }

	err = godotenv.Load("docker_vars.env")
	if err != nil {
		return errors.Wrap(err, "Could not load docker_vars.env file")
	}

	err = godotenv.Load("vk_info.env")
	if err != nil {
		return errors.Wrap(err, "Could not load vk_info.env file")
	}

	return nil
}
