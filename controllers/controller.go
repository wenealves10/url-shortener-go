package controllers

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/wenealves10/url-shortener-go/model"
)

var DbClient model.DbConnector
var Timeout time.Duration

func GetUrlByHash(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()
	hash := c.Param("hash")
	result, err := DbClient.FindOne(&ctx, hash)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("unable to find hash %s", hash))
		return
	}
	c.Redirect(http.StatusMovedPermanently, result.Url)
}

func CreateShortUrl(c *gin.Context) {
	var req model.ShortUrlReq
	var shortUrl model.ShortUrl
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	err := c.ShouldBindBodyWith(&req, binding.JSON)
	if err != nil {
		tempReqBody, _ := c.Get(gin.BodyBytesKey)
		c.JSON(http.StatusBadRequest, fmt.Sprintf("unable to parse body %s", string(tempReqBody.([]byte))))
		return
	}
	shortUrl = model.ShortUrl{
		Url:  req.Url,
		Hash: utilGenerateHash(req.Url),
	}
	err = DbClient.Insert(&ctx, &shortUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("unable to insert %s", req.Url))
		return
	}
	c.JSON(http.StatusOK, shortUrl)
}

func utilGenerateHash(url string) string {
	md5Instance := md5.New()
	md5Instance.Write([]byte(url))
	md5Hash := hex.EncodeToString(md5Instance.Sum(nil))
	return md5Hash
}
