package schemas

import "time"

type Kudosu struct {
	Total     int `json:"total"`
	Available int `json:"available"`
}

type UserCountry struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Cover struct {
	Id        string `json:"id"`
	URL       string `json:"url"`
	CustomURL string `json:"custom_url"`
}

type Badge struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	AwardedAt   time.Time `json:"awarded_at"`
	URL         string    `json:"url"`
}

type Group struct {
	Id          int    `json:"id"`
	Identifier  string `json:"identifier"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type MonthlyPlaycount struct {
	StartDate string `json:"start_date"`
	Count     int    `json:"count"`
}

type Page struct {
	Html string `json:"html"`
	Raw  string `json:"raw"`
}

type ReplaysWatchedCount struct {
	StartDate string `json:"start_date"`
	Count     int    `json:"count"`
}

type Level struct {
	Current  int     `json:"current"`
	Progress float64 `json:"progress"`
}

type PP struct {
	Raw float64 `json:"raw"`
}

type UserRank struct {
	Global  int `json:"global"`
	Country int `json:"country"`
}

type UserStatistics struct {
	HitAccuracy  float64     `json:"hit_accuracy"`
	IsRanked     bool        `json:"is_ranked"`
	Level        Level       `json:"level"`
	MaximumCombo int         `json:"maximum_combo"`
	PlayCount    int         `json:"play_count"`
	PlayTime     int         `json:"play_time"`
	PP           float64     `json:"pp"`
	GlobalRank   int         `json:"global_rank"`
	RankedScore  int64       `json:"ranked_score"`
	TotalHits    int64       `json:"total_hits"`
	TotalScore   int64       `json:"total_score"`
	GradeCounts  GradeCounts `json:"grade_counts"`
	Rank         UserRank    `json:"rank"`
}

type UserAchievement struct {
	AchievementId int       `json:"achievement_id"`
	EarnedAt      time.Time `json:"earned_at"`
}

type RankHistory struct {
	Mode string `json:"mode"`
	Data []int  `json:"data"`
}

type RankHightes struct {
	Date time.Time `json:"date"`
	Rank int       `json:"rank"`
}

type UserTeamBadge struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	URL         string `json:"url"`
}

type User struct {
	AvatarURL                string                `json:"avatar_url"`
	DefaultGroup             string                `json:"default_group"`
	Id                       int                   `json:"id"`
	IsActive                 bool                  `json:"is_active"`
	IsBot                    bool                  `json:"is_bot"`
	IsDeleted                bool                  `json:"is_deleted"`
	IsOnline                 bool                  `json:"is_online"`
	IsSupporter              bool                  `json:"is_supporter"`
	LastVisit                *time.Time            `json:"last_visit"`
	PmFriendsOnly            bool                  `json:"pm_friends_only"`
	ProfileColour            *string               `json:"profile_colour"`
	Username                 string                `json:"username"`
	CoverURL                 string                `json:"cover_url"`
	Discord                  *string               `json:"discord"`
	HasSupported             bool                  `json:"has_supported"`
	Interests                *string               `json:"interests"`
	JoinDate                 time.Time             `json:"join_date"`
	Kudosu                   Kudosu                `json:"kudosu"`
	Location                 *string               `json:"location"`
	MaxBlocks                int                   `json:"max_blocks"`
	MaxFriends               int                   `json:"max_friends"`
	Occupation               *string               `json:"occupation"`
	Playmode                 string                `json:"playmode"`
	Playstyle                []string              `json:"playstyle"`
	PostCount                int                   `json:"post_count"`
	ProfileHue               int                   `json:"profile_hue"`
	ProfileOrder             []string              `json:"profile_order"`
	Title                    *string               `json:"title"`
	Twitter                  *string               `json:"twitter"`
	Website                  *string               `json:"website"`
	Country                  UserCountry           `json:"country"`
	Cover                    Cover                 `json:"cover"`
	IsRestricted             bool                  `json:"is_restricted"`
	AccountHistory           []any                 `json:"account_history"`
	ActiveTournamentBanner   any                   `json:"active_tournament_banner"`
	Badges                   []Badge               `json:"badges"`
	FavouriteBeatmapsetCount int                   `json:"favourite_beatmapset_count"`
	FollowerCount            int                   `json:"follower_count"`
	GraveyardBeatmapsetCount int                   `json:"graveyard_beatmapset_count"`
	Groups                   []Group               `json:"groups"`
	LovedBeatmapsetCount     int                   `json:"loved_beatmapset_count"`
	MonthlyPlaycounts        []MonthlyPlaycount    `json:"monthly_playcounts"`
	Page                     Page                  `json:"page"`
	PendingBeatmapsetCount   int                   `json:"pending_beatmapset_count"`
	PreviousUsernames        []string              `json:"previous_usernames"`
	RankedBeatmapsetCount    int                   `json:"ranked_beatmapset_count"`
	ReplaysWatchedCounts     []ReplaysWatchedCount `json:"replays_watched_counts"`
	ScoresFirstCount         int                   `json:"scores_first_count"`
	Statistics               UserStatistics        `json:"statistics"`
	SupportLevel             int                   `json:"support_level"`
	UserAchievements         []UserAchievement     `json:"user_achievements"`
	RankHistory              RankHistory           `json:"rank_history"`
	RankHighest              RankHistory           `json:"rank_highest"`
	Team                     *UserTeamBadge        `json:"team,omitempty"`
}

type UserTiny struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	Color       string `json:"profile_colour"`
	CountryCode string `json:"country_code"`
	IsActive    bool   `json:"is_active"`
	IsBot       bool   `json:"is_bot"`
	IsDeleted   bool   `json:"is_deleted"`
	IsOnline    bool   `json:"is_online"`
	IsSupporter bool   `json:"is_supporter"`
}

type UserBasic struct {
	Id            int           `json:"id"`
	Username      string        `json:"username"`
	Country       UserCountry   `json:"country"`
	Team          UserTeamBadge `json:"team"`
	DefaultGroup  string        `json:"default_group"`
	ProfileColor  string        `json:"profile_colour"`
	IsActive      bool          `json:"is_active"`
	IsBot         bool          `json:"is_bot"`
	IsDeleted     bool          `json:"is_deleted"`
	IsOnline      bool          `json:"is_online"`
	IsSupporter   bool          `json:"is_supporter"`
	LastVisit     string        `json:"last_visit"`
	PMFriendsOnly bool          `json:"pm_friends_only"`
	Cover         struct {
		Id        string `json:"id"`
		Url       string `json:"url"`
		CustomUrl string `json:"custom_url"`
	} `json:""`
}
