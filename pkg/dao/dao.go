package dao

import (
	"context"
	"github.com/TBXark/chat-bot-go/pkg/dao/ent"
	"github.com/TBXark/chat-bot-go/pkg/dao/ent/chathistory"
	"github.com/TBXark/chat-bot-go/pkg/dao/ent/openaitoken"
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

func (d *Dao) FinaAllActiveOpenAIToken() ([]*ent.OpenAIToken, error) {
	return d.DB.OpenAIToken.Query().Where(openaitoken.IsActive(true)).All(context.Background())
}

func (d *Dao) CreateOpenAIToken(token string) error {
	_, err := d.DB.OpenAIToken.Create().SetToken(token).SetIsActive(true).Save(context.Background())
	return err
}

func (d *Dao) DeleteOpenAIToken(token string) error {
	_, err := d.DB.OpenAIToken.Update().Where(openaitoken.Token(token)).SetIsActive(false).Save(context.Background())
	return err
}

func (d *Dao) DeleteOpenAITokenById(id int) error {
	_, err := d.DB.OpenAIToken.Update().Where(openaitoken.ID(id)).SetIsActive(false).Save(context.Background())
	return err
}
