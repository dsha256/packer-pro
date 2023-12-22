package packer

import (
	"context"
	_ "embed"
	"log"
	"time"

	"encore.dev/storage/cache"
	"encore.dev/storage/sqldb"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/dsha256/packer-pro/internal/entity"
)

const sortedSizesKeyName = "SortedSizes"

// SortedSizesKey is a Redis key for SortedSizes keyspace.
type SortedSizesKey struct {
	Name string
}

var (
	// Cache is Redis cluster of this app.
	Cache = cache.NewCluster("packer-cache", cache.ClusterConfig{
		EvictionPolicy: cache.AllKeysLRU,
	})

	// SortedSizes is a Redis Keyspace.
	SortedSizes = cache.NewListKeyspace[SortedSizesKey, int](Cache, cache.KeyspaceConfig{
		KeyPattern:    "sortedSizes/:Name",
		DefaultExpiry: cache.ExpireIn(1 * time.Hour),
	})

	// sortedSizesKey...
	sortedSizesKey = SortedSizesKey{Name: sortedSizesKeyName}

	// DB ...
	DB = sqldb.NewDatabase("packer", sqldb.DatabaseConfig{
		Migrations: "./db/migrations",
	})
)

//encore:service
type Packer struct {
	entity *entity.Client
}

//lint:ignore U1000 This function is used by Encore to init the service.
func initPacker() (*Packer, error) {
	packer := Packer{}
	ctx := context.Background()

	dbDriver := entsql.OpenDB(dialect.Postgres, DB.Stdlib())
	entClient := entity.NewClient(entity.Driver(dbDriver))
	packer.entity = entClient

	err := refreshSortedSizesCacheFromDB(ctx, packer.entity)
	if err != nil {
		return &Packer{}, err
	}

	return &packer, nil
}

// Shutdown contains a grace-full shutdown scenario.
func (packer *Packer) Shutdown(force context.Context) {
	err := packer.entity.Close()
	if err != nil {
		log.Fatalln("can not close the DB connection:", err)
	}
}
