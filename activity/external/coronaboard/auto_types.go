package coronaboard

type CoronaBoardData struct {
	ChartForGlobal struct {
		Kr struct {
			ConfirmedAcc []int64 `json:"confirmed_acc"`
		} `json:"KR"`
	} `json:"chartForGlobal"`
}
