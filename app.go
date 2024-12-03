package main

import (
	"context"
	"cs50NameChanger/models"
	"errors"
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
	err := a.CreateDB()
	if err != nil {
		log.Println(err)
	}
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
	if _, err := os.Stat("cs50x.db"); err == nil {
		return nil
	}
	db, err := a.ConnectDB()
	if err != nil {
		log.Println(err)
		return err
	}
	if err := a.migrateDB(db); err != nil {
		log.Println(err)
		return err
	}
	return a.populateDB()
}

func (a *App) migrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Program{}, &models.Cycle{}, &models.Week{}, &models.Group{},
		&models.ProgramsResponse{}, &models.CyclesResponse{}, &models.WeeksResponse{},
		&models.GroupsResponse{},
	)
}

func (a *App) populateDB() error {
	return a.executeInTransaction(func(tx *gorm.DB) error {
		programs := []string{"CS50", "Web50"}
		for _, program := range programs {
			if err := a.createProgramWithCyclesAndGroups(tx, program); err != nil {
				return err
			}
		}
		return nil
	})
}

func (a *App) createProgramWithCyclesAndGroups(tx *gorm.DB, program string) error {
	programDB, err := a.getOrCreateProgram(tx, program)
	if err != nil {
		return err
	}
	if err := a.createProgramResponse(tx, programDB); err != nil {
		return err
	}

	cycles := []string{"Y25C1", "Y25C2"}
	for _, cycleName := range cycles {
		cycleDB, err := a.getOrCreateCycle(tx, cycleName, programDB)
		if err != nil {
			return err
		}
		if err := a.createCycleResponse(tx, cycleDB, programDB.Name); err != nil {
			return err
		}

		if err := a.createWeeks(tx, cycleDB); err != nil {
			return err
		}
		if err := a.createGroups(tx, program, cycleDB); err != nil {
			return err
		}
	}
	return nil
}

