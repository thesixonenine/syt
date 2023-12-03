package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const WishHistoryFilePath = "C:\\Program Files\\Star Rail\\Game\\StarRail_Data\\webCaches\\2.18.0.0\\Cache\\Cache_Data\\data_2"
const JSONFilePath = "../../src/assets/data/star-rail-wish.json"

var re = regexp.MustCompile(`\p{C}`)

var gachaTypeMap = map[string]string{
	"11": "角色活动跃迁",
	"12": "光锥活动跃迁",
	"1":  "常驻跃迁",
	// "2": "新手跃迁",
}
var absParams = []string{"authkey", "authkey_ver", "sign_type", "game_biz", "lang", "auth_appid", "size", "gacha_type", "page", "end_id"}

func main() {
	// 使用抽卡 URL 进行循环查询抽卡历史, 一但发现已经存在于历史 JSON 文件中, 则停止查询
	FetchWishes(FindURLFromLogFile(), LocalHistoryJSONFileToMap())
}

// FindURLFromLogFile 查询日志文件中的抽卡 URL
func FindURLFromLogFile() UrlParam {
	content, err := os.ReadFile(WishHistoryFilePath)
	if err != nil {
		log.Fatalf("日志文件[%s]未找到", WishHistoryFilePath)
	}
	lastUrl := ""
	nMap := map[string]string{}
	for _, s := range strings.Split(string(content), "1/0/") {
		t := re.ReplaceAllString(strings.TrimSpace(s), "")
		if !strings.HasPrefix(s, "http") && !strings.Contains(t, "getGachaLog") {
			continue
		}
		u, parseErr := url.Parse(t)
		if parseErr != nil {
			continue
		}
		queryMap := ParseQuery(u.RawQuery)

		for k, v := range queryMap {
			if slices.Contains(absParams, k) {
				nMap[k] = v
			}
		}
		lastUrl = u.Scheme + "://" + u.Host + u.Path
	}
	return UrlParam{lastUrl, nMap}
}

// ParseQuery 解析 URL 参数为 map
func ParseQuery(q string) map[string]string {
	m := map[string]string{}
	for _, s := range strings.Split(q, "&") {
		split := strings.Split(s, "=")
		if len(split) >= 2 {
			m[split[0]] = split[1]
		}
	}
	return m
}

// LocalHistoryJSONFileToMap 将本地抽卡历史 JSON 文件转为 map 对象, 如果文件不存在则创建
func LocalHistoryJSONFileToMap() map[string][]HKRPGWish {
	historyFile, err := os.OpenFile(JSONFilePath, syscall.O_RDWR|syscall.O_CREAT, os.ModePerm)
	if err != nil {
		log.Fatalf("打开文件[%s]异常,err[%s]\n", JSONFilePath, err.Error())
	}
	defer func() {
		if err := historyFile.Close(); err != nil {
			log.Fatalf("关闭文件[%s]异常,err[%s]\n", JSONFilePath, err.Error())
		}
	}()
	historyFileContent, err := io.ReadAll(historyFile)
	if err != nil {
		log.Fatalf("读取文件[%s]异常,err[%s]\n", JSONFilePath, err.Error())
	}
	history := map[string][]HKRPGWish{}
	_ = json.Unmarshal(historyFileContent, &history)
	return history
}

// FetchWishes 从指定 URL 及参数中拉取抽卡参数, 并追加到 Map 中
func FetchWishes(urlParam UrlParam, localHistoryMap map[string][]HKRPGWish) {
	paramMap := urlParam.ParamMap
	// 循环抽卡类型
	for k, v := range gachaTypeMap {
		localIdList := MapToId(localHistoryMap[k])
		fmt.Printf("开始获取[%s]\n", v)
		page := 1
		size := 5
		paramMap["gacha_type"] = k
		paramMap["page"] = strconv.Itoa(page)
		paramMap["end_id"] = "0"
		paramMap["size"] = strconv.Itoa(size)
		for {
			wishList := FetchData(urlParam.BaseUrl + "?" + MapToStr(paramMap)).Data.List
			isContains := false
			for _, wish := range wishList {
				if slices.Contains(localIdList, wish.Id) {
					isContains = true
					continue
				}
				localHistoryMap[k] = append(localHistoryMap[k], wish)
				fmt.Println(wish.String())
			}
			if isContains {
				break
			}
			dataLen := len(wishList)
			if dataLen == 0 {
				break
			}
			paramMap["end_id"] = wishList[dataLen-1].Id
			if dataLen < size {
				break
			}
			page++
		}
	}
	StoreToFile(localHistoryMap)
}

func MapToId(wishes []HKRPGWish) []string {
	var idList []string
	for i := range wishes {
		idList = append(idList, wishes[i].Id)
	}
	return idList
}

func StoreToFile(allList map[string][]HKRPGWish) {
	allList = SortMapWish(allList)
	marshal, err := json.Marshal(allList)
	if err != nil {
		fmt.Printf("JSON序列化异常[%s]\n", err.Error())
		return
	}
	WriteToFile(JSONIndent(marshal))
}

func JSONIndent(marshal []byte) []byte {
	var out bytes.Buffer
	err := json.Indent(&out, marshal, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	return out.Bytes()
}

func WriteToFile(marshal []byte) {
	err := os.WriteFile(JSONFilePath, marshal, syscall.O_RDWR|syscall.O_CREAT)
	if err != nil {
		fmt.Printf("写入文件异常[%s]\n", err.Error())
	}
}

func SortMapWish(ist map[string][]HKRPGWish) map[string][]HKRPGWish {
	list := map[string][]HKRPGWish{}
	for k, v := range ist {
		sort.Slice(v, func(i, j int) bool {
			return v[i].Id < v[j].Id
		})
		list[k] = v
	}
	return list
}

func MapToStr(m map[string]string) string {
	var sl []string
	for k, v := range m {
		sl = append(sl, k+"="+v)
	}
	return strings.Join(sl, "&")
}

func FetchData(link string) Page[HKRPGWish] {
	time.Sleep(5 * time.Second)
	resp, err := http.Get(link)
	if err != nil {
		log.Fatalf("HTTP请求异常,err[%s]", err.Error())
	}
	body := resp.Body
	defer func() {
		if err := body.Close(); err != nil {
			log.Fatalf("关闭HTTP请求异常,err[%s]", err.Error())
		}
	}()
	bodyByte, httpReadErr := io.ReadAll(resp.Body)
	if httpReadErr != nil {
		log.Fatalf("读取HTTP Body异常,err[%s]", err.Error())
	}
	p := Page[HKRPGWish]{}
	_ = json.Unmarshal(bodyByte, &p)
	return p
}

type HKRPGWish struct {
	Uid       string `json:"uid"`
	GachaId   string `json:"gacha_id"`
	GachaType string `json:"gacha_type"`
	ItemId    string `json:"item_id"`
	Count     string `json:"count"`
	Time      string `json:"time"`
	Name      string `json:"name"`
	Lang      string `json:"lang"`
	ItemType  string `json:"item_type"`
	RankType  string `json:"rank_type"`
	Id        string `json:"id"`
}

func (wish HKRPGWish) String() string {
	return fmt.Sprintf("%s %s %s", wish.Time, gachaTypeMap[wish.GachaType], wish.Name)
}

type Page[T any] struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		Page           int    `json:"page"`
		Size           int    `json:"size"`
		List           []T    `json:"list"`
		Region         string `json:"region"`
		RegionTimeZone int    `json:"region_time_zone"`
	} `json:"data"`
}

type UrlParam struct {
	BaseUrl  string
	ParamMap map[string]string
}
