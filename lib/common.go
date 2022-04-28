package lib

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

const BookmarkPath = "./data/bookmark.yaml"
const SearchPath = "./data/search.yaml"
const ConfigPath = "./data/config.yaml"
const RoutesPath = "./data/routes.yaml"

func LoadSkins() ([]string, error) {
	files, err := ioutil.ReadDir("./skin")
	if err != nil {
		return []string{}, err
	}

	var folders []string
	for _, file := range files {
		if file.IsDir() {
			folders = append(folders, file.Name())
		}
	}
	return folders, nil
}

func LoadYamlFile(filePath string, T interface{}) (err error) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return gin.Error{Err: err, Type: 1}
	}
	if err != nil {
		return gin.Error{Err: err, Type: 2}
	}
	err = yaml.Unmarshal(file, T)
	if err != nil {
		log.Printf("Unmarshal err: %v ", err.Error())
	}
	return nil
}

func WriteYamlFile(filePath string, T interface{}) (err error) {
	if !IsFileExists(filePath) {
		f, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer f.Close()
	}
	data, err := yaml.Marshal(T)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, data, 0777)
	if err != nil {
		return err
	}
	return nil
}

func WriteFile(filePath string, data []byte) (err error) {
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, data, 0777)
	if err != nil {
		return err
	}
	return nil
}

func InArray(target interface{}, array []interface{}) bool {
	for _, element := range array {
		if target == element {
			return true
		}
	}
	return false
}

func InIntArray(target int, array []int) bool {
	for _, element := range array {
		if target == element {
			return true
		}
	}
	return false
}

func IsFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
