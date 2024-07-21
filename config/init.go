package config

import (
	"strings"

	"github.com/anandMohanan/qufetch/app"
	"github.com/anandMohanan/qufetch/filesystem"
	"github.com/anandMohanan/qufetch/where"
	"github.com/spf13/viper"
)

// ConfigFormat is the format of the config file
// Available options are: json, yaml, toml
const ConfigFormat = "toml"

var EnvKeyReplacer = strings.NewReplacer(".", "_")

func Init() error {
	viper.SetConfigName(app.Name)
	viper.SetConfigType(ConfigFormat)
	viper.SetFs(filesystem.Api())
	viper.AddConfigPath(where.Config())
	viper.SetTypeByDefaultValue(true)
	viper.SetEnvPrefix(app.Name)
	viper.SetEnvKeyReplacer(EnvKeyReplacer)

	setDefaults()

	err := viper.ReadInConfig()

	switch err.(type) {
	case viper.ConfigFileNotFoundError:
		// Use defaults then
		return nil
	default:
		return err
	}
}
