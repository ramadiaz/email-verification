package repositories

import (
	"database/sql"
	"email-verification/config"
	"log"
)

type CompRepositories interface {
	InsertToken(email string, token string) error
}

type compRepositories struct {
	DB *sql.DB
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
            Token VARCHAR(6),
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