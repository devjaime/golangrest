package product

import (
	"github.com/devjaime/golangrest/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type Product struct {
	gorm.Model
	Sku                 int    `json:"sku"`
	Descripcion         string `json:"descripcion"`
	Ean_description     string `json:"ean_description"`
	Precio              int    `json:"precio"`
	Stock               int    `json:"stock"`
	Store_number        int    `json:"store_number"`
	Item_number         int    `json:"item_number"`
	Catalogo_extendido  string `json:"catalogo_extendido"`
	Cod_segmento        int    `json:"cod_segmento"`
	Segmento            string `json:"segmento"`
	Cod_subgrupo        int    `json:"cod_subgrupo"`
	Subgrupo            string `json:"subgrupo"`
	Cod_depto_comercial int    `json:"cod_depto_comercial"`
	Depto_comercial     string `json:"depto_comercial"`
	Cod_categoria       int    `json:"cod_categoria"`
	Categoria           string `json:"categoria"`
	Cod_subcategoria    int    `json:"cod_subcategoria"`
	Subcategoria        string `json:"subcategoria"`
	Pasillo             int    `json:"pasillo"`
	Habilitado_tienda   string `json:"habilitado_tienda"`
	Habilitado_web      string `json:"habilitado_web"`
	Imagen              string `json:"imagen"`
}

func GetProducts(c *fiber.Ctx) {
	db := database.DBConn
	var products []Product
	db.Find(&products)
	c.JSON(products)
}

func GetProduct(c *fiber.Ctx) {
	item_number := c.Params("item_number")
	db := database.DBConn
	var product Product
	db.Find(&product, "item_number = ?", item_number)
	c.JSON(product)
}

func NewProduct(c *fiber.Ctx) {
	db := database.DBConn
	var product Product
	product.Sku = 1
	product.Descripcion = "prueba"
	product.Ean_description = "prueba"
	product.Precio = 1
	product.Stock = 1
	product.Store_number = 1
	product.Item_number = 1
	product.Catalogo_extendido = "prueba"
	product.Cod_segmento = 1
	product.Segmento = "prueba"
	product.Cod_subgrupo = 1
	product.Subgrupo = "prueba"
	product.Cod_depto_comercial = 1
	product.Depto_comercial = "prueba"
	product.Cod_categoria = 1
	product.Categoria = "prueba"
	product.Cod_subcategoria = 1
	product.Subcategoria = "prueba"
	product.Pasillo = 1
	product.Habilitado_tienda = "prueba"
	product.Habilitado_web = "prueba"
	product.Imagen = "prueba"
	db.Create(&product)
	c.JSON(product)
}
