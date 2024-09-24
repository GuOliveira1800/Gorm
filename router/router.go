package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"store-backend/handler/HandlerAuth"
	"store-backend/handler/HandlerProduct"
	"store-backend/handler/HandlerUser"
	"store-backend/middleware"
)

func SetupRoutes(app *fiber.App) {
	// User
	user := app.Group("/user", logger.New())
	user.Post("/login", HandlerAuth.Login)
	user.Post("/cadastrar", HandlerUser.CreateUser)
	user.Delete("/:id", middleware.Protected(), HandlerUser.DeleteUser)

	// // Product
	product := app.Group("/product")
	product.Post("/cadastrar", HandlerProduct.CreateProduct)
	product.Put("/atualizar", HandlerProduct.UpdateProduct)
	product.Get("/pegar", HandlerProduct.GetAllProduct)
	product.Get("/pegar/:id", HandlerProduct.GetByIdProduct)
	product.Delete("/deletar/:id", HandlerProduct.DeleteProduct)
}
