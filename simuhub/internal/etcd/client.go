package etcd

import (
	"errors"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"go.etcd.io/etcd/client/v3"
)

var (
	Client *clientv3.Client
	Key    string
)

func Connect() error {
	uri, ok := os.LookupEnv("ETCD_URI")

	if !ok {
		return errors.New("`ETCD_URI` not found in environment variables")
	}

	key, ok := os.LookupEnv("ETCD_KEY")

	if !ok {
		return errors.New("`ETCD_KEY` not found in environment variables")
	}

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{uri},
		DialTimeout: 3 * time.Second,
	})

	if err != nil {
		return err
	}

	log.Info().Msg("Connected to Etcd")

	Client = client
	Key = key

	return nil
}
