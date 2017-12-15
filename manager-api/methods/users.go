/*
 * Copyright (C) 2017 Nethesis S.r.l.
 * http://www.nethesis.it - info@nethesis.it
 *
 * This file is part of Windmill project.
 *
 * NethServer is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License,
 * or any later version.
 *
 * NethServer is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with NethServer.  If not, see COPYING.
 *
 * author: Edoardo Spadoni <edoardo.spadoni@nethesis.it>
 */

package methods

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"manager-api/database"
	"manager-api/models"
	"manager-api/utils"
)

func CreateUser(c *gin.Context) {
	firstName := c.PostForm("firstname")
	lastName := c.PostForm("lastname")
	username := c.PostForm("username")
	email := c.PostForm("email")
	accountType := c.PostForm("type")
	kbpsDown := c.PostForm("kbps_down")
	kbpsUp := c.PostForm("kbps_up")
	validFrom := c.PostForm("valid_from")
	validUntil := c.PostForm("valid_until")

	hotspotId := 1

	user := models.User{
		HotspotId:   hotspotId,
		FirstName:   firstName,
		LastName:    lastName,
		UserName:    username,
		Email:       email,
		AccountType: accountType,
		ValidFrom:   validFrom,
		ValidUntil:  validUntil,
		Created:     time.Now().String(),
	}

	if kbpsDownInt, err := strconv.Atoi(kbpsDown); err == nil {
		user.KbpsDown = kbpsDownInt
	}
	if kbpsUpInt, err := strconv.Atoi(kbpsUp); err == nil {
		user.KbpsUp = kbpsUpInt
	}

	db := database.Database()
	db.Save(&user)

	db.Close()

	c.JSON(http.StatusCreated, gin.H{"id": user.Id, "status": "success"})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	userId := c.Param("user_id")

	firstName := c.PostForm("firstname")
	lastName := c.PostForm("lastname")
	email := c.PostForm("email")
	kbpsDown := c.PostForm("kbps_down")
	kbpsUp := c.PostForm("kbps_up")
	validFrom := c.PostForm("valid_from")
	validUntil := c.PostForm("valid_until")

	db := database.Database()
	db.Where("id = ?", userId).First(&user)

	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No user found!"})
		return
	}

	user.FirstName = firstName
	user.LastName = lastName
	user.Email = email

	if kbpsDownInt, err := strconv.Atoi(kbpsDown); err == nil {
		user.KbpsDown = kbpsDownInt
	}
	if kbpsUpInt, err := strconv.Atoi(kbpsUp); err == nil {
		user.KbpsUp = kbpsUpInt
	}

	user.ValidFrom = validFrom
	user.ValidUntil = validUntil

	db.Save(&user)

	db.Close()
}

func GetUsers(c *gin.Context) {
	var users []models.User

	page := c.Query("page")
	limit := c.Query("limit")

	offsets := utils.OffsetCalc(page, limit)

	db := database.Database()
	db.Offset(offsets[0]).Limit(offsets[1]).Find(&users)

	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No users found!"})
		return
	}

	db.Close()

	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	var user models.User
	userId := c.Param("user_id")

	db := database.Database()
	db.Where("id = ?", userId).First(&user)

	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No user found!"})
		return
	}

	db.Close()

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	userId := c.Param("user_id")

	db := database.Database()
	db.Where("id = ?", userId).First(&user)

	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No user found!"})
		return
	}

	db.Delete(&user)

	db.Close()

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully!"})
}