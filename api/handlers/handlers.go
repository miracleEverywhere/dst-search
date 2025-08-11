package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"sync"
	"time"
)

type SearchResponse struct {
	GET []SearchRoomItem `json:"GET"`
}

type SearchRoomItem struct {
	Addr            string               `json:"__addr"`
	RowID           string               `json:"__rowId"`
	Host            string               `json:"host"`
	ClanOnly        bool                 `json:"clanonly"`
	Platform        int                  `json:"platform"`
	Mods            bool                 `json:"mods"`
	Name            string               `json:"name"`
	PvP             bool                 `json:"pvp"`
	Session         string               `json:"session"`
	Fo              bool                 `json:"fo"`
	Password        bool                 `json:"password"`
	GUID            string               `json:"guid"`
	MaxConnections  int                  `json:"maxconnections"`
	Dedicated       bool                 `json:"dedicated"`
	ClientHosted    bool                 `json:"clienthosted"`
	Connected       int                  `json:"connected"`
	Mode            string               `json:"mode"`
	Port            int                  `json:"port"`
	Version         int                  `json:"v"`
	Tags            string               `json:"tags"`
	Season          string               `json:"season"`
	LanOnly         bool                 `json:"lanonly"`
	Intent          string               `json:"intent"`
	AllowNewPlayers bool                 `json:"allownewplayers"`
	ServerPaused    bool                 `json:"serverpaused"`
	SteamID         string               `json:"steamid"`
	SteamRoom       string               `json:"steamroom"`
	Secondaries     map[string]Secondary `json:"secondaries"`
	Region          string               `json:"region"`
}

type Secondary struct {
	Addr    string `json:"__addr"`
	ID      string `json:"id"`
	SteamID string `json:"steamid"`
	Port    int    `json:"port"`
}

type DetailResponse struct {
	GET []ServerDetail `json:"GET"`
}

type ServerDetail struct {
	Addr            string `json:"__addr"`
	RowID           string `json:"__rowId"`
	Host            string `json:"host"`
	ClanOnly        bool   `json:"clanonly"`
	Platform        int    `json:"platform"`
	Mods            bool   `json:"mods"`
	Name            string `json:"name"`
	PvP             bool   `json:"pvp"`
	Session         string `json:"session"`
	Fo              bool   `json:"fo"`
	Password        bool   `json:"password"`
	GUID            string `json:"guid"`
	MaxConnections  int    `json:"maxconnections"`
	Dedicated       bool   `json:"dedicated"`
	ClientHosted    bool   `json:"clienthosted"`
	Connected       int    `json:"connected"`
	Mode            string `json:"mode"`
	Port            int    `json:"port"`
	Version         int    `json:"v"`
	Tags            string `json:"tags"`
	Season          string `json:"season"`
	LanOnly         bool   `json:"lanonly"`
	Intent          string `json:"intent"`
	AllowNewPlayers bool   `json:"allownewplayers"`
	ServerPaused    bool   `json:"serverpaused"`
	SteamID         string `json:"steamid"`
	SteamRoom       string `json:"steamroom"`
	Data            string `json:"data"`
	WorldGen        string `json:"worldgen"`
	Players         string `json:"players"`
	ModsInfo        []any  `json:"mods_info"`
	Desc            string `json:"desc"`
	Tick            int    `json:"tick"`
	ClientModsOff   bool   `json:"clientmodsoff"`
	NAT             int    `json:"nat"`
}

func Search(c *gin.Context) {
	type SearchForm struct {
		Text    string   `form:"text" json:"text"`
		Regions []string `form:"regions" json:"regions"`
	}

	var searchForm SearchForm
	if err := c.ShouldBindJSON(&searchForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pattern := `.*` + searchForm.Text + `.*`
	re := regexp.MustCompile(pattern)

	// 使用 WaitGroup 等待所有 goroutine 完成
	var wg sync.WaitGroup
	// 使用 mutex 保护共享数据
	var mu sync.Mutex
	// 结果收集
	var data []SearchRoomItem

	// 为每个 region 启动一个 goroutine
	for _, region := range searchForm.Regions {
		wg.Add(1)
		go func(region string) {
			defer wg.Done()

			url := fmt.Sprintf("https://lobby-v2-cdn.klei.com/%s-Steam.json.gz", region)
			client := &http.Client{
				Timeout: 30 * time.Second,
			}

			httpResponse, err := client.Get(url)
			if err != nil {
				return
			}
			defer httpResponse.Body.Close()

			if httpResponse.StatusCode != http.StatusOK {
				return
			}

			var searchResponse SearchResponse
			if err := json.NewDecoder(httpResponse.Body).Decode(&searchResponse); err != nil {
				return
			}

			// 过滤匹配的服务器
			var matchedItems []SearchRoomItem
			for _, item := range searchResponse.GET {
				if re.MatchString(item.Name) {
					item.Region = region
					matchedItems = append(matchedItems, item)
				}
			}

			// 将结果添加到共享数据
			if len(matchedItems) > 0 {
				mu.Lock()
				data = append(data, matchedItems...)
				mu.Unlock()
			}
		}(region)
	}

	// 等待所有 goroutine 完成
	wg.Wait()

	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "success", "data": data})
}

func Detail(c *gin.Context) {
	type DetailForm struct {
		RowID  string `form:"rowId" json:"rowId"`
		Region string `form:"region" json:"region"`
	}

	var detailForm DetailForm
	if err := c.ShouldBindJSON(&detailForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url := fmt.Sprintf("https://lobby-v2-%s.klei.com/lobby/read", detailForm.Region)

	payload := map[string]interface{}{
		"__token": "pds-g^KU_JpEFW0P3^GLDPpmMYhgm/Bd254Xwi9MqRouHutahdIfGzv4Nv+zE=",
		"query": map[string]string{
			"__rowId": detailForm.RowID,
		},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 201, "message": "请求科雷服务器失败", "data": nil})
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 201, "message": "请求科雷服务器失败", "data": nil})
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 201, "message": "请求科雷服务器失败", "data": nil})
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusOK, gin.H{"code": 201, "message": "请求科雷服务器失败", "data": nil})
		return
	}

	var detailResponse DetailResponse
	if err := json.NewDecoder(resp.Body).Decode(&detailResponse); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 201, "message": "请求科雷服务器失败", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "success", "data": detailResponse.GET[0]})
}
