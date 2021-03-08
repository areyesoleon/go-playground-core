package core

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MODEL

//Res model
type Res struct {
	Ok      bool        `json:"ok"`
	Res     interface{} `json:"res,omitempty"`
	Message string      `json:"message,omitempty"`
}

//User ...
type User struct {
	ID       primitive.ObjectID `bson:"_id, omitempty" json:"id,omitempty"`
	Name     string             `json:"name"`
	LastName string             `json:"lastName"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Admin    bool               `json:"admin"`
	Create   time.Time          `json:"create"`
	Update   time.Time          `json:"update"`
}

//Shop ...
type Shop struct {
	ID      primitive.ObjectID `json:"id"`
	Name    string             `json:"name"`
	Address string             `json:"address"`
	Coin    string             `json:"coin"`
	Create  time.Time          `json:"create"`
	Update  time.Time          `json:"update"`
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
