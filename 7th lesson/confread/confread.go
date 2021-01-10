package confread

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type ConfigStruct struct {
	Port        int    `yaml:"port"`
	DbURL       string `yaml:"db_url"`
	JaegerURL   string `yaml:"jaeger_url"`
	SentryURL   string `yaml:"sentry_url"`
	KafkaBroker string `yaml:"kafka_broker"`
	SomeAppID   string `yaml:"some_app_id"`
	SomeAppKey  string `yaml:"some_app_key"`
}

func checkconfurls(urlstring, prefix string) {
	_, err := url.ParseRequestURI(urlstring)
	if err != nil {
		fmt.Print(err)
	}
	if !strings.HasPrefix(urlstring, prefix) {
		fmt.Printf("Не задан префикс протокола:%v\n", urlstring)
	}
}
func ReadConf() {
	// слайс для чтения конфига
	buffer := make([]byte, 2000)

	// Проверяем, что файл существует
	if _, err := os.Stat("./conf.yaml"); err != nil {
		log.Fatalf("Конфиг файл не существует: %v", err)
	}

	//Открываем конфиг с чеком на ошибки
	file, err := os.Open("./conf.yaml")
	if err != nil {
		log.Fatal(err)
	}
	//Читаем файл в буфер зачем-то, пока непонятно зачем :)
	_, err = file.Read(buffer)
	//Задаем закрытие файла в случае падения\завершение программы
	defer func() {
		err = file.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	confiFile, err := ioutil.ReadFile("./conf.yaml")

	appconfig := ConfigStruct{}
	err = yaml.Unmarshal(confiFile, &appconfig)
	if err != nil {
		log.Printf("Не могу декодировать yaml-файл в структуру: %v", err)
	}
	checkconfurls(appconfig.JaegerURL, "http")
	checkconfurls(appconfig.SentryURL, "http")
	checkconfurls(appconfig.DbURL, "postgres")

	fmt.Printf("\nConfigfile:\n %+v\n", appconfig)
}
