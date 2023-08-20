package repositories

import (
	"gotgpeon/logger"
	"gotgpeon/models/entity"
	"strconv"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BotConfigRepository interface {
	GetAllowlist() map[int64]byte
	SetAllowlistCache(allowlistSet map[int64]byte) error
	SetAllowlistDBWithUserId(userId int64, isEnable byte) error
}

type botConfigRepository struct {
	BaseRepository
}

func NewBotConfigRepo(dbConn *gorm.DB, redisConn *redis.Client) BotConfigRepository {
	return &botConfigRepository{
		BaseRepository: BaseRepository{DbConn: dbConn, RedisConn: redisConn},
	}
}

func (repo *botConfigRepository) GetAllowlist() map[int64]byte {
	namespace := getNamespace("allowlist")
	result := make(map[int64]byte)
	// Attempt read from cache.
	rdb := repo.GetRedis()
	resp, err := rdb.SMembers(baseCtx, namespace).Result()
	if err == nil && len(resp) > 0 {
		// Generate result.
		for _, v := range resp {
			kV, err := strconv.Atoi(v)
			if err != nil {
				continue
			}
			result[int64(kV)] = 1
		}
		return result
	}

	// Data doesn't exist in the cache. Attempt read from database.
	tableName := entity.PeonMemberAllowlist{}.TableName()
	var rows []entity.PeonMemberAllowlist

	// Read from database.
	err = repo.GetDB().Table(tableName).Where("status = ?", 1).Find(&rows).Error
	if err != nil {
		logger.Error("Bot GetAllowlist Err:" + err.Error())
		return result
	}

	// Generate Result.
	for _, v := range rows {
		if v.Status == 1 {
			result[v.MemberId] = 1
		}
	}

	// Save to cache.
	repo.SetAllowlistCache(result)
	return result
}

func (repo *botConfigRepository) SetAllowlistCache(allowlistSet map[int64]byte) error {
	namespace := getNamespace("allowlist")

	// Convert mapKey to interface slice.
	keySlice := mapKeyToSlice(allowlistSet)
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

func (repo *botConfigRepository) SetAllowlistDBWithUserId(memberId int64, isEnable byte) error {
	// Create struct.
	value := entity.PeonMemberAllowlist{
		MemberId: int64(memberId),
		Status:   int8(isEnable),
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

// / ====
// / Support Function
// / ====
func sliceToMapSet[T int | int64 | string](slice []T) map[T]byte {
	var result map[T]byte
	for _, v := range slice {
		result[v] = 1
	}
	return result
}

func mapKeyToSlice[T int | int64 | string](m map[T]byte) []T {
	result := []T{}
	for k := range m {
		result = append(result, k)
	}
	return result
}

func getNamespace(keyword string) string {
	return "bot:" + keyword
}
