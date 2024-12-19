package main

import (
	"base-practice/flowy"
	"log"
)

func main() {
	//_chan.Run()
	//_chan.RunCase()
	//_chan.Run2()
	//_chan.Run3()

	//_json.RunStruct2Json()
	//_json.RunJson2Struct()
	//_json.Nested()

	//_ctx.Run()

	//flag.RunFlags()

	//log.RunNormal()
	//log.RunError()
	//log.RunCustom()

	//_sync.RunOnce()
	//_sync.RunAtomic()
	//_sync.RunChan()

	//bufBytes.Writer()
	//bufBytes.Reader()
	//bufBytes.Reset()
	//bufBytes.Print()

	p, err := flowy.GetUserDataDir()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(p)
}
