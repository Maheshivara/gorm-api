package routers_test

import (
	"encoding/json"
	mockdata "gormCompose/src/mockData"
	"gormCompose/src/routers"
	mock_services "gormCompose/src/services/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestFoodRouter_Create(t *testing.T) {
	testCases := createFoodTest{}
	t.Run("Test when create food is successful", testCases.whenCreateFoodSuccess)
	t.Run("Test when create food has no name", testCases.whenCreateFoodHasNoName)
	t.Run("Test when create food has no price", testCases.whenCreateFoodHasNoPrice)
	t.Run("Test when create food has invalid name", testCases.whenCreateFoodHasInvalidName)
	t.Run("Test when create food has invalid price", testCases.whenCreateFoodHasInvalidPrice)
	t.Run("Test when create food has internal error", testCases.whenCreateFoodHasInternalError)
}

type createFoodTest struct{}

func (*createFoodTest) whenCreateFoodSuccess(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	foodService := mock_services.NewMockFood(controller)
	foodService.EXPECT().Create(gomock.Any()).Return(mockdata.SuccessFoodCreateResponse, nil)
	foodInfo, _ := json.Marshal(mockdata.ValidFoodRequest)
	requestBody := strings.NewReader(string(foodInfo))
	req, err := http.NewRequest("POST", "/api/foods", requestBody)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	foodRouter := routers.FoodRouter(foodService)

	router := gin.New()
	router.POST("/api/foods", foodRouter.Create)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	expectedResponse, _ := json.Marshal(mockdata.SuccessFoodCreateResponse)
	assert.Equal(t, string(expectedResponse), recorder.Body.String())
}

func (*createFoodTest) whenCreateFoodHasNoName(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	foodService := mock_services.NewMockFood(controller)
	foodService.EXPECT().Create(gomock.Any()).Times(0)
	foodInfo, _ := json.Marshal(mockdata.NoNameFoodRequest)
	requestBody := strings.NewReader(string(foodInfo))
	req, err := http.NewRequest("POST", "/api/foods", requestBody)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	foodRouter := routers.FoodRouter(foodService)

	router := gin.New()
	router.POST("/api/foods", foodRouter.Create)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
	assert.Equal(t, mockdata.NoNameFoodErrorResponse, recorder.Body.String())
}

func (*createFoodTest) whenCreateFoodHasNoPrice(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	foodService := mock_services.NewMockFood(controller)
	foodService.EXPECT().Create(gomock.Any()).Times(0)
	foodInfo, _ := json.Marshal(mockdata.NoPriceFoodRequest)
	requestBody := strings.NewReader(string(foodInfo))
	req, err := http.NewRequest("POST", "/api/foods", requestBody)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	foodRouter := routers.FoodRouter(foodService)

	router := gin.New()
	router.POST("/api/foods", foodRouter.Create)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
	assert.Equal(t, mockdata.NoPriceFoodErrorResponse, recorder.Body.String())
}

func (*createFoodTest) whenCreateFoodHasInvalidName(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	foodService := mock_services.NewMockFood(controller)
	foodService.EXPECT().Create(gomock.Any()).Times(0)
	foodInfo, _ := json.Marshal(mockdata.InvalidNameFoodRequest)
	requestBody := strings.NewReader(string(foodInfo))
	req, err := http.NewRequest("POST", "/api/foods", requestBody)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	foodRouter := routers.FoodRouter(foodService)

	router := gin.New()
	router.POST("/api/foods", foodRouter.Create)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
	assert.Equal(t, mockdata.InvalidNameErrorResponse, recorder.Body.String())
}

func (*createFoodTest) whenCreateFoodHasInvalidPrice(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	foodService := mock_services.NewMockFood(controller)
	foodService.EXPECT().Create(gomock.Any()).Times(0)
	foodInfo, _ := json.Marshal(mockdata.InvalidPriceFoodRequest)
	requestBody := strings.NewReader(string(foodInfo))
	req, err := http.NewRequest("POST", "/api/foods", requestBody)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	foodRouter := routers.FoodRouter(foodService)

	router := gin.New()
	router.POST("/api/foods", foodRouter.Create)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
	assert.Equal(t, mockdata.InvalidPriceErrorResponse, recorder.Body.String())
}

