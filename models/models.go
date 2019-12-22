package models

import (
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Username string     `json:"username"`
	Password string     `json:"password"`
	Sessions []*Session `gorm:"many2many:user_sessions;"`
}

type Message struct {
	gorm.Model
	SessionID    uint
	SessionRefer uint
	FromId       int
	ToId         int
	Text         string
}

type Session struct {
	gorm.Model
	Users    []User `gorm:"many2many:user_sessions"`
	Messages []Message
	Conn     websocket.Conn `gorm:"-"`
}
