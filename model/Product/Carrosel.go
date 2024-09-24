package Product

import "gorm.io/gorm"

type Carrosel struct {
	gorm.Model
	Imagem    string `gorm:"not null" json:"imagem"`
	ProdutoID uint   `gorm:"not null;" json:"produto_id"` // Chave estrangeira para Produto
}