func (*createFoodTest) whenCreateFoodHasInternalError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	foodService := mock_services.NewMockFood(controller)
	foodService.EXPECT().Create(gomock.Any()).Return(nil, mockdata.ErrInternal)
	foodInfo, _ := json.Marshal(mockdata.ValidFoodRequest)
	requestBody := strings.NewReader(string(foodInfo))
	req, err := http.NewRequest("POST", "/api/foods", requestBody)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	foodRouter := routers.FoodRouter(foodService)

	router := gin.New()
	router.POST("/api/foods", foodRouter.Create)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Equal(t, mockdata.InternalErrorResponse, recorder.Body.String())
}

func TestFoodRouter_Update(t *testing.T) {
	testCases := updateFoodTest{}
	t.Run("Test when update food is successful", testCases.whenUpdateFoodSuccess)
	t.Run("Test when update food has no name", testCases.whenUpdateFoodHasNoName)
	t.Run("Test when update food has no price", testCases.whenUpdateFoodHasNoPrice)
	t.Run("Test when update food has invalid name", testCases.whenUpdateFoodHasInvalidName)
	t.Run("Test when update food has invalid price", testCases.whenUpdateFoodHasInvalidPrice)
	t.Run("Test when update food has internal error", testCases.whenUpdateFoodHasInternalError)
	t.Run("Test when update food has invalid id", testCases.whenUpdateFoodHasInvalidId)
	t.Run("Test when update food is not found", testCases.whenUpdateFoodIsNotFound)
}

type updateFoodTest struct{}

func (*updateFoodTest) whenUpdateFoodSuccess(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	foodService := mock_services.NewMockFood(controller)
	foodService.EXPECT().Update(gomock.Any()).Return(mockdata.SuccessFoodCreateResponse, nil)
	foodInfo, _ := json.Marshal(mockdata.ValidFoodRequest)
	requestBody := strings.NewReader(string(foodInfo))
	requestUrl := "/api/foods/" + mockdata.FoodDefaultId.String()
	req, err := http.NewRequest("PUT", requestUrl, requestBody)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	foodRouter := routers.FoodRouter(foodService)

	router := gin.New()
	router.PUT("/api/foods/:id", foodRouter.Update)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	expectedResponse, _ := json.Marshal(mockdata.SuccessFoodCreateResponse)
	assert.Equal(t, string(expectedResponse), recorder.Body.String())
}

func (*updateFoodTest) whenUpdateFoodHasNoName(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	foodService := mock_services.NewMockFood(controller)
	foodService.EXPECT().Update(gomock.Any()).Times(0)
	foodInfo, _ := json.Marshal(mockdata.NoNameFoodRequest)
	requestBody := strings.NewReader(string(foodInfo))
	requestUrl := "/api/foods/" + mockdata.FoodDefaultId.String()
	req, err := http.NewRequest("PUT", requestUrl, requestBody)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	foodRouter := routers.FoodRouter(foodService)

	router := gin.New()
	router.PUT("/api/foods/:id", foodRouter.Update)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
	assert.Equal(t, mockdata.NoNameFoodErrorResponse, recorder.Body.String())
}

func (*updateFoodTest) whenUpdateFoodHasNoPrice(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	foodService := mock_services.NewMockFood(controller)
	foodService.EXPECT().Update(gomock.Any()).Times(0)
	foodInfo, _ := json.Marshal(mockdata.NoPriceFoodRequest)
	requestBody := strings.NewReader(string(foodInfo))
	requestUrl := "/api/foods/" + mockdata.FoodDefaultId.String()
	req, err := http.NewRequest("PUT", requestUrl, requestBody)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	foodRouter := routers.FoodRouter(foodService)

	router := gin.New()
	router.PUT("/api/foods/:id", foodRouter.Update)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
	assert.Equal(t, mockdata.NoPriceFoodErrorResponse, recorder.Body.String())
}

