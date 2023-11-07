package bark

type BarkParamsLevel string

const (
	Active        BarkParamsLevel = "active"
	TimeSensitive BarkParamsLevel = "timeSensitive"
	Passive       BarkParamsLevel = "passive"
)

type BarkParams struct {
	Title string          `json:"title"`
	Body  string          `json:"body"`
	Level BarkParamsLevel `json:"Level"`
}
