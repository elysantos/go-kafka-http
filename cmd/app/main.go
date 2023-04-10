package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/elysantos/go-api-messages/internal/infra/akafka"
	"github.com/elysantos/go-api-messages/internal/infra/repository"
	"github.com/elysantos/go-api-messages/internal/infra/web"
	"github.com/elysantos/go-api-messages/internal/usecase"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306/products")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := repository.NewProductRepositoryMysql(db)
	createProductsUsecase := usecase.NewCreateProductUseCase(repository)
	listProdutctUseCase := usecase.NewListProductUseCase(repository)

	productHandlers := web.NewProductHandlers(createProductsUsecase, listProdutctUseCase)

	r := chi.NewRouter()
	r.Post("/products", productHandlers.CreateProductHandler)
	r.Get("/products", productHandlers.ListProductHandler)

	go http.ListenAndServe(":8000", r)

	msgChan := make(chan *kafka.Message)

	go akafka.Consume([]string{"products"}, "host.docker.internal:9094", msgChan)

	for msg := range msgChan {
		request := usecase.CreateProductRequest{}
		err := json.Unmarshal(msg.Value, &request)
		if err != nil {
			log.Error("Failed to unmarshal")
		}

		_, err = createProductsUsecase.Execute(request)
	}

}
