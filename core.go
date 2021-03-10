package core

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var singleTonMaster *MasterStruct
var once sync.Once

//MODEL

//Res model
type Res struct {
	Ok      bool        `json:"ok"`
	Res     interface{} `json:"res,omitempty"`
	Message string      `json:"message,omitempty"`
}

// TimeData ...
type TimeData struct {
	Create time.Time `json:"create"`
	Update time.Time `json:"update"`
}

//User ...
type User struct {
	ID       primitive.ObjectID `bson:"_id, omitempty" json:"id,omitempty"`
	Name     string             `json:"name"`
	LastName string             `json:"lastName"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Admin    bool               `json:"admin"`
}

//Place ...
type Place struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Coin    string `json:"coin"`
}

//Shop ...
type Shop struct {
	ID primitive.ObjectID `bson:"_id, omitempty" json:"id,omitempty"`
	Place
	TimeData
}

//Kiosk ...
type Kiosk struct {
	ID primitive.ObjectID `bson:"_id, omitempty" json:"id,omitempty"`
	Place
	IDshop primitive.ObjectID `json:"idShop"`
	TimeData
}

//MasterStruct ...
type MasterStruct struct {
	IDuser  primitive.ObjectID `json:"idUser"`
	IDshop  primitive.ObjectID `json:"idShop"`
	IDkiosk primitive.ObjectID `json:"idKiosk"`
}

//SingleTonMaster ...
func SingleTonMaster() *MasterStruct {

	once.Do(func() {
		singleTonMaster = &MasterStruct{}
	})

	return singleTonMaster
}

//SetIDUser set user id
func (m *MasterStruct) SetIDUser(id primitive.ObjectID) {
	m.IDuser = id
}

//GetIDUser user id
func (m *MasterStruct) GetIDUser() primitive.ObjectID {
	return m.IDuser
}

//SetIDShop shop user id
func (m *MasterStruct) SetIDShop(id primitive.ObjectID) {
	m.IDshop = id
}

//GetIDShop shop id
func (m *MasterStruct) GetIDShop() primitive.ObjectID {
	return m.IDshop
}

//SetIDkiosk set kiosk id
func (m *MasterStruct) SetIDkiosk(id primitive.ObjectID) {
	m.IDkiosk = id
}

//GetIDkiosk kiosk id
func (m *MasterStruct) GetIDkiosk() primitive.ObjectID {
	return m.IDkiosk
}

//Db ...
func Db(collection string) *mongo.Collection {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.puckd.mongodb.net/test"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client.Database("minim").Collection(collection)
}
