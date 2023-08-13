package schedule

import "gotgpeon/logger"

func (s *peonSchedule) CacheToDB() {

	chatIdList := s.ChatRepo.GetAvaliableChatIds()
	for _, chatId := range chatIdList {
		// Save chatConfig to database
		chatCfg, err := s.ChatRepo.GetChatConfig(chatId)
		if err != nil {
			logger.Errorf("CacheToDB getChatConfig err: %s", err.Error())
			continue
		}
		s.ChatRepo.SetConfigDb(chatId, chatCfg)

		// Save userRecord to database and promote level

	}

}
