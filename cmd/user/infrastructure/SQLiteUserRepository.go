package infrastructure

import (
	"database/sql"
	"fmt"
	"os"
	"restaurant-management-backend/cmd/user/domain"
	"restaurant-management-backend/cmd/user/domain/types"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type sqliteUserRepository struct {
	db *sql.DB
}

func NewSQLiteUserRepository(url string) *sqliteUserRepository {
	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}

	return &sqliteUserRepository{
		db: db,
	}
}

func (r *sqliteUserRepository) Create(user *domain.User) error {
	stmt, err := r.db.Prepare("INSERT INTO user (username, password, role) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username.Value, user.Password.Value, user.Role.Value)
	if err != nil {
		return err
	}

	return nil
}

func (this *sqliteUserRepository) GetAll() ([]*domain.User, error) {
	stmt, err := this.db.Prepare("SELECT * FROM user")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User

	for rows.Next() {
		var user domain.UserPrimitive

		if err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Role); err != nil {
			return nil, err
		}

		userId, _ := types.NewUserId(user.Id)
		userUsername, _ := types.NewUserUsername(user.Username)
		userPassword, _ := types.NewUserPassword(user.Password)
		userRole, _ := types.NewUserRole(user.Role)

		newUser := domain.NewUser(userId, userUsername, userPassword, userRole)

		users = append(users, newUser)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
	}

	return users, nil
}

func (r *sqliteUserRepository) GetById(userId *types.UserId) (*domain.User, error) {
	stmt, err := r.db.Prepare("SELECT * FROM user WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(userId.Value)

	var user domain.UserPrimitive

	err = row.Scan(&user.Id, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}

	userUsername, _ := types.NewUserUsername(user.Username)
	userPassword, _ := types.NewUserPassword(user.Password)
	userRole, _ := types.NewUserRole(user.Role)

	newUser := domain.NewUser(userId, userUsername, userPassword, userRole)

	return newUser, nil
}

func (r *sqliteUserRepository) Edit(user *domain.User) error {
	stmt, err := r.db.Prepare("UPDATE user SET username = ?, password = ?, role = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username.Value, user.Password.Value, user.Role.Value, user.Id.Value)
	if err != nil {
		return err
	}

	return nil
}

func (r *sqliteUserRepository) Delete(id *types.UserId) error {
	stmt, err := r.db.Prepare("DELETE FROM user WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id.Value)
	if err != nil {
		return err
	}

	return nil
}
