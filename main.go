package main

import "github.com/nandobas/gohr/pattern"

func main() {
	//lib.LoadFromFile("files/config.json")

	/*err, x1, x2 := formulas.Baskara(2, 10, 8)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	fmt.Printf("x1 = %f", x1)
	fmt.Printf("x2 = %f", x2)*/

	//formulas.Closure()

	/*f := formulas.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}*/

	/*p1 := entities.Person{
		Name: "nando",
		Age:  42,
	}

	p2 := entities.Person{
		Name: "luisa",
		Age:  7,
	}
	fmt.Println(p1, p2)*/

	/*hosts := map[string]entities.IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}*/

	//lib.Pic()
	//formulas.SyncOutput()
	//formulas.TimeAfterOutput()
	//formulas.TimeAfterOutput2()
	//entities.WorkOutput()
	//lib.ReflectOutput()

	//quiz.Area()
	//quiz.LenOffBytes()
	//quiz.IotaOutput()
	//quiz.TimeAddSub()
	//quiz.ChannelOutputDeadLock()
	//quiz.ChannelOutput()
	//quiz.ChannelOutputClose()
	//quiz.SumInt8()
	//quiz.Search()
	//entities.PrintGenerics()
	//iosystem.Serv()
	pattern.TestCarDrive()
}
