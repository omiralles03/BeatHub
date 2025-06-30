package schemas

type ScoreCategory string

const (
	ScoreCategoryBest   ScoreCategory = "best"
	ScoreCategoryRecent ScoreCategory = "recent"
	ScoreCategoryFirst  ScoreCategory = "first"
	ScoreCategoryPinned ScoreCategory = "pinned"
)

type ScoreOpts struct {
	Mode   Mode
	Offset int
	Limit  int
}

type ScoreHitCounts struct {
	Geki int `json:"count_geki"`
	X300 int `json:"count_300"`
	Katu int `json:"count_katu"`
	X100 int `json:"count_100"`
	X50  int `json:"count_50"`
	Miss int `json:"count_miss"`
}

type Score struct {
	Id                    int                   `json:"id"`
	Accuracy              float64               `json:"accuracy"`
	BestId                int                   `json:"best_id"`
	CreatedAt             string                `json:"created_at"`
	MaxCombo              int                   `json:"max_combo"`
	Mode                  Mode                  `json:"mode"`
	ModeInt               ModeInt               `json:"mode_int"`
	Mods                  []string              `json:"mods"`
	Passed                bool                  `json:"passed"`
	Perfect               bool                  `json:"perfect"`
	PP                    float64               `json:"pp"`
	Rank                  string                `json:"rank"`
	Replay                bool                  `json:"replay"`
	Score                 int                   `json:"score"`
	HitCounts             ScoreHitCounts        `json:"statistics"`
	Type                  string                `json:"type"`
	UserId                int                   `json:"user_id"`
	CurrentUserAttributes CurrentUserAttributes `json:"current_user_attributes"`
	Beatmap               Beatmap               `json:"beatmap"`
	Beatmapset            Beatmapset            `json:"beatmapset"`
	RankGlobal            int                   `json:"rank_global"`
	User                  UserBasic             `json:"user"`
}

type CurrentUserAttributes struct {
	Pin Pin `json:"pin"`
}

type Pin struct {
	IsPinned  bool   `json:"is_pinned"`
	ScoreID   int    `json:"score_id"`
	ScoreType string `json:"score_type"`
}
