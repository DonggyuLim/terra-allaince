package main

import (
	"sync"

	"github.com/DonggyuLim/Alliance-Rank/data"
	"github.com/DonggyuLim/Alliance-Rank/db"
	"github.com/DonggyuLim/Alliance-Rank/rest"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	db.Connect()
	defer db.Close()
	db.NewDB()
	defer db.BadgerClose()

	go data.Main(wg)
	//  data.Main(wg,Harkonnen)
	//  data.Main(wg,CORRINO)
	//  data.Main(wg,ORDOS)

	go rest.Start(wg)
	wg.Wait()
}
