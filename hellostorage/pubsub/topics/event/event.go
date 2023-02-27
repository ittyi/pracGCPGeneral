package event

type UpdateLabelRelationJobEventMessage struct {
	Event   string `json:"event"`
	Action  string `json:"action"`
	TeamID  string `json:"teamId"`
	LabelID string `json:"labelId"`
}
