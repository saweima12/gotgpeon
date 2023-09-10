package controller

import (
	"gotgpeon/pkg/services"

	"github.com/gin-gonic/gin"
)

type DataViewController struct {
	PeonServ    services.PeonService
	DeletedServ services.DeletedService
}

func NewDataViewController(peonServ services.PeonService, deletedServ services.DeletedService) *DataViewController {
	return &DataViewController{
		PeonServ:    peonServ,
		DeletedServ: deletedServ,
	}
}

func (con *DataViewController) GetAllChats(ctx *gin.Context) {

}

func (con *DataViewController) GetChatMembers(ctx *gin.Context) {

}

func (con *DataViewController) GetChatDeleteMsgs(ctx *gin.Context) {

}