func (a *App) getOrCreateProgram(tx *gorm.DB, name string) (*models.Program, error) {
	var program models.Program
	err := tx.Where("UPPER(name) = ?", strings.ToUpper(name)).First(&program).Error
	if err == nil {
		return &program, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	program = models.Program{Name: name}
	if err := tx.Create(&program).Error; err != nil {
		return nil, err
	}
	return &program, nil
}

func (a *App) createProgramResponse(tx *gorm.DB, program *models.Program) error {
	return tx.Create(&models.ProgramsResponse{
		ID:   program.ID,
		Name: program.Name,
	}).Error
}

func (a *App) getOrCreateCycle(tx *gorm.DB, name string, program *models.Program) (*models.Cycle, error) {
	var cycle models.Cycle
	err := tx.Where("UPPER(name) = ? AND program_id = ?", strings.ToUpper(name), program.ID).First(&cycle).Error
	if err == nil {
		return &cycle, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	cycle = models.Cycle{Name: name, ProgramID: program.ID}
	if err := tx.Create(&cycle).Error; err != nil {
		return nil, err
	}
	return &cycle, nil
}

func (a *App) createCycleResponse(tx *gorm.DB, cycle *models.Cycle, programName string) error {
	return tx.Create(&models.CyclesResponse{
		ID:          cycle.ID,
		Name:        cycle.Name,
		ProgramID:   cycle.ProgramID,
		ProgramName: programName,
	}).Error
}

func (a *App) createWeeks(tx *gorm.DB, cycle *models.Cycle) error {
	for i := 1; i <= 20; i++ {
		weekName := fmt.Sprintf("S%02d", i)
		week := models.Week{Name: weekName, CycleID: cycle.ID}
		if err := tx.Create(&week).Error; err != nil {
			return err
		}
		if err := tx.Create(&models.WeeksResponse{
			ID:        week.ID,
			Name:      week.Name,
			CycleID:   week.CycleID,
			CycleName: cycle.Name,
		}).Error; err != nil {
			return err
		}
	}
	return nil
}

func (a *App) createGroups(tx *gorm.DB, program string, cycle *models.Cycle) error {
	var groups []string
	if program == "CS50" {
		groups = []string{"A", "B", "C", "D", "E", "F"}
	} else {
		groups = []string{"A", "B", "C"}
	}
	for _, group := range groups {
		groupDB := models.Group{Name: group, CycleID: cycle.ID}
		if err := tx.Create(&groupDB).Error; err != nil {
			return err
		}
		if err := tx.Create(&models.GroupsResponse{
			ID:        groupDB.ID,
			Name:      groupDB.Name,
			CycleID:   groupDB.CycleID,
			CycleName: cycle.Name,
		}).Error; err != nil {
			return err
		}
	}
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

func (a *App) ListWeeks(cycle string, program string) []string {
	db, err := a.ConnectDB()
	if err != nil {
		return nil
	}
	var programDB models.Program
	db.First(&programDB, "name = ?", program)
	var weeks []models.WeeksResponse
	db.Joins("Cycle").Find(&weeks, "cycle.name = ? AND cycle.program_id = ?", cycle, programDB.ID)
	var weekNames []string
	for _, week := range weeks {
		weekNames = append(weekNames, week.Name)
	}
	if len(weekNames) == 0 {
		return []string{"No hay semanas"}
	}
	return weekNames
}

func (a *App) ListGroups(cycle string, program string) []string {
	db, err := a.ConnectDB()
	if err != nil {
		return nil
	}
	var programDB models.Program
	db.First(&programDB, "name = ?", program)
	var groups []models.GroupsResponse
	db.Joins("Cycle").Find(&groups, "cycle.name = ? AND cycle.program_id = ?", cycle, programDB.ID)
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
	return a.executeInTransaction(func(tx *gorm.DB) error {
		// Check for duplicate programs
		var existingProgram models.Program
		if err := tx.Where("UPPER(name) = ?", strings.ToUpper(name)).First(&existingProgram).Error; err == nil {
			return fmt.Errorf("el programa %s ya existe", name)
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		program := models.Program{Name: name}
		if err := tx.Create(&program).Error; err != nil {
			return err
		}

		programResponse := models.ProgramsResponse{
			ID:   program.ID,
			Name: program.Name,
		}
		if err := tx.Create(&programResponse).Error; err != nil {
			return err
		}

		return nil
	})
}

func (a *App) CreateCycle(name string, program string) error {
	return a.executeInTransaction(func(tx *gorm.DB) error {
		// Check if the program exists
		var programDB models.Program
		if err := tx.First(&programDB, "name = ?", program).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("el programa %s no existe", program)
			}
			return err
		}

		// Ensure the cycle doesn't already exist
		var cycle models.Cycle
		if err := tx.Where("UPPER(name) = ? AND program_id = ?", strings.ToUpper(name), programDB.ID).
			First(&cycle).Error; err == nil {
			return fmt.Errorf("el ciclo %s ya existe", name)
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err // Return any other errors
		}

		// Create the cycle
		cycle = models.Cycle{Name: name, ProgramID: programDB.ID}
		if err := tx.Create(&cycle).Error; err != nil {
			return err
		}

		// Create the CyclesResponse
		cycleResponse := models.CyclesResponse{
			ID:          cycle.ID,
			Name:        cycle.Name,
			ProgramID:   cycle.ProgramID,
			ProgramName: program,
		}
		if err := tx.Create(&cycleResponse).Error; err != nil {
			return err
		}

		return nil
	})
}

func (a *App) CreateWeek(name string, cycle string) error {
	return a.executeInTransaction(func(tx *gorm.DB) error {
		var cycleDB models.Cycle
		if err := tx.Where("name = ?", cycle).First(&cycleDB).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("el ciclo %s no existe", cycle)
			}
			return err
		}

		var existingWeek models.Week
		if err := tx.Where("UPPER(name) = ? AND cycle_id = ?", strings.ToUpper(name), cycleDB.ID).First(&existingWeek).Error; err == nil {
			return fmt.Errorf("la semana %s ya existe", name)
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		week := models.Week{Name: name, CycleID: cycleDB.ID}
		if err := tx.Create(&week).Error; err != nil {
			return err
		}

		weekResponse := models.WeeksResponse{
			ID:        week.ID,
			Name:      week.Name,
			CycleID:   week.CycleID,
			CycleName: cycle,
		}
		if err := tx.Create(&weekResponse).Error; err != nil {
			return err
		}

		return nil
	})
}

func (a *App) CreateGroup(name string, cycle string) error {

	return a.executeInTransaction(func(tx *gorm.DB) error {
		var cycleDB models.Cycle
		if err := tx.Where("name = ?", cycle).First(&cycleDB).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("el ciclo %s no existe", cycle)
			}
			return err
		}

		var existingGroup models.Group
		if err := tx.Where("UPPER(name) = ? AND cycle_id = ?", strings.ToUpper(name), cycleDB.ID).First(&existingGroup).Error; err == nil {
			return fmt.Errorf("el grupo %s ya existe", name)
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		group := models.Group{Name: name, CycleID: cycleDB.ID}
		if err := tx.Create(&group).Error; err != nil {
			return err
		}

		// Create the GroupsResponse
		groupResponse := models.GroupsResponse{
			ID:        group.ID,
			Name:      group.Name,
			CycleID:   group.CycleID,
			CycleName: cycle,
		}
		if err := tx.Create(&groupResponse).Error; err != nil {
			return err
		}

		return nil
	})
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
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Images (*.png;*.jpg;*.jpeg)",
				Pattern:     "*.png;*.jpg;*.jpeg",
			},
		},
	})
	if err != nil {
		return nil, err
	}
	if len(filepath) == 0 {
		return nil, nil
	}
	return filepath, nil
}

