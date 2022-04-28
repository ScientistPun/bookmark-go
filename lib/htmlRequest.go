package lib

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"
)

var cookie string = ""

var userAgents = []string{"Mozilla/5.0 (compatible, MSIE 10.0, Windows NT, DigExt)",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, 360SE)",
	"Mozilla/4.0 (compatible, MSIE 8.0, Windows NT 6.0, Trident/4.0)",
	"Mozilla/5.0 (compatible, MSIE 9.0, Windows NT 6.1, Trident/5.0,",
	"Opera/9.80 (Windows NT 6.1, U, en) Presto/2.8.131 Version/11.11",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, TencentTraveler 4.0)",
	"Mozilla/5.0 (Windows, U, Windows NT 6.1, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Macintosh, Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
	"Mozilla/5.0 (Macintosh, U, Intel Mac OS X 10_6_8, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Linux, U, Android 3.0, en-us, Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
	"Mozilla/5.0 (iPad, U, CPU OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, Trident/4.0, SE 2.X MetaSr 1.0, SE 2.X MetaSr 1.0, .NET CLR 2.0.50727, SE 2.X MetaSr 1.0)",
	"Mozilla/5.0 (iPhone, U, CPU iPhone OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"MQQBrowser/26 Mozilla/5.0 (Linux, U, Android 2.3.7, zh-cn, MB200 Build/GRJ22, CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1"}

func getRandomUserAgent() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return userAgents[r.Intn(len(userAgents))]
}

func getUrlResponse(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Cookie", cookie)
	req.Header.Add("Agent", getRandomUserAgent())
	return req, nil
}

func getHtmlResponse(url string) (string, error) {
	req, err := getUrlResponse(url)
	if err != nil {
		return "", err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respHtml := string(respByte)
	return respHtml, nil
}

func getDesc(htmlResp string) (string, error) {
	regx := regexp.MustCompile(`(?i)<title>([\s\S]*?)</title>`)
	if regx == nil { //解释失败，返回nil
		return "", &GetHtmlInfoError{name: "Title获取失败"}
	}
	match := regx.FindAllStringSubmatch(htmlResp, -1)
	if len(match) < 1 || len(match[0]) < 2 {
		return "", &GetHtmlInfoError{name: "Title不匹配"}
	}
	desc := strings.TrimSpace(match[0][1])
	desc = strings.ReplaceAll(desc, "&nbsp;", " ")
	return desc, nil
}

func getDomain(url string) (string, string) {
	urlArr := strings.Split(url, "//")
	protocol := urlArr[0]
	addr := urlArr[1]
	domain := strings.Split(addr, "/")[0]

	return protocol, domain
}

func getIcon(htmlResp string, url string) (string, error) {
	regx := regexp.MustCompile(`(?i)<link rel="(shortcut )?(apple-touch-)?icon"(.*?)>`)
	if regx == nil { //解释失败，返回nil
		return "", &GetHtmlInfoError{name: "Icon获取失败"}
	}
	match := regx.FindAllStringSubmatch(htmlResp, -1)
	if len(match) < 1 || len(match[0]) < 2 {
		return "", &GetHtmlInfoError{name: "Icon不匹配"}
	}

	hrefRegx := regexp.MustCompile(`(?i)href="(.*?)"`)
	match = hrefRegx.FindAllStringSubmatch(match[0][0], -1)
	if len(match[0]) < 2 {
		return "", &GetHtmlInfoError{name: "Icon不匹配"}
	}
	icon := match[0][1]

	protocol, domain := getDomain(url)

	if !strings.Contains(icon, protocol) {
		if !strings.Contains(icon, "//") {
			icon = fmt.Sprintf("%v//%v/%v", protocol, domain, icon)
		} else {
			icon = fmt.Sprintf("%v%v", protocol, icon)
		}
	}

	return icon, nil
}

func GetHtmlInfo(url string) (desc string, icon string, err error) {
	resp, err := getHtmlResponse(url)
	if err != nil {
		return "", "", err
	}

	desc, err = getDesc(resp)
	if err != nil { //解释失败，返回nil
		return "", "", err
	}

	icon, err = getIcon(resp, url)
	if err != nil { //解释失败，返回nil
		return "", "", err
	}

	return desc, icon, nil
}

type GetHtmlInfoError struct {
	name string
}

func (e *GetHtmlInfoError) Error() string {
	return e.name
}
