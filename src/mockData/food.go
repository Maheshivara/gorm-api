package mockdata

import (
	"errors"
	"fmt"
	"gormCompose/src/messages"
	"gormCompose/src/models"
	"gormCompose/src/request"
	"gormCompose/src/response"
	"time"

	"github.com/google/uuid"
)

var (
	FoodDefaultId    = uuid.MustParse("cb5cf6bb-f7ae-4847-a7e9-a11f6e9ee34a")
	DefaultDate      = time.Date(2025, time.February, 11, 0, 0, 0, 0, time.UTC)
	ValidFoodRequest = &request.Food{Name: "Macarroon", Price: 15.90}

	NoNameFoodRequest       = &request.Food{Price: 15.90}
	NoPriceFoodRequest      = &request.Food{Name: "Macarroon"}
	InvalidNameFoodRequest  = &request.Food{Name: "A", Price: 15.90}
	InvalidPriceFoodRequest = &request.Food{Name: "Macarrão", Price: -15.90}

	SuccessFoodCreateResponse = &models.Food{ID: FoodDefaultId, Name: "Macarroon", Price: 15.90, CreatedAt: DefaultDate, UpdatedAt: DefaultDate, DeletedAt: nil}
	SuccessFoodListResponse   = &response.SearchResult{Data: []*models.Food{SuccessFoodCreateResponse}, PerPage: 20, Total: 1, Page: 1}
	NoNameFoodErrorResponse   = `{"error":"Key: 'Food.Name' Error:Field validation for 'Name' failed on the 'required' tag"}`
	NoPriceFoodErrorResponse  = `{"error":"Key: 'Food.Price' Error:Field validation for 'Price' failed on the 'required' tag"}`
	InvalidNameErrorResponse  = `{"error":"Key: 'Food.Name' Error:Field validation for 'Name' failed on the 'min' tag"}`
	InvalidPriceErrorResponse = `{"error":"Key: 'Food.Price' Error:Field validation for 'Price' failed on the 'gt' tag"}`
	InternalErrorResponse     = fmt.Sprintf(`{"error":"%s"}`, messages.INTERNAL_SERVER_ERROR)
	NotFoundErrorResponse     = fmt.Sprintf(`{"error":"%s"}`, messages.NOT_FOUND_ERROR)

	ErrInternal = errors.New("internal error")
)
