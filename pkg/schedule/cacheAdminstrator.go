package schedule

import (
	"fmt"
	"gotgpeon/logger"
	"strconv"
)

func (s *peonSchedule) CacheAdminstrator() {
	logger.Info("CacheAdminstrator Job Start")

	chatIds := s.ChatRepo.GetAvaliableChatIds()
	for _, chatIdStr := range chatIds {
		chatId, _ := strconv.Atoi(chatIdStr)
		// members, err := s.BotAPI.GetChatAdministrators(tgbotapi.ChatAdministratorsConfig{
		// 	ChatConfig: tgbotapi.ChatConfig{ChatID: int64(chatId)},
		// })
		// if err != nil {
		// 	logger.Errorf("CacheAdminstrator err: %s || ChatId: %s", err.Error(), chatId)
		// 	continue
		// }
		fmt.Println(chatId)
		// process adminstrator

	}

	logger.Info("CacheAdminstrator Job End")
}
