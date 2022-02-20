package handler

import (
	"fmt"

	"github.com/indes/flowerss-bot/internal/model"
	"go.uber.org/zap"
	tb "gopkg.in/telebot.v3"
)

type Start struct {
}

func (s *Start) Command() string {
	return "/start"
}

func (s *Start) Description() string {
	return "开始使用"
}

func (s *Start) Handle(ctx tb.Context) error {
	user, _ := model.FindOrCreateUserByTelegramID(ctx.Chat().ID)
	zap.S().Infof("/start user_id: %d telegram_id: %d", user.ID, user.TelegramID)
	return ctx.Send(fmt.Sprintf("你好，欢迎使用flowerss。"))
}