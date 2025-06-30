package api

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"BeatHub-Backend/internal/schemas"
)

//////////////////// BEATMAPS \\\\\\\\\\\\\\\\\\\\\\\\\\\\\

// Return beatmap
func LookupBeatmap(id, checksum, filename string) (schemas.Beatmap, error) {
	query := make(map[string]string)

	if id != "" {
		query["id"] = id
	} else if checksum != "" {
		query["checksum"] = checksum
	} else if filename != "" {
		query["filename"] = filename
	} else {
		return schemas.Beatmap{}, fmt.Errorf("At least one field must be provided")
	}

	return Fetch[schemas.Beatmap](FetchOptions{
		Url:    "https://osu.ppy.sh/api/v2/beatmaps/lookup",
		Method: http.MethodGet,
		Token:  ApplicationToken,
		Query:  query,
	})
}

// Return a User's score on a Beatmap
// GET /beatmaps/{beatmap}/scores/users/{user}
func GetUserBeatmapScore(beatmapID, userID string, mode schemas.Mode, mods []string, legacyOnly bool) (schemas.Score, error) {
	query := make(map[string]string)

	if mode != "" {
		query["mode"] = string(mode)
	}

	if len(mods) > 0 {
		query["mods"] = strings.Join(mods, ",")
	}

	if legacyOnly {
		query["legacy_only"] = "1"
	} else {
		query["legacy_only"] = "0"
	}

	type UserScoreResponse struct {
		Score schemas.Score `json:"score"`
	}

	res, err := Fetch[UserScoreResponse](FetchOptions{
		Url:    fmt.Sprintf("https://osu.ppy.sh/api/v2/beatmaps/%s/scores/users/%s", beatmapID, userID),
		Method: http.MethodGet,
		Token:  ApplicationToken,
		Query:  query,
	})
	if err != nil {
		return schemas.Score{}, err
	}
	return res.Score, nil
}

// Returns the top scores for a beatmap.
// Depending on user preferences, this may only show legacy scores.
// GET /beatmaps/{beatmap}/scores
func GetBeatmapScores(beatmapID string, mode schemas.Mode, mods []string, scoreType string, legacyOnly bool) ([]schemas.Score, error) {
	query := make(map[string]string)

	if mode != "" {
		query["mode"] = string(mode)
	}

	if len(mods) > 0 {
		query["mods"] = strings.Join(mods, ",")
	}

	if legacyOnly {
		query["legacy_only"] = "1"
	} else {
		query["legacy_only"] = "0"
	}

	type BeatmapScoresResponse struct {
		Scores []schemas.Score `json:"scores"`
	}

	res, err := Fetch[BeatmapScoresResponse](FetchOptions{
		Url:    fmt.Sprintf("https://osu.ppy.sh/api/v2/beatmaps/%s/scores", beatmapID),
		Method: http.MethodGet,
		Token:  ApplicationToken,
		Query:  query,
	})
	if err != nil {
		return nil, err
	}
	return res.Scores, nil
}

// Returns a list of beatmaps.
// GET /beatmaps?ids[]=...
func GetBeatmaps(beatmapsID []int) ([]schemas.Beatmap, error) {
	if len(beatmapsID) == 0 {
		return nil, fmt.Errorf("No Beatmap IDs provided")
	}

	if len(beatmapsID) > 50 {
		return nil, fmt.Errorf("Cannot request more than 50 beatmaps at once")
	}

	queryParams := url.Values{}
	for _, id := range beatmapsID {
		queryParams.Add("ids[]", strconv.Itoa(id))
	}

	type BeatmapsResponse struct {
		Beatmaps []schemas.Beatmap `json:"beatmaps"`
	}

	res, err := Fetch[BeatmapsResponse](FetchOptions{
		Url:    "https://osu.ppy.sh/api/v2/beatmaps?" + queryParams.Encode(),
		Method: http.MethodGet,
		Token:  ApplicationToken,
	})

	if err != nil {
		return nil, err
	}
	return res.Beatmaps, nil
}

// Gets beatmap data for the specified beatmap ID.
// GET /beatmaps/{beatmap}
func GetBeatmap(beatmapID string) (schemas.Beatmap, error) {
	return Fetch[schemas.Beatmap](FetchOptions{
		Url:    fmt.Sprintf("https://osu.ppy.sh/api/v2/beatmaps/%s", beatmapID),
		Method: http.MethodGet,
		Token:  ApplicationToken,
	})
}

////////////////// BEATMAPSETS \\\\\\\\\\\\\\\\\\\\\\\\\\

// Return a Beatmapset for the specified ID
// GET /beatmapsets/{beatmapset}
func GetBeatmapset(beatmapsetID string) (schemas.Beatmapset, error) {
	return Fetch[schemas.Beatmapset](FetchOptions{
		Url:    fmt.Sprintf("https://osu.ppy.sh/api/v2/beatmapsets/%s", beatmapsetID),
		Method: http.MethodGet,
		Token:  ApplicationToken,
	})
}

// Returns a list of Beatmapsets
// GET /beatmapsets/search
func SearchBeatmapsets(queryStr string, status schemas.BeatmapsetStatus, mode schemas.Mode, limit, offset int) ([]schemas.Beatmapset, error) {
	query := map[string]string{
		"q": queryStr,
	}

	if status != "" {
		query["s"] = string(status)
	}

	if mode != "" {
		query["m"] = string(mode)
	}

	if limit > 0 {
		query["limit"] = strconv.Itoa(limit)
	}

	if offset > 0 {
		query["offset"] = strconv.Itoa(offset)
	}

	type SearchResponse struct {
		Beatmapsets []schemas.Beatmapset `json:"beatmapsets"`
	}

	res, err := Fetch[SearchResponse](FetchOptions{
		Url:    "https://osu.ppy.sh/api/v2/beatmapsets/search",
		Method: http.MethodGet,
		Token:  ApplicationToken,
		Query:  query,
	})

	if err != nil {
		return nil, err
	}

	return res.Beatmapsets, nil
}
