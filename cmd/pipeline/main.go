package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bosamatheus/go-etl-pipeline/configs"
	"github.com/bosamatheus/go-etl-pipeline/internal/infrastructure/repository"
	"github.com/bosamatheus/go-etl-pipeline/internal/usecase/pipeline/extractor"
	"github.com/bosamatheus/go-etl-pipeline/internal/usecase/pipeline/job"
	"github.com/bosamatheus/go-etl-pipeline/internal/usecase/pipeline/loader"
	"github.com/bosamatheus/go-etl-pipeline/internal/usecase/pipeline/publisher"
	"github.com/bosamatheus/go-etl-pipeline/internal/usecase/pipeline/task"
	"github.com/bosamatheus/go-etl-pipeline/internal/usecase/pipeline/transformer"
)

func main() {
	db := newDB()
	defer db.Close()

	reviewRepo := repository.NewReviewPostgres(db)
	filenameExtract := "data/olist_order_reviews_dataset.csv"
	filenameLoad := "tmp/pipeline/reviews/temp.csv"

	task := newTask(filenameExtract, filenameLoad)
	publisher := publisher.NewPostgresPublisher(filenameLoad, true, &reviewRepo)
	job := job.NewDefaultJob(task, publisher)

	job.Launch()
}

func newDB() *sql.DB {
	connString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true",
		configs.UserDB,
		configs.PasswordDB,
		configs.HostDB,
		configs.NameDB,
	)
	db, err := sql.Open(configs.DriverNameDB, connString)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

func newTask(filenameExtract, filenameLoad string) task.Task {
	extractor := extractor.NewCSVExtractor(filenameExtract)
	transformer := transformer.NewReviewTransformer()
	loader := loader.NewCSVLoader(filenameLoad)
	return task.NewDefaultTask(extractor, transformer, loader)
}
