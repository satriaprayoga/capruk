package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Server struct {
	Host         string        `mapstructure:"host"`
	HTTPPort     int           `mapstructure:"http_port"`
	ReadTimeOut  time.Duration `mapstructure:"read_timeout"`
	WriteTimeOut time.Duration `mapstructure:"write_timeout"`
}

type Database struct {
	Type        string `mapstructure:"type"`
	Host        string `mapstructure:"host"`
	Port        string `mapstructure:"port"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Name        string `mapstructure:"name"`
	TablePrefix string `mapstructure:"table_prefix"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	Key      string `mapstructure:"key"`
	Password string `mapstructure:"password"`
}

type SMTP struct {
	Server   string `mapstructure:"server"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Passwd   string `mapstructure:"passwd"`
	Identity string `mapstructure:"identity"`
	Sender   string `mapstructure:"sender"`
}

type AppConfig struct {
	AppName           string    `mapstructure:"app_name"`
	Debug             bool      `mapstructure:"debug"`
	Server            *Server   `mapstructure:"server"`
	Database          *Database `mapstructure:"database"`
	Redis             *Redis    `mapstructure:"redis"`
	SMTP              *SMTP     `mapstructure:"smtp"`
	JwtSecret         string    `mapstructure:"jwt_secret"`
	JWTExpired        int       `mapstructure:"expire_jwt"`
	PageSize          int       `mapstructure:"page_size"`
	PrefixURL         string    `mapstructure:"prefix_url"`
	RuntimeRootPath   string    `mapstructure:"runtime_root_path"`
	ImageSavePath     string    `mapstructure:"image_save_path"`
	ImageMaxSize      int       `mapstructure:"image_size"`
	ImageAllowExts    []string  `mapstructure:"image_allow_ext"`
	ExportSavePath    string    `mapstructure:"export_save_path"`
	QrCodeSavePath    string    `mapstructure:"qr_code"`
	LogSavePath       string    `mapstructure:"log_save_path"`
	LogSaveName       string    `mapstructure:"log_save_name"`
	LogFileExt        string    `mapstructure:"log_file_ext"`
	TimeFormat        string    `mapstructure:"time_format"`
	Issuer            string    `mapstructure:"issuer"`
	UrlForgotPassword string    `mapstructure:"url_forgot_password"`
	UrlVerityUser     string    `mapstructure:"url_verity_user"`
}

func LoadConfig() *AppConfig {
	var config = &AppConfig{}
	now := time.Now()
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fail to parse config file")
	}
	err = viper.Unmarshal(config)
	if err != nil {
		log.Fatalf("Fail to Unmarshall 'config.json': %v", err)
	}
	timeSpent := time.Since(now)
	log.Printf("Config setting is ready in %v", timeSpent)
	return config

}
