package lib

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Link struct {
	Title string `yaml:"title"`
	Desc  string `yaml:"desc"`
	Url   string `yaml:"url"`
	Icon  string `yaml:"icon"`
}

type Tag struct {
	Icon  string `yaml:"icon"`
	Title string `yaml:"title"`
	Name  string `yaml:"name"`
	Links []Link
}

type LinkRange struct {
	Sort     int
	TagTitle string
	TagName  string
	Link     Link
}

type TagRange struct {
	Sort int
	Tag  Tag
}

func InitBookmark() []Tag {
	sampleLink := Link{
		Title: "bookmark",
		Desc:  "test sample",
		Url:   "http://localhost:8080",
		Icon:  "",
	}

	bookmark := Tag{
		Title: "常用标签",
		Name:  "common",
		Icon:  "fa-tags",
		Links: []Link{sampleLink},
	}
	tags := []Tag{bookmark}
	WriteYamlFile(BookmarkPath, tags)
	return tags
}

func LoadBookmark() (tagArray []Tag, err error) {
	var tags []Tag

	if !IsFileExists(BookmarkPath) {
		tags = InitBookmark()
	} else {
		err = LoadYamlFile(BookmarkPath, &tags)
		if err != nil {
			return nil, err
		}
	}
	return tags, nil
}

func SetBookmarkText(text string) (err error) {
	content := []byte(text)
	var tags []Tag
	err = yaml.Unmarshal(content, &tags)
	if err != nil {
		return err
	}
	err = WriteFile(BookmarkPath, content)
	if err != nil {
		return err
	}
	return nil
}

func GetBookmarkTextFrom(bookmark []Tag) (text string, err error) {
	txt := ""
	for _, tag := range bookmark {
		txt += fmt.Sprintf("- title: %v\n  name: %v\n  links: \n", tag.Title, tag.Name)
		for _, link := range tag.Links {
			txt += fmt.Sprintf("  - title: %v\n    desc: %v\n    url: %v\n    icon: %v\n", link.Title, link.Desc, link.Url, link.Icon)
		}
	}

	return txt, nil
}
