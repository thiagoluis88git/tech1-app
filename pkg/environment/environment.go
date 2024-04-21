package environment

import "flag"

var (
	DbHost *string = flag.String("dbHost", "database", "Database host")
)
