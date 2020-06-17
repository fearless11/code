// 展示如何编写基础单元测试
// 一组参数和结果来测试
package unit

import (
	"net/http"
	"testing"
)

// 运行查看详细信息
//  go test -v

const checkMark = "\u2713"
const ballotX = "\u2717"

// TestDownload 确认http包的Get函数可以下载内容
func TestDownload(t *testing.T) {
	url := "http://www.npr.org/rss/rss.php?id=1001"
	statusCode := 200

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", url, statusCode)
		resp, err := http.Get(url)
		{
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.", ballotX, err)
			}
			t.Log("\t\tShould be able to make the Get call.", checkMark)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%d\" status. %v", statusCode, checkMark)
			} else {
				// 报告测试失败但不停止当前测试函数的执行
				t.Errorf("\t\tShould receive a \"%d\" status. %v %v", statusCode, ballotX, resp.StatusCode)
			}
		}

	}
}
