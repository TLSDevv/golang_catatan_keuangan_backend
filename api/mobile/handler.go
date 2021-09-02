package mobile

import (
	"encoding/json"
	"net/http"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/api/mobile/controller"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/exception"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/web"
	"github.com/go-chi/chi"
)

type Handler struct {
	Route *chi.Mux
}

func (h *Handler) NewHandler(uc controller.UserControllerInterface, cc controller.CategoryControllerInterface, tc controller.TransactionControllerInterface) {
	h.Route = chi.NewRouter()
	h.Route.Use(exception.Recover)
	h.Route.NotFound(func(w http.ResponseWriter, r *http.Request) {
		panic(exception.NewNotFoundError("WRONG URL"))
	})
	h.Route.Route("/api/v1", func(r chi.Router) {
		r.Route("/", func(r chi.Router) {
			r.Get("/", Home) // testing api
		})
		r.Route("/user", func(r chi.Router) {
			r.Get("/{id}", uc.GetUser)
			r.Post("/", uc.CreateUser)
			r.Put("/", uc.UpdateUser)
		})
		r.Route("/category", func(r chi.Router) {
			r.Get("/:id", cc.GetCategory)
			r.Get("/", cc.ListCategory)
			r.Post("/", cc.CreateCategory)
			r.Put("/:id", cc.UpdateCategory)
			r.Delete("/:id", cc.DeleteCategory)
		})
		r.Route("/transaction", func(r chi.Router) {
			r.Get("/:id", tc.GetTransaction)
			r.Get("/", tc.ListTransaction)
			r.Post("/", tc.CreateTransaction)
			r.Put("/:id", tc.UpdateTransaction)
			r.Delete("/:id", tc.DeleteTransaction)
		})
	})
}

func Home(w http.ResponseWriter, r *http.Request) {
	res := web.WebResponse{
		Code:   http.StatusOK,
		Data:   nil,
		Status: "Ini Halaman Home",
	}
	response, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
