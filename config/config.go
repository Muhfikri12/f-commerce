package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	AppName   string
	Port      string
	Debug     bool
	Migration bool
	Seeder    bool
	Key       Key
	Database  Database
	Redis     Redis
	SmtpEmail SmtpEmail
}

type Database struct {
	DBName         string
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBConnection   int
	DBTimezone     string
	DBMaxIdleConst int
	DBMaxOpenConst int
	DBMaxIdleTime  int
	DBMaxLifeTime  int
}

type Redis struct {
	Url      string
	Password string
	Prefix   string
}

type Key struct {
	PublicKey  string
	PrivateKey string
}

type SmtpEmail struct {
	SmtpHost  string
	SmtpPort  string
	SmtpUser  string
	FromEmail string
	ApiKey    string
}

func SetConfig() (*Config, error) {

	log, _ := zap.NewProduction()
	defer log.Sync()
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("../..")
	viper.SetConfigType("dotenv")
	viper.SetConfigName(".env")

	viper.SetDefault("DBHost", "localhost")
	viper.SetDefault("DBPort", "5432")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Error("Error reading config file: %s, using default values or environment variables", zap.Error(err))
	}

	config := Config{
		AppName:   viper.GetString("APP_NAME"),
		Port:      viper.GetString("PORT"),
		Debug:     viper.GetBool("DEBUG"),
		Migration: viper.GetBool("AUTO_MIGRATE"),
		Seeder:    viper.GetBool("SEEDER"),

		Database: Database{
			DBName:         viper.GetString("DB_NAME"),
			DBHost:         viper.GetString("DB_HOST"),
			DBPort:         viper.GetString("DB_PORT"),
			DBUser:         viper.GetString("DB_USER"),
			DBPassword:     viper.GetString("DB_PASSWORD"),
			DBConnection:   viper.GetInt("DB_ConnectTimeOut"),
			DBTimezone:     viper.GetString("DB_TIMEZONE"),
			DBMaxIdleConst: viper.GetInt("DB_MAX_IDLE_CONNS"),
			DBMaxOpenConst: viper.GetInt("DB_MAX_OPEN_CONNS"),
			DBMaxIdleTime:  viper.GetInt("DB_MAX_IDLE_TIME"),
			DBMaxLifeTime:  viper.GetInt("DB_MAX_LIFE_TIME"),
		},

		Redis: Redis{
			Url:      viper.GetString("REDIS_URL"),
			Password: viper.GetString("REDIS_PASSWORD"),
			Prefix:   viper.GetString("REDIS_PREFIX"),
		},

		Key: Key{
			PublicKey:  viper.GetString("PUBLIC_KEY"),
			PrivateKey: viper.GetString("PRIVATE_KEY"),
		},

		SmtpEmail: SmtpEmail{
			SmtpHost:  viper.GetString("SMTP_HOST"),
			SmtpPort:  viper.GetString("SMTP_PORT"),
			SmtpUser:  viper.GetString("SMTP_USER"),
			FromEmail: viper.GetString("FROM_EMAIL"),
			ApiKey:    viper.GetString("SMTP_API_KEY"),
		},
	}

	return &config, nil
}
