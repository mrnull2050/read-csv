package main

import (
	"context"
	"database/sql"

	_ "modernc.org/sqlite"

	"fmt"

	_ "github.com/mattn/go-sqlite3"

	repo "poj/Repo"
	"poj/model"
	"poj/reader"
	"time"
)

func main() {
	users, err := reader.ReadUserCSVFormat("data/data.csv")
	if err != nil {
		panic(err)
	}
	for _, u := range users {
		fmt.Println(u)
	}
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		panic(err)
	}
	repo := repo.UserRepo{DB: db}

	users, _ = reader.ReadUserCSVFormat("data/data.csv")

	jobs := make(chan Job, len(users))
	results := make(chan error, len(users))

	numWorkers := 10
	for i := 0; i < numWorkers; i++ {
		go StartWorking(&repo, jobs, results)
	}

	for _, u := range users {
		jobs <- Job{User: u}
	}
	close(jobs)
	for i := 0; i < len(users); i++ {
		if err := <-results; err != nil {
			fmt.Println("❌ Error:", err)
		}
	}
	fmt.Println("ALL USERS SAVED SUCCESSFULLY")

}

type Job struct {
	User model.UserModel
}

func StartWorking(repo *repo.UserRepo, jobs <-chan Job, result chan<- error) {
	ctx, err := context.WithTimeout(context.Background(), 60*time.Second)
	if err != nil {
		fmt.Println("error while start working")
	}
	for job := range jobs {
		err := repo.Save(ctx, job.User)
		result <- err
	}
}
