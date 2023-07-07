package repositories

import (
	"fmt"
	"gotgpeon/logger"
	"gotgpeon/models"
	"gotgpeon/models/entity"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BotConfigRepository interface {
	GetWhiteList() map[string]byte
	SetWhiteListCache(whitelistSet map[string]byte) error
	SetWhiteListDBWithUserId(userId string, isEnable byte) error
}

type botConfigRepository struct {
	BaseRepository
}

func NewBotConfigRepo(dbConn *gorm.DB, redisConn *redis.Client) BotConfigRepository {
	return &botConfigRepository{
		BaseRepository: BaseRepository{DbConn: dbConn, RedisConn: redisConn},
	}
}

func (repo *botConfigRepository) GetWhiteList() map[string]byte {
	namespace := getNamespace("whitelist")
	result := make(map[string]byte)
	// Attempt read from cache.
	rdb := repo.GetRedis()
	resp, err := rdb.SMembers(baseCtx, namespace).Result()
	if err == nil && len(resp) > 0 {
		// Generate result.
		for _, v := range resp {
			result[v] = 1
		}
		return result
	}

	// Data doesn't exist in the cache. Attempt read from database.
	tableName := entity.PeonUserWhitelist{}.TableName()
	var rows []entity.PeonUserWhitelist

	// Read from database.
	err = repo.GetDB().Table(tableName).Where("status = ?", "ok").Find(&rows).Error
	fmt.Println(rows)
	if err != nil {
		logger.Error("Bot GetWhitelist Err:" + err.Error())
		return result
	}

	// Generate Result.
	for _, v := range rows {
		if v.Status == "ok" {
			result[v.UserId] = 1
		}
	}

	// Save to cache.
	repo.SetWhiteListCache(result)
	return result
}

func (repo *botConfigRepository) SetWhiteListCache(whitelistSet map[string]byte) error {
	namespace := getNamespace("whitelist")

	// Convert mapKey to interface slice.
	keySlice := mapKeyToSlice(whitelistSet)
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

func (repo *botConfigRepository) SetWhiteListDBWithUserId(userId string, isEnable byte) error {

	isOK := models.OK
	if isEnable != 1 {
		isOK = models.NG
	}

	value := entity.PeonUserWhitelist{
		UserId: userId,
		Status: isOK,
	}

	err := repo.GetDB().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "chat_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"config_json", "attatch_json"}),
	}).Create(&value).Error

	if err != nil {
		return err
	}

	return nil
}

func (botconfigrepositroy *botConfigRepository) GetViolateRecord(chatId string, userId string) (int, error) {
	return 0, nil // TODO Implement
}

func (botconfigrepositroy *botConfigRepository) SetViolateCache(chatId string, userId string, num int) {
	panic("not implemented") // TODO: Implement
}

/// ====
/// Support Function
/// ====
func sliceToMapSet(slice []string) map[string]string {
	var result map[string]string
	for _, v := range slice {
		result[v] = "ok"
	}
	return result
}

func mapKeyToSlice(m map[string]byte) []string {
	result := []string{}
	for k := range m {
		result = append(result, k)
	}
	return result
}

func getNamespace(keyword string) string {
	return "bot:" + keyword
}
