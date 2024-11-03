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

func (a *App) executeInTransaction(fn func(*gorm.DB) error) error {
	db, err := a.ConnectDB()
	if err != nil {
		return err
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (a *App) CreateProgram(name string) error {
	a.executeInTransaction(func(tx *gorm.DB) error {
		var count int64
		err := tx.Table("programs").Where("UPPER(name) = ?", strings.ToUpper(name)).Count(&count).Error
		if err != nil {
			return err
		}
		if count > 0 {
			return fmt.Errorf("el programa %s ya existe", name)
		}
		program := models.Program{Name: name}
		err = tx.Create(&program).Error
		if err != nil {
			return err
		}
		tx.Create(&models.ProgramsResponse{
			ID:   program.ID,
			Name: program.Name,
		})
		return nil
	},
	)
	return nil
}

func (a *App) CreateCycle(name string, program string) error {
	a.executeInTransaction(func(tx *gorm.DB) error {
		var count int64
		err := tx.Table("cycles").Where("UPPER(name) = ? AND program_id = ?", strings.ToUpper(name), program).Count(&count).Error
		if err != nil {
			return err
		}
		if count > 0 {
			return fmt.Errorf("el ciclo %s ya existe", name)
		}
		var programDB models.Program
		tx.First(&programDB, "name = ?", program)
		cycle := models.Cycle{Name: name, ProgramID: programDB.ID}
		err = tx.Create(&cycle).Error
		if err != nil {
			return err
		}
		tx.Create(&models.CyclesResponse{
			ID:          cycle.ID,
			Name:        cycle.Name,
			ProgramID:   cycle.ProgramID,
			ProgramName: program,
		})
		return nil
	},
	)
	return nil
}

func (a *App) CreateWeek(name string, cycle string) error {
	a.executeInTransaction(func(tx *gorm.DB) error {
		var count int64
		err := tx.Table("weeks").Where("UPPER(name) = ? AND cycle_id = ?", strings.ToUpper(name), cycle).Count(&count).Error
		if err != nil {
			return err
		}
		if count > 0 {
			return fmt.Errorf("la semana %s ya existe", name)
		}
		var cycleDB models.Cycle
		tx.First(&cycleDB, "name = ?", cycle)
		week := models.Week{Name: name, CycleID: cycleDB.ID}
		err = tx.Create(&week).Error
		if err != nil {
			return err
		}
		tx.Create(&models.WeeksResponse{
			ID:        week.ID,
			Name:      week.Name,
			CycleID:   week.CycleID,
			CycleName: cycle,
		})
		return nil
	},
	)
	return nil
}

func (a *App) CreateGroup(name string, cycle string) error {
	a.executeInTransaction(func(tx *gorm.DB) error {
		var count int64
		err := tx.Table("groups").Where("UPPER(name) = ? AND cycle_id = ?", strings.ToUpper(name), cycle).Count(&count).Error
		if err != nil {
			return err
		}
		if count > 0 {
			return fmt.Errorf("el grupo %s ya existe", name)
		}
		var cycleDB models.Cycle
		tx.First(&cycleDB, "name = ?", cycle)
		group := models.Group{Name: name, CycleID: cycleDB.ID}
		err = tx.Create(&group).Error
		if err != nil {
			return err
		}
		tx.Create(&models.GroupsResponse{
			ID:        group.ID,
			Name:      group.Name,
			CycleID:   group.CycleID,
			CycleName: cycle,
		})
		return nil
	},
	)
	return nil
}

func (a *App) DeleteProgram(id int) error {
	a.executeInTransaction(func(tx *gorm.DB) error {
		var program models.Program
		err := tx.First(&program, "id = ?", id).Error
		if err != nil {
			return err
		}
		err = tx.Unscoped().Delete(&program).Error
		if err != nil {
			return err
		}
		tx.Unscoped().Delete(&models.ProgramsResponse{}, "id = ?", program.ID)
		var cycles []models.Cycle
		tx.Find(&cycles, "program_id = ?", program.ID)
		for _, cycle := range cycles {
			tx.Unscoped().Delete(&cycle)
			tx.Unscoped().Delete(&models.CyclesResponse{}, "id = ?", cycle.ID)
			var weeks []models.Week
			tx.Find(&weeks, "cycle_id = ?", cycle.ID)
			for _, week := range weeks {
				tx.Unscoped().Delete(&week)
				tx.Unscoped().Delete(&models.WeeksResponse{}, "id = ?", week.ID)
				var groups []models.Group
				tx.Find(&groups, "cycle_id = ?", cycle.ID)
				for _, group := range groups {
					tx.Unscoped().Delete(&group)
					tx.Unscoped().Delete(&models.GroupsResponse{}, "id = ?", group.ID)
				}
			}
		}
		return nil
	},
	)
	return nil
}

func (a *App) DeleteCycle(id int) error {
	a.executeInTransaction(func(tx *gorm.DB) error {
		var cycle models.Cycle
		err := tx.First(&cycle, "id = ?", id).Error
		if err != nil {
			return err
		}
		err = tx.Unscoped().Delete(&cycle).Error
		if err != nil {
			return err
		}
		tx.Unscoped().Delete(&models.CyclesResponse{}, "id = ?", cycle.ID)
		var weeks []models.Week
		tx.Find(&weeks, "cycle_id = ?", cycle.ID)
		for _, week := range weeks {
			tx.Unscoped().Delete(&week)
			tx.Unscoped().Delete(&models.WeeksResponse{}, "id = ?", week.ID)
		}
		var groups []models.Group
		tx.Find(&groups, "cycle_id = ?", cycle.ID)
		for _, group := range groups {
			tx.Unscoped().Delete(&group)
			tx.Unscoped().Delete(&models.GroupsResponse{}, "id = ?", group.ID)
		}
		return nil
	},
	)
	return nil
}

func (a *App) DeleteWeek(id int) error {
	a.executeInTransaction(func(tx *gorm.DB) error {
		var week models.Week
		err := tx.First(&week, "id = ?", id).Error
		if err != nil {
			return err
		}
		err = tx.Unscoped().Delete(&week).Error
		if err != nil {
			return err
		}
		tx.Unscoped().Delete(&models.WeeksResponse{}, "id = ?", week.ID)
		return nil
	},
	)
	return nil
}

func (a *App) DeleteGroup(id int) error {
	a.executeInTransaction(func(tx *gorm.DB) error {
		var group models.Group
		err := tx.First(&group, "id = ?", id).Error
		if err != nil {
			return err
		}
		err = tx.Unscoped().Delete(&group).Error
		if err != nil {
			return err
		}
		tx.Unscoped().Delete(&models.GroupsResponse{}, "id = ?", group.ID)
		return nil
	},
	)
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

	a.OpenFolder(newPath)

	return nil
}

func (a *App) OpenFolder(path string) error {
	runtime.BrowserOpenURL(a.ctx, "file://"+path)
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
		var program models.Program
		db.First(&program, "id = ?", cycle.ProgramID)
		cyclesResponse = append(cyclesResponse, models.CyclesResponse{
			ID:          cycle.ID,
			Name:        cycle.Name,
			ProgramID:   cycle.ProgramID,
			ProgramName: program.Name,
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
		var cycle models.Cycle
		db.First(&cycle, "id = ?", week.CycleID)
		weeksResponse = append(weeksResponse, models.WeeksResponse{
			ID:        week.ID,
			Name:      week.Name,
			CycleID:   week.CycleID,
			CycleName: cycle.Name,
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
		var cycle models.Cycle
		db.First(&cycle, "id = ?", group.CycleID)
		groupsResponse = append(groupsResponse, models.GroupsResponse{
			ID:        group.ID,
			Name:      group.Name,
			CycleID:   group.CycleID,
			CycleName: cycle.Name,
		})
	}
	return groupsResponse
}
