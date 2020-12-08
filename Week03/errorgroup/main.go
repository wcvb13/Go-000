package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return startHttpServer(ctx,":8080",&serverHandler{})
	})
	g.Go(func() error {
		return handelSignal(ctx)
	})
	err := g.Wait()
	if err != nil {
		fmt.Printf("quit group with err: \"%s\"",err)
	}
}

// http server 的启动和关闭
func startHttpServer(ctx context.Context,addr string, handle http.Handler) error {
	s := http.Server{
		Addr:    addr,
		Handler: handle,
	}

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("http server quite")
			s.Shutdown(context.Background())
		}
	}()
	fmt.Println("http server start")
	return s.ListenAndServe()
}

// linux signal 信号的注册和处理
func handelSignal(ctx context.Context) error {
	ch := make(chan os.Signal)
	signal.Notify(ch,syscall.SIGINT,syscall.SIGTERM)
	fmt.Println("signal catch start")
	select {
	case s := <-ch:
		fmt.Println("catch system signal, quit")
		return fmt.Errorf("quit with signal: %s",s)
	case <-ctx.Done():
		return fmt.Errorf("quit all group")
	}
}

type serverHandler struct {
}

func (s *serverHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("Hello word!"))
}