package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wenealves10/url-shortener-go/controllers"
	"github.com/wenealves10/url-shortener-go/mongo"
	"github.com/wenealves10/url-shortener-go/server"
	"golang.org/x/sync/errgroup"
)

var (
	address         = flag.String("address", ":8080", "Address to listen to")
	mongoConfigFile = flag.String("mongo-conf", "mongo/config.json", "MongoDB config file")
	paramTimeout    = flag.Int("timeout", 10, "Timeout in seconds to connect to MongoDB")
)

var timeout = time.Duration(*paramTimeout) * time.Second

func MainRouter() http.Handler {
	r := gin.New()
	r.Use(gin.Recovery(), server.MiddlewareReqHandler())
	v1 := r.Group("/api/v1")
	shortUrl := v1.Group("/short-url")
	{
		shortUrl.GET("/:hash", controllers.GetUrlByHash)
		shortUrl.POST("/", controllers.CreateShortUrl)
	}
	return r
}

func initializeMongoConnection() {
	mongoConf, err := mongo.ParseConfig(*mongoConfigFile)
	if err != nil {
		log.Fatalf("unable to parse mongo config file %s", err)
	}
	mongo.ConnectDB(mongoConf.Uri, timeout)
	controllers.DbClient = mongo.GetMongoDbConnector(mongoConf.Db, mongoConf.Collection)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()
	timeout = time.Duration(*paramTimeout) * time.Second
	controllers.Timeout = timeout
	initializeMongoConnection()
	var g errgroup.Group
	mainServer := &http.Server{
		Addr:         *address,
		Handler:      MainRouter(),
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}
	g.Go(func() error {
		return mainServer.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatalf("unable to start server %s", err)
	}
}
