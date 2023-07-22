package config

import "github.com/spf13/viper"

type Config struct {
	Port            string `mapstructure:"PORT"`
	DBUrl           string `mapstructure:"DB_URL"`
	JWTSecretKey    string `mapstructure:"JWT_SECRET_KEY"`
	Issuer          string `mapstructure:"ISSUER"`
	ExpirationHours int    `mapstructure:"EXPIRATION_HOURS"`
	UrlSvcPort      string `mapstructure:"URL_SVC_PORT"`
	AuthSvcPort     string `mapstructure:"AUTH_SVC_PORT"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
