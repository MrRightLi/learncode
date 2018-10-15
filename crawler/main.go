package main

import (
	"project/learncode/crawler/engine"
	"project/learncode/crawler/persist"
	"project/learncode/crawler/scheduler"
	"project/learncode/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
		ItemChan: persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	},
	)
}
