package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/salamandery/imersao12-go-esquenta/internal/infra/akafka"
	"github.com/salamandery/imersao12-go-esquenta/internal/infra/repository"
	"github.com/salamandery/imersao12-go-esquenta/internal/infra/web"
	"github.com/salamandery/imersao12-go-esquenta/internal/usecase"
)

func main() {
	fmt.Println("Aplicação rodando...")
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/products")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := repository.NewProductRepositoryMysql(db)
	createProductUsecase := usecase.NewCreateProductUseCase(repository)
	listProductsUseCase := usecase.NewListProductsUseCase(repository)

	productsHandler := web.NewProductHandlers(createProductUsecase, listProductsUseCase)

	r := chi.NewRouter()
	r.Post("/products", productsHandler.CreateProductHandler)
	r.Get("/products", productsHandler.ListProductHandler)

	go http.ListenAndServe(":8000", r)

	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"products"}, "host.docker.internal:9094", msgChan)

	for msg := range msgChan {
		dto := usecase.CreateProductInputDto{}
		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {
			panic(err)
		}
		_, err = createProductUsecase.Execute(dto)
		fmt.Println(string(msg.Value))
	}
}