func (a *App) ChangeFileNames(files []string, newName string, start_index int) error {
	exePath, err := os.Executable()
	if err != nil {
		return err
	}
	newPath := filepath.Join(filepath.Dir(exePath), newName)

	if _, err := os.Stat(newPath); os.IsNotExist(err) {
		// If it doesn't exist, create it
		err = os.Mkdir(newPath, 0755)
		if err != nil {
			return err
		}
	}

	for i := start_index; i < len(files)+start_index; i++ {
		file := files[i-start_index]
		fileExt := filepath.Ext(file)
		newFileName := fmt.Sprintf("%s%d%s", newName, i, fileExt)
		newFilePath := filepath.Join(newPath, newFileName)
		err := os.Rename(file, newFilePath)
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

func (a *App) WatchLastCreatedIndex(newName string) int {
	exePath, err := os.Executable()
	if err != nil {
		return 0
	}
	newPath := filepath.Join(filepath.Dir(exePath), newName)
	filesInDir, err := os.ReadDir(newPath)
	if err != nil {
		return 0
	}
	maxIndex := 0
	for _, f := range filesInDir {
		fileName := f.Name()
		var index int
		_, err := fmt.Sscanf(fileName, newName+"%d", &index)
		if err == nil && index > maxIndex {
			maxIndex = index
		}
	}
	return maxIndex
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
		var program models.Program
		db.First(&program, "id = ?", cycle.ProgramID)
		weeksResponse = append(weeksResponse, models.WeeksResponse{
			ID:        week.ID,
			Name:      week.Name,
			CycleID:   week.CycleID,
			CycleName: fmt.Sprintf("%s (%s)", cycle.Name, program.Name),
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
		var program models.Program
		db.First(&program, "id = ?", cycle.ProgramID)
		groupsResponse = append(groupsResponse, models.GroupsResponse{
			ID:        group.ID,
			Name:      group.Name,
			CycleID:   group.CycleID,
			CycleName: fmt.Sprintf("%s (%s)", cycle.Name, program.Name),
		})
	}
	return groupsResponse
}

func (a *App) EditProgram(id int, nuevo_nombre string) error {
	db, err := a.ConnectDB()
	if err != nil {
		return err
	}
	var existingProgram models.Program
	if err = db.First(&existingProgram, "name = ?", nuevo_nombre).Error; err == nil {
		return fmt.Errorf("el programa %s ya existe", nuevo_nombre)
	}

	var program models.Program
	err = db.First(&program, "id = ?", id).Error
	if err != nil {
		return err
	}

	if program.Name == nuevo_nombre {
		return nil
	}
	program.Name = nuevo_nombre
	err = db.Save(&program).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *App) EditCycle(id int, nuevo_nombre string) error {
	db, err := a.ConnectDB()
	if err != nil {
		return err
	}

	var cycle models.Cycle
	err = db.First(&cycle, "id = ?", id).Error
	if err != nil {
		return err
	}

	if cycle.Name == nuevo_nombre {
		return nil
	}

	var existingCycle models.Cycle
	if err = db.First(&existingCycle, "name = ? AND program_id = ?", nuevo_nombre, cycle.ProgramID).Error; err == nil {
		return fmt.Errorf("el ciclo %s ya existe", nuevo_nombre)
	}

	cycle.Name = nuevo_nombre
	err = db.Save(&cycle).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *App) EditWeek(id int, nuevo_nombre string) error {
	db, err := a.ConnectDB()
	if err != nil {
		return err
	}

	var week models.Week
	err = db.First(&week, "id = ?", id).Error
	if err != nil {
		return err
	}

	if week.Name == nuevo_nombre {
		return nil
	}

	var existingWeek models.Week
	if err = db.First(&existingWeek, "name = ? AND cycle_id = ?", nuevo_nombre, week.CycleID).Error; err == nil {
		return fmt.Errorf("la semana %s ya existe", nuevo_nombre)
	}

	week.Name = nuevo_nombre
	err = db.Save(&week).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *App) EditGroup(id int, nuevo_nombre string) error {
	db, err := a.ConnectDB()
	if err != nil {
		return err
	}

	var group models.Group
	err = db.First(&group, "id = ?", id).Error
	if err != nil {
		return err
	}

	if group.Name == nuevo_nombre {
		return nil
	}

	var existingGroup models.Group
	if err = db.First(&existingGroup, "name = ? AND cycle_id = ?", nuevo_nombre, group.CycleID).Error; err == nil {
		return fmt.Errorf("el grupo %s ya existe", nuevo_nombre)
	}

	group.Name = nuevo_nombre
	err = db.Save(&group).Error
	if err != nil {
		return err
	}
	return nil
}
