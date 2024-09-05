// C:\GoProject\src\eShop\db\security.go

package db

import (
	"eShop/configs"
	"fmt"
	"os"
)

func securityConfig() string {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		configs.AppSettings.PostgresParams.Host,
		configs.AppSettings.PostgresParams.Port,
		configs.AppSettings.PostgresParams.User,
		configs.AppSettings.PostgresParams.Database,
		os.Getenv("DB_PASSWORD"))

	return dsn
}
