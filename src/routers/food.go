package routers

import (
	"gormCompose/src/messages"
	"gormCompose/src/models"
	"gormCompose/src/request"
	"gormCompose/src/response"
	s "gormCompose/src/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Food interface {
	Create(c *gin.Context)
	List(c *gin.Context)
}
type food struct{}

func FoodRouter() Food {
	return &food{}
}

// CreateFood
//
//	@Summary		Adiciona Comida.
//	@Description	Cadastra uma nova Comida.
//	@Tags			Comida
//	@Accept			json
//	@Produce		json
//	@Param			json	body		request.Food	true	"Informações da nova Comida"
//	@Success		201		{object}	models.Food
//	@Failure		422		{object}	response.ApiError
//	@Failure		500		{object}	response.ApiError
//	@Router			/foods [post]
func (*food) Create(c *gin.Context) {
	var requestFood request.Food

	err := c.ShouldBindJSON(&requestFood)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	newFood, err := s.FoodService().Create(&models.Food{Name: requestFood.Name, Price: requestFood.Price})
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": messages.INTERNAL_SERVER_ERROR})
		return
	}
	c.JSON(http.StatusCreated, newFood)
}

// ListFood
//
//	@Summary		Lista Comidas.
//	@Description	Retorna a lista de comidas cadastradas.
//	@Tags			Comida
//	@Produce		json
//	@Param			page	query		integer	false	"Número da página".
//	@Param			perPage	query		integer	false	"Número de resultados por página".
//	@Success		200		{object}	response.SearchResult
//	@Failure		500		{object}	response.ApiError
//	@Router			/foods [get]
func (*food) List(c *gin.Context) {
	pageInfo := &response.Pagination{
		Page:    1,
		PerPage: 20,
	}
	queryPage := c.Query("page")

	if len(queryPage) > 0 {
		page, err := strconv.Atoi(queryPage)
		if err == nil && page > 0 {
			pageInfo.Page = page
		}
	}

	queryPerPage := c.Query("perPage")

	if len(queryPerPage) > 0 {
		perPage, err := strconv.Atoi(queryPerPage)
		if err == nil && perPage > 0 {
			pageInfo.PerPage = perPage
		}
	}

	foodList, err := s.FoodService().List(pageInfo)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": messages.INTERNAL_SERVER_ERROR})
		return
	}
	c.JSON(http.StatusOK, foodList)
}
