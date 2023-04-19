package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type ChatHistory struct {
	ent.Schema
}

func (ChatHistory) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("chat_id"),
		field.Text("content"),
	}
}

func (ChatHistory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("chat_id"),
	}
}

type ChatConfig struct {
	ent.Schema
}

func (ChatConfig) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("chat_id"),
		field.Text("json"),
	}
}

func (ChatConfig) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("chat_id"),
	}
}

type OpenAIToken struct {
	ent.Schema
}

func (OpenAIToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").MaxLen(52),
		field.Bool("is_active").Default(true),
	}
}

func (OpenAIToken) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("token", "is_active"),
	}
}
