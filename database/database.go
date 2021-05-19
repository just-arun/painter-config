package database

import (
	aero "github.com/aerospike/aerospike-client-go"
	"github.com/just-arun/painter-config/config"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	AeroClient  *aero.Client
	MongoClient *mongo.Client
)

func InitConnection() {
	aerospike := config.AppConfig.Database.Aerospike
	GetAerospikeInstance(aerospike.Host, aerospike.Port)
	GetMongoInstance()
}
