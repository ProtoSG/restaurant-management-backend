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
	stmt, err := r.db.Prepare("INSERT INTO users (username, password, role) VALUES (?, ?, ?)")
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

func (r *sqliteUserRepository) GetAll() ([]*domain.User, error) {
	stmt, err := r.db.Prepare("SELECT * FROM user")
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
			fmt.Println("Error scanning row:", err)
			return nil, err
		}

		userId, _ := types.NewUserId(user.Id)
		userUsername, _ := types.NewUserUsername(user.Username)
		userPassword, _ := types.NewUserPassword(user.Password)
		userRole, _ := types.NewUserRole(user.Role)

		newUser := domain.NewUser(userId, userUsername, userPassword, userRole)

		users = append(users, newUser)
		fmt.Println(user.Id, user.Username)
	}

	fmt.Println(users)

	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
	}

	return users, nil
}

func (r *sqliteUserRepository) GetById(userId *types.UserId) (*domain.User, error) {
	stmt, err := r.db.Prepare("SELECT id, username, password, role FROM users WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(userId.Value)

	var id int
	var username string
	var password string
	var role string

	err = row.Scan(&id, &username, &password, &role)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	userUsername, _ := types.NewUserUsername(username)
	userPassword, _ := types.NewUserPassword(password)
	userRole, _ := types.NewUserRole(role)

	user := domain.NewUser(userId, userUsername, userPassword, userRole)

	return user, nil
}

func (r *sqliteUserRepository) Edit(user *domain.User) error {
	stmt, err := r.db.Prepare("UPDATE users SET username = ?, password = ?, role = ? WHERE id = ?")
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
	stmt, err := r.db.Prepare("DELETE FROM users WHERE id = ?")
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
