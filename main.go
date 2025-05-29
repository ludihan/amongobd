package main

import (
	"context"
	"fmt"

	"amongodb/utils"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)


func main() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	utils.Check(err)

	uri := viper.GetString("MONGO_URI")
	if uri == "" {
		panic("uri do mongodb não está presente no .env")
	}
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client := utils.Must(mongo.Connect(opts))

	defer func() {
		utils.Check(client.Disconnect(context.TODO()))
	}()

	// Send a ping to confirm a successful connection
	utils.Check(client.Ping(context.TODO(), readpref.Primary()))
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}
