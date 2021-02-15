package controller

import (
	"gin-gonic/entity"
	"gin-gonic/service"
	"gin-gonic/validators"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}
func (c *controller) Save(ctx *gin.Context) error {
	var vid entity.Video
	err := ctx.BindJSON(&vid)
	if err != nil {
		return err
	}
	err = validate.Struct(vid)
	if err != nil {
		return err
	}
	c.service.Save(vid)
	return nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	vid := c.service.FindAll()
	data := gin.H{
		"title":  "Video Page",
		"videos": vid,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
