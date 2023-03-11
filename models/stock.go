package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Stock struct {
	ID        string    `gorm:"column:FCSKID" json:"fcskid"`
	FCWHOUSE  string    `gorm:"column:FCWHOUSE" json:"fcwhouse" form:"fcwhouse"`
	FCPROD    string    `gorm:"column:FCPROD" json:"fcprod" form:"fcprod"`
	FNQTY     float64   `gorm:"column:FNQTY" json:"fnqty" form:"fnqty"`
	FNPRICE   float64   `gorm:"column:FNPRICE" json:"fnprice" form:"fnprice"`
	FTLASTUPD time.Time `gorm:"column:FTLASTUPD" json:"ftlastupd" form:"ftlastupd"`
	Whs       Whs       `gorm:"foreignKey:FCWHOUSE;references:FCSKID;" json:"whs,omitempty"`
	Product   Product   `gorm:"foreignKey:FCPROD;references:FCSKID;" json:"product,omitempty"`
	RefProd   RefProd   `gorm:"foreignKey:FCPROD;references:FCPROD;" json:"refprod,omitempty"`
}

func (Stock) TableName() string {
	return "STOCK"
}

func (obj *Stock) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.ID = id
	return
}
