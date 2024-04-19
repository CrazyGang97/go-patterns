package factory

import (
	"io"
)

type Store interface {
	Open(string) (io.ReadWriteCloser, error)
}

type StorageType int

const (
	DiskStorageType StorageType = 1 << iota
	MemoryStorageType
)

type Disk struct {
}

func (d *Disk) Open(string) (io.ReadWriteCloser, error) {
	return nil, nil
}

type Memory struct {
}

func (m *Memory) Open(string) (io.ReadWriteCloser, error) {
	return nil, nil
}

func NewStore(t StorageType) Store {
	switch t {
	case DiskStorageType:
		return &Disk{}
	case MemoryStorageType:
		return &Memory{}
	default:
		return nil
	}
}
