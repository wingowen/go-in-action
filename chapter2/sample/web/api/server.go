package api

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"go-in-action/chapter2/sample/search"
)

// SearchResult 定义了返回给前端的搜索结果结构
type SearchResult struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Source      string `json:"source"`
	Date        string `json:"date"`
}

// SearchHandler 处理搜索请求
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// 只允许GET请求
	if r.Method != http.MethodGet {
		http.Error(w, "只支持GET请求", http.StatusMethodNotAllowed)
		return
	}

	// 获取搜索关键词
	searchTerm := r.URL.Query().Get("q")
	if searchTerm == "" {
		http.Error(w, "请提供搜索关键词", http.StatusBadRequest)
		return
	}

	log.Printf("收到搜索请求: %s\n", searchTerm)

	// 创建结果通道
	results := make(chan *search.Result)

	// 执行搜索
	go func() {
		// 获取feed列表
		feeds, err := search.RetrieveFeeds()
		if err != nil {
			log.Printf("获取feed失败: %v\n", err)
			close(results)
			return
		}

		// 搜索所有feed
		var wg sync.WaitGroup
		for _, feed := range feeds {
			wg.Add(1)
			go func(feed *search.Feed) {
				defer wg.Done()
				matcher, exists := search.GetMatcher(feed.Type)
				if !exists {
					// 获取默认匹配器
					defaultMatcher, _ := search.GetMatcher("default")
					matcher = defaultMatcher
				}
				search.Match(matcher, feed, searchTerm, results)
			}(feed)
		}
		wg.Wait()
		close(results)
	}()

	// 收集结果
	var searchResults []SearchResult
	for result := range results {
		// 对于RSS匹配器，我们知道结果格式
		// 这里简化处理，直接使用结果
		searchResults = append(searchResults, SearchResult{
			Title:       result.Content,
			Description: "", // 在实际应用中，需要从feed项目中提取
			Source:      result.Field,
			Date:        "", // 在实际应用中，需要从feed项目中提取
		})
	}

	// 返回JSON结果
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(searchResults)
}

// GetMatcher 直接调用search包的GetMatcher函数
func GetMatcher(feedType string) (search.Matcher, bool) {
	return search.GetMatcher(feedType)
}

// StartServer 启动HTTP服务器
func StartServer() {
	http.HandleFunc("/api/search", SearchHandler)

	log.Println("服务器启动在 http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
