package controller

import (
	"github.com/divyam234/teldrive/pkg/services"
)

type Controller struct {
	FileService   *services.FileService
	UserService   *services.UserService
	UploadService *services.UploadService
	AuthService   *services.AuthService
}

func NewController(fileService *services.FileService,
	userService *services.UserService,
	uploadService *services.UploadService,
	authService *services.AuthService) *Controller {
	return &Controller{
		FileService:   fileService,
		UserService:   userService,
		UploadService: uploadService,
		AuthService:   authService,
	}
}
