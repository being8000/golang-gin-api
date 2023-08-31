package handlers

import (
	"fmt"
	"log"
	"net/http"
	"zehan/gin/handlers/vo"
	"zehan/gin/utils"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	Kit *utils.Kit
}

func (handler *UserHandler) Login(c *gin.Context) {
	username, password := c.PostForm("mobile"), c.PostForm("password")
	// If user has logged in, force him to log out firstly
	// for iter := utils.GlobalCache.Iterator(); iter.SetNext(); {
	// 	info, err := iter.Value()
	// 	if err != nil {
	// 		continue
	// 	}
	// 	if string(info.Value()) == username {
	// 		utils.GlobalCache.Delete(info.Key())
	// 		log.Printf("forced %s to log out\n", username)
	// 		break
	// 	}
	// }

	// Apparently we don't do this in real world :)
	if username == "alice" && password == "111" {
		log.Println(fmt.Sprintf("%s has logged in.", username))
	} else if username == "bob" && password == "123" {
		log.Println(fmt.Sprintf("%s has logged in.", username))
	} else {
		c.JSON(200, utils.RestResponse{Message: "no such account"})
		return
	}

	// Generate random session id
	u, err := uuid.NewRandom()
	if err != nil {
		log.Println(fmt.Errorf("failed to generate UUID: %w", err))
	}
	sessionId := fmt.Sprintf("%s-%s", u.String(), username)
	// Store current subject in cache
	err = utils.GlobalCache.Set(sessionId, []byte(username))
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to store current subject in cache: %w", err))
		return
	}
	// Send session id back to client in cookie
	c.SetCookie("current_subject", sessionId, 30*60, "/api", "", false, true)
	c.JSON(200, utils.RestResponse{Code: 1, Message: username + " logged in successfully"})
}

func (handler *UserHandler) AddPolicy(c *gin.Context) {

	// kit := handler.Kit
	// adapter := kit.Casbin.Adapter
	// e, _ := casbin.NewEnforcer("./app/conf/rbac_model.conf", adapter)
	e, err := casbin.NewEnforcer("./app/conf/rbac_pattern.conf", "./app/casbin/user_policy.csv")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	policy := &vo.Policy{}
	if err := c.ShouldBind(policy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	// err = e.GetAdapter().(*fileadapter.Adapter).AddPolicy("jack", "data2", "write")
	err = e.GetAdapter().(*gormadapter.Adapter).Transaction(e, func(e casbin.IEnforcer) error {
		_, err := e.AddPolicy(policy)
		if err != nil {
			return err
		}
		// _, err = e.AddPolicy("jack", "data2", "write")
		// if err != nil {
		// 	return err
		// }
		return nil
	})
	if err != nil {
		// handle if transaction failed
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
}
