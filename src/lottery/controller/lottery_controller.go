package controller

import (
	"github.com/gin-gonic/gin"
)

// 予約
//
// 実際に抽選を行うための準備を行う
// このタイミングでuuidを発行して、このuuidで抽選を実施する
//
// {
//   "uuid" : "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
// }
//
func LotteryReserve(c *gin.Context) {
	c.JSON(200, gin.H{
		"uuid": "LotteryReserve",
	})
}

// 抽選
// 抽選を実施する
// このタイミングで、指定したuuidで抽選を実施する
func LotteryPick(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "LotteryPick",
	})
}

// 抽選
// 抽選を実施する
// このタイミングで抽選結果とuuidを返却する
func LotteryOrder(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "LotteryOrder",
	})
}

func RegisterLotteryController(engine *gin.RouterGroup) {
	var group = engine.Group("/v1")
	group.GET("/lottery/ready", LotteryReserve)
	group.PUT("/lottery/pick", LotteryPick)
	group.POST("/lottery/order", LotteryOrder)
}
