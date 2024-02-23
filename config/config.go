package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

//bazı fonksiyonlar bazı fonksiyonların içinde kullanılacak.

// verilen path doğru mu?
func pathExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if !os.IsNotExist(err) {
		return false
	}
	return true
}

// dosyam var mı?
func fileExists(fileName string) (bool, string) {
	_, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return false, ""
		}
	}
	return true, fileName
}

type Config struct {
	ClientID     string `yaml:"clientID"`
	ClientSecret string `yaml:"clientSecret"`
}

// dosyayı oluşturmak
func CreateConfigFile(clientID string, clientSecret string) {

	configg := Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}

	yamlFile, err := yaml.Marshal(&configg)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("config.yaml", yamlFile, 0644)
	if err != nil {
		log.Fatal(err)
	}

}

func ReadConfigFile() (Config, error) {
	var c Config

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return c, err
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return c, err
	}
	return c, nil

}
