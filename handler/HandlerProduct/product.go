package HandlerProduct

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"store-backend/database"
	"store-backend/model/Product"
)

type NewProducts struct {
	Id        uint     `json:"id"`
	Nome      string   `json:"nome"`
	Descricao string   `json:"descricao"`
	Preco     float64  `json:"preco"`
	Imagens   []string `json:"listaImagem"`
}

func CreateProduct(c *fiber.Ctx) error {
	db := database.DB
	product := new(NewProducts)
	if err := c.BodyParser(product); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	dbProduct := Product.Products{
		Nome:      product.Nome,
		Descricao: product.Descricao,
		Preco:     product.Preco,
	}

	ret := db.Create(&dbProduct)

	if ret.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Erro ao cadastrar o produto", "data": ret.Error})
	}

	for _, imagen := range product.Imagens {
		carrosel := Product.Carrosel{
			Imagem:    imagen,
			ProdutoID: dbProduct.ID,
		}
		db.Create(&carrosel)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created product", "data": product})
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	cmdSql := fmt.Sprint("delete from carrosels where produto_id = ", id)
	err := db.Exec(cmdSql)

	if err.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Erro mysql", "data": err.Error})
	}

	cmdSql = fmt.Sprint("delete from products where id = ", id)
	err = db.Exec(cmdSql)

	if err.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Erro mysql", "data": err.Error})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Delete product!"})
}

func UpdateProduct(c *fiber.Ctx) error {
	db := database.DB
	product := new(NewProducts)
	if err := c.BodyParser(product); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	dbProduct := new(Product.Products)
	ret := db.Find(&dbProduct, product.Id)

	if ret.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Erro ao encontrar o produto", "data": ret.Error})
	}

	dbProduct.Nome = product.Nome
	dbProduct.Descricao = product.Descricao
	dbProduct.Preco = product.Preco

	ret = db.Save(&dbProduct)

	if ret.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Erro ao atualizar o produto", "data": ret.Error})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Update product", "data": product})
}

func GetAllProduct(c *fiber.Ctx) error {
	db := database.DB

	var products []Product.Products
	result := db.Preload("Carrosels").Find(&products)
	fmt.Println(products)

	if result.Error != nil {
		fmt.Println("Error fetching products:", result.Error)
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Update product", "data": products})
}

func GetByIdProduct(c *fiber.Ctx) error {
	idBusca := c.Params("id")
	db := database.DB
	var products Product.Products
	result := db.Preload("Carrosels").Find(&products, idBusca)
	fmt.Println(products)

	if result.Error != nil {
		fmt.Println("Error fetching products:", result.Error)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Update product", "data": products})
}
