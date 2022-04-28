package lib

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

type Search struct {
	Title  string `yaml:"title"`
	Query  string `yaml:"query"`
	Enable bool   `yaml:"enable"`
}

type SearchEngine struct {
	Title string
	Url   string
	Param string
}

func InitSearch() []Search {
	baidu := Search{
		Title:  "Baidu",
		Query:  "https://www.baidu.com/s?wd=",
		Enable: true,
	}
	bing := Search{
		Title:  "必应",
		Query:  "https://www.bing.com/search?q=",
		Enable: true,
	}
	searchArray := []Search{baidu, bing}
	WriteYamlFile(SearchPath, searchArray)
	return searchArray
}

func LoadSearch() (searchArray []Search, err error) {
	if !IsFileExists(SearchPath) {
		searchArray = InitSearch()
	} else {
		err = LoadYamlFile(SearchPath, &searchArray)
		if err != nil {
			return nil, err
		}
	}
	return searchArray, nil
}

func LoadSearchEngine() (engines []SearchEngine, err error) {
	array, err := LoadSearch()
	if err != nil {
		return nil, err
	}

	for _, element := range array {
		actions := strings.Split(element.Query, "?")
		engines = append(engines, SearchEngine{
			Title: element.Title,
			Url:   actions[0],
			Param: strings.Trim(actions[1], "="),
		})
	}
	return engines, nil
}

func GetSearchEngineTextFrom() (text string, err error) {
	array, err := LoadSearch()
	if err != nil {
		return "", err
	}
	txt := ""
	for _, element := range array {
		enable := "false"
		if element.Enable {
			enable = "true"
		}
		txt += fmt.Sprintf("- title: %v\n  query: %v\n  enable: %v\n", element.Title, element.Query, enable)
	}

	return txt, nil
}

func SetSearchEngineText(text string) (err error) {
	content := []byte(text)
	var array []Search
	err = yaml.Unmarshal(content, &array)
	if err != nil {
		return err
	}
	err = WriteFile(SearchPath, content)
	if err != nil {
		return err
	}
	return nil
}
