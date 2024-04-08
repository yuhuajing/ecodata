package config

import (
	"go.mongodb.org/mongo-driver/mongo"
	"main/common/ethconn"
)

var (
	EthServer = "https://eth-sepolia.g.alchemy.com/v2/t39LUOfkEMNMz_uxQk2gkwK3kJ1HwDZF"
	Client    = ethconn.ConnBlockchain(EthServer)
	TraceSC   = "0x6B423275216D3fcA49D1c8624286c47C44B96Dd5"
)
var (
	Mongodburl           = "mongodb://localhost:27017"
	Dbname               = "ecotra"
	DbcollectionUserInfo = "userinfo"
	DbcollectionEcoInfo  = "ecoinfo"
	DbcollectionTousu    = "tousu"
	Mongoclient          *mongo.Client
)
