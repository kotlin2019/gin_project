package routers

import "github.com/zhanghe06/gin_project/controllers"

func RegisterDailySentence() {
	Ver.GET("/daily_sentences", controllers.ListsDailySentenceHandler)
	Ver.GET("/daily_sentence/:id", controllers.GetDailySentenceHandler)
	Ver.POST("/daily_sentence", controllers.CreateDailySentenceHandler)
	Ver.POST("/daily_sentence/transaction", controllers.UpdateDailySentenceTransactionHandler)
	Ver.PUT("/daily_sentence/:id", controllers.UpdateDailySentenceHandler)
	Ver.PUT("/daily_sentence/:id/title", controllers.ReTitleDailySentenceHandler)
	Ver.DELETE("/daily_sentence/:id", controllers.DeleteDailySentenceHandler)
	Ver.PATCH("/daily_sentence/:id/score", controllers.ScoreDailySentenceHandler)
}
