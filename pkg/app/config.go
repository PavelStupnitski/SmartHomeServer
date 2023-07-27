package app

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var PathBinary = getPathBinary()

// ReadConfigFile - чтение конфигуарционного файла, возвращает настройку модуля
func readConfigFile() error {
	readByte, errRead := ioutil.ReadFile(PathBinary + "configs/app.yml")
	if errRead != nil {
		return errRead
	}

	var errYml = yaml.Unmarshal(readByte, GetInstance().config)
	if errYml != nil {
		log.Println("[ERROR] ", errYml)
		return errYml
	}
	GetInstance().config.checkReceivedConfig()
	return nil
}

// checkReceivedConfig - проверка на корректный ввод настройки модуля
func (config Config) checkReceivedConfig() {
	if len(config.Host) == 0 {
		log.Println("[WARN ] The opensearch parameter is not specified.")
	} else if len(config.User) == 0 {
		log.Println("[WARN ] The username parameter is not specified.")
	} else if len(config.Password) == 0 {
		log.Println("[WARN ] The password parameter is not specified.")
	} else if len(config.Cert) == 0 {
		log.Println("[WARN ] The cert parameter is not specified.")
	} else if len(config.Port) == 0 {
		log.Println("[WARN ] The port parameter is not specified.")
	}
}

// getPathBinary - возвращает путь к директории, где лежит бинарный файл
func getPathBinary() string {
	var path, errPath = os.Executable()
	if errPath != nil {
		log.Println("[ERROR] ", errPath)
	}
	return path[:strings.LastIndex(path, "/")]
}
