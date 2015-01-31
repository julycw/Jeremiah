package module

import (
	"code.google.com/p/mahonia"
	"errors"
	"fmt"
	"github.com/julycw/Jeremiah/models/mgo_models"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
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

	var computer mgo_models.Computer
	//匹配概览数据列表
	ptnContentInfoList := regexp.MustCompile(`(?s)<ul id="parameter2" class="p-parameter-list">(.*?)</ul>`)
	//匹配概览数据
	ptnContentInfoValue := regexp.MustCompile(`<li title='(.*?)'>(.*?)：(.*?)</li>`)

	//匹配详细数据表格
	ptnContentTable := regexp.MustCompile(`(?s)<table cellpadding="0" cellspacing="1" width="100%" border="0" class="Ptable">(.*)</table>`)
	//匹配各个模块信息，如主体、CPU、内存等
	ptnContentSection := regexp.MustCompile(`="2">(.*)</th><tr>(?s:.*?)<th class="tdTitle" colspan`)
	//匹配各个参数信息，如内存4G，硬盘500GB等
	ptnContentRow := regexp.MustCompile(`<td class="tdTitle">(.*)</td><td>(.*)</td>`)

	//匹配概览数据列
	matchList := ptnContentInfoList.FindStringSubmatch(content)
	if matchList != nil {
		listContent := matchList[1]
		matcheSections := ptnContentInfoValue.FindAllStringSubmatch(listContent, -1)
		for _, section := range matcheSections {
			switch strings.TrimSpace(strings.TrimSpace(section[2])) {
			case "商品名称":
				computer.Name = section[1]
			case "品牌":
				computer.Brand = section[1]
			case "触摸":
				computer.ScreenType = section[1]
			case "电脑类别":
				computer.Type = section[1]
			case "商品编号":
				computer.JDNumber = section[1]
			case "显示器尺寸":
				computer.ScreenSize = section[1]
			case "显卡":
				computer.GraphicsType = section[1]
			}
		}
	}
	//匹配表格
	matchTable := ptnContentTable.FindStringSubmatch(content)
	if matchTable != nil {
		//加这个后缀
		tableContent := matchTable[1] + "<th class=\"tdTitle\" colspan"
		matcheSections := ptnContentSection.FindAllStringSubmatch(tableContent, -1)
		for _, section := range matcheSections {
			matchRows := ptnContentRow.FindAllStringSubmatch(section[0], -1)
			switch strings.TrimSpace(section[1]) {
			case "主体":
				for _, row := range matchRows {
					switch strings.TrimSpace(row[1]) {
					case "系列":
						computer.Series = row[2]
					case "型号":
						computer.Model = row[2]
					case "平台":
						computer.Platform = row[2]
					case "操作系统":
						computer.OS = row[2]
					}
				}
			case "CPU":
				for _, row := range matchRows {
					switch strings.TrimSpace(row[1]) {
					case "类型":
						computer.CPUBrand = row[2]
					case "CPU型号":
						computer.CPUModel = row[2]
					case "速度":
						computer.CPUFrequency = row[2]
					case "核心数":
						computer.CPUCoreNum = row[2]
					case "二级缓存":
						computer.CPUSecondCache = row[2]
					case "三级缓存":
						computer.CPUThirdCache = row[2]
					}
				}
			case "显卡":
				for _, row := range matchRows {
					switch strings.TrimSpace(row[1]) {
					case "显卡品牌":
						computer.GraphicsBrand = row[2]
					case "显卡芯片", "显示芯片":
						computer.GraphicsModel = row[2]
					case "显卡容量", "显存容量":
						computer.GraphicsMemorySize = row[2]
					case "显存规格":
						computer.GraphicsMemoryType = row[2]
					}
				}
			case "主板":
				for _, row := range matchRows {
					switch strings.TrimSpace(row[1]) {
					case "芯片组":
						//none
					case "显卡类型":
						computer.GraphicsType = row[2]
					case "声卡":
						computer.AudioCard = row[2]
					case "网卡":
						computer.WebCard = row[2]
					}
				}
			case "内存":
				for _, row := range matchRows {
					switch strings.TrimSpace(row[1]) {
					case "容量":
						computer.MemorySize = row[2]
					case "速度":
						computer.MemoryType = row[2]
					}
				}
			case "硬盘":
				for _, row := range matchRows {
					switch strings.TrimSpace(row[1]) {
					case "容量":
						computer.DiskSize = row[2]
					case "类型":
						computer.DiskType = row[2]
					case "转速":
						computer.DiskSpeed = row[2]
					}
				}
			case "光驱":
				for _, row := range matchRows {
					switch strings.TrimSpace(row[1]) {
					case "类型":
						computer.CDRom = row[2]
					}
				}
			case "输入设备":
				// None
			case "规格":
				for _, row := range matchRows {
					switch strings.TrimSpace(row[1]) {
					case "电源":
						computer.Power = row[2]
					case "电源功率":
						computer.Power += " " + row[2]
					case "尺寸":
						computer.Size = row[2]
					case "重量":
						computer.Weight = row[2]
					}
				}
			case "特性":
				for _, row := range matchRows {
					switch strings.TrimSpace(row[1]) {
					case "特性":
						computer.Feature = row[2]
					}
				}
			}
		}
	}

	if matchList == nil && matchTable == nil {
		return nil, errors.New("无法从该页面获取信息")
	}

	computer.ID = bson.NewObjectId()
	computer.IDStr = computer.ID.Hex()

	//获取价格
	urlGetPrice := fmt.Sprintf("http://p.3.cn/prices/get?skuid=J_%s", computer.JDNumber)
	if c, code := readContent(urlGetPrice); code == 200 {
		ptnGetPrice := regexp.MustCompile(`"p":"(.*)","m":"(.*)"`)
		matchPrice := ptnGetPrice.FindStringSubmatch(c)
		if matchPrice != nil {
			temp, _ := strconv.ParseFloat(matchPrice[1], 32)
			computer.JDPrice = float32(temp)
			temp, _ = strconv.ParseFloat(matchPrice[2], 32)
			computer.Price = float32(temp)
		}
	}

	return computer, nil
}

func readContent(url string) (content string, code int) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
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
