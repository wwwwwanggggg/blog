package config

import (
	"blog/service/validator"
)

func init() {
	initConfig()
	initLogger()
	validator.InitValidator(Config.AppLanguage)
}
