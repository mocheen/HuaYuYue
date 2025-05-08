package conf

import (
	"bytes"
	_ "embed"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var (
	//go:embed conf.yaml
	configFile []byte
	Conf       *Config
	once       sync.Once
)

type Config struct {
	Server ServerConfig `mapstructure:"server"`
	MySQL  MySQLConfig  `mapstructure:"mysql"`
	Redis  RedisConfig  `mapstructure:"redis"`
	Etcd   EtcdConfig   `mapstructure:"etcd"`
}

type ServerConfig struct {
	Domain      string `mapstructure:"domain"`
	Version     string `mapstructure:"version"`
	JwtSecret   string `mapstructure:"jwtSecret"`
	GrpcAddress string `mapstructure:"grpcAddress"`
}

type MySQLConfig struct {
	DriverName string `mapstructure:"driverName"`
	Host       string `mapstructure:"host"`
	Port       string `mapstructure:"port"`
	Database   string `mapstructure:"database"`
	Username   string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	Charset    string `mapstructure:"charset"`
}

type RedisConfig struct {
	Address  string `mapstructure:"address"`
	Username string `mapstructure:"username"`
	DB       int    `mapstructure:"db"`
}

type EtcdConfig struct {
	Address string `mapstructure:"address"`
}

// GetConf 获取配置实例(单例模式)
func GetConf() *Config {
	once.Do(initConf)
	return Conf
}

func initConf() {
	// 1. 加载.env文件
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// 2. 初始化Viper
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(bytes.NewBuffer(configFile)); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// 3. 设置环境变量前缀和替换规则
	viper.SetEnvPrefix("HUA")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// 4. 处理环境变量替换
	for _, key := range viper.AllKeys() {
		value := viper.GetString(key)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			envVar := strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")
			viper.Set(key, os.Getenv(envVar))
		}
	}

	// 5. 解析到结构体
	Conf = new(Config)
	if err := viper.Unmarshal(Conf); err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}

	log.Println("User service configuration loaded successfully")
}
