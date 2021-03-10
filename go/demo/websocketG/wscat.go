package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	flagDebug = flag.Bool("debug", true, "Print message type")
)

func recvMeg(conn *websocket.Conn, id int) error {
	timeStart := time.Now()
	timePrint := time.Now()
	var numSum int64
	for {
		t, buf, err := conn.ReadMessage()
		if err == nil {
			if *flagDebug {
				if time.Since(timePrint)/time.Second > 1 {
					timePrint = time.Now()
					log.Printf("id = %d ------ type = %d, len = %d .\r\n", id, t, len(buf))
				}
				if len(buf) > 500 {
					numSum++
					if time.Since(timeStart)/time.Second > 0 {
						//log.Println(numSum / int64(time.Since(timeStart)/time.Second))
					}
				}
				continue
			}
			os.Stdout.Write(buf)
		} else {
			break
		}
	}
	return nil
}

func connect(path string, rd io.Reader, id int) error {
	u, err := url.Parse(path)
	if err != nil {
		panic(err)
	}
	h := http.Header{"Origin": {"http://" + u.Host}}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), h)

	if err != nil {
		panic(err)
	}

	if rd != nil {
		go recvMeg(conn, id)
		b := bufio.NewReader(rd)
		for {
			line, err := b.ReadString('\n')
			if err != nil {
				log.Println(err)
				break
			}
			if err := conn.WriteMessage(websocket.TextMessage, []byte(line)); err != nil {
				break
			}
		}
	} else {
		recvMeg(conn, id)
	}

	conn.Close()
	return nil
}

var u1 = "ws://127.0.0.1:8061/video?interval=3000&url=fake%3A%2F%2FD%3A%5Cwork%5Ctest%5Cmany_high.ts&limit=5&crop=face,full&name=portface1&analyze=true&more=1&extract=1&facemin=48&threshold=180,40,40,0.7&group=1,2,3,4,5,6"
var u2 = "ws://127.0.0.1:8061/video?interval=3000&url=fake%3A%2F%2FD%3A%5Cwork%5Ctest%5Cmany_high.ts&limit=5&crop=face,full&name=portface2&analyze=true&more=1&extract=1&facemin=48&threshold=180,40,40,0.7&group=1,2,3,4,5,6"
var u3 = "ws://127.0.0.1:8061/video?interval=3000&url=fake%3A%2F%2FD%3A%5Cwork%5Ctest%5Cmany_high.ts&limit=5&crop=face,full&name=portface3&analyze=true&more=1&extract=1&facemin=48&threshold=180,40,40,0.7&group=1,2,3,4,5,6"
var u4 = "ws://127.0.0.1:8061/video?interval=3000&url=fake%3A%2F%2FD%3A%5Cwork%5Ctest%5Cmany_high.ts&limit=5&crop=face,full&name=portface4&analyze=true&more=1&extract=1&facemin=48&threshold=180,40,40,0.7&group=1,2,3,4,5,6"
var u5 = "ws://localhost:8061/video_layout"

var u11 = "ws://127.0.0.1:8061/video?interval=3000&url=fake%3A%2F%2FD%3A%5Cwork%5Ctest%5Cmiddle_cropped.ts&limit=5&crop=face,full&name=portface1&analyze=true&more=1&extract=1&facemin=48&threshold=180,40,40,0.7&group=1,2,3,4,5,6"
var u12 = "ws://127.0.0.1:8061/video?interval=3000&url=fake%3A%2F%2FD%3A%5Cwork%5Ctest%5Cmiddle_cropped.ts&limit=5&crop=face,full&name=portface2&analyze=true&more=1&extract=1&facemin=48&threshold=180,40,40,0.7&group=1,2,3,4,5,6"
var u13 = "ws://127.0.0.1:8061/video?interval=3000&url=fake%3A%2F%2FD%3A%5Cwork%5Ctest%5Cmiddle_cropped.ts&limit=5&crop=face,full&name=portface3&analyze=true&more=1&extract=1&facemin=48&threshold=180,40,40,0.7&group=1,2,3,4,5,6"
var u14 = "ws://127.0.0.1:8061/video?interval=3000&url=fake%3A%2F%2FD%3A%5Cwork%5Ctest%5Cmiddle_cropped.ts&limit=5&crop=face,full&name=portface4&analyze=true&more=1&extract=1&facemin=48&threshold=180,40,40,0.7&group=1,2,3,4,5,6"

var t1 = "ws://127.0.0.1:8061/video_offline?interval=3000&url=D%3A%5Cwork%5Ctest%5Cmiddle_cropped.ts&limit=5&crop=face,full&name=offline1&analyze=true&more=1&extract=1&facemin=48&threshold=180,40,40,0.7&group=1,2,3,4,5,6"

var t2 = "ws://127.0.0.1:8061/video_offline?interval=3000&url=D%3A%5Cwork%5Ctest%5Cmany_high.ts&limit=5&crop=face,full&name=offline2&analyze=true&more=1&extract=1&facemin=48&threshold=180,40,40,0.7&group=1,2,3,4,5,6"

//websocket.exe
func main() {
	// s, _ := url.QueryUnescape("fake%3A%2F%2FD%3A%5Cwork%5Ctest%5Cmiddle_cropped.ts")
	// s = url.QueryEscape("D:\\work\\test\\middle_cropped.ts")
	// fmt.Println(s)

	fmt.Println("in")
	flag.Parse()
	fmt.Println(flag.Arg(0))

	var wg sync.WaitGroup
	if flag.Arg(0) == "z" {
		wg.Add(4)
		go connect(u11, nil, 1)
		go connect(u12, nil, 2)
		go connect(u13, nil, 3)
		go connect(u14, nil, 4)
		//go connect(u5, os.Stdin, 5)
	} else if flag.Arg(0) == "h" {
		wg.Add(4)
		go connect(u1, nil, 1)
		go connect(u2, nil, 2)
		go connect(u3, nil, 3)
		go connect(u4, nil, 4)
		//go connect(u5, os.Stdin, 5)
	} else if flag.Arg(0) == "t" {
		wg.Add(2)
		go connect(t1, nil, 1)
		go connect(t2, nil, 1)
	} else {
		panic("para is error")
	}
	wg.Wait()
	fmt.Println("end")
}
