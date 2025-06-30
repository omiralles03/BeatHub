package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"BeatHub-Backend/internal/api"
	"BeatHub-Backend/internal/schemas"
	"github.com/labstack/echo/v4"
)

// Getting more control over errors with it
func handleError(c echo.Context, err error, message string, statusCode int) error {
	c.Logger().Errorf("%s: %v", message, err)
	return c.JSON(statusCode, map[string]string{"error": message, "statusCode": strconv.Itoa(statusCode)})
}

// HandleGetBeatmap fetches a single beatmap by ID.
// GET /api/beatmaps/:beatmapId
func HandleGetBeatmap(c echo.Context) error {
	beatmapID := c.Param("beatmapId")
	if beatmapID == "" {
		return handleError(c, nil, "No Beatmap ID provided", http.StatusBadRequest)
	}

	beatmap, err := api.GetBeatmap(beatmapID)
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			return handleError(c, err, fmt.Sprintf("Beatmap %s not found", beatmapID), http.StatusNotFound)
		}
		return handleError(c, err, "Failed to retrieve beatmap data", http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, beatmap)
}

// HandleGetBeatmaps fetches a list of beatmaps by their IDs.
// GET /api/beatmaps?ids=1,2,3,4
func HandleGetBeatmaps(c echo.Context) error {
	idsParam := c.QueryParam("ids")
	if idsParam == "" {
		return handleError(c, nil, "No Beamap IDs provided", http.StatusBadRequest)
	}

	// Use coma separator for ids
	idStrings := strings.Split(idsParam, ",")

	var beatmapIDs []int
	for _, idStr := range idStrings {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return handleError(c, err, fmt.Sprintf("Invalid beatmap ID '%s'", idStr), http.StatusBadRequest)
		}
		beatmapIDs = append(beatmapIDs, id)
	}

	beatmaps, err := api.GetBeatmaps(beatmapIDs)
	if err != nil {
		return handleError(c, err, "Failed to retrieve beatmap list", http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, beatmaps)
}

// HandleLookupBeatmap fetches a beatmap by ID, checksum, or filename.
// GET /api/beatmaps/lookup?id=... OR checksum=... OR filename=...
func HandleLookupBeatmap(c echo.Context) error {
	id := c.QueryParam("id")
	checksum := c.QueryParam("checksum")
	filename := c.QueryParam("filename")

	if id == "" && checksum == "" && filename == "" {
		return handleError(c, nil, "No id/checksum/filename provided", http.StatusBadRequest)
	}

	// Parameters check just in case
	paramCount := 0
	if id != "" {
		paramCount++
	}
	if checksum != "" {
		paramCount++
	}
	if filename != "" {
		paramCount++
	}

	if paramCount > 1 {
		return handleError(c, nil, "Only one parameter can be provided for beatmap lookup", http.StatusBadRequest)
	}

	beatmap, err := api.LookupBeatmap(id, checksum, filename)
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			return handleError(c, err, "Beatmap not found with provided criteria", http.StatusNotFound)
		}
		return handleError(c, err, "Failed to lookup beatmap", http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, beatmap)
}

// HandleGetUserBeatmapScore fetches a user's score on a beatmap.
// GET /api/beatmaps/:beatmapId/scores/users/:userId
func HandleGetUserBeatmapScore(c echo.Context) error {
	beatmapID := c.Param("beatmapId")
	userID := c.Param("userId")

	if beatmapID == "" || userID == "" {
		return handleError(c, nil, "Beatmap ID and/or User ID missing", http.StatusBadRequest)
	}

	modeStr := c.QueryParam("mode")
	modsStr := c.QueryParam("mods")
	legacyOnlyStr := c.QueryParam("legacy_only")

	var mode schemas.Mode
	if modeStr != "" {
		mode = schemas.Mode(modeStr)
		if mode != schemas.ModeStandard && mode != schemas.ModeTaiko && mode != schemas.ModeFruits && mode != schemas.ModeMania {
			return handleError(c, nil, "Invalid mode provided", http.StatusBadRequest)
		}
	}

	var mods []string
	if modsStr != "" {
		mods = strings.Split(modsStr, ",")
	}

	legacyOnly := false
	if legacyOnlyStr == "1" {
		legacyOnly = true
	}

	score, err := api.GetUserBeatmapScore(beatmapID, userID, mode, mods, legacyOnly)
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			return handleError(c, err, "User was not found or no scores on the map", http.StatusNotFound)
		}
		return handleError(c, err, "Failed to retrieve user beatmap score", http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, score)
}

