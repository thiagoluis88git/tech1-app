package environment

import "flag"

var (
	DbHost          *string = flag.String("DB_HOST", "service-database-internal", "Database host")
	DbPort          *string = flag.String("DB_PORT", "5432", "Database Port")
	DbUser          *string = flag.String("POSTGRES_USER", "fastfood", "Database Port")
	DbName          *string = flag.String("POSTGRES_DB", "fastfood_db", "Database Port")
	DbPassword      *string = flag.String("POSTGRES_PASSWORD", "fastfood1234", "Database Port")
	RedocFolderPath *string = flag.String("PATH_REDOC_FOLDER", "/docs/swagger.json", "Swagger docs folder")
)
