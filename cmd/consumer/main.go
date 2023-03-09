package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/YvesJaques/gointensivo2/internal/infra/database"
	"github.com/YvesJaques/gointensivo2/internal/usecase"
	"github.com/YvesJaques/gointensivo2/pkg/kafka"
)

func main() {
	db, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := database.NewOrderRepository(db)
	usecase := usecase.CalculateFinalPrice{OrderRepository: repository}
	msgChanKafka := make(chan *kafka.Message)
}
