package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/BurntSushi/toml"
	"github.com/maximprokopchuk/address_service/internal/config"
	"github.com/maximprokopchuk/address_service/internal/grpcserver"
	"github.com/maximprokopchuk/address_service/internal/store"
	"github.com/maximprokopchuk/address_service/pkg/api"
	"google.golang.org/grpc"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config_path", "configs/config.toml", "path to config TOML file")
}

func run() error {
	ctx := context.Background()
	cfg := config.NewConfig()
	_, err := toml.DecodeFile(configPath, cfg)

	if err != nil {
		return err
	}

	st := store.New(cfg.Store)

	if err := st.Open(ctx); err != nil {
		return err
	}

	defer st.Close(ctx)

	s := grpc.NewServer()
	srv := grpcserver.New(st)
	api.RegisterAddressServiceServer(s, srv)

	l, err := net.Listen("tcp", ":"+cfg.BindUrl)

	if err != nil {
		return err
	}

	log.Println("LISTEN " + cfg.BindUrl)

	if err = s.Serve(l); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
