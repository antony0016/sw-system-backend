package v1

import (
	"fmt"
	"github.com/antony0016/sw-system-backend/core/model"
	"github.com/antony0016/sw-system-backend/pkg/httpresponse"
	"github.com/antony0016/sw-system-backend/services/pgdb"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//func FindAll(){
//	var categories []model.Category
//	err := db.Find(&categories).Error
//}

func DeleteCategoryRelate(c *gin.Context, response model.Response, categoryID int) error {
	db := pgdb.DB
	err := db.Unscoped().Where(
		"category_id=?", categoryID).Delete(&model.OrderItem{}).Error
	if err != nil {
		fmt.Println(err)
		response.ErrorCode = http.StatusInternalServerError
		httpresponse.MakeResponse(c, response)
		return err
	}
	err = db.Unscoped().Where(
		"category_id=?", categoryID).Delete(&model.Item{}).Error
	if err != nil {
		fmt.Println(err)
		response.ErrorCode = http.StatusInternalServerError
		httpresponse.MakeResponse(c, response)
	}
	return err
}

func AllCategory(c *gin.Context) {
	var categories []model.Category
	var response model.Response
	db := pgdb.DB
	if err := db.Find(&categories).Error; err != nil {
		fmt.Println(err)
		response.ErrorCode = http.StatusInternalServerError
	}
	response.Data = categories
	httpresponse.MakeResponse(c, response)
}

func OneCategory(c *gin.Context) {
	var categories []model.Category
	var response model.Response
	db := pgdb.DB
	id, _ := strconv.Atoi(c.Param("id"))
	//db.First(&categories, id)
	if err := db.First(&categories, id).Error; err != nil {
		fmt.Println(err)
		response.ErrorCode = http.StatusInternalServerError
	}
	response.Data = categories
	httpresponse.MakeResponse(c, response)
}

func CreateCategory(c *gin.Context) {
	var category model.Category
	var response model.Response
	db := pgdb.DB
	categoryName := c.PostForm("category_name")
	db.Where("category_name=?", categoryName).Find(&category)
	if category.ID != 0 {
		response.ErrorCode = http.StatusConflict
		response.ErrorMessage = "category already exists"
		httpresponse.MakeResponse(c, response)
		return
	}
	category.CategoryName = categoryName
	if err := db.Save(&category).Error; err != nil {
		fmt.Println(err)
		response.ErrorCode = http.StatusInternalServerError
		return
	}
	response.Message = "Create category successful."
	httpresponse.MakeResponse(c, response)
}

func UpdateCategory(c *gin.Context) {
	var category model.Category
	var response model.Response
	db := pgdb.DB
	id, _ := strconv.Atoi(c.Param("id"))
	db.First(&category, id)
	if category.ID == 0 {
		response.ErrorCode = http.StatusNoContent
		response.ErrorMessage = "Category not found."
		httpresponse.MakeResponse(c, response)
		return
	}
	err := db.Model(&category).Update(
		"category_name", c.PostForm("category_name"),
	).Error
	if err != nil {
		fmt.Println(err)
		response.ErrorCode = http.StatusInternalServerError
	}
	response.Message = "Category updated successfully."
	httpresponse.MakeResponse(c, response)
}

func DeleteCategory(c *gin.Context) {
	var category model.Category
	var response model.Response
	db := pgdb.DB
	id, _ := strconv.Atoi(c.Param("id"))
	db.First(&category, id)
	if category.ID == 0 {
		response.ErrorCode = http.StatusNoContent
		response.ErrorMessage = "Category not found."
		httpresponse.MakeResponse(c, response)
		return
	}
	if err := DeleteCategoryRelate(c, response, id); err != nil {
		fmt.Println(err)
		return
	}
	if err := db.Unscoped().Delete(&category).Error; err != nil {
		fmt.Println(err)
		response.ErrorCode = http.StatusInternalServerError
	}
	response.Message = "Category deleted successfully."
	httpresponse.MakeResponse(c, response)
}
