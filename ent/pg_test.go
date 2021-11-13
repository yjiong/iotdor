package ent

import (
	"context"
	"dcmbroker/ent/migrate"
	"testing"

	// postgres driver
	_ "github.com/lib/pq"
)

func TestENT(t *testing.T) {
	t.Log("++++++++++++++++++test ent++++++++++++++++++++")
	client, err := Open("postgres", "host=mqtt.yaojiong.top port=5432 user=iotd dbname=iotd password=yaojiong")
	if err != nil {
	}
	if client.Schema.Create(context.Background(), migrate.WithDropColumn(true), migrate.WithDropIndex(true)) != nil {
		t.Fail()
	}
	defer client.Close()
}
