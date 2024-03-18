package models

type Vehicle struct {
	ID        string `json:"ID"`
	JourneyID string `json:"JourneyIdentifier"`
	AVMKey    string `json:"AVMKey"`

	MOTCode        int    `json:"MOTCode"`
	MOTSubcode     int    `json:"MOTSubcode"`
	MOTDescription string `json:"MOTDescr"`

	X         string `json:"X"`
	Y         string `json:"Y"`
	PreviousX string `json:"XPrevious"`
	PreviousY string `json:"YPrevious"`

	Timestamp         string `json:"Timestamp"`
	PreviousTimestamp string `json:"TimestampPrevious"`

	LineText   string `json:"LineText"`
	LineNumber string `json:"LineNumber"`

	CurrentStop string `json:"CurrentStop"`
	NextStop    string `json:"NextStop"`
	Direction   string `json:"DirectionText"`

	TrainNumber string `json:"TrainNumber"`
	TrainName   string `json:"TrainName"`
	TrainType   string `json:"TrainType"`
	Operator    string `json:"Operator"`
	IsAtStop    bool   `json:"IsAtStop"`

	RealtimeAvailable int    `json:"RealtimeAvailable"`
	ProductID         string `json:"ProductIdentifier"`
	OperationDay      string `json:"DayOfOperation"`
	Delay             int    `json:"Delay"`
}
