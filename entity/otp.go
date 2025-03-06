package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Otp struct {
	ID        string     `form:"id" json:"id" gorm:"type:uuid;primary_key;column:id"`
	NoHp      string     `form:"no_hp" json:"no_hp" gorm:"column:no_hp"`
	Used      bool       `form:"used" json:"used" gorm:"column:used"`
	Otp       int        `form:"otp" json:"otp" gorm:"column:otp"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
}

func (Otp) TableName() string {
	return "otp"
}

func (b *Otp) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}

type Token struct {
	ID        string     `form:"id" json:"id" gorm:"type:uuid;primary_key;column:id"`
	NoHp      string     `form:"no_hp" json:"no_hp" gorm:"column:no_hp"`
	Used      bool       `form:"used" json:"used" gorm:"column:used"`
	Token       string        `form:"token" json:"token" gorm:"column:token"`
	TempPass string `json:"temp_pass" gorm:"-;type:json"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
}

func (Token) TableName() string {
	return "token"
}

func (b *Token) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}