package main

import (
	"encoding/json"
	"fmt"
	"io"
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

const WishHistoryFilePath = "C:\\Program Files\\Star Rail\\Game\\StarRail_Data\\webCaches\\2.15.0.0\\Cache\\Cache_Data\\data_2"
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
	FindAllWish(FindURL(WishHistoryFilePath))
}

func FindURL(filePath string) (string, map[string]string) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", nil
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
	return lastUrl, nMap
}

func ParseQuery(q string) map[string]string {
	m := map[string]string{}
	for _, s := range strings.Split(q, "&") {
		split := strings.Split(s, "=")
		m[split[0]] = split[1]
	}
	return m
}

func FindAllWish(domainLink string, m map[string]string) {
	allList := map[string][]HKRPGWish{}
	for k, v := range gachaTypeMap {
		var gachaList []HKRPGWish
		fmt.Printf("开始获取[%s]\n", v)
		page := 1
		size := 5
		pageStr := strconv.Itoa(page)
		m["gacha_type"] = k
		m["page"] = pageStr
		m["end_id"] = "0"
		m["size"] = strconv.Itoa(size)
		for {
			fetchData, _ := FetchData(domainLink + "?" + MapToStr(m))
			dataList := fetchData.Data.List
			for _, wish := range dataList {
				gachaList = append(gachaList, wish)
				fmt.Println(wish.String())
			}
			dataLen := len(dataList)
			if dataLen == 0 {
				break
			}
			m["end_id"] = dataList[dataLen-1].Id
			if dataLen < size {
				break
			}
			page++
		}
		allList[k] = gachaList
	}
	MergeToFile(allList)
}

func MergeToFile(allList map[string][]HKRPGWish) {
	file, err := os.OpenFile(JSONFilePath, syscall.O_RDWR|syscall.O_CREAT, os.ModePerm)
	if err != nil {
		fmt.Printf("打开文件失败: %s\n", err.Error())
	}
	defer file.Close()
	all, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("读取文件失败: %s\n", err.Error())
	}
	fmt.Printf("已经存储的数据:\n%s\n", string(all))
	ist := map[string][]HKRPGWish{}
	_ = json.Unmarshal(all, &ist)
	for s, wishes := range allList {
		// Simple 如果不同则append
		istWishs := ist[s]
		idList := []string{}
		for i := range istWishs {
			idList = append(idList, istWishs[i].Id)
		}
		for w := range wishes {
			if slices.Contains(idList, wishes[w].Id) {
				continue
			}
			ist[s] = append(ist[s], wishes[w])
		}
	}
	StoreToFile(ist)
}

func StoreToFile(allList map[string][]HKRPGWish) {
	allList = SortMapWish(allList)
	marshal, err := json.Marshal(allList)
	if err != nil {
		fmt.Printf("JSON序列化失败[%s]\n", err.Error())
		return
	}
	err = os.WriteFile(JSONFilePath, marshal, syscall.O_RDWR|syscall.O_CREAT)
	if err != nil {
		fmt.Printf("写入文件失败[%s]\n", err.Error())
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

func FetchData(link string) (Page[HKRPGWish], error) {
	time.Sleep(8 * time.Second)
	resp, err := http.Get(link)
	if err != nil {
		return Page[HKRPGWish]{}, err
	}
	body := resp.Body
	defer body.Close()
	bodyByte, httpReadErr := io.ReadAll(resp.Body)
	if httpReadErr != nil {
		return Page[HKRPGWish]{}, httpReadErr
	}
	p := Page[HKRPGWish]{}
	_ = json.Unmarshal(bodyByte, &p)
	return p, nil
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
