package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	Port string `envconfig:"PORT" default:"8088"`
}

func main() {
	config := loadConfig()

	r := gin.Default()

	r.GET("/labels", func(c *gin.Context) {
		url := c.Query("url")
		c.JSON(200, gin.H{
			"url": url,
			"labels": []gin.H{
				{
					"label":       "canoe",
					"probability": 0.3231,
				},
				{
					"label":       "canoe",
					"probability": 0.3231,
				},
			},
		})
	})

	err := r.Run("0.0.0.0:" + config.Port) // listen and serve
	if err != nil {
		log.Fatal(err)
	}
}

func loadConfig() *Config {
	var config Config
	err := envconfig.Process("analyzer", &config)
	if err != nil {
		log.Fatal(err)
	}
	return &config
}
