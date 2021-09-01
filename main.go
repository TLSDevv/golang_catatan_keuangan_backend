package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/api/mobile"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/api/mobile/controller"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/api/mobile/service"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/api/repository"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool("debug") {
		fmt.Println("Service RUN on Debug mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	dbDriver := viper.GetString(`database.driver`)
	serverPort := viper.GetString(`server.port`)

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	db, err := sql.Open(dbDriver, dsn)

	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db)
	categoryController := controller.NewCategoryController(categoryService)

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db)
	userController := controller.NewUserController(userService)

	transactionRepository := repository.NewTransactionRepository()
	transactionService := service.NewTransactioService(transactionRepository, db)
	transactionController := controller.NewTransactionController(transactionService)

	handler := &mobile.Handler{}
	handler.NewHandler(userController, categoryController, transactionController)

	fmt.Printf("Listening to port %s", serverPort)
	http.ListenAndServe(serverPort, handler.Route)
}
