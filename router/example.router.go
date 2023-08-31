package router

import (
	"fmt"
	"log"
	"net/http"
	"zehan/gin/handlers"
	"zehan/gin/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (r *Router) Router_Example(k *utils.Kit) {
	handler := &handlers.ExampleHandler{
		Kit: k,
	}
	k.App.MaxMultipartMemory = 8 << 20
	app := k.App.Group("example")
	{
		app.GET("/someJSON", func(c *gin.Context) {
			data := map[string]interface{}{
				"lang": "GO语言",
				"tag":  "<br>",
			}

			// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
			c.AsciiJSON(http.StatusOK, data)
		})

		app.GET("/getb", handler.GetDataB)
		app.GET("/getc", handler.GetDataC)
		app.GET("/getd", handler.GetDataD)
		app.GET("/bind-data", handler.BindData)
		app.GET("/uri/:name/:id", handler.BindUri)

		app.GET("/moreJSON", func(c *gin.Context) {
			// You also can use a struct
			var msg struct {
				Name    string `json:"user"`
				Message string
				Number  int
			}
			msg.Name = "Lena"
			msg.Message = "hey"
			msg.Number = 123
			// Note that msg.Name becomes "user" in the JSON
			// Will output  :   {"user": "Lena", "Message": "hey", "Number": 123}
			c.JSON(http.StatusOK, msg)
		})

		app.GET("/someXML", func(c *gin.Context) {
			c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
		})

		app.GET("/someYAML", func(c *gin.Context) {
			c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
		})

		app.GET("/someProtoBuf", func(c *gin.Context) {
			reps := []int64{int64(1), int64(2)}
			label := "test"
			var test struct {
				Label string
				Reps  []int64
			}
			// The specific definition of protobuf is written in the testdata/protoexample file.
			test.Label = label
			test.Reps = reps
			// Note that data becomes binary data in the response
			// Will output protoexample.Test protobuf serialized data
			c.ProtoBuf(http.StatusOK, test)
		})
		app.GET("/someDataFromReader", func(c *gin.Context) {
			response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
			if err != nil || response.StatusCode != http.StatusOK {
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

		app.GET("/cookie", func(c *gin.Context) {

			cookie, err := c.Cookie("gin_cookie")

			if err != nil {
				cookie = "NotSet"
				c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
			}

			zap.S().Infof("Cookie value: %s \n", cookie)
		})
		app.POST("/upload", func(c *gin.Context) {
			// single file
			file, err := c.FormFile("file")
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err})
				return
			}
			log.Println(file.Filename)

			// Upload the file to specific dst.
			c.SaveUploadedFile(file, "./dist/"+file.Filename)

			c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
		})

		app.POST("/mupload", func(c *gin.Context) {
			// Multipart form
			form, err := c.MultipartForm()
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err})
				return
			}
			files := form.File["upload"]

			log.Println(files)
			for _, file := range files {
				log.Println(file.Filename)

				// 上传文件至指定目录
				c.SaveUploadedFile(file, "./dist/"+file.Filename)
			}
			c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
		})
	}

}
