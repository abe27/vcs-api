package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Employee struct {
	FCSKID     string    `gorm:"primaryKey;column:FCSKID;type:char;size:8;not null;index;" json:"fcskid"  form:"fcskid"` // FCSKID char(8) COLLATE Thai_BIN DEFAULT ' ' NOT NULL,
	FCUDATE    string    `gorm:"column:FCUDATE" json:"fcudate"  form:"fcudate"`                                          // FCUDATE char(2) COLLATE Thai_BIN DEFAULT ' ' NOT NULL,
	FCUTIME    string    `gorm:"column:FCUTIME" json:"fcutime"  form:"fcutime"`                                          // FCUTIME char(2) COLLATE Thai_BIN DEFAULT ' ' NOT NULL,
	FCLUPDAPP  string    `gorm:"column:FCLUPDAPP" json:"fclupdapp"  form:"fclupdapp"`                                    // FCLUPDAPP char(2) COLLATE Thai_BIN DEFAULT ' ' NOT NULL,
	FCLOGIN    string    `gorm:"column:FCLOGIN" json:"fclogin"  form:"fclogin"`                                          // FCLOGIN char(20) COLLATE Thai_BIN DEFAULT ' ' NOT NULL,
	FCPW       string    `gorm:"column:FCPW" json:"fcpw"  form:"fcpw"`                                                   // FCPW char(20) COLLATE Thai_BIN DEFAULT ' ' NOT NULL,
	FCHPW      string    `gorm:"column:FCHPW" json:"fchpw"  form:"fchpw"`                                                // FCHPW char(40) COLLATE Thai_BIN DEFAULT ' ' NOT NULL,
	FCHHPW     string    `gorm:"column:FCHHPW" json:"fchhpw"  form:"fchhpw"`                                             // FCHHPW char(80) COLLATE Thai_BIN DEFAULT ' ' NOT NULL,
	FCRCODE    string    `gorm:"column:FCRCODE" json:"fcrcode"  form:"fcrcode"`                                          // FCRCODE char(10) COLLATE Thai_BIN DEFAULT ' ' NOT NULL,
	FCUSERTEM  string    `gorm:"column:FCUSERTEM" json:"fcusertem"  form:"fcusertem"`                                    // FCUSERTEM char(8) COLLATE Thai_BIN DEFAULT ' ' NOT NULL,
	FTDATETIME time.Time `gorm:"column:FTDATETIME" json:"ftdatetime"  form:"ftdatetime" default:"now"`                   // FTDATETIME datetime DEFAULT getdate() NOT NULL,
	FIMILLISEC int64     `gorm:"column:FIMILLISEC" json:"fimillisec"  form:"fimillisec"`                                 // FIMILLISEC int DEFAULT 0 NOT NULL,
	FCSELTAG   string    `gorm:"column:FCSELTAG" json:"fcseltag"  form:"fcseltag"`                                       // FCSELTAG char(1) COLLATE Thai_BIN DEFAULT ' ' NOT NULL,
	FNSTATUS   float64   `gorm:"column:FNSTATUS" json:"fnstatus"  form:"fnstatus"`                                       // FNSTATUS decimal(1,0) DEFAULT 0 NOT NULL,
	FCSCRNLANG string    `gorm:"column:FCSCRNLANG" json:"fcscrnlang"  form:"fcscrnlang"`                                 // FCSCRNLANG char(2) COLLATE Thai_BIN NULL,
	FCMENULANG string    `gorm:"column:FCMENULANG" json:"fcmenulang"  form:"fcmenulang"`                                 // FCMENULANG char(2) COLLATE Thai_BIN NULL,
	FCREPOLANG string    `gorm:"column:FCREPOLANG" json:"fcrepolang"  form:"fcrepolang"`                                 // FCREPOLANG char(2) COLLATE Thai_BIN NULL,
	FCFIELLANG string    `gorm:"column:FCFIELLANG" json:"fcfiellang"  form:"fcfiellang"`                                 // FCFIELLANG char(2) COLLATE Thai_BIN NULL,
	FCYEARMODE string    `gorm:"column:FCYEARMODE" json:"fcyearmode"  form:"fcyearmode"`                                 // FCYEARMODE char(2) COLLATE Thai_BIN NULL,
	FCPTHAICOD string    `gorm:"column:FCPTHAICOD" json:"fcpthaicod"  form:"fcpthaicod"`                                 // FCPTHAICOD char(3) COLLATE Thai_BIN NULL,
	FCCAPMODE  string    `gorm:"column:FCCAPMODE" json:"fccapmode"  form:"fccapmode"`                                    // FCCAPMODE char(2) COLLATE Thai_BIN DEFAULT ' ' NOT NULL,
	FTLASTUPD  time.Time `gorm:"column:FTLASTUPD" json:"ftlastupd"  form:"ftlastupd" default:"now"`                      // FTLASTUPD datetime DEFAULT getdate() NULL,
	FTLASTEDIT time.Time `gorm:"column:FTLASTEDIT" json:"ftlastedit"  form:"ftlastedit" default:"now"`                   // FTLASTEDIT datetime NULL,
	FCCREATEAP string    `gorm:"column:FCCREATEAP" json:"fccreateap"  form:"fccreateap"`                                 // FCCREATEAP char(2) COLLATE Thai_BIN DEFAULT ' ' NOT NULL,
	FNMISSPASS float64   `gorm:"column:FNMISSPASS" json:"fnmisspass"  form:"fnmisspass"`                                 // FNMISSPASS decimal(2,0) DEFAULT 0 NOT NULL,
	FDLSTCHGPA time.Time `gorm:"column:FDLSTCHGPA" json:"fdlstchgpa"  form:"fdlstchgpa" default:"now"`                   // FDLSTCHGPA datetime NULL,
	FDLSTLOGIN time.Time `gorm:"column:FDLSTLOGIN" json:"fdlstlogin"  form:"fdlstlogin" default:"now"`                   // FDLSTLOGIN datetime NULL,
	FCU1ACC    string    `gorm:"column:FCU1ACC" json:"fcu1acc"  form:"fcu1acc"`                                          // FCU1ACC char(20) COLLATE Thai_BIN DEFAULT ' ' NOT NULL,
	FCDATAIMP  string    `gorm:"column:FCDATAIMP" json:"fcdataimp"  form:"fcdataimp"`                                    // FCDATAIMP char(1) COLLATE Thai_BIN DEFAULT ' ' NOT NULL,
	FCCREATEBY string    `gorm:"column:FCCREATEBY" json:"fccreateby"  form:"fccreateby"`                                 // FCCREATEBY char(8) COLLATE Thai_BIN DEFAULT ' ' NOT NULL,
	FCCORRECTB string    `gorm:"column:FCCORRECTB" json:"fccorrectb"  form:"fccorrectb"`                                 // FCCORRECTB char(8) COLLATE Thai_BIN DEFAULT ' ' NOT NULL,
	FCHPW01    string    `gorm:"column:FCHPW01" json:"fchpw01"  form:"fchpw01"`                                          // FCHPW01 char(60) COLLATE Thai_BIN DEFAULT ' ' NOT NULL,
	FCHHPW01   string    `gorm:"column:FCHHPW01" json:"fchhpw01"  form:"fchhpw01"`                                       // FCHHPW01 char(120) COLLATE Thai_BIN DEFAULT ' ' NOT NULL,
	FCENCOTYPE string    `gorm:"column:FCENCOTYPE" json:"fcencotype"  form:"fcencotype"`                                 // FCENCOTYPE char(10) COLLATE Thai_BIN DEFAULT ' ' NOT NULL,
	FMEXTRATAG string    `gorm:"column:FMEXTRATAG" json:"fmextratag"  form:"fmextratag"`                                 // FMEXTRATAG text COLLATE Thai_BIN DEFAULT NULL NULL,
	FCORGCODE  string    `gorm:"column:FCORGCODE" json:"fcorgcode"  form:"fcorgcode"`                                    // FCORGCODE varchar(128) COLLATE Thai_BIN DEFAULT '' NOT NULL,
	FCCUACC    string    `gorm:"column:FCCUACC" json:"fccuacc"  form:"fccuacc"`                                          // FCCUACC varchar(128) COLLATE Thai_BIN DEFAULT '' NOT NULL,
	FCAPPNAME  string    `gorm:"column:FCAPPNAME" json:"fcappname"  form:"fcappname"`                                    // FCAPPNAME varchar(128) COLLATE Thai_BIN DEFAULT '' NOT NULL,
}

func (Employee) TableName() string {
	return "EMPLR"
}

func (obj *Employee) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New(8)
	obj.FCSKID = id
	return
}

type FrmLogin struct {
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type AuthSession struct {
	Header   string      `json:"header"`
	JwtType  string      `json:"jwt_type,omitempty"`
	JwtToken string      `json:"jwt_token,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}
