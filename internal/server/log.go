package server

import (
	"fmt"
	"sync"
)

type Log struct {
	mu sync.Mutex
	records []Record
}
type Record struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"`
}


// create a new instance of our Log
func NewLog() *Log{
	return &Log{}
}

//append a record to our log
func (l *Log) Append(record Record)(uint64, error) {
	//lock our thread
	l.mu.Lock()
	defer l.mu.Unlock()
	record.Offset = uint64(len(l.records))
	l.records = append(l.records,record)
	return record.Offset, nil
}

//read a record from our log
func (l *Log) Read(offset uint64)(Record, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if offset >= uint64(len(l.records)) {
		return Record{}, fmt.Errorf("offset not found")
	}

	return l.records[offset], nil

}

