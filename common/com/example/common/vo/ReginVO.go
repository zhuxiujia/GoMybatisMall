package vo

type ReginVO struct {
	Value    string    `json:"value"`
	Label    string    `json:"label"`
	Children []ReginVO `json:"children"`
}
