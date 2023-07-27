package server

import (
	"fmt"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"log"
	"os"
	"time"
)

func run() {
	fmt.Println("[INFO] Start http port")
	err := GetInstance().client.ListenAndServe()
	if err != nil {
		fmt.Println("[FATAL ERROR] Unable to start http port: " + err.Error())
		log.Fatal("[FATAL ERROR] Unable to start http port: " + err.Error())
	}
}

func Init() {
	if err := rotationLog(); err != nil {
		log.Fatal("[ERROR]", err)
	}
	err := readConfigFile()
	if err != nil {
		log.Fatal("[ERROR]", err)
	}
	run()
}

// rotationLog - иницилизация и приминения ротации логов
func rotationLog() error {
	if _, err := os.Stat(PathBinary + "/logs"); os.IsNotExist(err) {
		err = os.Mkdir(PathBinary+"/logs", 0755)
		if err != nil {
			return err
		}
	}

	writer, err := rotatelogs.New(
		PathBinary+"/logs/opensearch_%W.log",
		rotatelogs.WithMaxAge(730*time.Hour),
		rotatelogs.WithRotationTime(168*time.Hour),
	)
	if err != nil {
		return err
	} else {
		log.SetOutput(writer)
	}
	return nil
}
