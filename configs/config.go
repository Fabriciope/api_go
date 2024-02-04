package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var Cfg *Config

type Config struct {
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBHost     string `mapstruture:"DB_HOST"`
	DBPort     int    `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`

	WebServerPort int `mapstructure:"WEB_SERVER_PORT"`

	JWTSecret    string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn int    `mapstructure:"JWT_EXPIRES_IN"`
	JWTTokenAuth *jwtauth.JWTAuth
	//JWTTokenAuthKey string
}

func LoadConfig(path string) (*Config, error) {
	var cfg *Config

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
		return nil, err
	}

	err = vipConfig.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	cfg.JWTTokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)

	return cfg, nil
}
