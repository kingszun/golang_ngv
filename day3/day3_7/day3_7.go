package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func PrintEverySecond(ctx context.Context) {
	//일정한 주기를 가지고 작업 수행
	tick := time.Tick(time.Second)
	for {
		select {
		case <-tick: //시그널은 신호만 있음, 채널은 데이터, 신호 둘다 있음
			fmt.Println("일정한 주기를 갖는 작업")
		case <-ctx.Done(): //context를 만들 때 리턴된 함수를 호출하면 수행
			wg.Done()
			return
		}
	}
}

func main() {
	//수행중인 goroutine의 개수를 1증가
	wg.Add(1)

	//goroutine을 종료할 수 있는 시그널을 전송하는 context 생성
	// ctx, cancle := context.WithCancel(context.Background())

	//시간동한 수행
	ctx, cancle := context.WithTimeout(
		context.Background(), 3*time.Second)

	//고루틴 실행
	go PrintEverySecond(ctx)
	//5초 대기
	time.Sleep(5 * time.Second)

	cancle()
	//수행중인 goroutine의 개수가 0보다 크면 대기하고 0이되면 수행
	wg.Wait()
}
