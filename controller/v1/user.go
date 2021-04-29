package v1

import "C"
import (
	"fmt"
	"github.com/antony0016/sw-system-backend/core/model"
	"github.com/antony0016/sw-system-backend/pkg/httpresponse"
	"github.com/antony0016/sw-system-backend/services/pgdb"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func AllUser(c *gin.Context) {
	var users []model.User
	var response model.Response
	db := pgdb.DB
	if err := db.Find(&users).Error; err != nil {
		fmt.Println(err)
		response.ErrorCode = http.StatusInternalServerError
	}
	response.Data = users
	httpresponse.MakeResponse(c, response)
}

func OneUser(c *gin.Context) {
	var users []model.User
	var response model.Response
	db := pgdb.DB
	userID, _ := strconv.Atoi(c.Param("id"))
	if err := db.First(&users, userID); err != nil {
		response.ErrorCode = http.StatusInternalServerError
	}
	response.Data = users
	httpresponse.MakeResponse(c, response)
}

func CreateUser(c *gin.Context) {
	var user model.User
	var response model.Response
	db := pgdb.DB
	mileage, _ := strconv.Atoi(c.PostForm("mileage"))
	licensePlate := strings.ToUpper(c.PostForm("license_plate"))
	licensePlate = strings.ReplaceAll(licensePlate, " ", "")
	db.Where("license_plate=?", c.PostForm("license_plate")).First(&user)
	if user.ID != 0 {
		response.ErrorCode = http.StatusConflict
		response.ErrorMessage = "User already exists"
		return
	}
	user = model.User{
		Username:     c.PostForm("username"),
		Phone:        c.PostForm("phone"),
		LicensePlate: licensePlate,
		Mileage:      mileage,
	}
	if err := db.Save(&user).Error; err != nil {
		fmt.Println(err)
		response.ErrorCode = http.StatusInternalServerError
		return
	}
	response.Message = "Create user successful."
	httpresponse.MakeResponse(c, response)
}

func UpdateUser(c *gin.Context) {
	var user model.User
	var response model.Response
	var err error
	db := pgdb.DB
	id, _ := strconv.Atoi(c.Param("id"))
	db.First(&user, id)
	if user.ID == 0 {
		response.ErrorCode = http.StatusNoContent
		response.ErrorMessage = "User not found."
		httpresponse.MakeResponse(c, response)
		return
	}
	err = db.Model(&user).Update("license_plate", c.PostForm("license_plate")).Error
	err = db.Model(&user).Update("phone", c.PostForm("phone")).Error
	err = db.Model(&user).Update("username", c.PostForm("username")).Error
	err = db.Model(&user).Update("mileage", c.PostForm("mileage")).Error
	if err != nil {
		fmt.Println(err)
		response.ErrorCode = http.StatusInternalServerError
	}
	response.Message = "user updated successfully."
	httpresponse.MakeResponse(c, response)
}

func DeleteUser(c *gin.Context) {
	var user model.User
	var response model.Response
	db := pgdb.DB
	id, _ := strconv.Atoi(c.Param("id"))
	db.First(&user, id)
	if user.ID == 0 {
		response.ErrorCode = http.StatusNoContent
		response.ErrorMessage = "User not found."
		httpresponse.MakeResponse(c, response)
		return
	}
	if err := db.Unscoped().Delete(&user).Error; err != nil {
		fmt.Println(err)
		response.ErrorCode = http.StatusInternalServerError
	}
	response.Message = "user deleted successfully."
	httpresponse.MakeResponse(c, response)
}
