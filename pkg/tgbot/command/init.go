package command

import (
	"gotgpeon/pkg/services"
	"gotgpeon/utils"
)

type CommandHandler interface {
	Invoke(helper *utils.MessageHelper)
}

type CommandMap struct {
	PeonService   services.PeonService
	RecordService services.RecordService
	BotService    services.BotService
	commandMap    map[string]func(helper *utils.MessageHelper)
}

func (h *CommandMap) Invoke(helper *utils.MessageHelper) {
	cmdStr := helper.Command()
	cmdFunc, ok := h.commandMap[cmdStr]

	if ok {
		cmdFunc(helper)
	}
}

func (h *CommandMap) Init() {
	h.commandMap = map[string]func(helper *utils.MessageHelper){
		"start": h.handleStartCommand,
		"point": h.handlePointCommand,
		// "setlevel": h.handleSetLevelCommand,
		// "save":     h.handleSaveCommand,
	}
}
