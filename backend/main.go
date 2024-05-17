package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/", fs)

	// Route for Excel template downloas
	http.HandleFunc("/download-template", func(w http.ResponseWriter, r *http.Request) {
		filePath := "./data/templates/template.xlsx"
		fileName := "template.xlsx"
		w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
		w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		http.ServeFile(w, r, filePath)
	})

	// Route for file upload
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Créez un dossier pour stocker les fichiers uploadés
		os.MkdirAll("./uploads", os.ModePerm)
		dst, err := os.Create(filepath.Join("./uploads", "uploaded_file.xlsx"))
		if err != nil {
			http.Error(w, "Error saving the file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			http.Error(w, "Error saving the file", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"success": true}`))
	})

	port := "8080"
	fmt.Printf("Listening on http://localhost:%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
