package observer_go

type Client struct {
	ID             int    `json:"id,omitempty"`
	FamilyName     string `json:"familyName,omitempty"`
	FirstName      string `json:"firstName,omitempty"`
	PatronymicName string `json:"patronymicName,omitempty"`
	LastNPSMessage string `json:"lastNpsMessage,omitempty"`
	BonusesSum     int    `json:"bonusesSum,omitempty"`
	Birthday       string `json:"birthday,omitempty"`
	Phone          string `json:"phone,omitempty"`
	FormatedPhone  string `json:"formatedPhone,omitempty"`
	UnFormatedPhone  string `json:"formatedPhone,omitempty"`

type Deal struct {
	ID                  int     `json:"id,omitempty"`
	CreatedAt           string  `json:"createdAt,omitempty"`
	FinishedAt          string  `json:"finishedAt,omitempty"`
	FinishExpupectingAt string  `json:"finishExpectingAt,omitempty"`
	DepartmentName      string  `json:"departmentName,omitempty"`
	Client              Client  `json:"client,omitempty"`
	Type                string  `json:"type,omitempty"`
	Status              string  `json:"status,omitempty"`
	Sum                 float64 `json:"sum,omitempty"`
	PaidSum             float64 `json:"paidSum,omitempty"`
	ReceiptPrinted      bool    `json:"receiptPrinted,omitempty"`
	CreatedById         int     `json:"createdById,omitempty"`
	Redo                bool    `json:"redo,omitempty"`
	ReadyStatusInStore  bool    `json:"readyStatusInStore,omitempty"`
	Company             string  `json:"dealCompany,omitempty"`
	CreatedUserName     string  `json:"createdUser_name,omitempty"`
	NPSStatus           string  `json:"npsStatus,omitempty"`
	RISStatus           string  `json:"risStatus,omitempty"`
	DBStatus            string  `json:"dbStatus,omitempty"`

	Error error
}
