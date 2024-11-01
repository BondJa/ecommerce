package products

import (
	"fmt"
	"net/http"

	"github.com/HimandriSharma/ecommerce/types"
	"github.com/HimandriSharma/ecommerce/utils"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}
func (h *Handler) RegisterRouter(router *mux.Router) {
	router.HandleFunc("/products", h.handleCreateProduct).Methods(http.MethodGet)
	router.HandleFunc("/products", h.handleCreateProduct).Methods(http.MethodPost)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var payload types.CreateProductPayload
		if err := utils.ParseJSON(r, &payload); err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
		if err := utils.Validate.Struct(payload); err != nil {
			errors := err.(validator.ValidationErrors)
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload for %v", errors))
			return
		}
		err := h.store.CreateProduct(types.Product{
			Name:        payload.Name,
			Description: payload.Description,
			Image:       payload.Image,
			Price:       payload.Price,
			Quantity:    payload.Quantity,
		})
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		utils.WriteJSON(w, http.StatusCreated, nil)
	} else {
		ps, err := h.store.GetProducts()
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		utils.WriteJSON(w, http.StatusOK, ps)
	}
}
