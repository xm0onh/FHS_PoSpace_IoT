package main

import (
	"github.com/xm0onh/FHS_PoSpace_IoT"
	"github.com/xm0onh/FHS_PoSpace_IoT/benchmark"
	"github.com/xm0onh/FHS_PoSpace_IoT/db"
)

// Database implements FHS_PoSpace_IoT.DB interface for benchmarking
type Database struct {
	FHS_PoSpace_IoT.Client
}

func (d *Database) Init() error {
	return nil
}

func (d *Database) Stop() error {
	return nil
}

func (d *Database) Write(k int, v []byte) error {
	key := db.Key(k)
	err := d.Put(key, v)
	return err
}

func main() {
	FHS_PoSpace_IoT.Init()

	d := new(Database)
	d.Client = FHS_PoSpace_IoT.NewHTTPClient()
	b := benchmark.NewBenchmark(d)
	b.Run()
}
