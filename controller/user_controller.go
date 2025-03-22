package controller

import (
    "github.com/gin-gonic/gin"
    "deployapi/models"
    "net/http"
)

var users = []models.User{
    {ID: 1, Name: "John Doe", Email: "john@example.com"},
    {ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
}

// GetUsers mengembalikan daftar pengguna statis
func GetUsers(c *gin.Context) {
    c.JSON(http.StatusOK, users)
}

// AddUser menambahkan pengguna baru ke dalam array
func AddUser(c *gin.Context) {
    var newUser models.User
    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    newUser.ID = len(users) + 1
    users = append(users, newUser)
    c.JSON(http.StatusCreated, newUser)
}
