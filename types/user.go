package types

type User struct {
	ID        int    `bson:"_id" json:"id,omitempty"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
}
