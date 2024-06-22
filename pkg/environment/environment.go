package environment

import (
	"flag"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	DbHost          *string = flag.String("DB_HOST", "service-database-internal", "Database host")
	DbPort          *string = flag.String("DB_PORT", "5432", "Database Port")
	DbUser          *string = flag.String("POSTGRES_USER", "fastfood", "Database Port")
	DbName          *string = flag.String("POSTGRES_DB", "fastfood_db", "Database Port")
	DbPassword      *string = flag.String("POSTGRES_PASSWORD", "fastfood1234", "Database Port")
	RedocFolderPath *string = flag.String("PATH_REDOC_FOLDER", "/docs/swagger.json", "Swagger docs folder")

	localDev = flag.String("localDev", "false", "local development")

	singleton *Environment
)

const (
	QRCodeGatewayRootURL = "QR_CODE_GATEWAY_ROOT_URL"
	QRCodeGatewayToken   = "QR_CODE_GATEWAY_TOKEN"
)

type Environment struct {
	qrCodeGatewayRootURL string
	qrCodeGatewayToken   string
}

func LoadEnvironmentVariables() {
	flag.Parse()

	if localFlag := *localDev; localFlag != "false" {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading .env file", err.Error())
		}
	}

	qrCodeGatewayRootURL := getEnvironmentVariable(QRCodeGatewayRootURL)
	qrCodeGatewayToken := getEnvironmentVariable(QRCodeGatewayToken)

	once := &sync.Once{}

	once.Do(func() {
		singleton = &Environment{
			qrCodeGatewayRootURL: qrCodeGatewayRootURL,
			qrCodeGatewayToken:   qrCodeGatewayToken,
		}
	})
}

func getEnvironmentVariable(key string) string {
	value, hashKey := os.LookupEnv(key)

	if !hashKey {
		log.Fatalf("There is no %v environment variable", key)
	}

	return value
}

func GetQRCodeGatewayRootURL() string {
	if singleton != nil {
		return singleton.qrCodeGatewayRootURL
	}

	return getEnvironmentVariable(QRCodeGatewayRootURL)
}

func GetQRCodeGatewayToken() string {
	if singleton != nil {
		return singleton.qrCodeGatewayToken
	}

	return getEnvironmentVariable(QRCodeGatewayToken)
}
