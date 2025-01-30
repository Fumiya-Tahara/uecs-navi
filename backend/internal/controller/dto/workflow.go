package dto

type Workflow struct {
	Name     string `json:"name"`
	M304ID   int    `json:"m304_id"`
	DeviceID int    `json:"device_id"`
}

func NewWorkflow(name string, m304ID, deviceID int) *Workflow {
	return &Workflow{
		Name:     name,
		M304ID:   m304ID,
		DeviceID: deviceID,
	}
}
