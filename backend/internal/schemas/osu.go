package schemas

type Mode string
type ModeInt int

const (
	ModeStandard    Mode    = "osu"
	ModeStandardInt ModeInt = 0
	ModeTaiko       Mode    = "taiko"
	ModeTaikoInt    ModeInt = 1
	ModeFruits      Mode    = "fruits"
	ModeFruitsInt   ModeInt = 2
	ModeMania       Mode    = "mania"
	ModeManiaInt    ModeInt = 3
)

type Language struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type GradeCounts struct {
	A   int `json:"a"`
	S   int `json:"s"`
	SH  int `json:"sh"`
	SS  int `json:"ss"`
	SSH int `json:"ssh"`
}
