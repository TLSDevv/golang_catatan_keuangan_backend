package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/api/mobile/controller"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/api/mobile/service"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/api/repository"
	config "github.com/TLSDevv/golang_catatan_keuangan_backend/config"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/exception"
	pkgconstant "github.com/TLSDevv/golang_catatan_keuangan_backend/pkg/constant"
	pkgdb "github.com/TLSDevv/golang_catatan_keuangan_backend/pkg/db"
	"github.com/go-playground/validator"
	"github.com/rs/cors"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.New()

	if err != nil {
		log.Fatal(err)
	}

	serve(config)
}

func serve(config *config.Config) {
	addr := fmt.Sprintf("%s:%d", config.API.Host, config.API.Port)

	server := &http.Server{
		Addr:    addr,
		Handler: newHandler(config),
	}

	fmt.Println("server running on", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func newHandler(config *config.Config) http.Handler {
	db, err := initDB(config)

	if err != nil {
		log.Fatal(err)
	}

	api := NewAPI(config)
	api.initMobileRoute(config, db)

	return api.router
}

type API struct {
	router   *chi.Mux
	validate *validator.Validate
}

func NewAPI(config *config.Config) *API {
	validate := validator.New()
	router := chi.NewRouter()

	router.Use(initCors(config).Handler)
	router.Use(exception.Recover)

	//swagger ui, aksesnya pake /docs/
	fs := http.FileServer(http.Dir("./dist"))
	router.Handle("/docs/*", http.StripPrefix("/docs", fs))

	return &API{
		router:   router,
		validate: validate,
	}
}

func (api *API) initMobileRoute(config *config.Config, db *sql.DB) {
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo, db, api.validate)
	userController := controller.NewUserController(userService)

	transactionRepo := repository.NewTransactionRepository()
	transactionService := service.NewTransactioService(transactionRepo, db, api.validate)
	transactionController := controller.NewTransactionController(transactionService)

	categoryRepo := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepo, db, api.validate)
	categoryController := controller.NewCategoryController(categoryService)

	// route

	api.router.Route("/api/v1", func(r chi.Router) {
		r.Route("/", func(r chi.Router) {
			r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
				rw.Write([]byte("works"))
			}) // testing api
		})
		r.Route("/users", func(r chi.Router) {
			r.Get("/{id}", userController.GetUser)
			r.Post("/", userController.CreateUser)
			r.Put("/{id}", userController.UpdateUser)
		})
		r.Route("/categories", func(r chi.Router) {
			r.Get("/{id}", categoryController.GetCategory)
			r.Get("/", categoryController.ListCategory)
			r.Post("/", categoryController.CreateCategory)
			r.Put("/{id}", categoryController.UpdateCategory)
			r.Delete("/{id}", categoryController.DeleteCategory)
		})
		r.Route("/transactions", func(r chi.Router) {
			r.Get("/{id}", transactionController.GetTransaction)
			r.Get("/", transactionController.ListTransaction)
			r.Post("/", transactionController.CreateTransaction)
			r.Put("/{id}", transactionController.UpdateTransaction)
			r.Delete("/{id}", transactionController.DeleteTransaction)
		})
	})

}

func initDB(config *config.Config) (*sql.DB, error) {
	dbParam := pkgdb.DBParam{
		Host:     config.DB.Host,
		Port:     config.DB.Port,
		User:     config.DB.User,
		Password: config.DB.Password,
		Name:     config.DB.Name,
	}

	return pkgdb.InitDB(dbParam)
}

func initCors(config *config.Config) *cors.Cors {
	// TODO : add api config

	return cors.New(cors.Options{
		AllowedOrigins: pkgconstant.DefaultAllowdOrigins,
		AllowedMethods: pkgconstant.DefaultAllowedMethods,
		AllowedHeaders: pkgconstant.DefaultAllowedHeaders,
	})
}
