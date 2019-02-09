package infrastructure

import (
	"github.com/gin-gonic/gin"
	"lottery/controller"
)

type Launch struct {
}

func routing(engine *gin.RouterGroup) {
	controller.RegisterLotteryController(engine)
}

func server_launch() {
	var config Configure
	config = LoadConfig()
	r := gin.Default()
	api := r.Group("/api")
	routing(api)
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	var bind = ":8088"
	if config.Binding != "" {
		bind = config.Binding
	}

	r.Run(bind)
}

func (b *Launch) Help() string {
	return "Launch Server"
}

func (b *Launch) Run(args []string) int {
	server_launch()
	return 0
}

func (b *Launch) Synopsis() string {
	return "Launch Server"
}
