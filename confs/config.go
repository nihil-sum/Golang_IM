package confs

import (
	"fmt"

	"github.com/spf13/viper"
)

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     uint64 `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

func LoadDBConfig() (*DBConfig, error) {
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var db DBConfig
	err = viper.UnmarshalKey("DB", &db)
	return &db, nil
}
func (cfg DBConfig) DSN() string {

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
}
