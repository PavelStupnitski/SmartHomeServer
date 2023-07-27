package mqtt

import (
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"log"
	"os"
	"time"
)

func Init() {
	if err := rotationLog(); err != nil {
		log.Fatal("[ERROR]", err)
	}

	err := readConfigFile()
	if err != nil {
		log.Fatal("[ERROR]", err)
	}

	var check bool
	check, err = connectionToMQTT()
	if err != nil {
		log.Println(err)
	}
	if !check {
		log.Println("[ERROR] Stopping...")
		os.Exit(0)
	}

	for {
		select {
		case object := <-GetInstance().channelData:
			publish(object)
		case <-GetInstance().channelExit:
			return
		}
	}
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
