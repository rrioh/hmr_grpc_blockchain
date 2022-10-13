package model

type ServerConfig struct {
	Blockchain BlockchainConfig `toml:"blockchain"`
}

type BlockchainConfig struct {
	Address    string `toml:"address"`
	Difficulty int    `toml:"difficulty"`
}