func (*updateFoodTest) whenUpdateFoodHasInvalidName(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	foodService := mock_services.NewMockFood(controller)
	foodService.EXPECT().Update(gomock.Any()).Times(0)
	foodInfo, _ := json.Marshal(mockdata.InvalidNameFoodRequest)
	requestBody := strings.NewReader(string(foodInfo))
	requestUrl := "/api/foods/" + mockdata.FoodDefaultId.String()
	req, err := http.NewRequest("PUT", requestUrl, requestBody)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	foodRouter := routers.FoodRouter(foodService)

	router := gin.New()
	router.PUT("/api/foods/:id", foodRouter.Update)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
	assert.Equal(t, mockdata.InvalidNameErrorResponse, recorder.Body.String())
}

func (*updateFoodTest) whenUpdateFoodHasInvalidPrice(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	foodService := mock_services.NewMockFood(controller)
	foodService.EXPECT().Update(gomock.Any()).Times(0)
	foodInfo, _ := json.Marshal(mockdata.InvalidPriceFoodRequest)
	requestBody := strings.NewReader(string(foodInfo))
	requestUrl := "/api/foods/" + mockdata.FoodDefaultId.String()
	req, err := http.NewRequest("PUT", requestUrl, requestBody)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	foodRouter := routers.FoodRouter(foodService)

	router := gin.New()
	router.PUT("/api/foods/:id", foodRouter.Update)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
	assert.Equal(t, mockdata.InvalidPriceErrorResponse, recorder.Body.String())
}

func (*updateFoodTest) whenUpdateFoodHasInternalError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	foodService := mock_services.NewMockFood(controller)
	foodService.EXPECT().Update(gomock.Any()).Return(nil, mockdata.ErrInternal)
	foodInfo, _ := json.Marshal(mockdata.ValidFoodRequest)
	requestBody := strings.NewReader(string(foodInfo))
	requestUrl := "/api/foods/" + mockdata.FoodDefaultId.String()
	req, err := http.NewRequest("PUT", requestUrl, requestBody)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	foodRouter := routers.FoodRouter(foodService)

	router := gin.New()
	router.PUT("/api/foods/:id", foodRouter.Update)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Equal(t, mockdata.InternalErrorResponse, recorder.Body.String())
}

func (*updateFoodTest) whenUpdateFoodHasInvalidId(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	foodService := mock_services.NewMockFood(controller)
	foodService.EXPECT().Update(gomock.Any()).Times(0)
	foodInfo, _ := json.Marshal(mockdata.ValidFoodRequest)
	requestBody := strings.NewReader(string(foodInfo))
	requestUrl := "/api/foods/" + "invalid-id"
	req, err := http.NewRequest("PUT", requestUrl, requestBody)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	foodRouter := routers.FoodRouter(foodService)

	router := gin.New()
	router.PUT("/api/foods/:id", foodRouter.Update)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
	assert.Equal(t, mockdata.NotFoundErrorResponse, recorder.Body.String())
}

func (*updateFoodTest) whenUpdateFoodIsNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	foodService := mock_services.NewMockFood(controller)
	foodService.EXPECT().Update(gomock.Any()).Return(nil, nil)
	foodInfo, _ := json.Marshal(mockdata.ValidFoodRequest)
	requestBody := strings.NewReader(string(foodInfo))
	requestUrl := "/api/foods/" + mockdata.FoodDefaultId.String()
	req, err := http.NewRequest("PUT", requestUrl, requestBody)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	foodRouter := routers.FoodRouter(foodService)

	router := gin.New()
	router.PUT("/api/foods/:id", foodRouter.Update)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
	assert.Equal(t, mockdata.NotFoundErrorResponse, recorder.Body.String())
}

func TestFoodRouter_List(t *testing.T) {
	testCases := listFoodTest{}
	t.Run("Test when list food is successful", testCases.whenListFoodSuccess)
	t.Run("Test when list food has internal error", testCases.whenListFoodHasInternalError)
}