// HandleGetBeatmapScores fetches top scores for a beatmap.
// GET /api/beatmaps/:beatmapId/scores
func HandleGetBeatmapScores(c echo.Context) error {
	beatmapID := c.Param("beatmapId")
	if beatmapID == "" {
		return handleError(c, nil, "No Beatmap ID provided", http.StatusBadRequest)
	}

	modeStr := c.QueryParam("mode")
	modsStr := c.QueryParam("mods")
	typeStr := c.QueryParam("type")
	legacyOnlyStr := c.QueryParam("legacy_only")

	var mode schemas.Mode
	if modeStr != "" {
		mode = schemas.Mode(modeStr)
		if mode != schemas.ModeStandard && mode != schemas.ModeTaiko && mode != schemas.ModeFruits && mode != schemas.ModeMania {
			return handleError(c, nil, "Invalid mode provided", http.StatusBadRequest)
		}
	}

	var mods []string
	if modsStr != "" {
		mods = strings.Split(modsStr, ",")
	}

	legacyOnly := false
	if legacyOnlyStr == "1" {
		legacyOnly = true
	}

	scores, err := api.GetBeatmapScores(beatmapID, mode, mods, typeStr, legacyOnly)
	if err != nil {
		return handleError(c, err, "Failed to retrieve beatmap scores", http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, scores)
}

// HandleGetBeatmapset fetches a single beatmapset by ID.
// GET /api/beatmapsets/:beatmapsetId
func HandleGetBeatmapset(c echo.Context) error {
	beatmapsetID := c.Param("beatmapsetId")
	if beatmapsetID == "" {
		return handleError(c, nil, "No Beatmapset ID provided", http.StatusBadRequest)
	}

	beatmapset, err := api.GetBeatmapset(beatmapsetID)
	if err != nil {
		if strings.Contains(err.Error(), "status 404") {
			return handleError(c, err, fmt.Sprintf("Beatmapset %s not found", beatmapsetID), http.StatusNotFound)
		}
		return handleError(c, err, "Failed to retrieve beatmapset data", http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, beatmapset)
}

// HandleSearchBeatmapsets searches for beatmapsets.
// GET /api/beatmapsets/search?q=songtitle&status=ranked&mode=osu&limit=25&offset=0
func HandleSearchBeatmapsets(c echo.Context) error {
	query := c.QueryParam("q")
	statusStr := c.QueryParam("status")
	modeStr := c.QueryParam("mode")
	limitStr := c.QueryParam("limit")
	offsetStr := c.QueryParam("offset")

	if query == "" {
		return handleError(c, nil, "No search query 'q' provided", http.StatusBadRequest)
	}

	var status schemas.BeatmapsetStatus
	if statusStr != "" {
		switch schemas.BeatmapsetStatus(statusStr) {
		case schemas.BeatmapsetStatusRanked, schemas.BeatmapsetStatusApproved, schemas.BeatmapsetStatusQualified,
			schemas.BeatmapsetStatusLoved, schemas.BeatmapsetStatusPending, schemas.BeatmapsetStatusWIP,
			schemas.BeatmapsetStatusGraveyard:

			status = schemas.BeatmapsetStatus(statusStr)
		default:
			return handleError(c, nil, "Invalid beatmapset status provided", http.StatusBadRequest)
		}
	}

	var mode schemas.Mode
	if modeStr != "" {
		switch schemas.Mode(modeStr) {
		case schemas.ModeStandard, schemas.ModeTaiko, schemas.ModeFruits, schemas.ModeMania:

			mode = schemas.Mode(modeStr)
		default:
			return handleError(c, nil, "Invalid game mode provided", http.StatusBadRequest)
		}
	}

	limit := 50
	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err != nil || parsedLimit <= 0 {
			return handleError(c, err, "Invalid limit parameter, must be a positive integer", http.StatusBadRequest)
		}
		parsedLimit = min(50, parsedLimit)
	}

	offset := 0
	if offsetStr != "" {
		parsedOffset, err := strconv.Atoi(offsetStr)
		if err != nil || parsedOffset < 0 {
			return handleError(c, err, "Invalid offset parameter, must be a non-negative integer", http.StatusBadRequest)
		}
		offset = parsedOffset
	}

	beatmapsets, err := api.SearchBeatmapsets(query, status, mode, limit, offset)
	if err != nil {
		return handleError(c, err, "Failed to search beatmapsets", http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, beatmapsets)
}
