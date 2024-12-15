package models

import (
	"context"
	"database/sql"
	"time"
)

// DBModel is the type for database connection values
type DBMOdel struct {
	DB *sql.DB
}

// Models is the wrapper for all models
type Models struct {
	DB DBMOdel
}

// Returns a model type with adabase connection pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBMOdel{DB: db},
	}
}

// Widget is the type for all widgets
type Widget struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	InventoryLevel int       `json:"inventory_level"`
	Price          int       `json:"price"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
}

func (m *DBMOdel) GetWidget(id int) (Widget, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var widget Widget

	row := m.DB.QueryRowContext(ctx, "select id, name from widgets where id = ?", id)
	err := row.Scan(&widget.ID, &widget.Name)
	if err != nil {
		return widget, err
	}

	return widget, nil
}
