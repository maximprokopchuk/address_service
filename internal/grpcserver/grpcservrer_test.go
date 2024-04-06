package grpcserver_test

import (
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/maximprokopchuk/address_service/internal/config"
	"github.com/maximprokopchuk/address_service/internal/grpcserver"
	"github.com/maximprokopchuk/address_service/internal/store"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	cfg := config.NewConfig()
	_, err := toml.DecodeFile("../../configs/config.test.toml", cfg)
	assert.Nil(t, err)
	st := store.New(cfg.Store)
	server := grpcserver.New(st)
	assert.NotNil(t, server)
	assert.Equal(t, st, server.Store, "should include store")
	assert.NotNil(t, server.CreateAddress)
	assert.NotNil(t, server.DeleteAddress)
	assert.NotNil(t, server.GetAddressById)
	assert.NotNil(t, server.ListAddressesByParentAndType)
}
