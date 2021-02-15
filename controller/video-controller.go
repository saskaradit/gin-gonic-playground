package controller

import (
	"gin-gonic/entity"
	"gin-gonic/service"
	"gin-gonic/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
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
