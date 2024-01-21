package server

import (
	productpsql "app/internal/product/repository/psql"
	productservice "app/internal/product/transport/http"
	productUC "app/internal/product/usecase"
	"database/sql"

	storagepsql "app/internal/storage/repository/psql"
	storageservice "app/internal/storage/transport/http"
	storageUC "app/internal/storage/usecase"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

type App struct {
	port         string
	server       *rpc.Server
	router       *mux.Router
	databasePool *sql.DB
}

func NewApp(port string) *App {
	databasePool, err := initConnPool()
	if err != nil {
		log.Fatalf("Can't connect to database %v", err)
	}

	// storage service
	storageRepo := storagepsql.NewProductRepository(databasePool)
	storageUseCase := storageUC.NewProductUseCase(storageRepo)
	storageService := storageservice.NewHandler(storageUseCase)

	// product service
	productRepo := productpsql.NewProductRepository(databasePool, storageRepo)
	productUseCase := productUC.NewProductUseCase(productRepo)
	productService := productservice.NewHandler(productUseCase)

	// rpc server
	server := rpc.NewServer()
	server.RegisterCodec(json.NewCodec(), "application/json charset=UTF-8")
	server.RegisterService(productService, "ProductService")
	server.RegisterService(storageService, "StorageService")

	router := mux.NewRouter()
	router.Handle("/rpc", server)

	return &App{
		port:         port,
		router:       router,
		databasePool: databasePool,
	}
}

func (app *App) Run() error {
	defer app.databasePool.Close()
	log.Printf("Web started at %v", app.port)
	err := http.ListenAndServe(":"+app.port, app.router)
	return err
}
