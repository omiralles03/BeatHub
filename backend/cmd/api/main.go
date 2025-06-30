package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"BeatHub-Backend/internal/api"
	"BeatHub-Backend/internal/config"
	"BeatHub-Backend/internal/handlers"
)

func main() {

	appCfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load app config: %w", err)
	}

	PORT := appCfg.PORT
	fmt.Println("\nPORT: ", PORT)

	log.Println("\nAuthenticating osu! API...")
	clientAuth, err := api.GetClientAuth(appCfg.OSU_CLIENT_ID, appCfg.OSU_CLIENT_SECRET)
	if err != nil {
		log.Println("\nCould not authenticate: %w", err)
	}
	api.ApplicationToken = clientAuth.AccessToken
	log.Println("\nSuccessfully authenticated")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// 10 requests per second
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Main Page")
	})

	apiGroup := e.Group("/api")

	// Beatmap Routes
	apiGroup.GET("/beatmaps/:beatmapId", handlers.HandleGetBeatmap)
	apiGroup.GET("/beatmaps", handlers.HandleGetBeatmaps)
	apiGroup.GET("/beatmaps/lookup", handlers.HandleLookupBeatmap)
	apiGroup.GET("/beatmaps/:beatmapId/scores", handlers.HandleGetBeatmapScores)
	apiGroup.GET("/beatmaps/:beatmapId/scores/users/:userId", handlers.HandleGetUserBeatmapScore)

	// Beatmapsets Routes
	apiGroup.GET("/beatmapsets/:beatmapsetId", handlers.HandleGetBeatmapset)
	apiGroup.GET("/beatmapsets/search", handlers.HandleSearchBeatmapsets)

	log.Fatal(e.Start(":" + PORT))
}
