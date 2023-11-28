package bark

type BarkParamsLevel string

const (
	LevelActive        BarkParamsLevel = "active"
	LevelTimeSensitive BarkParamsLevel = "timeSensitive"
	LevelPassive       BarkParamsLevel = "passive"
)

type BarkParams struct {
	Title string          `json:"title"`
	Body  string          `json:"body"`
	Level BarkParamsLevel `json:"level"`
	Icon  string          `json:"icon"`
}
