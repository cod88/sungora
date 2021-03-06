package core

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
)

const (
	extToml = "toml"
	extYaml = "yaml"
)

// SearchConfigPath поиск конфигурации
func SearchConfigPath(serviceName string) (path, ext string) {
	if serviceName == "" {
		serviceName = filepath.Base(os.Args[0])
		serviceName = strings.Split(serviceName, filepath.Ext(serviceName))[0]
	}
	//
	sep := string(os.PathSeparator)
	path = filepath.Dir(filepath.Dir(os.Args[0]))
	if path == "." {
		path, _ = os.Getwd()
		path = filepath.Dir(path)
	}
	path += sep + "config" + sep + serviceName
	//
	if inf, err := os.Stat(path + extYaml); err == nil {
		if inf.Mode().IsRegular() == true {
			return path, extYaml
		}
	}
	if inf, err := os.Stat(path + extToml); err == nil {
		if inf.Mode().IsRegular() == true {
			return path, extToml
		}
	}
	return
}

// LoadConfig загрузка конфигурации
func LoadConfig(path string, cfg interface{}) (err error) {
	var data []byte
	l := strings.SplitAfter(path, ".")
	ext := l[len(l)-1]
	switch ext {
	case extToml:
		_, err = toml.DecodeFile(path, cfg);
	case extYaml:
		if data, err = ioutil.ReadFile(path); err != nil {
			return
		}
		if err = yaml.Unmarshal(data, cfg); err != nil {
			return
		}
	default:
		return errors.New("undefined config: " + path)
	}
	return
}
