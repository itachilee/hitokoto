package setting

import (
	"time"

	"github.com/gookit/goutil/dump"
	"github.com/spf13/viper"
)

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerConfig = new(Server)

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseConfig = new(Database)

type App struct {
	QrCodeSavePath  string `yaml:"QrCodeSavePath"`
	RuntimeRootPath string `yaml:"RuntimeRootPath"`
	ImageMaxSize    int    `yaml:"ImageMaxSize"`
	TimeFormat      int    `yaml:"TimeFormat"`
	ImageSavePath   string `yaml:"ImageSavePath"`
	LogFileExt      string `yaml:"LogFileExt"`
	ExportSavePath  string `yaml:"ExportSavePath"`
	LogSaveName     string `yaml:"LogSaveName"`
	PrefixUrl       string `yaml:"PrefixUrl"`
	ImageAllowExts  string `yaml:"ImageAllowExts"`
	FontSavePath    string `yaml:"FontSavePath"`
	LogSavePath     string `yaml:"LogSavePath"`
	PageSize        int    `yaml:"PageSize"`
	JwtSecret       int    `yaml:"JwtSecret"`
}

var AppCinfig = new(App)

func Setup(path string) {
	viper := viper.New()

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	cfgServer := viper.Sub("server")

	ServerConfig = &Server{
		RunMode:      cfgServer.GetString("RunMode"),
		HttpPort:     cfgServer.GetInt("HttpPort"),
		ReadTimeout:  cfgServer.GetDuration("ReadTimeout"),
		WriteTimeout: cfgServer.GetDuration("WriteTimeout"),
	}

	dump.P(ServerConfig)

	cfgDatabase := viper.Sub("database")

	DatabaseConfig = &Database{
		Type:        cfgDatabase.GetString("Type"),
		User:        cfgDatabase.GetString("User"),
		Password:    cfgDatabase.GetString("Password"),
		Host:        cfgDatabase.GetString("Host"),
		Name:        cfgDatabase.GetString("Name"),
		TablePrefix: cfgDatabase.GetString("TablePrefix"),
	}

	dump.P(DatabaseConfig)

	cfgApp := viper.Sub("app")
	AppCinfig = &App{
		PageSize:  cfgApp.GetInt("PageSize"),
		JwtSecret: cfgApp.GetInt("JwtSecret"),
		PrefixUrl: cfgApp.GetString("PrefixUrl"),

		RuntimeRootPath: cfgApp.GetString("RuntimeRootPath"),

		ImageSavePath: cfgApp.GetString("ImageSavePath"),

		ImageMaxSize:   cfgApp.GetInt("ImageMaxSize"),
		ImageAllowExts: cfgApp.GetString("ImageAllowExts"),

		ExportSavePath: cfgApp.GetString("ExportSavePath"),
		QrCodeSavePath: cfgApp.GetString("QrCodeSavePath"),
		FontSavePath:   cfgApp.GetString("FontSavePath"),

		LogSavePath: cfgApp.GetString("LogSavePath"),
		LogSaveName: cfgApp.GetString("LogSaveName"),
		LogFileExt:  cfgApp.GetString("LogFileExt"),
		TimeFormat:  cfgApp.GetInt("TimeFormat"),
	}

	dump.P(AppCinfig)
}
