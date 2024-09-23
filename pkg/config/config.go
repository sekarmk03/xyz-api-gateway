package config

import "github.com/spf13/viper"

type Config struct {
	Port              string `mapstructure:"PORT"`
	ConsumerSvcUrl    string `mapstructure:"CONSUMER_SVC_URL"`
	TransactionSvcUrl string `mapstructure:"TRANSACTION_SVC_URL"`
	AuthSvcUrl        string `mapstructure:"AUTH_SVC_URL"`
	FeOriginUrl       string `mapstructure:"FE_ORIGIN_URL"`
	SSLCert           string `mapstructure:"SSL_CERT"`
	SSLKey            string `mapstructure:"SSL_KEY"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
