package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        	string 		`yaml:"env" env:"ENV" env-required:"true"`
	MySQL      	MySQL		`yaml:"mysql" env-required:"true"`
	HTTPServer 	HTTPServer	`yaml:"http_server" env-required:"true"`
}
type MySQL struct {
	Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
	Port     int    `yaml:"port" env:"PORT" env-default:"3306"`
	User     string `yaml:"user" env:"USER" env-default:"root"`
	Password string `yaml:"password" env:"PASSWORD"`
	DBName   string `yaml:"dbname" env-required:"true"`
}
type HTTPServer struct {
	Address string	`yaml:"address"`
}

func MustLoad() *Config{
	configpath := os.Getenv("CONFIG_PATH")
	//checking whether the configuration file exists
	if configpath == ""{
		flags:= flag.String("config", "", "path to the configuration file") //for the cli if the user didnt set the config path
		flag.Parse()
		configpath = *flags
		
		if configpath == ""{
			log.Fatal("Config File Not Set!")
		}
	}

	//checking if the config file is present in the path or has some error
	if _,err:= os.Stat(configpath); os.IsNotExist(err){
		log.Fatalf("Config File Not Found!: %s", configpath)
	}

	//loading the file into the struct
	var cfg Config
	err:= cleanenv.ReadConfig(configpath, &cfg)
	if err!= nil{
		log.Fatalf("Cannot Read The File: %s",err.Error())
	}
	return &cfg
}