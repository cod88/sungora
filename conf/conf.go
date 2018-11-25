package conf

import (
	"os"
	"time"
)

func init() {
	DirWork, _ = os.Getwd()
	sep := string(os.PathSeparator)
	DirConfig = DirWork + sep + "config"
	DirStatic = DirWork + sep + "www"
}

type ConfigMain struct {
	Name           string
	TimeZone       string
	DriverDB       string // Драйвер DB
	SessionTimeout time.Duration
	Host           string
	Port           int
	Mode           string // Режим работы приложения
}

type ConfigMysql struct {
	Host     string // протокол, хост и порт подключения
	Name     string // Имя базы данных
	Login    string // Логин к базе данных
	Password string // Пароль к базе данных
	Charset  string // Кодировка данных (utf-8 - по умолчанию)
}

type ConfigPostgresql struct {
	Host     string // Хост базы данных (localhost - по умолчанию)
	Port     int64  // Порт подключения по протоколу tcp/ip (3306 по умолчанию)
	Name     string // Имя базы данных
	Login    string // Логин к базе данных
	Password string // Пароль к базе данных
	Charset  string // Кодировка данных (utf-8 - по умолчанию)
}

var DirWork string
var DirConfig string
var DirStatic string
var TimeLocation *time.Location
