package Product

import "gorm.io/gorm"

// User struct
type Products struct {
	gorm.Model
	Nome      string     `gorm:"not null;size:200" json:"nome"`
	Descricao string     `gorm:"not null;size:200" json:"descricao"`
	Preco     float64    `gorm:"not null;" json:"preco"`
	Carrosels []Carrosel `gorm:"foreignKey:ProdutoID" json:"listaImagens"`
}
