package database

import (
	aero "github.com/aerospike/aerospike-client-go"
	painterconfig "github.com/just-arun/painter-config"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	AeroClient  *aero.Client
	MongoClient *mongo.Client
)

func InitConnection() {
	aerospike := painterconfig.AppConfig.Database.Aerospike
	GetAerospikeInstance(aerospike.Host, aerospike.Port)
	GetMongoInstance()
}
