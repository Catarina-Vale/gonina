package repository

import (
    "database/sql"
    "yourapp/model"
)

type UserRepo struct {
    DB *sql.DB
}

func (repo *UserRepo) GetAll() ([]model.User, error) {
    rows, err := repo.DB.Query("SELECT id, email, name, age FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []model.User
    for rows.Next() {
        var u model.User
        err := rows.Scan(&u.ID, &u.Email, &u.Name, &u.Age)
        if err != nil {
            return nil, err
        }
        users = append(users, u)
    }
    return users, nil
}

func (repo *UserRepo) Create(user model.User) error {
    _, err := repo.DB.Exec("INSERT INTO users (email, name, age) VALUES ($1, $2, $3)",
        user.Email, user.Name, user.Age)
    return err
}

func (repo *UserRepo) Update(user model.User) error {
    _, err := repo.DB.Exec("UPDATE users SET email=$1, name=$2, age=$3 WHERE id=$4",
        user.Email, user.Name, user.Age, user.ID)
    return err
}

func (repo *UserRepo) Delete(id int) error {
    _, err := repo.DB.Exec("DELETE FROM users WHERE id=$1", id)
    return err
}
