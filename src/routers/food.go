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
	"github.com/google/uuid"
)

type Food interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
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

// UpdateFood
//
//	@Summary		Atualiza Comida.
//	@Description	Atualiza uma Comida existente.
//	@Tags			Comida
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string			true	"ID da Comida"	Format(uuid)
//	@Param			json	body		request.Food	true	"Novas Informações da Comida"
//	@Success		200		{object}	models.Food
//	@Failure		404		{object}	response.ApiError
//	@Failure		422		{object}	response.ApiError
//	@Failure		500		{object}	response.ApiError
//	@Router			/foods/{id} [put]
func (*food) Update(c *gin.Context) {
	idParam := c.Params.ByName("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": messages.NOT_FOUND_ERROR})
		return
	}

	var requestFood request.Food
	err = c.ShouldBindJSON(&requestFood)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	updatedFood, err := s.FoodService().Update(&models.Food{ID: id, Name: requestFood.Name, Price: requestFood.Price})
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": messages.INTERNAL_SERVER_ERROR})
		return
	}
	c.JSON(http.StatusOK, updatedFood)
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

// UpdateFood
//
//	@Summary		Remove Comida.
//	@Description	Deleta uma Comida existente.
//	@Tags			Comida
//	@Param			id	path	string	true	"ID da Comida"	Format(uuid)
//	@Produce		json
//	@Success		204
//	@Failure		404	{object}	response.ApiError
//	@Failure		422	{object}	response.ApiError
//	@Failure		500	{object}	response.ApiError
//	@Router			/foods/{id} [delete]
func (*food) Delete(c *gin.Context) {
	idParam := c.Params.ByName("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": messages.NOT_FOUND_ERROR})
		return
	}

	deleted, err := s.FoodService().Delete(&models.Food{ID: id})
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": messages.INTERNAL_SERVER_ERROR})
		return
	}
	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": messages.NOT_FOUND_ERROR})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
