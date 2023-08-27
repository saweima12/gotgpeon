package command

import (
	"gotgpeon/pkg/services"
	"gotgpeon/pkg/tgbot/core"
)

type CommandHandler interface {
	Invoke(helper *core.MessageHelper)
}

type CommandMap struct {
	PeonService   services.PeonService
	RecordService services.RecordService
	BotService    services.BotService
	commandMap    map[string]func(helper *core.MessageHelper)
}

func (h *CommandMap) Invoke(helper *core.MessageHelper) {
	cmdStr := helper.Command()
	cmdFunc, ok := h.commandMap[cmdStr]

	if ok {
		cmdFunc(helper)
	}
}

func (h *CommandMap) Init() {
	h.commandMap = map[string]func(helper *core.MessageHelper){
		"start":    h.handleStartCmd,
		"point":    h.handlePointCmd,
		"del":      h.handleDelCmd,
		"setlevel": h.handleSetLevelCmd,
		// "save":     h.handleSaveCommand,
	}
}
