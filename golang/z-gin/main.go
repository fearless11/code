package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
	// json中打tag的校验有用到
	"gopkg.in/go-playground/validator.v8"
)

func main() {
	// quickStart()
	basicAuth()
	// httpParameters()
	// httpQuerystringParameters()
	// httpEncoded()
	// httpQueryPostForm()
	// httpMapQuerystringPostform()
	// httpUploadFile()
	// httpUploadMultiplefiles()
	// httpGroupingRoutes()
	// httpMiddleware()
	// ginLogFile()
	// ginCustomLogFormat()
	// ginModelBingValidation()
	// ginCustomValidators()
	// ginBindQueryString()
	// ginJSONP()
	// ginStatic()
	// ginDataForm()
	// ginHTMLRendering()
	// testCNext()
}

func quickStart() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

// Authorized group (uses gin.BasicAuth() middleware)
// Same than:
// authorized := r.Group("/")
// authorized.Use(gin.BasicAuth(gin.Credentials{
//	  "foo":  "bar",
//	  "manu": "123",
//}))
func basicAuth() {
	var db = make(map[string]string)

	r := gin.Default()
	// 本质：对比http头部header中的Authorization字段的值,看与设置的值是否匹配
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar",
		"manu": "123",
	}))

	authorized.POST("admin", func(c *gin.Context) {
		// 在预先的认证处理后,没有Get到值会panic
		user := c.MustGet(gin.AuthUserKey).(string)
		var json struct {
			Value string `json:"value" binding:"required"`
		}
		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{
				"user":   user,
				"value":  json.Value,
				"status": "ok",
			})
		}
	})

	r.Run(":8080")
}

// Using GET,POST,PUT,PATCH,DELETE and OPTIONS
/*
func httpRequestMethod() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	router.PATCH("/somePatch", patching)
	router.HEAD("/someHead", head)
	router.OPTIONS("/someOptions", options)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}
*/

// Parameters in path
func httpParameters() {
	router := gin.Default()

	// 精确路由
	// handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 模糊路由
	// this one will match /user/john/ and also /user/john/send
	// If no other router match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.Run()
}

// Querystring parameters
// what is querystring?
// http://localhost/welcome?firstname=vera&lastname=li
// firstname=vera&lastname=li is querystring and must be http Get method
func httpQuerystringParameters() {
	router := gin.Default()

	router.GET("/welcome", func(c *gin.Context) {
		// 设置默认值
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for  c.Request.URL.Query().Get("lastname")
		c.String(http.StatusOK, "hello %s %s", firstname, lastname)
	})

	router.Run()
}

// Multipart/Urlencoded Form
// post + http请求头部
// Content-Type:multipart/form-data; boundary=--------------------------725072089786234412149248
func httpEncoded() {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anoymous")
		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	router.Run()
}

// query + post form
// localhost/post?id=1&page=100
// Content-Type:application/x-www-form-urlencoded
// message:something
// name:vera
func httpQueryPostForm() {
	router := gin.Default()
	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		c.JSON(http.StatusOK, gin.H{
			"stat":    "posted",
			"message": message,
			"name":    name,
			"id":      id,
			"page":    page,
		})
	})

	router.Run(":80")
}

// Map as querystring or postform parameters
func httpMapQuerystringPostform() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})

	router.Run(":80")
}

// Upload files
//  curl -X POST http://localhost/upload -F "file=@./a.txt" -H "Content-Type:multipart/form-data"
func httpUploadFile() {
	router := gin.Default()

	// set a lower memory limit for multipart forms(default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 * 2^20 = 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, fmt.Sprintf("%s's uploaded!", file.Filename))
	})

	router.Run(":80")
}

// Multiple files
//  curl -X POST http://localhost/upload -F "file=@./a.txt" -F "file=@./b.txt" -H "Content-Type:multipart/form-data"
func httpUploadMultiplefiles() {
	router := gin.Default()

	router.MaxMultipartMemory = 8 << 20
	router.POST("/upload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["file"]

		for _, file := range files {
			c.SaveUploadedFile(file, file.Filename)
		}

		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded :)", len(files)))
	})

	router.Run(":80")
}

// Grouping routes
func httpGroupingRoutes() {
	router := gin.Default()

	v1 := router.Group("v1")
	{
		v1.GET("login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"login": "v1 success",
			})
		})
	}

	v2 := router.Group("v2")
	{
		v2.GET("login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"login": "v2 success",
			})
		})
	}

	router.Run(":80")
}

// Using middleware
// Blank Gin without middleware by default
// gin.Default()  // default with the logger and recovery middleware already attached
// middleware 中间件的理解是在做某些操作前进行的一系列准备操作
var db = make(map[string]string)

