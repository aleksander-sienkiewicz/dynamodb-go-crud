package entities

import (
	"time" //for createdat/updated at functions

	"github.com/google/uuid"
)

/*
base.go contains base things, id, created at, updated at, very related to base is entity/product
*/
type Interface interface { //interfaces are the craziest thing ever? not acc but pre cool how much time it saves
	GenerateID()
	SetCreatedAt()
	SetUpdatedAt()
	TableName() string
	GetMap() map[string]interface{}
	GetFilterId() map[string]interface{} //since these all share same methods, we can group them in the same interface
} // now if we call type interface, we can pass thru any of the objects containing the above methods.

type Base struct { //level of abstraction, go cant handle json so we gotta play around with it a lil bti <3
	ID        uuid.UUID `json:"_id"`       //id, type, pass str
	CreatedAt time.Time `json:"createdAt"` //created at
	UpdatedAt time.Time `json:"updatedAt"` //updated at
} //for like go, python ,ruby we gotta feed in the strings , jscript will support it

// (b *Base is struct method)
func (b *Base) GenerateID() {
	b.ID = uuid.New()
}

func (b *Base) SetCreatedAt() {
	b.CreatedAt = time.Now()
}

func (b *Base) SetUpdatedAt() {
	b.UpdatedAt = time.Now()
}

func GetTimeFormat() string { //returns str
	return "2006-01-02T15:04:05-0700" //regular time stamp, thats the formal were lookin fore
}
