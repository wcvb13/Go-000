package di

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	v1 "projecttest/api/app/user/v1"
	"projecttest/internal/app/service"
	"syscall"
)

type App struct {
	server *service.UserService
}

func NewApp(srv *service.UserService) *App  {
	return &App{server: srv}
}

func (app *App) Start() error {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		return err
	}
	gs := grpc.NewServer()
	v1.RegisterUserServiceServer(gs, app.server)
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		go func() {
			<-ctx.Done()
			gs.GracefulStop()
			fmt.Println("Shutdown grpc server.")
		}()
		return gs.Serve(listener)
	})
	g.Go(func() error {
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
	})
	return g.Wait()
}
