// C:\GoProject\src\eShop\db\connection.go

// package db

// import (
// 	"eShop/configs"
// 	"eShop/logger"
// 	"fmt"
// 	"os"
// 	"time"

// 	_ "github.com/jinzhu/gorm/dialects/postgres"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var dbConn *gorm.DB

// // ConnectToDB устанавливает соединение с базой данных PostgreSQL и учитывает временную зону...
// func ConnectToDB() error {
// 	// Формируем строку подключения с указанием временной зоны Asia/Tashkent, так как для Душанбе в Windows временная зона отсутствует...
// 	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s TimeZone=Asia/Dushanbe",
// 		configs.AppSettings.PostgresParams.Host,
// 		configs.AppSettings.PostgresParams.Port,
// 		configs.AppSettings.PostgresParams.User,
// 		configs.AppSettings.PostgresParams.Database,
// 		os.Getenv("DB_PASSWORD"))

// 	// Логируем попытку подключения к базе данных...
// 	logger.Info.Println("Connecting to database...")

// 	// Открываем соединение с базой данных с использованием GORM...
// 	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{NowFunc: func() time.Time {
// 		return time.Now().Local()
// 	}})
// 	if err != nil {
// 		// Логируем ошибку, если не удалось подключиться к базе данных...
// 		logger.Error.Printf("Failed to connect to database: %v", err)
// 		return err
// 	}

// 	// Логируем успешное подключение к базе данных...
// 	logger.Info.Println("Successfully connected to database!")

// 	// Сохраняем подключение для дальнейшего использования...
// 	dbConn = db
// 	return nil
// }

// func CloseDBConn() error {
// 	sqlDB, err := dbConn.DB()
// 	if err != nil {
// 		logger.Error.Printf("Failed to retrieve SQL DB from gorm.DB: %v", err)
// 		return err
// 	}
// 	err = sqlDB.Close()
// 	if err != nil {
// 		logger.Error.Printf("Failed to close database connection: %v", err)
// 		return err
// 	}
// 	logger.Info.Println("\nDatabase connection closed successfully!")
// 	return nil
// }

// func GetDBConn() *gorm.DB {
// 	return dbConn
// }

package db

import (
	"eShop/configs"
	"eShop/logger"
	"fmt"
	"os"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

// EnsureDatabaseExists проверяет наличие базы данных и создает её, если она не существует.
func EnsureDatabaseExists() error {
	// Формируем строку подключения к базе данных postgres для создания новой базы данных...
	createDBConnStr := fmt.Sprintf("host=%s port=%s user=%s dbname=postgres password=%s TimeZone=Asia/Dushanbe",
		configs.AppSettings.PostgresParams.Host,
		configs.AppSettings.PostgresParams.Port,
		configs.AppSettings.PostgresParams.User,
		os.Getenv("DB_PASSWORD"))

	// Логируем попытку подключения для создания базы данных...
	logger.Info.Println("Connecting to database to check if the database exists...")

	// Открываем соединение с базой данных postgres...
	tempDB, err := gorm.Open(postgres.Open(createDBConnStr), &gorm.Config{NowFunc: func() time.Time {
		return time.Now().Local()
	}})
	if err != nil {
		logger.Error.Printf("Failed to connect to database 'postgres': %v", err)
		return err
	}

	// Проверяем, существует ли база данных...
	var exists bool
	checkDBQuery := fmt.Sprintf("SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = '%s')", configs.AppSettings.PostgresParams.Database)
	tempDB.Raw(checkDBQuery).Scan(&exists)

	// Если базы данных не существует, создаем её...
	if !exists {
		createDBQuery := fmt.Sprintf("CREATE DATABASE %s", configs.AppSettings.PostgresParams.Database)
		if err := tempDB.Exec(createDBQuery).Error; err != nil {
			logger.Error.Printf("Failed to create database '%s': %v", configs.AppSettings.PostgresParams.Database, err)
			return err
		}
		logger.Info.Printf("Database '%s' created successfully!", configs.AppSettings.PostgresParams.Database)
	} else {
		logger.Info.Printf("Database '%s' already exists, skipping creation.", configs.AppSettings.PostgresParams.Database)
	}

	// Закрываем временное соединение...
	sqlTempDB, _ := tempDB.DB()
	sqlTempDB.Close()

	return nil
}

// ConnectToDB устанавливает соединение с базой данных PostgreSQL и учитывает временную зону...
func ConnectToDB() error {
	// Сначала убеждаемся, что база данных существует.
	if err := EnsureDatabaseExists(); err != nil {
		return err
	}

	// Теперь подключаемся к базе данных.
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s TimeZone=Asia/Dushanbe",
		configs.AppSettings.PostgresParams.Host,
		configs.AppSettings.PostgresParams.Port,
		configs.AppSettings.PostgresParams.User,
		configs.AppSettings.PostgresParams.Database,
		os.Getenv("DB_PASSWORD"))

	logger.Info.Println("Connecting to the newly created or existing database...")

	// Открываем постоянное соединение с новой базой данных...
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{NowFunc: func() time.Time {
		return time.Now().Local()
	}})
	if err != nil {
		logger.Error.Printf("Failed to connect to the new database: %v", err)
		return err
	}

	logger.Info.Println("Successfully connected to the database!")

	dbConn = db
	return nil
}

func CloseDBConn() error {
	sqlDB, err := dbConn.DB()
	if err != nil {
		logger.Error.Printf("Failed to retrieve SQL DB from gorm.DB: %v", err)
		return err
	}
	err = sqlDB.Close()
	if err != nil {
		logger.Error.Printf("Failed to close database connection: %v", err)
		return err
	}
	logger.Info.Println("\nDatabase connection closed successfully!")
	return nil
}

func GetDBConn() *gorm.DB {
	return dbConn
}
