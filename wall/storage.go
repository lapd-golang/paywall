package wall

import "github.com/mongodb/mongo-go-driver/mongo"

func connect() (*mongo.Client, error) {
	return mongo.NewClient("mongodb://localhost:27017")
}
