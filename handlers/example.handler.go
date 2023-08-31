package handlers

import (
	"log"
	"net/http"
	"time"
	"zehan/gin/utils"

	"github.com/gin-gonic/gin"
)

type ExampleHandler struct {
	Kit *utils.Kit
}
type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}

type StructC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_c"`
}

type StructD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func (e *ExampleHandler) GetDataB(c *gin.Context) {
	var b StructB
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}

func (e *ExampleHandler) GetDataC(c *gin.Context) {
	var b StructC
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStructPointer,
		"c": b.FieldC,
	})
}

func (e *ExampleHandler) GetDataD(c *gin.Context) {
	var b StructD
	c.Bind(&b)
	c.JSON(200, gin.H{
		"x": b.NestedAnonyStruct,
		"d": b.FieldD,
	})
}

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func (e *ExampleHandler) BindData(c *gin.Context) {
	var person Person
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	if err := c.ShouldBind(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	log.Println(person.Name)
	log.Println(person.Address)
	log.Println(person.Birthday)
	c.String(200, "Success")
}

type PersonUri struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func (e *ExampleHandler) BindUri(c *gin.Context) {
	var person PersonUri
	if err := c.ShouldBindUri(&person); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
}
