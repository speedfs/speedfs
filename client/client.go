package client

type Client struct {
	config  *Config
	tracker *TrackerClient
	storage *StorageClient
}

func New(config *Config) *Client {
	return &Client{
		config:  config,
		tracker: NewTrackerClient(config),
		storage: NewStorageClient(config),
	}
}
