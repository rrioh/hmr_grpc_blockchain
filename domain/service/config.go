package service

import (
	"hmr_grpc_blockchain/domain/model"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

/*
tomlファイルのconfigをもとにmodel.ServerConfig構造体を作成し、返却する。
*/
func NewServerConfig() (*model.ServerConfig, error) {
	var config model.ServerConfig

	_, err := toml.DecodeFile(filepath.Join(os.Getenv("GOPATH"), "src/hmr_grpc_blockchain/conf/server.toml"), &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
