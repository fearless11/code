// 展示如何内部模仿HTTP GET调用
package mock

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 运行查看效果
//  go test -v

// mock测试场景：当需要以来第三方服务而无法操作第三方服务，如数据库、网络
const checkMark = "\u2713"
const ballotX = "\u2717"

// feed模仿mock期望接收的XML文档
var feed = `<?xml version="1.0" encodig="UTF-8"?>
<rss>
<channel>
	<title>Going Go Programming</title>
	<description>Golang: https://github.com/goingo</description>
	<link>http://www.goinggo.net/</link>
	<item>
		<pubDate>Sun, 15 Mar 2015 15:04:00 +0000<pubDate>
		<title>Object Oriented Programming Mechanics</title>
		<description>Go is an object oriented language.</description>
		<link>http://www.goinggo.net/2015/03/object-oriented</link>
	</item>
</channel>
</rss>
`

func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Sprintln(w, feed)
	}
	return httptest.NewServer(http.HandlerFunc(f))
}

func TestDownload(t *testing.T) {
	statusCode := http.StatusOK

	// 生成模拟服务器
	server := mockServer()
	defer server.Close()

	t.Log("Given the need to test downloading content")
	{
		{
			t.Logf("\tWhen checking \"%s\" for status code \"%d\"", server.URL, statusCode)
			resp, err := http.Get(server.URL)
			{
				if err != nil {
					t.Fatal("\t\tShould be able to make the Get call.", ballotX, err)
				}
				t.Log("\t\tShould be able to make the Get call.", checkMark)

				defer resp.Body.Close()

				if resp.StatusCode != statusCode {
					t.Errorf("\t\tShould receive a \"%d\" status. %v %v", statusCode, ballotX, resp.StatusCode)
				}
				t.Logf("\t\tShould receive a \"%d\" status. %v", statusCode, checkMark)

			}

		}
	}
}
