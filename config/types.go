package config

type Configuration struct {
	Common CommonConfig `yaml:"common"`
	TgBot  TgBotConfig  `yaml:"tgbot"`
}

type CommonConfig struct {
	RedisUri string `yaml:"redis_uri"`
	DBUri    string `yaml:"db_uri"`
}

type TgBotConfig struct {
	BotToken string `yaml:"bot_token"`
	HookURL  string `yaml:"hook_url"`
}
