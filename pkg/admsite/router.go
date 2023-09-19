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
	recordRepo := repositories.NewRecordRepository(db, rdb)

	// declare services.
	deletedServ := services.NewDeletedService(deletedMsgRepo)
	dataviewServ := services.NewDataviewService(chatRepo, recordRepo)

	// declare controller
	dataviewCon := controller.NewDataViewController(deletedServ, dataviewServ)

	// define router.
	dataview := x.Group("/dataview")
	dataview.GET("/chats", dataviewCon.GetChatList)
	dataview.GET("/chats/:chat_id/members", dataviewCon.GetChatMembers)
	dataview.GET("/chats/:chat_id/deletedmsg", dataviewCon.GetChatDeleteMsgs)
}
