package client

type StorageClient struct {
	config *Config
}

func NewStorageClient(config *Config) *StorageClient {
	return &StorageClient{
		config: config,
	}
}
