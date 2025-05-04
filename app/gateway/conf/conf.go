package conf

import (
	"github.com/spf13/viper"
	"os"
)

var Conf *Config

type Config struct {
	Server   *Server             `yaml:"server"`
	MySQL    *MySQL              `yaml:"mysql"`
	Redis    *Redis              `yaml:"redis"`
	Etcd     *Etcd               `yaml:"etcd"`
	Services map[string]*Service `yaml:"services"`
	Domain   map[string]*Domain  `yaml:"domain"`
}

type Server struct {
	Port      string `yaml:"port"`
	Version   string `yaml:"version"`
	JwtSecret string `yaml:"jwtSecret"`
}

type MySQL struct {
	DriverName string `yaml:"driverName"`
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	Database   string `yaml:"database"`
	UserName   string `yaml:"username"`
	Password   string `yaml:"password"`
	Charset    string `yaml:"charset"`
}

type Redis struct {
	UserName string `yaml:"userName"`
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
}

type Etcd struct {
	Address string `yaml:"address"`
}

type Service struct {
	Name        string   `yaml:"name"`
	LoadBalance bool     `yaml:"loadBalance"`
	Addr        []string `yaml:"addr"`
}

type Domain struct {
	Name string `yaml:"name"`
}

// 加载配置
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/conf") //配置文件路径
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
}
