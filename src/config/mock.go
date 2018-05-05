/*
 * Radon
 *
 * Copyright 2018 The Radon Authors.
 * Code is licensed under the GPLv3.
 *
 */

package config

var (
	// MockSchemaConfig config.
	MockSchemaConfig = &SchemaConfig{
		DB:     "sbtest",
		Tables: MockTablesConfig,
	}

	// MockTablesConfig config.
	MockTablesConfig = []*TableConfig{
		&TableConfig{
			Name:       "A",
			ShardKey:   "id",
			Partitions: MockPartitionAConfig,
		},
		&TableConfig{
			Name:       "B",
			ShardKey:   "id",
			Partitions: MockPartitionBConfig,
		},
	}

	// MockPartitionAConfig config.
	MockPartitionAConfig = []*PartitionConfig{
		&PartitionConfig{
			Table:   "A1",
			Segment: "0-2",
			Backend: "backend1",
		},
		&PartitionConfig{
			Table:   "A2",
			Segment: "2-4",
			Backend: "backend1",
		},
		&PartitionConfig{
			Table:   "A3",
			Segment: "4-8",
			Backend: "backend2",
		},
		&PartitionConfig{
			Table:   "A4",
			Segment: "8-16",
			Backend: "backend2",
		},
	}

	// MockPartitionBConfig config.
	MockPartitionBConfig = []*PartitionConfig{
		&PartitionConfig{
			Table:   "B1",
			Segment: "0-4",
			Backend: "backend2",
		},
		&PartitionConfig{
			Table:   "B2",
			Segment: "4-8",
			Backend: "backend1",
		},
		&PartitionConfig{
			Table:   "B3",
			Segment: "8-16",
			Backend: "backend2",
		},
	}

	// MockBackends config.
	MockBackends = []*BackendConfig{
		&BackendConfig{
			Name:           "backend1",
			Address:        "127.0.0.1:3304",
			User:           "root",
			Password:       "",
			MaxConnections: 1024,
		},
	}

	// MockBackup config.
	MockBackup = &BackendConfig{
		Name:           "backupnode",
		Address:        "127.0.0.1:3304",
		User:           "root",
		Password:       "",
		MaxConnections: 1024,
	}

	// MockProxyConfig config.
	MockProxyConfig = &ProxyConfig{
		Endpoint:            ":5566",
		MaxConnections:      1024,
		MetaDir:             "/tmp/radonmeta",
		PeerAddress:         ":8080",
		BackupDefaultEngine: "TokuDB",
	}

	// MockLogConfig config.
	MockLogConfig = &LogConfig{
		Level: "DEBUG",
	}
)
