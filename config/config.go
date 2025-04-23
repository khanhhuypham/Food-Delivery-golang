package config

import "github.com/spf13/viper"

type Config struct {
	App   AppConfig
	Mysql MysqlConfig
	Aws   AwsConfig
	Redis RedisConfig
	Nat   NatConfig
	Grpc  GrpcConfig
}

type AppConfig struct {
	Version      string
	Mode         string
	Port         string
	Secret       string
	MigrationURL string
}

type MysqlConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}
type AwsConfig struct {
	Region    string
	APIKey    string
	SecretKey string
	S3Domain  string
	S3Bucket  string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       string
}

type NatConfig struct {
	Url string
}

type GrpcConfig struct {
	Url string
}

func LoadConfig(fileName string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(fileName)
	var config Config
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, err
		}
		return nil, err
	}
	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil

}
