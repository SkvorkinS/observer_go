package observer_go

type Client struct {
	ID             int    `json:"id"`
	FamilyName     string `json:"familyName"`
	FirstName      string `json:"firstName"`
	PatronymicName string `json:"patronymicName"`
	LastNPSMessage string `json:"lastNpsMessage"`
	BonusesSum     int    `json:"bonusesSum"`
	Birthday       string `json:"birthday"`
	Phone          string `json:"phone"`
	FormatedPhone  string `json:"formatedPhone"`
}

type Deal struct {
	ID                  int     `json:"id"`
	CreatedAt           string  `json:"createdAt"`
	FinishedAt          string  `json:"finishedAt"`
	FinishExpupectingAt string  `json:"finishExpectingAt"`
	DepartmentName      string  `json:"departmentName"`
	Client              Client  `json:"client"`
	Type                string  `json:"type"`
	Status              string  `json:"status"`
	Sum                 float64 `json:"sum"`
	PaidSum             float64 `json:"paidSum"`
	ReceiptPrinted      bool    `json:"receiptPrinted"`
	CreatedById         int     `json:"createdById"`
	Redo                bool    `json:"redo"`
	ReadyStatusInStore  bool    `json:"readyStatusInStore"`
	Company             string  `json:"dealCompany"`
	CreatedUserName     string  `json:"createdUser_name"`
	NPSStatus           string  `json:"npsStatus"`
	RISStatus           string  `json:"risStatus"`
	DBStatus            string  `json:"dbStatus"`

	Error error
}
