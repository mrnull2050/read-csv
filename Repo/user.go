package repo

import (
	"context"
	"database/sql"
	"poj/model"
)

type UserRepo struct {
	DB *sql.DB
}

func (r *UserRepo) Save(ctx context.Context, u model.UserModel) error {
	Query := "INSERT INTO usersN (name, email) VALUES (?, ?)"
	_, err := r.DB.ExecContext(ctx, Query, u.Name, u.Email)
	return err
}
func (r *UserRepo) GetAll(ctx context.Context) ([]model.UserModel, error) {
	rows, err := r.DB.Query("SELECT WHERE userN (name , email)")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.UserModel
	for rows.Next() {
		var u model.UserModel
		rows.Scan(&u.ID, &u.Name, &u.Email)
		users = append(users, u)
	}
	return users, nil
}
func(r *UserRepo) GetUserById(id int) (model.UserModel , error){
	var u model.UserModel
	err := r.DB.QueryRow(`SELECT id , name , email FROM usersN WHERE id=$1`, id).Scan(&u.ID, &u.Name, &u.Email)
	
	return u , err 

}

func(r *UserRepo) Update(u model.UserModel) error{
	_, err := r.DB.Exec(`UPDATE users SET name=$1, email=$2, age=$3 WHERE id=$4
    `, u.Name, u.Email, u.ID)

	return err
}
func (r *UserRepo) Delete(id int) error {
	_, err := r.DB.Exec(`DELETE FROM users WHERE id=$1`, id)
	return err
}
