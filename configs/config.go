package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var Cfg *config

type config struct {
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBHost     string `mapstruture:"DB_HOST"`
	DBPort     uint   `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`

	WebServerPort uint `mapstructure:"WEB_SERVER_PORT"`

	JWTSecret    string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn uint   `mapstructure:"JWT_EXPIRES_IN"`
	JWTTokenAuth *jwtauth.JWTAuth
	//JWTTokenAuthKey string
}

func LoadConfig(path string) error {
	var c config

	vipConfig := viper.New()
	vipConfig.SetConfigName("app_config")
	vipConfig.SetConfigType("env")
	vipConfig.SetConfigFile(".env")
	vipConfig.AddConfigPath(path)
	vipConfig.AddConfigPath("./../../")
	vipConfig.AddConfigPath("./../")
	vipConfig.AddConfigPath(".")
	vipConfig.AutomaticEnv()
	err := vipConfig.ReadInConfig()
	if err != nil {
		return err
	}

	err = vipConfig.Unmarshal(&c)
	if err != nil {
		return err
	}

	c.JWTTokenAuth = jwtauth.New("HS256", []byte(c.JWTSecret), nil)

	Cfg = &c
	return nil
}
