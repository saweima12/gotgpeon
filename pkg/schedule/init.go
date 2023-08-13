package schedule

import (
	"gotgpeon/db"
	"gotgpeon/pkg/repositories"
	"gotgpeon/pkg/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/robfig/cron/v3"
)

type PeonSchedule interface {
	Run()
}

type peonSchedule struct {
	Croniter      *cron.Cron
	BotAPI        *tgbotapi.BotAPI
	ChatRepo      repositories.ChatRepository
	RecordService services.RecordService
}

func NewPeonSchedule(botAPI *tgbotapi.BotAPI) (PeonSchedule, error) {
	croniter := cron.New(cron.WithSeconds())
	dbConn := db.GetDB()
	rdb := db.GetCache()

	// declare repository & service
	chatRepo := repositories.NewChatRepo(dbConn, rdb)
	recordRepo := repositories.NewRecordRepository(dbConn, rdb)

	recordService := services.NewRecordService(recordRepo)

	// declare schedule.
	sch := &peonSchedule{
		Croniter:      croniter,
		BotAPI:        botAPI,
		ChatRepo:      chatRepo,
		RecordService: recordService,
	}

	// Startup job.
	sch.CacheAdminstrator()

	// Add Job
	croniter.AddFunc("0 30 * * * *", sch.CacheAdminstrator)
	croniter.AddFunc("0 17 * * * *", sch.CacheToDB)
	return sch, nil
}

func (s *peonSchedule) Run() {
	s.Croniter.Run()
}