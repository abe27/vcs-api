package models

import (
	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Product struct {
	FCSKID      string      `gorm:"primaryKey;column:FCSKID;size:8;" json:"fcskid"  form:"fcskid"`  //FCSKID char(8) COLLATE Thai_BIN DEFAULT " " NOT NULL,
	FCTYPE      string      `gorm:"column:FCTYPE;size:1;" json:"fctype"  form:"fctype"`             //FCTYPE char(1) COLLATE Thai_BIN DEFAULT " " NOT NULL,
	FCCODE      string      `gorm:"column:FCCODE;unique;size:35;" json:"fccode"  form:"fccode"`     //FCCODE char(35) COLLATE Thai_BIN DEFAULT " " NOT NULL,
	FCSNAME     string      `gorm:"column:FCSNAME;size:140;" json:"fcsname"  form:"fcsname"`        //FCSNAME char(140) COLLATE Thai_BIN DEFAULT " " NOT NULL,
	FCSNAME2    string      `gorm:"column:FCSNAME2" json:"fcsname2"  form:"fcsname2"`               //FCSNAME2 char(140) COLLATE Thai_BIN DEFAULT " " NOT NULL,
	FCNAME      string      `gorm:"column:FCNAME" json:"fcname"  form:"fcname"`                     //FCNAME char(140) COLLATE Thai_BIN DEFAULT " " NOT NULL,
	FCNAME2     string      `gorm:"column:FCNAME2" json:"fcname2"  form:"fcname2"`                  //FCNAME2 char(140) COLLATE Thai_BIN DEFAULT " " NOT NULL,
	FNAVGCOST   float64     `gorm:"column:FNAVGCOST" json:"fnavgcost"  form:"fnavgcost"`            //FNAVGCOST decimal(9,2) DEFAULT 0 NOT NULL,
	FNSTDCOST   float64     `gorm:"column:FNSTDCOST" json:"fnstdcost"  form:"fnstdcost"`            //FNSTDCOST decimal(15,4) DEFAULT 0 NOT NULL,
	FCSTATUS    string      `gorm:"column:FCSTATUS;size:1;type:char" json:"status" form:"fcstatus"` //FCSTATUS
	ProductType ProductType `gorm:"foreignKey:FCTYPE;references:FCCODE;" json:"product_type,omitempty"`
	// Stock       Stock       `json:"stock"`
	// Stock       []Stock     `gorm:"foreignKey:FCSKID;references:FCPROD;" json:"stock,omitempty"`
	// FTDATETIME time.Time `gorm:"column:FTDATETIME" json:"ftdatetime"  form:"ftdatetime" default:"now"` //FTDATETIME datetime DEFAULT getdate() NOT NULL,
	// FTLASTUPD  time.Time `gorm:"column:FTLASTUPD" json:"ftlastupd"  form:"ftlastupd" default:"now"`    //FTLASTUPD datetime DEFAULT getdate() NULL,
	// FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT" json:"ftlastedit"  form:"ftlastedit" default:"now"` //FTLASTEDIT datetime NULL,
}

func (Product) TableName() string {
	return "PROD"
}

func (obj *Product) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}

type ProductType struct {
	ID      string `gorm:"primaryKey;column:FCSKID;size:8;" json:"fcskid"  form:"fcskid"` //FCSKID char(8) COLLATE Thai_BIN DEFAULT " " NOT NULL,
	FCCODE  string `gorm:"column:FCCODE;unique;size:35;" json:"fccode"  form:"fccode"`    //FCSNAME2 char(140) COLLATE Thai_BIN DEFAULT " " NOT NULL,
	FCNAME  string `gorm:"column:FCNAME" json:"fcname"  form:"fcname"`                    //FCNAME char(140) COLLATE Thai_BIN DEFAULT " " NOT NULL,
	FCNAME2 string `gorm:"column:FCNAME2" json:"fcname2"  form:"fcname2"`
	// FTDATETIME time.Time `gorm:"column:FTDATETIME" json:"ftdatetime"  form:"ftdatetime" default:"now"` //FTDATETIME datetime DEFAULT getdate() NOT NULL,
	// FTLASTUPD  time.Time `gorm:"column:FTLASTUPD" json:"ftlastupd"  form:"ftlastupd" default:"now"`    //FTLASTUPD datetime DEFAULT getdate() NULL,
	// FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT" json:"ftlastedit"  form:"ftlastedit" default:"now"` //FTLASTEDIT datetime NULL,
}

func (ProductType) TableName() string {
	return "PRODTYPE"
}

func (obj *ProductType) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.ID = id
	return
}
