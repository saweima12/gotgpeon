package schedule

import (
	"gotgpeon/data/models"
	"gotgpeon/logger"
	"time"
)

func (s *peonSchedule) CacheToDB() {

	logger.Info("CacheToDb Job Start")
	now := time.Now().UTC()

	chatIdList, err := s.ChatRepo.GetAvaliableChatList()
	if err != nil {
		logger.Error("CacheToDB db err:" + err.Error())
	}

	for chatId := range chatIdList {
		// Save chatConfig to database
		chatCfg, err := s.ChatRepo.GetChatCfg(chatId)
		if err != nil {
			logger.Errorf("CacheToDB getChatConfig err: %s", err.Error())
			continue
		}

		chatJobCfg, err := s.ChatRepo.GetChatJobCfg(chatId)
		if err != nil {
			logger.Errorf("CacheToDB getChatJobConfig err: %s", err.Error())
			continue
		}

		err = s.ChatRepo.UpdateChatCfgDB(chatId, chatCfg, chatJobCfg)
		if err != nil {
			logger.Errorf("CacheToDB setChatConfig err: %s", err.Error())
			continue
		}

		// Save userRecord to database and promote level
		users := s.RecordService.GetAllCacheByChat(chatId)

		delList := make([]int64, 0)
		for _, user := range users {
			if user.MemberLevel < models.JUNIOR {
				dayCheck := now.Sub(user.CreatedTime) >= time.Hour*24*time.Duration(chatJobCfg.JuniorDay)
				if user.Point >= int(chatJobCfg.JuniorLowest) && dayCheck {
					user.MemberLevel = models.JUNIOR
				}
			}
			delList = append(delList, user.MemberId)
			s.RecordService.SetUserRecordDB(chatId, user)
		}

		// delete saved userRecord.
		s.RecordService.DelCacheByMemberIds(chatId, delList)
	}

	logger.Info("CacheToDb Job End")
}
