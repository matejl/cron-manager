package model

type schedule struct {
	StringSpec string `json:"stringSpec"`
}

func NewSchedule(stringSpec string) schedule {
	return schedule{
		StringSpec: stringSpec,
	}
}