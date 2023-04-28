package product

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/aleksander-sienkiewicz/dynamodb-go-crud/internal/entities"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// define a product struct
type Product struct {
	entities.Base        //itll have entities.base (id, createdat, updatedat)
	Name          string `json:"name"` //will have its own name - in ` ` is how it looks in json
}

// helps us do a lot of marshaling and unmarshalling from json, to something understandable to golang
func InterfaceToModel(data interface{}) (instance *Product, err error) { //we will take data as an interface which we declared in base
	//gonna return our product, err
	bytes, err := json.Marshal(data) //take bytes and err, call json.marshal. pass data to marshal
	if err != nil {                  //handle error
		return instance, err //return instance, err
	}

	return instance, json.Unmarshal(bytes, &instance) //return instance and readable data
}

// getfilterid happens when we listone. define that here
func (p *Product) GetFilterId() map[string]interface{} { //map type interface
	//method for that structure
	return map[string]interface{}{"_id": p.ID.String()}
}

func (p *Product) TableName() string {
	return "products" //return name of table that we want to send data to in dynamoDB
} // we dont have to do this, but makes it clean, provides abstraction. generally a yummy project structure imo

func (p *Product) Bytes() ([]byte, error) { //return slice of bytes
	return json.Marshal(p) //product
}

func (p *Product) GetMap() map[string]interface{} {
	return map[string]interface{}{
		"_id":       p.ID.String(),
		"name":      p.Name,
		"createdAt": p.CreatedAt.Format(entities.GetTimeFormat()),
		"updatedAt": p.UpdatedAt.Format(entities.GetTimeFormat()),
	}
}

// confusing ahh function.
func ParseDynamoAtributeToStruct(response map[string]*dynamodb.AttributeValue) (p Product, err error) {
	if response == nil || (response != nil && len(response) == 0) {
		return p, errors.New("Item not found")
	}
	for key, value := range response {
		if key == "_id" {
			p.ID, err = uuid.Parse(*value.S)
			if p.ID == uuid.Nil {
				err = errors.New("Item not found")
			}
		}
		if key == "name" {
			p.Name = *value.S
		}
		if key == "createdAt" {
			p.CreatedAt, err = time.Parse(entities.GetTimeFormat(), *value.S)
		}
		if key == "updatedAt" {
			p.UpdatedAt, err = time.Parse(entities.GetTimeFormat(), *value.S)
		}
		if err != nil {
			return p, err
		}
	}

	return p, nil
}
