package packer

import (
	"context"
	_ "embed"
	"log"

	"encore.dev/storage/sqldb"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/dsha256/packer-pro/internal/entity"
)

var DB = sqldb.NewDatabase("packer", sqldb.DatabaseConfig{
	Migrations: "./db/migrations",
})

//encore:service
type Packer struct {
	entity *entity.Client
}

func initPacket() (*Packer, error) {
	packer := Packer{}

	dbDriver := entsql.OpenDB(dialect.Postgres, DB.Stdlib())
	entClient := entity.NewClient(entity.Driver(dbDriver))
	packer.entity = entClient

	return &packer, nil
}

func (packer *Packer) Shutdown(force context.Context) {
	err := packer.entity.Close()
	if err != nil {
		log.Fatalln("can not close the DB connection")
	}
}
