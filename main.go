package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type Categoria struct {
	ID     string `json:"categoria_id"`
	Nombre string `json:"nombre"`
}

func getAlbums(db *sql.DB) []Categoria {
	rows, err := db.Query("SELECT * FROM Categoria")

	if err != nil {
		fmt.Printf("Error querying database: %q\n", err)
		os.Exit(1)
	}

	defer rows.Close()

	var categorias []Categoria

	var categoria Categoria

	for rows.Next() {
		if err := rows.Scan(&categoria.ID, &categoria.Nombre); err != nil {
			fmt.Printf("Error scanning row: %q\n", err)
			return nil
		}

		categorias = append(categorias, categoria)
		fmt.Printf("ID: %s, Nombre: %s\n", categoria.ID, categoria.Nombre)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("Error iterating rows: %q\n", err)
		return nil
	}

	return categorias
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	DATABASE_URL := os.Getenv("DATABASE_URL")

	if DATABASE_URL == "" {
		fmt.Println("DATABASE_URL is not set")
		os.Exit(1)
	}

	TOKEN := os.Getenv("TOKEN")

	if TOKEN == "" {
		fmt.Println("TOKEN is not set")
		os.Exit(1)
	}

	url := DATABASE_URL + "?authToken=" + TOKEN
	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Printf("Error opening database: %q\n", err)
		os.Exit(1)
	}

	defer db.Close()

	router := gin.Default()
	router.GET("/categorias", func(c *gin.Context) {
		categorias := getAlbums(db)
		c.JSON(200, categorias)
	})
	router.Run("localhost:8080")
}
