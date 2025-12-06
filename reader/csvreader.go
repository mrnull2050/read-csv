package reader

import (
	"encoding/csv"
	"os"
	"poj/model"
	"strconv"
)

func ReadUserCSVFormat(path string) ([]model.UserModel, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	var users []model.UserModel
	for i := 1; i < len(rows); i++ {
		line := rows[i]

		id, _ := strconv.Atoi(line[0])
		name := line[1]
		email := line[2]

		users = append(users, model.UserModel{
			ID:    id,
			Name:  name,
			Email: email,
		})
	}

	return users, nil
}