func httpMiddleware() {
	db["vera"] = "vera"
	r := gin.New()

	// logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one
	r.Use(gin.Recovery())

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// per group middleware! in this case we use the custom created
	// gin.BasicAuth is middleware just in the 'authorized' group
	authorized := r.Group("/")
	authorized.Use(gin.BasicAuth(gin.Accounts{
		"hi": "123", // user:hi  password:123
	}))

	authorized.GET("/admin", func(c *gin.Context) {
		c.String(http.StatusOK, "hi")
	})

	r.Run(":80")
}

// How to write log file
func ginLogFile() {
	// Diable console color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	f, _ := os.Create("gin.log")

	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":80")
}

// Custom Log Format
func ginCustomLogFormat() {
	// Default With the Logger and Recovery middleware already attached
	// r := gin.Default()

	// 没有使用中间件
	router := gin.New()

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s \n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Path,
		)
	}))

	router.Use(gin.Recovery())
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":80")
}

//Login Model binding and validation
// purpose : to bind a request body into a type, use model binding. We currently support binding
// of JSON, XML, YAML and standard form vaules(foo=bar&boo=baz)
// two sets of methods for binding:
// 		Methods:  BindJSON 、 BindXML
//		Behavior:  error方法
// 通过绑定实现请求与对象的转换
// curl -v -X POST \
//   http://localhost:80/loginJSON \
//   -H 'content-type: application/json' \
//   -d '{ "user": "manu" }'
type Login struct {
	User     string `form:"user" json:"user"  xml:"user" binding:"required" `
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func ginModelBingValidation() {
	router := gin.Default()

	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "manu" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	router.Run(":80")
}

// Booking Custom Validators
// Booking contains binded and validated data.
// 自定义校验, 没懂怎么实现的？？？
type Booking struct {
	CheckIn  time.Time `form:"check_in"  binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

func bookableDate(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value,
	fileType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func ginCustomValidators() {
	route := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}
	route.GET("/bookable", getBookable)
	route.Run(":80")
}

// Only Bind Query String
func ginBindQueryString() {
	route := gin.Default()
	route.Any("/testing", startPage)
	route.Run(":80")
}

func startPage(c *gin.Context) {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		c.String(http.StatusOK, person.Name+person.Address)
	}
}

// Person for test
type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

// JSONP
// https://baike.baidu.com/item/jsonp/493658
// JSONP是JSON的一种使用模式,解决主流浏览器跨域数据访问的问题。
// 同源策略：使得server1.com的网页无法与不是server1.com的服务器沟通
// JSONP：使用HTML的script元素的开放策略,获取其他源的JSON资料
// JSONP抓到的资料并不是JSON,而是任意的JavaScript,用JavaScript直译器执行而不是JSON解析器解析
// Using JSONP to request data from a server in a different domain.
// Add callback to response body if the query parameter callback exists.
// 应用:
//  如何在服务器段实现对JSONP的支持?
//   本质是服务器返回的非JSON,而是把JSON数据转换成想要的script tags的形式
//   callback仅仅是JSONP的简单实现,可以根据具体需要实现更复杂的功能,比如在客户端动态集成更多变量数据未完成的分页功能
//  curl http://127.0.0.1:8080/JSONP?callback=aaa   response: aaa({"foo":"bar"})
func ginJSONP() {
	r := gin.Default()
	r.GET("/JSONP", func(c *gin.Context) {
		data := gin.H{
			"foo": "bar",
		}
		c.JSONP(http.StatusOK, data)
	})

	r.Run(":8080")
}

// Serving static files
// 作为文件服务器StaticFS
func ginStatic() {
	r := gin.Default()
	r.Static("/assets", "./www")
	r.StaticFile("/favicon.ico", "./favicon.ico")
	r.StaticFS("/files", http.Dir("../quickstart"))
	r.Run(":80")
}

// Serving data form reader
// 例子：提供图片下载
func ginDataForm() {
	r := gin.Default()
	r.GET("/dataform", func(c *gin.Context) {
		response, err := http.Get("https://activity.zhenai.com/admin/img/home_bride.d6e809dc.png")
		if err != nil {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")
		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}
		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})
	r.Run(":80")
}

// HTML rendering
func ginHTMLRendering() {
	r := gin.Default()
	// Glob 正则
	r.LoadHTMLGlob("templates/*")
	r.GET("/index", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "welcome website.",
		})
	})
	r.Run(":80")
}

// test c.Next()
// https://segmentfault.com/q/1010000020256918
// c.Next()决定了执行下一个中间件的时机
func testCNext() {
	r := gin.New()

	mid1 := func(c *gin.Context) {
		start := time.Now()
		fmt.Println("middleware1 start")
		// 注释 or 不注释
		c.Next()
		// 注释后要执行完mid1的代码才执行下一个中间件
		// 不注释则将mid1挂起执行完mid2才执行后面的输出
		fmt.Println(time.Since(start))
		fmt.Println("middleware1 ending")
	}

	mid2 := func(c *gin.Context) {
		fmt.Println("middleware2 start")
		c.Next()
		fmt.Println("middleware2 ending")
	}

	r.Use(mid1, mid2)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hi")
	})
	r.Run()
}
