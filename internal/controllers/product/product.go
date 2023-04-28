package product

import (
	"time"

	"github.com/aleksander-sienkiewicz/dynamodb-go-crud/internal/entities/product"
	"github.com/aleksander-sienkiewicz/dynamodb-go-crud/internal/repository/adapter"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/google/uuid"
)

//we got rroutes in a list, and we have functions, and handlers
//these functions have a 1:1 relationship with the routes and handlers
//routes we have routes post, get, put, etc., we will have funcs for those routes

// create controller
type Controller struct {
	repository adapter.Interface
}

// create interface
type Interface interface { //our beautiful interface, I <3 abstraction
	ListOne(ID uuid.UUID) (entity product.Product, err error) //list 1 prod
	ListAll() (entities []product.Product, err error)         //lists db
	Create(entity *product.Product) (uuid.UUID, error)        //return id, maybe err
	Update(ID uuid.UUID, entity *product.Product) error       //maybe er
	Remove(ID uuid.UUID) error                                //maybe return err
}

// create new controlleer for us,
func NewController(repository adapter.Interface) Interface {
	return &Controller{repository: repository}
}

// pass id, return product or error
func (c *Controller) ListOne(id uuid.UUID) (entity product.Product, err error) {

	entity.ID = id //pass id to entity.id
	response, err := c.repository.FindOne(entity.GetFilterId(), entity.TableName())

	if err != nil { //if theres err
		return entity, err //error code

	} //this whole func is elaborated upon in entities/product
	return product.ParseDynamoAtributeToStruct(response.Item) //return one item
}

// list our 1:1 functions mapping routes to handles
// for reading we can either listall and read all or listone
// this is listall (in case u cant read but can also just only read green)
func (c *Controller) ListAll() (entities []product.Product, err error) {
	entities = []product.Product{} //entities == slice of products
	var entity product.Product     //var entity is single entity in that slice

	//create filter. we will filter based on name, and also that name should not be empty
	filter := expression.Name("name").NotEqual(expression.Value(""))
	condition, err := expression.NewBuilder().WithFilter(filter).Build()
	//pass filter here, build condition. use condition to run query on data base
	if err != nil { //might have error here so check
		return entities, err //handle errors
	}
	//use respositry to run the database levle function, findall. iwhich is located in internal/repository
	response, err := c.repository.FindAll(condition, entity.TableName()) //find all is the database level func
	if err != nil {                                                      //check for err
		return entities, err //cry, always check for errors when u run a func or ur goofy tbh
	}
	//ya so that uses dynamoDB func Findall to search the table of tablename using conditions u defined jsut above into a
	//particular struct golang understands

	if response != nil { //if we got a response, check if empty
		for _, value := range response.Items { //for response items, if not empty
			entity, err := product.ParseDynamoAtributeToStruct(value) //entity captures parse values. the dynamodb func
			if err != nil {                                           //handle error here too                       //if err is encountered u cry
				return entities, err
			}
			entities = append(entities, entity) // append data to list of entities
		}
	}

	return entities, nil //return list, nil for error if all goes well
}

// accept complete db, after creating entity, return id. type uuid, return error if theres an issue
func (c *Controller) Create(entity *product.Product) (uuid.UUID, error) {
	entity.CreatedAt = time.Now()                                              //when this is called, we want a timestamp to know when exactly the entry was created
	_, err := c.repository.CreateOrUpdate(entity.GetMap(), entity.TableName()) //pass table name, entity, call createorupdate in respository
	//and creates that new entry in the data base
	return entity.ID, err //return the entity id of the entry. nil for err idealy
}

// send id of product u want to replace, to locate the index. pass new entity to replace- return error if error
func (c *Controller) Update(id uuid.UUID, entity *product.Product) error {
	found, err := c.ListOne(id) // call listOne on id, capture into found
	if err != nil {             //if something goes wrong
		return err //return error code
	}
	found.ID = id                                                            //id that user passed
	found.Name = entity.Name                                                 //name user passed
	found.UpdatedAt = time.Now()                                             //update timestamp for update log
	_, err = c.repository.CreateOrUpdate(found.GetMap(), entity.TableName()) //call respository pkg, create or updayte func
	// pass found map and entity table name
	return err //nil for err or err code
}

// find id and remove.
func (c *Controller) Remove(id uuid.UUID) error {
	entity, err := c.ListOne(id) //call listOne, pass thru id
	if err != nil {              //handle err
		return err //error code
	}
	_, err = c.repository.Delete(entity.GetFilterId(), entity.TableName()) //for c.respository, delete it
	//we dont wanna capture a deleted variable so capture it in _, err if it doesnt work
	return err //return err code if didnt work
}
