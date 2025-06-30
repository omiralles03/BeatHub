package schemas

type RankingCategory string

const (
	RankingCategoryScore RankingCategory = "score"
	RankingCategoryPP    RankingCategory = "performance"
)

type RankingFilter string

const (
	RankingFilterAll     RankingFilter = "all"
	RankingFilterFriends RankingFilter = "friends"
)

type RankingOpts struct {
	Cursor  string
	Filter  RankingFilter
	Country string
	Variant string
}

type Ranking struct {
	Count300       int         `json:"count_300"`
	Count100       int         `json:"count_100"`
	Count50        int         `json:"count_50"`
	CountMiss      int         `json:"count_miss"`
	Level          Level       `json:"level"`
	GlobalRank     int         `json:"global_rank"`
	GlobalRankExp  any         `json:"global_rank_exp"`
	PP             int         `json:"pp"`
	PPExp          int         `json:"pp_exp"`
	RankedScore    int         `json:"ranked_score"`
	HitAccuracy    int         `json:"hit_accuracy"`
	PlayCount      int         `json:"play_count"`
	PlayTime       int         `json:"play_time"`
	TotalScore     int         `json:"total_score"`
	TotalHits      int         `json:"total_hits"`
	MaxCombo       int         `json:"maximum_combo"`
	ReplaysWatched int         `json:"replays_watched_by_others"`
	IsRanked       bool        `json:"is_ranked"`
	GradeCounts    GradeCounts `json:"grade_counts"`
	RankChange     int         `json:"rank_change_since_30_days"`
	User           UserBasic   `json:"user"`
}

type RankingsResponse struct {
	Rankings []Ranking `json:"ranking"`
	Total    int       `json:"total"`
	Cursor   struct {
		Page string `json:"page"`
	} `json:"country"`
}
