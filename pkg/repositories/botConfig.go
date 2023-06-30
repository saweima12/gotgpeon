package repositories

import (
	"gotgpeon/logger"
	"gotgpeon/models/entity"
	"gotgpeon/utils/maputil"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BotConfigRepository interface {
	GetWhiteList(chatId string) map[string]struct{}
	SetWhiteListCache(chatId string) error
	SetWhiteListDB(chatId string) error
	GetViolateRecord(chatId string, userId string)
	SetViolateCache(chatId string, userId string)
}

type botConfigRepository struct {
	BaseRepository
}

func NewBotConfigRepo(dbConn *gorm.DB, redisConn *redis.Client) BotConfigRepository {
	return &botConfigRepository{
		BaseRepository: BaseRepository{DbConn: dbConn, RedisConn: redisConn},
	}
}

func (repo *botConfigRepository) GetWhiteList(chatId string) map[string]string {
	namespace := getNamespace("whitelist")

	result := make(map[string]string)
	// Attempt read from cache.
	resp, err := repo.GetRedis().SMembers(baseCtx, namespace).Result()
	if err == nil && len(resp) >= 1 {
		// Generate result.
		for _, v := range resp {
			result[v] = "ok"
		}
		return result
	}

	// Data doesn't exist in the cache. Attempt read from database.
	tableName := entity.PeonUserWhitelist{}.TableName()
	var rows []entity.PeonUserWhitelist

	// Read from database.
	err = repo.GetDB().Table(tableName).Where("status = ?", "ok").Find(&rows).Error
	if err != nil {
		logger.Error("Bot GetWhitelist Err:" + err.Error())
		return result
	}

	// Generate Result.
	for _, v := range rows {
    if v.Status == "ok" {
      result[v.UserId] = "ok"
    }
	}

	return result
}

func (repo *botConfigRepository) SetWhiteListCache(chatId string, whitelistSet map[string]struct{}) error {
	namespace := getNamespace("whitelist")

	// Convert mapKey to interface slice.
	keySlice := maputil.GetMapSetKeys(whitelistSet)
	iSlice := make([]interface{}, len(keySlice))
	for i, v := range keySlice {
		iSlice[i] = v
	}

	rdb := repo.GetRedis()
	// Try refresh key to set.
	rdb.Del(baseCtx, namespace)
	err := rdb.SAdd(baseCtx, namespace, iSlice...).Err()
	if err != nil {
		return err
	}
	return nil
}

func (repo *botConfigRepository) SetWhiteListDB(chatId string, whitelistSet map[string]struct{}) error {
  tableName := entity.PeonUserWhitelist{}.TableName()
  db := repo.GetDB()

  for uid, status := range whitelistSet {
    newItem := entity.PeonUserWhitelist {
      UserId: uid,
      Status: status,
    }

    err := db.Clauses(clause.OnConflict{
      Columns: clause.Column{ Name: "user_id" },
      DoUpdates: clause.Assignment(map[string]interface{}{"status" : status})
    }).Create(newItem).Error


    if err != nil {
      logger.Errorf("SetWhiteListDB error: %s", err.Error())
    }
  }

}

func (botconfigrepositroy *botConfigRepository) GetViolateRecord(chatId string, userId string) {
	panic("not implemented") // TODO: Implement
}

func (botconfigrepositroy *botConfigRepository) SetViolateCache(chatId string, userId string) {
	panic("not implemented") // TODO: Implement
}

func getNamespace(keyword string) string {
	return "bot:" + keyword
}