type listFoodTest struct{}

func (*listFoodTest) whenListFoodSuccess(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	foodService := mock_services.NewMockFood(controller)
	foodService.EXPECT().List(gomock.Any()).Return(mockdata.SuccessFoodListResponse, nil)
	req, err := http.NewRequest("GET", "/api/foods", nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	foodRouter := routers.FoodRouter(foodService)

	router := gin.New()
	router.GET("/api/foods", foodRouter.List)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	expectedResponse, _ := json.Marshal(mockdata.SuccessFoodListResponse)
	assert.Equal(t, string(expectedResponse), recorder.Body.String())
}

func (*listFoodTest) whenListFoodHasInternalError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	foodService := mock_services.NewMockFood(controller)
	foodService.EXPECT().List(gomock.Any()).Return(nil, mockdata.ErrInternal)
	req, err := http.NewRequest("GET", "/api/foods", nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	foodRouter := routers.FoodRouter(foodService)

	router := gin.New()
	router.GET("/api/foods", foodRouter.List)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Equal(t, mockdata.InternalErrorResponse, recorder.Body.String())
}

func TestFoodRouter_Delete(t *testing.T) {
	testCases := deleteFoodTest{}
	t.Run("Test when delete food is successful", testCases.whenDeleteFoodSuccess)
	t.Run("Test when delete food has internal error", testCases.whenDeleteFoodHasInternalError)
	t.Run("Test when delete food is not found", testCases.whenDeleteFoodIsNotFound)
	t.Run("Test when delete food has invalid id", testCases.whenDeleteFoodHasInvalidId)
}

type deleteFoodTest struct{}

func (*deleteFoodTest) whenDeleteFoodSuccess(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	foodService := mock_services.NewMockFood(controller)
	foodService.EXPECT().Delete(gomock.Any()).Return(true, nil)
	requestUrl := "/api/foods/" + mockdata.FoodDefaultId.String()
	req, err := http.NewRequest("DELETE", requestUrl, nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	foodRouter := routers.FoodRouter(foodService)

	router := gin.New()
	router.DELETE("/api/foods/:id", foodRouter.Delete)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusNoContent, recorder.Code)
}

func (*deleteFoodTest) whenDeleteFoodHasInternalError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	foodService := mock_services.NewMockFood(controller)
	foodService.EXPECT().Delete(gomock.Any()).Return(false, mockdata.ErrInternal)
	requestUrl := "/api/foods/" + mockdata.FoodDefaultId.String()
	req, err := http.NewRequest("DELETE", requestUrl, nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	foodRouter := routers.FoodRouter(foodService)

	router := gin.New()
	router.DELETE("/api/foods/:id", foodRouter.Delete)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Equal(t, mockdata.InternalErrorResponse, recorder.Body.String())
}

func (*deleteFoodTest) whenDeleteFoodIsNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	foodService := mock_services.NewMockFood(controller)
	foodService.EXPECT().Delete(gomock.Any()).Return(false, nil)
	requestUrl := "/api/foods/" + mockdata.FoodDefaultId.String()
	req, err := http.NewRequest("DELETE", requestUrl, nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	foodRouter := routers.FoodRouter(foodService)

	router := gin.New()
	router.DELETE("/api/foods/:id", foodRouter.Delete)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
	assert.Equal(t, mockdata.NotFoundErrorResponse, recorder.Body.String())
}

func (*deleteFoodTest) whenDeleteFoodHasInvalidId(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	foodService := mock_services.NewMockFood(controller)
	foodService.EXPECT().Delete(gomock.Any()).Times(0)
	requestUrl := "/api/foods/" + "invalid-id"
	req, err := http.NewRequest("DELETE", requestUrl, nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	foodRouter := routers.FoodRouter(foodService)

	router := gin.New()
	router.DELETE("/api/foods/:id", foodRouter.Delete)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
	assert.Equal(t, mockdata.NotFoundErrorResponse, recorder.Body.String())
}
