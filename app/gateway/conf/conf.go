package conf

import (
	"github.com/spf13/viper"
	"os"
)

// 加载配置
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/conf") //配置文件路径
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
}
