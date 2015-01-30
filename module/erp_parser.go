package module

import (
	"code.google.com/p/mahonia"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	// "github.com/julycw/Jeremiah/models/mgo_models"
)

var (
	ptnBrTag   = regexp.MustCompile(`<br>`)
	ptnHTMLTag = regexp.MustCompile(`(?s)</?.*?>`)
	ptnSpace   = regexp.MustCompile(`(^\s+)|( )`)
)

func ConvertGBKToUTF8(str string) string {
	enc := mahonia.NewDecoder("gbk")
	return enc.ConvertString(str)
}

func ParseJD(url string) (interface{}, error) {
	content, code := readContent(url)
	if code != 200 {
		return nil, errors.New(fmt.Sprintf("Error:%d\n", code))
	}

	ptnContentDetail := regexp.MustCompile(`(?s).*<table cellpadding="0" cellspacing="1" width="100%" border="0" class="Ptable">(.*)</table>.*`)
	ptnContentRow := regexp.MustCompile(`<td class="tdTitle">(.*)</td><td>(.*)</td>`)

	match := ptnContentDetail.FindStringSubmatch(content)
	if match != nil {
		content = match[1]
		// matchMult := ptnContentRow.FindAllStringSubmatch(content, 100000)
		// fmt.Println(len(matchMult))
		// for _, v := range matchMult {
		// 	if v != nil {
		// 		fmt.Println(v[1], v[2])
		// 	} else {
		// 		fmt.Println("none")
		// 	}
		// }
		// re := regexp.MustCompile("a(x*)b")
		// fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-", -1))
		// fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-", -1))
		// fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-axb-", -1))
		// fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-ab-", -1))

		matches := ptnContentRow.FindAllStringSubmatch(content, -1)
		for _, item := range matches {
			fmt.Printf("参数:%s / 值:%s\n", item[1], item[2])
		}
	}

	return content, nil
}

func readContent(url string) (content string, code int) {
	response, err := http.Get(url)
	if err != nil {
		return "", -100
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", -100
	}
	content = ConvertGBKToUTF8(string(body))
	content = ptnBrTag.ReplaceAllString(content, "\r\n")
	code = response.StatusCode
	return
}
