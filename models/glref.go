package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type GlRef struct {
	ID         string    `gorm:"column:FCSKID" json:"fcskid"`
	FCCODE     string    `gorm:"column:FCCODE" json:"fccode"`
	FCREFNO    string    `gorm:"column:FCREFNO" json:"fcrefno"`
	FNAMT      string    `gorm:"column:FNAMT" json:"fnamt"`
	FCFRWHOUSE string    `gorm:"column:FCFRWHOUSE" json:"fcfrwhouse"`
	FCTOWHOUSE string    `gorm:"column:FCTOWHOUSE" json:"fctowhouse"`
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD" json:"ftlastupd" form:"ftlastupd"`
	FromWhs    Whs       `gorm:"foreignKey:FCFRWHOUSE;references:FCSKID;" json:"from_whs,omitempty"`
	ToWhs      Whs       `gorm:"foreignKey:FCTOWHOUSE;references:FCSKID;" json:"to_whs,omitempty"`
}

func (GlRef) TableName() string {
	return "GLREF"
}

func (obj *GlRef) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.ID = id
	return
}
