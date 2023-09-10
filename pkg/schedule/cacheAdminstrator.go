package schedule

import (
	"gotgpeon/logger"
	"gotgpeon/data/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (s *peonSchedule) CacheAdminstrator() {
	logger.Info("CacheAdminstrator Job Start")

	chatIds := s.ChatRepo.GetAvaliableChatIds()
	for _, chatId := range chatIds {
		// Get ChatAdminstrators from telegram
		members, err := s.BotAPI.GetChatAdministrators(tgbotapi.ChatAdministratorsConfig{
			ChatConfig: tgbotapi.ChatConfig{ChatID: chatId},
		})
		if err != nil {
			logger.Errorf("CacheAdminstrator err: %s || ChatId: %s", err.Error(), chatId)
			continue
		}

		memberIdList := make([]int64, 0, len(members))
		memberIdNameMap := make(map[int64]string)
		// process memberId & memberName
		for _, member := range members {
			if member.User.IsBot {
				continue
			}

			memberId := member.User.ID
			memberName := member.User.FirstName + " " + member.User.LastName
			// Add to list & map
			memberIdList = append(memberIdList, memberId)
			memberIdNameMap[memberId] = memberName
		}

		chatInfo, err := s.BotAPI.GetChat(tgbotapi.ChatInfoConfig{
			ChatConfig: tgbotapi.ChatConfig{ChatID: chatId},
		})

		if err != nil {
			logger.Errorf("CacheAdminstrator err: %s || ChatId: %s", err.Error(), chatId)
			continue
		}

		// Update chatConfig.
		chatCfg, err := s.ChatRepo.GetChatCfg(chatId)
		if err != nil {
			logger.Errorf("CacheGroupAdmin getChatConfig err: %s", err.Error())
			continue
		}
		chatCfg.Adminstrators = memberIdList
		chatCfg.ChatName = chatInfo.Title
		s.ChatRepo.SetChatCfgCache(chatId, chatCfg)

		// Update UserRecord
		for mId, mName := range memberIdNameMap {
			query := models.MessageRecord{
				MemberId: mId,
				FullName: mName,
			}
			record := s.RecordService.GetUserRecordByCaht(chatId, &query)
			record.FullName = mName

			err = s.RecordService.SetUserRecordCache(chatId, record)
			if err != nil {
				logger.Errorf("CacheGroupAdmin SetUserRecordCache err: %s", err.Error())
				continue
			}
		}
	}

	logger.Info("CacheAdminstrator Job End")
}
