package router

import (
	"github.com/bugoutianzhen123/TruthOrDare/handler"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func InitEngine(h handler.Handler) *gin.Engine {
	g := gin.Default()

	pprof.Register(g)

	user := g.Group("/user")
	{
		user.POST("/register", h.CreateUser)
		user.POST("/login", h.Login)
	}

	//friend := g.Group("/friend")
	//{}

	//group := g.Group("/group")
	//{
	//	group.GET("/Chat", h.HandleWebSocket)
	//	group.POST("/create", h.CreateGroup)
	//	group.GET("/ws", h.HandleWebSocket)
	//}

	game := g.Group("/game")
	{
		//game.GET("/ws")
		game.GET("/start", h.StartGame)
		game.POST("/createCard", h.CreateCard)
		game.DELETE("/deleteCard", h.RemoveCard)
		game.POST("/batchCreateCards", h.BatchCreateCards)
		game.POST("/saveGameHistory", h.SaveGameHistory)
		game.GET("/allGameHistories", h.GetAllGameHistories)
		game.GET("/userGameHistories", h.GetGameHistoriesByUserID)
	}

	return g
}
