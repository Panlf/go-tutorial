package main

func main() {

	//进入等待状态
	/*
	a := make(chan bool, 100)
	b := make(chan bool, 100)
	go func() {
		time.Sleep(time.Minute)
		for i := 0; i < 10; i++ {
			a <- true
			b <- true
		}
	}()

	for i := 0; i < 10; i++ {
		select {
		case <-a:
			print("< a")
		case <-b:
			print("< b")
		}
	}*/

}