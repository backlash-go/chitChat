package db

import (
	"flag"
	"github.com/spf13/viper"
	"log"
)

func Init()  {
	configFile := flag.String("conf", "config/config.yaml", "path of config file")
	flag.Parse()
	viper.SetConfigFile(*configFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("viper read config is failed, err is %v configFile is %s ", err, configFile)
	}
	log.Println("logger init ")
	//init mysql
	dbConf := viper.GetStringMapString("database")
	InitDB(dbConf["user"], dbConf["password"], dbConf["host"], dbConf["port"], dbConf["name"])
	dbRedisConf := viper.GetStringMapString("authRedis")
	InitRedis(dbRedisConf["host"],dbRedisConf["password"],dbRedisConf["db"])
}
