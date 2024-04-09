package repositories

import (
	"database/sql"
	"email-verification/config"
	"email-verification/entities"
	errorHandlers "email-verification/errorHandlers"
	"log"
)

type CompRepositories interface {
	InsertToken(email string, token string) error
	GetUser(email string) (*entities.Users, error)
	VerifyEmail(email string) error
	RegistUser(email string, password string) error
}

type compRepositories struct {
	DB *sql.DB
}

type CustomError struct {
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

func NewCompRepositories(DB *sql.DB) *compRepositories {
	db := config.InitDB()

	_, err := db.Exec(`
    BEGIN TRY
        CREATE TABLE Users (
            Id INT IDENTITY(1,1) PRIMARY KEY,
            Email NVARCHAR(255) NOT NULL UNIQUE,
			Password VARCHAR(MAX) NOT NULL,
			IsVerified BIT DEFAULT 0,
            Token VARCHAR(6) DEFAULT 0,
            CreatedAt DATE DEFAULT CURRENT_TIMESTAMP
        )
		END TRY
		BEGIN CATCH
			-- Ignore the error if the table already exists
		END CATCH
		`)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	return &compRepositories{
		DB: DB,
	}
}

func (r *compRepositories) InsertToken(email string, token string) error {
	_, err := r.DB.Exec("UPDATE Users SET Token = @p1 WHERE Email = @p2", token, email)
	if err != nil {
		return err
	}

	return nil
}

func (r *compRepositories) GetUser(email string) (*entities.Users, error) {
    rows, err := r.DB.Query("SELECT * FROM Users WHERE Email = @p1", email)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var user *entities.Users

    for rows.Next() {
        var d entities.Users
        err = rows.Scan(&d.Id, &d.Email, &d.Password, &d.IsVerified, &d.Token, &d.CreatedAt)
        if err != nil {
            return nil, err
        }

        user = &d
    }

    return user, nil
}

func (r *compRepositories) VerifyEmail(email string) error {
	_, err := r.DB.Exec("UPDATE Users SET IsVerified = 'true' WHERE Email = @p1", email)
	if err != nil {
		return err
	}

	return nil
}

func (r *compRepositories) RegistUser(email string, password string) error {
	_, err := r.DB.Exec("INSERT INTO Users (Email, Password) VALUES(@p1, @p2)", email, password)

	

	if err != nil {
		err = &errorHandlers.CustomError{
			Message: "Email already registered!",
		}
		return err
	}
	
	return nil
}
