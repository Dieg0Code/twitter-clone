package middlew

import (
	"net/http"

	"github.com/dieg0code/twitter-clone/db"
)

// CheckDB middleware to check the database status
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Conexion perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
