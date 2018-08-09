/*获取多个URL内容*/
package main

import (
	"time"
	"os"
	"net/http"
	"io/ioutil"
	"fmt"
	"io"
	)

func main() {
	start := time.Now()
	ch := make(chan string )
	for _ , url := range os.Args[1:]{
		go fetch(url,ch)
	}
	for range os.Args[1:]{
		fmt.Printf(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n",time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string){
	start := time.Now()
	sep,_ := http.Get(url)
	b,_ := io.Copy(ioutil.Discard,sep.Body)
	sep.Body.Close()
	secs := time.Since(start).Seconds()
	ch<- fmt.Sprintf("%.2fs  %7d  %s",secs,b,url)
}
