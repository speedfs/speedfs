package client

type TrackerClient struct {
	config *Config
}

func NewTrackerClient(config *Config) *TrackerClient {
	return &TrackerClient{
		config: config,
	}
}
