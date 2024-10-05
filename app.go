package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	// sqlite ORM
	// "gorm.io/gorm"
	// "gorm.io/driver/sqlite"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	names := [2]string{"Mario", "Luis"}
	return fmt.Sprintf("Hello %s, It's show time!", names[1])
}

func (a *App) ListPrograms() []string {
	programs := []string{"cs50", "web50", "CodeKids"}
	return programs
}

func (a *App) ListCycles() []string {
	cycles := []string{"Y24C1", "Y24C2", "Y24C3"}
	return cycles
}

func (a *App) ListWeeks() []string {
	weeks := []string{"S01", "S02", "S03"}
	return weeks
}

func (a *App) ListGroups() []string {
	groups := []string{"A", "B", "C"}
	return groups
}

func (a *App) MakeNewName(program string, cycle string, week string, group string) string {
	if program == "" || cycle == "" || week == "" || group == "" {
		return "Nombre del archivo"
	}
	return fmt.Sprintf("%s_%s_%s_grupo%s_foto", program, cycle, week, group)
}

func (a *App) OpenFileDialog() ([]string, error) {
	filepath, err := runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Ingresa los archivos que desee modificar",
	})
	if err != nil {
		return nil, err
	}
	return filepath, nil
}

func (a *App) ChangeFileNames(files []string, newName string) error {
	for i := 0; i < len(files); i++ {
		file := files[i]
		dir := filepath.Dir(file)
		fileName := filepath.Base(file)
		newFile := filepath.Join(dir, newName+strconv.Itoa(i+1)+filepath.Ext(fileName))
		err := os.Rename(file, newFile)
		if err != nil {
			return err
		}
	}
	/*
		for _, file := range files {
			dir, filename := filepath.Split(file)
			newFile := filepath.Join(dir, newName+filepath.Ext(filename))
			err := os.Rename(file, newFile)
			if err != nil {
				return err
			}
		}
	*/
	return nil
}
