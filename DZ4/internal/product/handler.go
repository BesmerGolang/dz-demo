package product

import (
	"dz4/pkg/req"
	"dz4/pkg/res"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type ProductHandler struct {
	ProdRepo *ProdRepo
}
type ProductHandlerDeps struct {
	ProdRepo *ProdRepo
}

func NewProdHandler(router *http.ServeMux, deps ProductHandlerDeps) {
	handler := &ProductHandler{
		ProdRepo: deps.ProdRepo,
	}
	router.HandleFunc("POST /product", handler.Create())
	router.HandleFunc("PATCH /product/{id}", handler.Update())
	router.HandleFunc("DELETE /product/{id}", handler.Delete())
	router.HandleFunc("GET /product/{id}", handler.GetById())
}
func (handler *ProductHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		existedProduct, err := handler.ProdRepo.GetById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if existedProduct == nil {
			res.Json(w, nil, 404)
		}
		res.Json(w, existedProduct, 200)
	}
}
func (handler *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[req.ProductReq](w, r)
		if err != nil {
			return
		}
		product := &Product{
			Name:        body.Name,
			Description: body.Description,
			Images:      body.Images,
		}
		createdProduct, err := handler.ProdRepo.Create(product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, createdProduct, 201)
	}
}
func (handler *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[req.ProductReq](w, r)
		if err != nil {
			return
		}
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		product, err := handler.ProdRepo.Update(&Product{
			Model:       gorm.Model{ID: uint(id)},
			Name:        body.Name,
			Description: body.Description,
			Images:      body.Images,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		res.Json(w, product, 201)
	}
}

func (handler *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = handler.ProdRepo.GetById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		err = handler.ProdRepo.Delete(&Product{}, uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Json(w, nil, 200)

	}
}
