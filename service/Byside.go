package service

import (
	BysideApi "Byside/service/api/Byside"
	"Byside/service/controller/aclCtrl"
	"Byside/service/internal/config"
	"Byside/service/internal/database"
	"context"
	"fmt"
	"go.uber.org/dig"
	"net/http"
)

func Byside() Service {
	once.Do(func() {
		srv = &BysideServer{}
	})

	return srv
}

type BysideServer struct{}

func (srv *BysideServer) Run() {

	container := dig.New()
	srv.provideConfig(container)
	srv.provideService(container)
	srv.provideController(container)
	srv.provideCore(container)

	srv.invokeApiRoutes(container)

	if err := container.Invoke(srv.run); err != nil {
		panic(err)
	}

}

func (srv *BysideServer) provideConfig(container *dig.Container) {

	if err := container.Provide(config.NewByside); err != nil {
		panic(err)
	}
}

func (srv *BysideServer) provideService(container *dig.Container) {
	if err := container.Provide(func() context.Context {
		return context.TODO()
	}); err != nil {
		panic(err)
	}

	if err := container.Provide(database.NewByside); err != nil {
		panic(err)
	}

	if err := container.Provide(BysideApi.NewServer); err != nil {
		panic(err)
	}

	if err := container.Provide(BysideApi.NewRouterRoot); err != nil {
		panic(err)
	}

	if err := container.Provide(BysideApi.NewGinEngine); err != nil {
		panic(err)
	}

}

func (srv *BysideServer) provideController(container *dig.Container) {
	if err := container.Provide(aclCtrl.NewAcl); err != nil {
		panic(err)
	}
}

func (srv *BysideServer) invokeApiRoutes(container *dig.Container) {
	if err := container.Invoke(BysideApi.NewServer); err != nil {
		panic(err)
	}

	if err := container.Invoke(BysideApi.NewGinEngine); err != nil {
		panic(err)
	}
	if err := container.Invoke(BysideApi.NewAcl); err != nil {
		panic(err)
	}
}

func (srv *BysideServer) provideCore(container *dig.Container) {}

func (srv *BysideServer) run(server *http.Server) {
	fmt.Printf("Byside starts at %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
