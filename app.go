package main

import (
	"context"
	"cs50NameChanger/models"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

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
func (a *App) shutdown(ctx context.Context) {
	// Wait for any running operations to complete or handle the context's cancellation
	select {
	case <-ctx.Done():
		// Context canceled, proceed with shutdown
		log.Println("Shutdown context received, closing resources...")
	default:
		// Perform any pre-shutdown actions if necessary
	}

	// Close the database connection
	a.CloseDB()

	// Any additional cleanup can be done here
	log.Println("Application has shut down.")
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

func (a *App) CloseDB() error {
	db, err := a.ConnectDB()
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
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

	var count int64
	err = db.Table("programs").Where("UPPER(name) = ?", strings.ToUpper(name)).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("el programa %s ya existe", name)
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
	var count int64
	err = db.Table("cycles").Where("UPPER(name) = ? AND program_id = ?", strings.ToUpper(name), program).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("el ciclo %s ya existe", name)
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
	var count int64
	err = db.Table("weeks").Where("UPPER(name) = ? AND cycle_id = ?", strings.ToUpper(name), cycle).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("la semana %s ya existe", name)
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
	var count int64
	err = db.Table("groups").Where("UPPER(name) = ? AND cycle_id = ?", strings.ToUpper(name), cycle).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("el grupo %s ya existe", name)
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
	exePath, err := os.Executable()
	if err != nil {
		return err
	}
	newPath := filepath.Join(filepath.Dir(exePath), newName)

	// Check if the directory already exists
	if _, err := os.Stat(newPath); os.IsNotExist(err) {
		// If it doesn't exist, create it
		err = os.Mkdir(newPath, 0755)
		if err != nil {
			return err
		}
	}

	// Get the current max index of the files already in the directory
	maxIndex := 0
	filesInDir, err := os.ReadDir(newPath)
	if err != nil {
		return err
	}

	for _, f := range filesInDir {
		fileName := f.Name()
		var index int
		// Check if the filename matches the pattern "newName_X.ext" and extract the index
		_, err := fmt.Sscanf(fileName, newName+"%d", &index)
		if err == nil && index > maxIndex {
			maxIndex = index
		}
	}

	// Start renaming from the next index
	for i := 0; i < len(files); i++ {
		file := files[i]
		newFileName := fmt.Sprintf("%s%d%s", newName, maxIndex+i+1, filepath.Ext(file))
		newFilePath := filepath.Join(newPath, newFileName)
		err = os.Rename(file, newFilePath)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) WatchPrograms() []models.ProgramsResponse {
	db, err := a.ConnectDB()
	if err != nil {
		return []models.ProgramsResponse{}
	}
	var programs []models.Program
	db.Find(&programs)
	if len(programs) == 0 {
		return []models.ProgramsResponse{}
	}
	var programsResponse []models.ProgramsResponse
	for _, program := range programs {
		programsResponse = append(programsResponse, models.ProgramsResponse{
			ID:   program.ID,
			Name: program.Name,
		})
	}
	return programsResponse
}

func (a *App) WatchCycles() []models.CyclesResponse {
	db, err := a.ConnectDB()
	if err != nil {
		return []models.CyclesResponse{}
	}
	var cycles []models.Cycle
	db.Find(&cycles)
	if len(cycles) == 0 {
		return []models.CyclesResponse{}
	}
	var cyclesResponse []models.CyclesResponse
	for _, cycle := range cycles {
		cyclesResponse = append(cyclesResponse, models.CyclesResponse{
			ID:        cycle.ID,
			Name:      cycle.Name,
			ProgramID: cycle.ProgramID,
		})
	}
	return cyclesResponse
}

func (a *App) WatchWeeks() []models.WeeksResponse {
	db, err := a.ConnectDB()
	if err != nil {
		return []models.WeeksResponse{}
	}
	var weeks []models.Week
	db.Find(&weeks)
	if len(weeks) == 0 {
		return []models.WeeksResponse{}
	}
	var weeksResponse []models.WeeksResponse
	for _, week := range weeks {
		weeksResponse = append(weeksResponse, models.WeeksResponse{
			ID:      week.ID,
			Name:    week.Name,
			CycleID: week.CycleID,
		})
	}
	return weeksResponse
}

func (a *App) WatchGroups() []models.GroupsResponse {
	db, err := a.ConnectDB()
	if err != nil {
		return []models.GroupsResponse{}
	}
	var groups []models.Group
	db.Find(&groups)
	if len(groups) == 0 {
		return []models.GroupsResponse{}
	}
	var groupsResponse []models.GroupsResponse
	for _, group := range groups {
		groupsResponse = append(groupsResponse, models.GroupsResponse{
			ID:      group.ID,
			Name:    group.Name,
			CycleID: group.CycleID,
		})
	}
	return groupsResponse
}
