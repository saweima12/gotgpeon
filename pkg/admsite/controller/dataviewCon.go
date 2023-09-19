package controller

import (
	"gotgpeon/pkg/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DataViewController struct {
	DeletedServ  services.DeletedService
	DataviewServ services.DataviewService
}

func NewDataViewController(deletedServ services.DeletedService, dataviewServ services.DataviewService) *DataViewController {
	return &DataViewController{
		DeletedServ:  deletedServ,
		DataviewServ: dataviewServ,
	}
}

func (con *DataViewController) GetChatList(ctx *gin.Context) {
	result := con.DataviewServ.GetChatList()
	ctx.JSON(200, result)
}

func (con *DataViewController) GetChatMembers(ctx *gin.Context) {
	chatIdStr := ctx.Param("chat_id")
	chatId, err := strconv.Atoi(chatIdStr)
	if err != nil {
		ctx.String(400, "Bad Reuqest"+err.Error())
	}

	result := con.DataviewServ.GetMemberListByChat(int64(chatId))
	ctx.JSON(200, result)

}

func (con *DataViewController) GetChatDeleteMsgs(ctx *gin.Context) {

}
