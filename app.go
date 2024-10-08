package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"cs50NameChanger/models"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

func (a *App) ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("cs50x.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(
		&models.Program{}, &models.Cycle{}, &models.Week{}, &models.Group{},
		&models.ProgramsResponse{}, &models.CyclesResponse{}, &models.WeeksResponse{},
		&models.GroupsResponse{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (a *App) CreateDB() error {
	db, err := a.ConnectDB()
	if err != nil {
		return err
	}
	db.AutoMigrate(
		&models.Program{}, &models.Cycle{}, &models.Week{}, &models.Group{},
		&models.ProgramsResponse{}, &models.CyclesResponse{}, &models.WeeksResponse{},
		&models.GroupsResponse{})
	return nil
}

func (a *App) ListPrograms() []string {
	db, err := a.ConnectDB()
	if err != nil {
		return nil
	}
	var programs []models.ProgramsResponse
	db.Find(&programs)
	var programNames []string
	for _, program := range programs {
		programNames = append(programNames, program.Name)
	}
	if len(programNames) == 0 {
		return []string{"No hay programas"}
	}
	return programNames
}

func (a *App) ListCycles(program string) []string {
	db, err := a.ConnectDB()
	if err != nil {
		return nil
	}
	var cycles []models.CyclesResponse
	db.Joins("Program").Find(&cycles, "program.name = ?", program)
	var cycleNames []string
	for _, cycle := range cycles {
		cycleNames = append(cycleNames, cycle.Name)
	}
	if len(cycleNames) == 0 {
		return []string{"No hay ciclos"}
	}
	return cycleNames
}

func (a *App) ListWeeks(cycle string) []string {
	db, err := a.ConnectDB()
	if err != nil {
		return nil
	}
	var weeks []models.WeeksResponse
	db.Joins("Cycle").Find(&weeks, "cycle.name = ?", cycle)
	var weekNames []string
	for _, week := range weeks {
		weekNames = append(weekNames, week.Name)
	}
	if len(weekNames) == 0 {
		return []string{"No hay semanas"}
	}
	return weekNames
}

func (a *App) ListGroups(cycle string) []string {
	db, err := a.ConnectDB()
	if err != nil {
		return nil
	}
	var groups []models.GroupsResponse
	db.Joins("Cycle").Find(&groups, "cycle.name = ?", cycle)
	var groupNames []string
	for _, group := range groups {
		groupNames = append(groupNames, group.Name)
	}
	if len(groupNames) == 0 {
		return []string{"No hay grupos"}
	}
	return groupNames
}

func (a *App) CreateProgram(name string) error {
	db, err := a.ConnectDB()
	if err != nil {
		return err
	}
	program := models.Program{Name: name}
	err = db.Create(&program).Error
	if err != nil {
		return err
	}
	db.Create(&models.ProgramsResponse{
		ID:   program.ID,
		Name: program.Name,
	})
	return nil
}

func (a *App) CreateCycle(name string, program string) error {
	db, err := a.ConnectDB()
	if err != nil {
		return err
	}
	var programDB models.Program
	db.First(&programDB, "name = ?", program)
	cycle := models.Cycle{Name: name, ProgramID: programDB.ID}
	err = db.Create(&cycle).Error
	if err != nil {
		return err
	}
	db.Create(&models.CyclesResponse{
		ID:        cycle.ID,
		Name:      cycle.Name,
		ProgramID: cycle.ProgramID,
	})
	return nil
}

func (a *App) CreateWeek(name string, cycle string) error {
	db, err := a.ConnectDB()
	if err != nil {
		return err
	}
	var cycleDB models.Cycle
	db.First(&cycleDB, "name = ?", cycle)
	week := models.Week{Name: name, CycleID: cycleDB.ID}
	err = db.Create(&week).Error
	if err != nil {
		return err
	}
	db.Create(&models.WeeksResponse{
		ID:      week.ID,
		Name:    week.Name,
		CycleID: week.CycleID,
	})
	return nil
}

func (a *App) CreateGroup(name string, cycle string) error {
	db, err := a.ConnectDB()
	if err != nil {
		return err
	}
	var cycleDB models.Cycle
	db.First(&cycleDB, "name = ?", cycle)
	group := models.Group{Name: name, CycleID: cycleDB.ID}
	err = db.Create(&group).Error
	if err != nil {
		return err
	}
	db.Create(&models.GroupsResponse{
		ID:      group.ID,
		Name:    group.Name,
		CycleID: group.CycleID,
	})
	return nil
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
	return nil
}
