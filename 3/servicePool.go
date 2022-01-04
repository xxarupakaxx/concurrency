package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"testing"
	"time"
)


func ConnectToSercice() interface{}{
	time.Sleep(1 *time.Second)
	return struct{}{}
}

func startNetWorkDaemon()  *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func ()  {
		server,err := net.Listen("tcp","localhost:1111")
		if err!=nil {
			log.Fatal(err)
		}
		defer server.Close()

		wg.Done()

		for{
			conn,err := server.Accept()
			if err!=nil {
				log.Print(err)
				continue
				
			}

			ConnectToSercice()
			fmt.Fprintln(conn,"")
			conn.Close()
		}
	}()

	return &wg


}

func init()  {
	daemonStarted := startNetWorkDaemon()
	daemonStarted.Wait()
	
}
func BenchmarkNetworkRequest(b *testing.B){
	for i := 0; i < b.N; i++ {
		conn,err := net.Dial("tcp","localhost:1111")
		if err!=nil {
			b.Fatalf("cannnot dial host %v",err)
			
		}

		if _,err := ioutil.ReadAll(conn);err !=nil{
			b.Fatalf("cannot read: %v",err)
		}

		conn.Close()
	}
}

func warmServiceDeamon() *sync.Pool {
	p := &sync.Pool{
		New: ConnectToSercice,
	}
	for i := 0; i < 10; i++ {
		p.Put(p.New())
	}

	return p
}

func startNewtDameown() *sync.WaitGroup  {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		connPool := warmServiceDeamon()
		server,_ :=net.Listen("tcp","localhost:1111")
		defer server.Close()

		wg.Done()

		for{
			conn,_ := server.Accept()
			svcConn := connPool.Get()
			fmt.Fprintln(conn,"")
			connPool.Put(svcConn)
			conn.Close()
		}
	
	}()
	return &wg
}