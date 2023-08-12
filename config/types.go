package config

type Configuration struct {
	Common     CommonConfig `yaml:"common"`
	TgBot      TgBotConfig  `yaml:"tgbot"`
	IgnoreWord []string     `yaml:"ignore_word"`
}

type CommonConfig struct {
	Mode         string   `yaml:"mode"`
	ListenPort   string   `yaml:"listen_port"`
	RedisUri     string   `yaml:"redis_uri"`
	DBUri        string   `yaml:"db_uri"`
	AllowViaBots []string `yaml:"allow_viabots"`
}

type TgBotConfig struct {
	UpdateMode     string `yaml:"update_mode"`
	AutoSetWebhook bool   `yaml:"auto_set_webhook"`
	BotToken       string `yaml:"bot_token"`
	HookURL        string `yaml:"hook_url"`
}
