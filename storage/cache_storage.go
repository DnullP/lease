package storage

type DataWithTTL struct {
	Data string
	DDL  int64
}

type Cache map[string]DataWithTTL

var NodeCache Cache
