package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Whs struct {
	ID     string `gorm:"primaryKey;column:FCSKID;size:8;" json:"fcskid"  form:"fcskid"` //FCSKID char(8) COLLATE Thai_BIN DEFAULT " " NOT NULL,
	FCCODE string `gorm:"column:FCCODE;unique;size:35;" json:"code"  form:"fccode"`      //FCSNAME2 char(140) COLLATE Thai_BIN DEFAULT " " NOT NULL,
	FCNAME string `gorm:"column:FCNAME" json:"name"  form:"fcname"`                      //FCNAME char(140) COLLATE Thai_BIN DEFAULT " " NOT NULL,
}

func (Whs) TableName() string {
	return "WHOUSE"
}

func (obj *Whs) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.ID = id
	return
}
