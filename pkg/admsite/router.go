package admsite

import (
	"gotgpeon/db"
	"gotgpeon/pkg/admsite/controller"
	"gotgpeon/pkg/repositories"
	"gotgpeon/pkg/services"

	"github.com/gin-gonic/gin"
)

func InitRouter(x *gin.Engine) {

	rdb := db.GetCache()
	db := db.GetDB()
	// declare repository
	chatRepo := repositories.NewChatRepo(db, rdb)
	deletedMsgRepo := repositories.NewDeletedMsgRepository(db, rdb)
	botRepo := repositories.NewBotConfigRepo(db, rdb)

	// declare services.
	peonServ := services.NewPeonService(chatRepo, botRepo)
	deletedServ := services.NewDeletedService(deletedMsgRepo)

	// declare controller
	dataviewCon := controller.NewDataViewController(peonServ, deletedServ)

	// define router.
	dataview := x.Group("/dataview")
	dataview.GET("/chats", dataviewCon.GetAllChats)
	dataview.GET("/chats/:chat_id/members", dataviewCon.GetChatMembers)
	dataview.GET("/chats/:chat_id/deletedmsg", dataviewCon.GetChatDeleteMsgs)
}
