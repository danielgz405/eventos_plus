package structures

type InsertReserveRequest struct {
	EventId string `json:"event_id" bson:"event_id"`
}

type UpdateReserveRequest struct {
	Acceted       bool   `json:"acceted" bson:"acceted"`
	DateToAcceted string `json:"date_to_acceted" bson:"date_to_acceted"`
}
