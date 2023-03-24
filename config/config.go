package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	HTTPServer httpServer
	Mongo      mongo
	Redis      redis
	Kafka      kafka
	App        app
	Topic      topic
	Jwt        jwt
	Adapter    adapter
}

type app struct {
	ProcessedTransactionApi string `envconfig:"PROCESSED_TRANSACTION_API"`
}

type httpServer struct {
	Host string `envconfig:"HTTP_SERVER_HOST" default:"localhost"`
	Port int    `envconfig:"HTTP_SERVER_PORT" default:"8080"`
}

type mongo struct {
	URI      string `envconfig:"MONGO_URI" default:"mongodb://localhost:27017"`
	Database string `envconfig:"DB_NAME" default:"main"`
}

type redis struct {
	Host     string `envconfig:"REDIS_HOST" default:"localhost"`
	Port     int    `envconfig:"REDIS_PORT" default:"6379"`
	Username string `envconfig:"REDIS_USERNAME"`
	Password string `envconfig:"REDIS_PASSWORD"`
}

type kafka struct {
	Server   string `envconfig:"KAFKA_SERVER"`
	ClientID string `envconfig:"KAFKA_CLIENT_ID"`
	Verbose  bool   `envconfig:"KAFKA_VERBOSE" default:"false"`
	Username string `envconfig:"KAFKA_USERNAME"`
	Password string `envconfig:"KAFKA_PASSWORD"`
}

type topic struct {
	CreateChatTopic string `envconfig:"CREATE_CHAT_TOPIC"`
	UpdateChatTopic string `envconfig:"UPDATE_CHAT_TOPIC"`
}

type jwt struct {
	JwtSecret string `envconfig:"JWT_SECRET"`
}

type adapter struct {
	AuthenApi string `envconfig:"AUTHEN_API"`
}



var cfg config

func Init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		err = godotenv.Load("../.env")
	}

	if err != nil {
		log.Printf("load env error : %s", err.Error())
	}
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("read env error : %s", err.Error())
	}
}

func Get() config {
	return cfg
}
