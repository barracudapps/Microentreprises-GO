package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

// Supprimer le dossier dist s'il existe
func removeDistDir() {
	distDir := "./dist"
	if _, err := os.Stat(distDir); !os.IsNotExist(err) {
		err := os.RemoveAll(distDir)
		if err != nil {
			fmt.Println("Error removing dist directory:", err)
			os.Exit(1)
		}
	}
}

// Copier les fichiers frontend dans le dossier de distribution
func copyFrontend(srcDir, dstDir string) {
	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		// Exclure les fichiers Sass
		if filepath.Ext(path) == ".scss" {
			return nil
		}
		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dstDir, relPath)
		os.MkdirAll(filepath.Dir(dstPath), os.ModePerm)
		return copyFile(path, dstPath)
	})
	if err != nil {
		fmt.Println("Error copying frontend files:", err)
		os.Exit(1)
	}
}

// Copier les fichiers de template dans le dossier de distribution
func copyTemplate(srcDir, dstDir string) {
	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dstDir, relPath)
		os.MkdirAll(filepath.Dir(dstPath), os.ModePerm)
		return copyFile(path, dstPath)
	})
	if err != nil {
		fmt.Println("Error copying template files:", err)
		os.Exit(1)
	}
}

// Copier un fichier
func copyFile(src, dst string) error {
	from, err := os.Open(src)
	if err != nil {
		return err
	}
	defer from.Close()
	to, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer to.Close()
	_, err = io.Copy(to, from)
	return err
}

func main() {
	removeDistDir()

	// Compiler les fichiers Sass en CSS
	cmd := exec.Command("sass", "frontend/assets/styles/main.scss", "frontend/assets/styles/main.css")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error compiling Sass:", err)
		os.Exit(1)
	}

	// Copier les fichiers frontend
	copyFrontend("frontend", "dist")

	// Copier les fichiers template
	copyTemplate("data/templates", "dist/data/templates")

	fmt.Println("Build completed successfully.")
}
