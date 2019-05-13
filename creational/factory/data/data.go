package data

import "io"

type Store interface {
	Open(string) (io.ReadWriteCloser, error)
}

type StorageType int

const (
	DiskStorage StorageType = 1 << iota
	TempStorage
	MemoryStorage
)

func NewStore(t StorageType) Store {
	switch t {
	case MemoryStorage:
		memory := Memory{}
		return &memory
	case DiskStorage:
		memory := Memory{}
		return &memory
	default:
		temp := Temp{}
		return &temp
	}
}

type Memory struct {
}

func (m *Memory) Open(str string) (io.ReadWriteCloser, error) {
	return nil, nil
}

type Disk struct {
}

func (m *Disk) Open(str string) (io.ReadWriteCloser, error) {
	return nil, nil
}

type Temp struct {
}

func (m *Temp) Open(str string) (io.ReadWriteCloser, error) {
	return nil, nil
}
