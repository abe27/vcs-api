package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type RefProd struct {
	ID        string    `gorm:"column:FCSKID" json:"fcskid"`
	FDDATE    time.Time `gorm:"column:FDDATE" json:"fddate"`
	FCGLREF   string    `gorm:"column:FCGLREF" json:"fcglref"`
	FCPROD    string    `gorm:"column:FCPROD" json:"fcprod"`
	FCWHOUSE  string    `gorm:"column:FCWHOUSE" json:"fcwho"`
	FCIOTYPE  string    `gorm:"column:FCIOTYPE" json:"fciotype"`
	FNQTY     float64   `gorm:"column:FNQTY" json:"fnqty"`
	FNPRICE   float64   `gorm:"column:FNPRICE" json:"fnprice"`
	FTLASTUPD time.Time `gorm:"column:FTLASTUPD" json:"ftlastupd" form:"ftlastupd"`
	GlRef     GlRef     `gorm:"foreignKey:FCGLREF;references:FCSKID;" json:"glref,omitempty"`
	Product   Product   `gorm:"foreignKey:FCPROD;references:FCSKID;" json:"product,omitempty"`
	Whs       Whs       `gorm:"foreignKey:FCWHOUSE;references:FCSKID;" json:"whs,omitempty"`
}

func (RefProd) TableName() string {
	return "REFPROD"
}

func (obj *RefProd) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.ID = id
	return
}
