package shutdown

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// 优雅退出

type Shutdown struct {
	ch      chan os.Signal
	timeout time.Duration
}

func New(t time.Duration) *Shutdown {
	return &Shutdown{
		ch:      make(chan os.Signal),
		timeout: t,
	}
}

func (s *Shutdown) Add(signals ...os.Signal) {
	signal.Notify(s.ch, signals...)
}

func (s *Shutdown) Start(server *http.Server) {
	<-s.ch
	fmt.Println("start exist......")

	ctx, cannel := context.WithTimeout(context.Background(), s.timeout*time.Second)
	defer cannel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Graceful exit failed. err: ", err)
	}
	fmt.Println("Graceful exit success.")
}
