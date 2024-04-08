package structures

type InsertTypeEventRequest struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}

type UpdateTypeEventRequest struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}
