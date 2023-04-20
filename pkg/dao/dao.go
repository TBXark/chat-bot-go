package dao

import (
	"github.com/TBXark/chat-bot-go/pkg/dao/ent"
)

type Dao struct {
	DB *ent.Client
}

func NewDao(client *ent.Client) *Dao {
	return &Dao{client}
}
