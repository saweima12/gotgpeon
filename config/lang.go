package config

type TextLang struct {
	TipsStartCmd         string `yaml:"tips_start_cmd"`
	TipsSetPermissionCmd string `yaml:"tips_setpermission_cmd"`
	TipsDeleteMsg        string `yaml:"tips_delete_msg"`
	ErrContentNozhtw     string `yaml:"err_content_nozhtw"`
	ErrNameBlock         string `yaml:"err_name_block"`
	ErrTypeMedia         string `yaml:"err_media_content"`
	ErrForward           string `yaml:"err_forward"`
	ActVbanMarkup        string `yaml:"act_vban_markup"`
}
