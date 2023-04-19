package dao

import (
	"context"
	"github.com/TBXark/chat-bot-go/pkg/dao/ent"
	"github.com/TBXark/chat-bot-go/pkg/dao/ent/chathistory"
)

func (d *Dao) FindChatHistoryByChatId(chatId int64) (*ent.ChatHistory, error) {
	return d.DB.ChatHistory.Query().Where(chathistory.ChatID(chatId)).First(context.Background())
}

func (d *Dao) UpdateChatHistory(chatId int64, content string) error {
	exist, err := d.DB.ChatHistory.Query().Where(chathistory.ChatID(chatId)).Exist(context.Background())
	if err != nil {
		return err
	}
	if exist {
		_, err = d.DB.ChatHistory.Update().Where(chathistory.ChatID(chatId)).SetContent(content).Save(context.Background())
	} else {
		_, err = d.DB.ChatHistory.Create().SetChatID(chatId).SetContent(content).Save(context.Background())
	}
	return err
}
