package middleware

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/icehardstone/fmanager/serror"
)

const (
	HeaderAuthorization = "Authorization"
	QueryTaken          = "token"
)

// 用户信息
type UserInfo struct {
	Name              *string `json:"name"`
	Sub               *string `json:"sub"`
	GiveName          *string `json:"given_name"`
	FamilyName        *string `json:"family_name"`
	WebSite           *string `json:"website"`
	PreferredUserName *string `json:"preferred_username"`
}

// OAuth2Middleware 验证OAuth 2.0 token的中间件
func OAuth2Middleware(c *gin.Context) {

	token := c.Request.Header.Get(HeaderAuthorization)
	if c.Request.URL.Query().Has(QueryTaken) {
		token = c.Request.URL.Query().Get(QueryTaken)
	}
	if len(token) == 0 {
		forbiden(c)
		return
	}
	// 验证token
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{
		Transport: tr,
	}

	req, err := http.NewRequest("GET", "https://43.134.115.113:5000/connect/userinfo", nil)
	if err != nil {
		log.Printf("new request error: %s", err)
		forbiden(c)
		return
	}

	req.Header.Add(HeaderAuthorization, token)
	resp, err := client.Do(req)

	if err != nil {
		forbiden(c)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)

		forbiden(c)
		return
	}

	var user UserInfo
	err = json.Unmarshal(body, &user)

	if err != nil || user.Name == nil {
		forbiden(c)
		return
	}

	log.Printf("response: %s", string(body))

	// 如果token有效，调用下一个处理程序
	c.Next()
}

func forbiden(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, &serror.APIError{
		ErrorCode:    401,
		ErrorMessage: "invalide token",
	})
	c.AbortWithStatus(http.StatusUnauthorized)
}
