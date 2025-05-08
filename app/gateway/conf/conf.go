package conf

import (
	"bytes"
	_ "embed"
	"log"
	"strconv"
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
	Server   ServerConfig             `mapstructure:"server"`
	Etcd     EtcdConfig               `mapstructure:"etcd"`
	Services map[string]ServiceConfig `mapstructure:"services"`
}

type ServerConfig struct {
	Port      string `mapstructure:"port"`
	Version   string `mapstructure:"version"`
	JwtSecret string `mapstructure:"jwtSecret"`
}

type EtcdConfig struct {
	Address string `mapstructure:"address"`
}

type ServiceConfig struct {
	Name        string   `mapstructure:"name"`
	LoadBalance bool     `mapstructure:"loadBalance"`
	Addr        []string `mapstructure:"addr"`
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
			viper.Set(key, viper.GetString(envVar))
		}
	}

	// 5. 解析到结构体
	Conf = new(Config)
	if err := viper.Unmarshal(Conf); err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}

	// 6. 特殊处理布尔值(因为环境变量替换后可能变为字符串)
	for serviceName, service := range Conf.Services {
		if lb, ok := viper.Get("services." + serviceName + ".loadBalance").(string); ok {
			service.LoadBalance, _ = strconv.ParseBool(lb)
		}
	}

	log.Println("Configuration loaded successfully")
}
