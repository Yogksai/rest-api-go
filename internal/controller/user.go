package controller

import (
	"strconv"

	"github.com/Yogksai/rest-api-go/internal/models"
	"github.com/Yogksai/rest-api-go/internal/repository"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Repo *repository.UserRepo
}

func NewController(repo *repository.UserRepo) *Controller {
	return &Controller{Repo: repo}
}

func (c Controller) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.Repo.Create(&user); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, user)
}

func (c Controller) GetUsers(ctx *gin.Context) {
	users, err := c.Repo.FindAll()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, users)
}

func (c Controller) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, err := c.Repo.FindByID(uint(idInt))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, user)
}

func (c Controller) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.Repo.Update(&user); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, user)
}

func (c Controller) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.Repo.Delete(uint(idInt)); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(204, nil)
}
