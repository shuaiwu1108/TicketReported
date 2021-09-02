package pkg

import "flag"
import "github.com/BurntSushi/toml"

type Config struct {
	MySql *MySql
}

type MySql struct {
	Host 	string
	UserName	string
	PassWord	string
	Port	int
	DbName	string
	Parameter	string
}

var (
	confPath string
	Conf = &Config{}
)

func init(){
	flag.StringVar(&confPath, "conf", "./conf/app.toml", "-conf path")
}

func Init() (err error){
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}
