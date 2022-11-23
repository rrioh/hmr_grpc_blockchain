package model

/*
サーバのConfig。
*/
type ServerConfig struct {
	// ブロックチェーンのConfig
	Blockchain BlockchainConfig `toml:"blockchain"`
}

/*
ブロックチェーンのConfig。
*/
type BlockchainConfig struct {
	// ブロックチェーンのアドレス
	Address string `toml:"address"`

	// マイニングの際のPOWの難易度
	Difficulty int `toml:"difficulty"`
}
