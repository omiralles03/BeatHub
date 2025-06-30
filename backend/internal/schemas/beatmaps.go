package schemas

import "time"

type BeatmapsetStatus string

const (
	BeatmapsetStatusRanked    BeatmapsetStatus = "ranked"
	BeatmapsetStatusApproved  BeatmapsetStatus = "approved"
	BeatmapsetStatusQualified BeatmapsetStatus = "qualified"
	BeatmapsetStatusLoved     BeatmapsetStatus = "loved"
	BeatmapsetStatusPending   BeatmapsetStatus = "pending"
	BeatmapsetStatusWIP       BeatmapsetStatus = "wip"
	BeatmapsetStatusGraveyard BeatmapsetStatus = "graveyard"
)

type Beatmap struct {
	Id               int              `json:"id"`
	BeatmapsetId     int              `json:"beatmapset_id"`
	DifficultyRating float64          `json:"difficulty_rating"`
	Mode             Mode             `json:"mode"`
	ModeInt          ModeInt          `json:"mode_int"`
	Status           BeatmapsetStatus `json:"status"`
	TotalLength      int              `json:"total_length"`
	UserId           int              `json:"user_id"`
	Version          string           `json:"version"`
	OD               float64          `json:"accuracy"` // Note: this might be 'drain' or 'overall_difficulty' in osu! API. Check documentation if issues
	AR               float64          `json:"ar"`
	CS               float64          `json:"cs"`
	HP               float64          `json:"drain"`
	BPM              float64          `json:"bpm"`
	Convert          bool             `json:"convert"`
	CountCircles     int              `json:"count_circles"`
	CountSliders     int              `json:"count_sliders"`
	CountSpinners    int              `json:"count_spinners"`
	DeletedAt        time.Time        `json:"deleted_at"`
	HitLength        int              `json:"hit_length"`
	IsScoreable      bool             `json:"is_scoreable"`
	LastUpdated      time.Time        `json:"last_updated"`
	Passcount        int              `json:"passcount"`
	Playcount        int              `json:"playcount"`
	Ranked           int              `json:"ranked"`
	URL              string           `json:"url"`
}

type Covers struct {
	Cover       string `json:"cover"`
	Cover2X     string `json:"cover@2x"`
	Card        string `json:"card"`
	Card2X      string `json:"card@2x"`
	List        string `json:"list"`
	List2X      string `json:"list@2x"`
	Slimcover   string `json:"slimcover"`
	Slimcover2X string `json:"slimcover@2x"`
}

type Nominator struct {
	BeatmapsetId int    `json:"beatmapset_id"`
	Rulesets     []Mode `json:"rulesets"`
	Reset        bool   `json:"reset"`
	UserId       int    `json:"user_id"`
}

type BeatmapDescription struct {
	Description string `json:"description"`
}

type BeatmapGenre struct {
	Name string `json:"name"`
}

type Beatmapset struct {
	Artist             string              `json:"artist"`
	ArtistUnicode      string              `json:"artist_unicode"`
	Covers             Covers              `json:"covers"`
	Creator            string              `json:"creator"`
	FavouriteCount     int                 `json:"favourite_count"`
	Id                 int                 `json:"id"`
	NSFW               bool                `json:"nsfw"`
	Offset             int                 `json:"offset"`
	PlayCount          int                 `json:"play_count"`
	SubmittedDate      time.Time           `json:"submitted_date"`
	RankedDate         *time.Time          `json:"ranked_date,omitempty"`
	PreviewURL         string              `json:"preview_url"`
	Source             string              `json:"source"`
	Status             BeatmapsetStatus    `json:"status"`
	Spotlight          bool                `json:"spotlight"`
	Title              string              `json:"title"`
	TitleUnicode       string              `json:"title_unicode"`
	UserId             int                 `json:"user_id"`
	Video              bool                `json:"video"`
	Tags               string              `json:"tags"`
	Beatmaps           []Beatmap           `json:"beatmaps,omitempty"`
	Converts           []Beatmap           `json:"converts,omitempty"`
	Description        *BeatmapDescription `json:"description,omitempty"`
	Genre              *BeatmapGenre       `json:"genre,omitempty"`
	HasFavourited      *bool               `json:"has_favourited,omitempty"`
	Language           *Language           `json:"language,omitempty"`
	CurrentNominations []Nominator         `json:"current_nominations,omitempty"`
	PackTags           []string            `json:"pack_tags,omitempty"`
	Ratings            []int               `json:"ratings,omitempty"`
	User               *UserTiny           `json:"user,omitempty"` // UserTiny from users.go
	TrackId            *int                `json:"track_id,omitempty"`
	// RelatedUsers       any                 `json:"related_users,omitempty"` // Uncomment if you uncommented in your friend's code
}
