package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Fabriciope/my-api/internal/dto"
	"github.com/Fabriciope/my-api/internal/infra/database/repositories"
	"github.com/Fabriciope/my-api/internal/services"
	"github.com/go-chi/chi/v5"
)

type productHandler struct {
	repository repositories.RepositoryInterface
	service    *services.ProductService
}

func newProductHandler(repository repositories.RepositoryInterface) *productHandler {
	return &productHandler{
		repository: repository,
		service:    services.NewProductService(repository),
	}
}

// Create godoc
//
//	@Summary		Create a product
//	@Description	Create a new product
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.CreateProductInput	true	"product request"
//	@Success		201		{object}	dto.DefaultOutput
//	@Failure		400		{object}	dto.DefaultOutput
//	@Router			/product/create [post]
//	@Security		ApiKeyAuth
func (h *productHandler) Create(w http.ResponseWriter, r *http.Request) {
	var productDTO dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&productDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorToJson("invalid parameters"))
		return
	}

	err = h.service.CreateProduct(&productDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorToJson(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(successToJson("product created"))
}

// Update godoc
//
//	@Summary		Update product
//	@Description	Update a product
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.UpdateProductInput	true	"update product request"
//	@Success		200		{object}	dto.DefaultOutput
//	@Failure		400		{object}	dto.DefaultOutput
//	@Router			/product/update/{id} [put]
//	@Security		ApiKeyAuth
func (h *productHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorToJson("id is required"))
		return
	}

	var productDTO dto.UpdateProductInput
	err := json.NewDecoder(r.Body).Decode(&productDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorToJson("invalid parameters"))
		return
	}

	err = h.service.UpdateProduct(id, &productDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorToJson(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(successToJson("product updated"))
}

// Delete godoc
//
//	@Summary		Delete product
//	@Description	Delete a product
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"product id"
//	@Success		200	{object}	dto.DefaultOutput
//	@Failure		400	{object}	dto.DefaultOutput
//	@Router			/product/delete/{id} [delete]
//	@Security		ApiKeyAuth
func (h *productHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if id := chi.URLParam(r, "id"); id != "" {
		err := h.service.DeleteProduct(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errorToJson(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(successToJson("product deleted"))
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write(errorToJson("id is required"))
}

// Get godoc
//
//	@Summary		Get a product
//	@Description	Get a product
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"product id"
//	@Success		200	{object}	models.Product
//	@Failure		404	{object}	dto.DefaultOutput
//	@Failure		400	{object}	dto.DefaultOutput
//	@Router			/product/{id} [get]
//	@Security		ApiKeyAuth
func (h *productHandler) Get(w http.ResponseWriter, r *http.Request) {
	if id := chi.URLParam(r, "id"); id != "" {
		productFound, err := h.repository.FindOneWhere("id", chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(errorToJson("product not found"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(successWithDataToJson("product found", productFound))
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write(errorToJson("id is required"))
}

// GetAll godoc
//
//	@Summary		Get all products
//	@Description	Get all products
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			page	path		string	true	"page"
//	@Param			limit	path		string	true	"limit"
//	@Param			sort	query		string	false	"sort"
//	@Success		200		{object}	dto.AllProductsOutput
//	@Failure		422		{object}	dto.DefaultOutput
//	@Failure		500		{object}	dto.DefaultOutput
//	@Router			/product/all/{page}/{limit} [get]
//	@Security		ApiKeyAuth
func (h *productHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// page, errPage := strconv.Atoi(r.URL.Query().Get("page"))
	// limit, errLimit := strconv.Atoi(r.URL.Query().Get("limit"))
	page, errPage := strconv.Atoi(chi.URLParam(r, "page"))
	limit, errLimit := strconv.Atoi(chi.URLParam(r, "limit"))
	if errPage != nil || errLimit != nil || page <= 0 || limit <= 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(errorToJson("the page and limit parameters must be numbers and above 0"))
		return
	}

	sort := r.URL.Query().Get("sort")
	products, err := h.service.GetAllWithPagination(page, limit, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errorToJson(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(successWithDataToJson(
		"products found",
		dto.AllProductsOutput{
			Page:     uint(page),
			Limit:    uint(limit),
			Sort:     sort,
			Products: products,
		},
	))
}
