package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/api/mobile/controller"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/api/mobile/service"
	config "github.com/TLSDevv/golang_catatan_keuangan_backend/config"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/exception"
	pkg "github.com/TLSDevv/golang_catatan_keuangan_backend/pkg"
	pkgconstant "github.com/TLSDevv/golang_catatan_keuangan_backend/pkg/constant"
	pkgdb "github.com/TLSDevv/golang_catatan_keuangan_backend/pkg/db"
	pkgjwt "github.com/TLSDevv/golang_catatan_keuangan_backend/pkg/jwt"
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
	if err := pkgjwt.Init(config); err != nil {
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

type middleware interface {
	Handler(next http.Handler) http.Handler
}

type API struct {
	router        *chi.Mux
	validate      *validator.Validate
	authorization middleware
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
		router:        router,
		validate:      validate,
		authorization: pkg.NewAuthenticator(),
	}
}

func (api *API) initMobileRoute(config *config.Config, db *sql.DB) {
	userService := service.NewUserService(db, api.validate)
	userController := controller.NewUserController(userService)

	transactionService := service.NewTransactioService(db, api.validate)
	transactionController := controller.NewTransactionController(transactionService)

	categoryService := service.NewCategoryService(db, api.validate)
	categoryController := controller.NewCategoryController(categoryService)

	authService := service.NewAuthService(db)
	authController := controller.NewAuth(authService)

	secureAccessMiddlewares := []func(http.Handler) http.Handler{
		api.authorization.Handler,
	}

	// route

	api.router.Route("/api/v1", func(r chi.Router) {

		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", authController.Login)
			r.With(secureAccessMiddlewares...).Delete("/session", authController.Logout)
			r.Post("/register", userController.CreateUser)
		})
		r.With(secureAccessMiddlewares...).Route("/users", func(r chi.Router) {
			r.Get("/{id}", userController.GetUser)
			r.Put("/{id}", userController.UpdateUser)
		})
		r.With(secureAccessMiddlewares...).Route("/categories", func(r chi.Router) {
			r.Get("/{id}", categoryController.GetCategory)
			r.Get("/", categoryController.ListCategory)
			r.Post("/", categoryController.CreateCategory)
			r.Put("/{id}", categoryController.UpdateCategory)
			r.Delete("/{id}", categoryController.DeleteCategory)
		})
		r.With(secureAccessMiddlewares...).Route("/transactions", func(r chi.Router) {
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
