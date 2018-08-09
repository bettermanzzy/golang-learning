package main

import "fmt"

type st struct {
	name string
}

type request struct {
	data []int
	rep chan int
}

func newrequest(data...int)*request{
	return &request{data,make(chan int ,1)}
}
func process(req *request){
	sum :=0
	for _,x :=range req.data{
		sum = sum +x
	}
	req.rep<-sum
}

func main(){
	chanint :=make(chan int ,10)
	chanint <- 10
	 a := <- chanint
	fmt.Println(a)

	 chanmap := make(chan map[string]string,10)
	 n :=map[string]string{"1":"lnxdle","2":"qbkdk","3":"bdkjwkd"}
	 m := make(map[string]string,20)
	 m["a"]="nxdsw"
	 m["b"]="ndjw"
	 m["c"]="ndwlhn"
	 chanmap <- m
	 chanmap <- n
	 b :=<-chanmap
	 d :=<-chanmap
	 fmt.Println(b,d)

	 chanstu := make(chan st , 10)
	 a1 := st{ "zhangsan",}
	 chanstu <- a1
	 b1 := <- chanstu
	 fmt.Println(b1)

	 chanstuid := make(chan *st,10)
	 a2 :=&st{"12365456"}
	 chanstuid <- a2
	 b2 := <- chanstuid
	 fmt.Println(*b2)

	 chaninter := make(chan interface{},10)
	 a3 := st{"123"}
	 chaninter  <- &a3
	 a4 := <- chaninter
	 fmt.Println("*********",a4)

	 var stt *st
	 stt ,ok :=a4.(*st)
	 if !ok{
	 	fmt.Println("cannot convert")
		 return
	 }
	 fmt.Println(*stt)

	 ch :=make(chan int ,11)
	 ch<-99
	 for i:=0;i<10;i++{
	 	ch <- i
	 }
	 fmt.Println(ch)
	 first_chan,ok := <- ch
	 if ok {
	 	fmt.Println("the first chan is :",first_chan)
	 }
	 ch<-10
	 for value :=range ch {
	 	fmt.Println(value)
	 	if value == 10{
			close(ch)
		}
	 }
	 fmt.Println("after range or close ch!")

	 //chan作为结构体成员进行，也可传参
	 ptr :=newrequest(10,20,30)
	 process(ptr)
	 fmt.Println(<-ptr.rep)

	 //单向管道，一端读一端写
	 chnn :=make(chan int ,10)
	 var send chan<- int = chnn
	 var recv <-chan int = chnn

	 send <- 1
	 val,ok :=<-recv
	 if ok {
	 	fmt.Println(val)
	 }
}
