package schedule

import (
	"gotgpeon/db"
	"gotgpeon/pkg/repositories"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robfig/cron/v3"
)

type PeonSchedule interface {
	Run()
}

type peonSchedule struct {
	Croniter   *cron.Cron
	BotAPI     *tgbotapi.BotAPI
	ChatRepo   repositories.ChatRepository
	RecordRepo repositories.RecordRepository
}

func NewPeonSchedule(botAPI *tgbotapi.BotAPI) (PeonSchedule, error) {
	croniter := cron.New()
	dbConn := db.GetDB()
	rdb := db.GetCache()

	// declare repository & service
	chatRepo := repositories.NewChatRepo(dbConn, rdb)
	recordRepo := repositories.NewRecordRepository(dbConn, rdb)

	// declare schedule.
	sch := &peonSchedule{
		Croniter:   croniter,
		BotAPI:     botAPI,
		ChatRepo:   chatRepo,
		RecordRepo: recordRepo,
	}

	// TestJob
	sch.CacheAdminstrator()

	// Add Job
	croniter.AddFunc("0 20 * * * *", sch.CacheAdminstrator)
	croniter.AddFunc("0 20 * * * *", sch.CacheToDB)
	return sch, nil
}

func (s *peonSchedule) Run() {
	// s.Croniter.Run()

}
