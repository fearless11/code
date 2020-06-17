// 表组测试
//  测试一组不同的输入并产生不同的输出结果
//  表组测试是利用一个测试函数测试多组值的好办法
package table

import (
	"net/http"
	"testing"
)

// 运行查看效果
//  go test -v

const checkMark = "\u2713"
const ballotX = "\u2717"

// TestDownload 确认http包的Get函数下载内容并正确处理不同的状态
//  扩展测试只需要将新的URL和statusCode加入表组就可以
func TestDownload(t *testing.T) {
	var urls = []struct {
		url        string
		statusCode int
	}{
		{
			"http://www.npr.org/rss/rss.php?id=1001",
			http.StatusOK,
		},
		{
			"http://feeds.nbcnews,com/feeds/topstories",
			http.StatusNotFound,
		},
	}

	t.Log("Given the need to test downloading different content.")
	{
		for _, u := range urls {
			t.Logf("\tWhen checking \"%s\" for status code \"%d\"", u.url, u.statusCode)
			{
				resp, err := http.Get(u.url)
				if err != nil {
					t.Fatal("\t\tShould be able to Get the url.", ballotX, err)
				}
				t.Log("\t\tShould be able to Get the url", checkMark)
				defer resp.Body.Close()

				if resp.StatusCode == u.statusCode {
					t.Logf("\t\tShould have a \"%d\" status. %v", u.statusCode, checkMark)
				} else {
					t.Errorf("\t\tShould have a \"%d\" status %v %v", u.statusCode, ballotX, resp.StatusCode)
				}
			}
		}
	}

}
