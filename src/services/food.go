package services

import (
	"gormCompose/src/driver"
	"gormCompose/src/models"
	"gormCompose/src/response"
	"log"
)

type Food interface {
	Create(food *models.Food) (*models.Food, error)
	Update(food *models.Food) (*models.Food, error)
	Delete(food *models.Food) (bool, error)
	List(pageInfo *response.Pagination) (*response.SearchResult, error)
}
type food struct{}

func FoodService() Food {
	return &food{}
}

func (*food) Create(food *models.Food) (*models.Food, error) {
	driver, err := driver.New()
	if err != nil {
		return nil, err
	}

	tx := driver.Create(food)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return food, nil
}

func (*food) List(pageInfo *response.Pagination) (*response.SearchResult, error) {
	offset := (pageInfo.Page - 1) * pageInfo.PerPage
	log.Print(offset)

	var foodList []*models.Food
	var total int64 = 0
	driver, err := driver.New()
	if err != nil {
		return nil, err
	}

	tx := driver.Model(&models.Food{}).Count(&total).Offset(offset).Limit(pageInfo.PerPage).Find(&foodList)
	if tx.Error != nil {
		return nil, tx.Error
	}
	res := &response.SearchResult{
		Data:    foodList,
		PerPage: pageInfo.PerPage,
		Total:   total,
		Page:    pageInfo.Page,
	}
	return res, nil
}

func (*food) Update(food *models.Food) (*models.Food, error) {
	driver, err := driver.New()
	if err != nil {
		return nil, err
	}

	tx := driver.Save(food)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return food, nil
}

func (*food) Delete(food *models.Food) (bool, error) {
	driver, err := driver.New()
	if err != nil {
		return false, err
	}

	tx := driver.Delete(food)
	if tx.Error != nil {
		return false, tx.Error
	}
	if tx.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}
