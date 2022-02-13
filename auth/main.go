package main

import (
	"context"
	"coolcar/shared/server"
	"flag"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	"os"
)

var addr = flag.String("addr", ":8081", "address is listen")
var mongoURI = flag.String("mongo_uri", "mongodb://localhost:27017", "mongo uri")
var privateKeyFile = flag.String("private_key_file", "auth/private.key", "private key file")
var wechatAppID = flag.String("wechat_app_id", "<APPID>", "wechat app id")
var wechatAppSecret = flag.String("wechat_app_secret", "<APPSERET>", "wechat app secret")

func main() {
	flag.Parse()
	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatal("cannot create logger: %v", err)
	}

	c := context.Background()
	mongoClient, err := mongo.Connect(c, options.Client().ApplyURI(*mongoURI))
	if err != nil {
		logger.Fatal("cannot connect mongodb", zap.Error(err))
	}

	pkFile, err := os.Open(*privateKeyFile)
	if err != nil {
		logger.Fatal("cannot open private key", zap.Error(err))
	}

	pkBates, err := ioutil.ReadAll(pkFile)
	if err != nil {
		logger.Fatal("cannot read private key", zap.Error(err))
	}

	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(pkBates)
	if err != nil {
		logger.Fatal("cannot parse private key", zap.Error(err))
	}

	logger.Sugar().Fatal(server.)
}
