// package main

// import (
// 	"core/framework"
// 	"net/http"
// )

// func main() {
// 	server := &http.Server{
// 		// 自定义的请求核心处理函数
// 		Handler: framework.NewCore(),
// 		// 请求监听地址
// 		Addr: ":8080",
// 	}
// 	server.ListenAndServe()

// }

package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 1 * time.Millisecond

func main() {
	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

}
