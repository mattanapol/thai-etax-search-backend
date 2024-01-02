package mongo

import (
	"context"
	"fmt"
	"github.com/mattanapol/thailand-etax-search-backend/internal/util/configuration"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func SetupMongoDb(mongoDbSetting mongoSetting,
) *mongo.Database {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?retryWrites=false",
		mongoDbSetting.User,
		mongoDbSetting.Password,
		mongoDbSetting.Host,
		mongoDbSetting.Port,
		mongoDbSetting.Name,
	)
	clientOpts := options.Client().
		//SetDirect(true). // for local test
		ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatalf("mongo open err: %v", err)
	}

	return client.Database(mongoDbSetting.Name)
}

type mongoSetting struct {
	Type     string `mapstructure:"type"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Name     string `mapstructure:"name"`
}

type setting struct {
	Database *mongoSetting `mapstructure:"mongodb"`
}

func newMongoDbConfiguration() (setting, error) {
	cfg := setting{
		Database: &mongoSetting{},
	}

	return configuration.NewConfiguration(&cfg)
}
