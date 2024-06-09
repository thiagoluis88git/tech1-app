package environment

import "flag"

var (
	DbHost          *string = flag.String("dbHost", "service-database-internal", "Database host")
	RedocFolderPath *string = flag.String("redocFolderPath", "/docs/swagger.json", "Swagger docs folder")
)
