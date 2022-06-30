package main

import (
	"github.com/15matias15/go-basic-udemy/controller"
	"github.com/15matias15/go-basic-udemy/model"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	log.Fatal(http.ListenAndServe("localhost:3000", mux))

}

//func swap(m1, m2 *int) {
//	var temp int
//	temp = *m2
//	*m2 = *m1
//	*m1 = temp
//}
//func main() {
//	m1, m2 := 2, 3
//	fmt.Println(m1, m2)
//	swap(&m1, &m2)
//	fmt.Println(m1, m2)
//}
//
//func main() {
//
//	c := make(chan int)
//
//	go func() {
//		c <- 1
//	}()
//
//	val := <-c
//	fmt.Println(val)
//
//	go func() {
//		c <- 2
//	}()
//
//	time.Sleep(time.Second * 2)
//	val = <-c
//	fmt.Println(val)
//}

//type Car struct {
//	Model string
//}
//
//func main() {
//
//	c := make(chan *Car, 3)
//
//	go func() {
//		c <- &Car{"a"}
//		c <- &Car{"b"}
//		c <- &Car{"c"}
//		c <- &Car{"d"}
//		c <- &Car{"e"}
//		close(c)
//	}()
//
//	for i := range c {
//		fmt.Println(i.Model)
//	}
//}

//func pinger(pinger <-chan int, ponger chan<- int) {
//	for {
//		<-pinger
//		fmt.Println("ping")
//		time.Sleep(time.Second)
//		ponger <- 1
//	}
//}
//
//func ponger(pinger chan<- int, ponger <-chan int) {
//	for {
//		<-ponger
//		fmt.Println("pong")
//		time.Sleep(time.Second)
//		pinger <- 1
//	}
//}
//
//func main() {
//	ping := make(chan int)
//	pong := make(chan int)
//
//	go pinger(ping, pong)
//	go ponger(ping, pong)
//
//	ping <- 1
//	//
//	//for {
//	//	time.Sleep(time.Second)
//	//}
//
//	select {}
//}
