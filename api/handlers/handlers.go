package handlers

import (
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
}
type Secondary struct {
	Addr    string `json:"__addr"`
	ID      string `json:"id"`
	SteamID string `json:"steamid"`
	Port    int    `json:"port"`
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

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}

func Detail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "快速入门",
	})
}
