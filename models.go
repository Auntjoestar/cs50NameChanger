package models

import (
	"gorm.io/gorm"
)

// Program struct
type Program struct {
	gorm.Model
	Name   string
	Cycles []Cycle `gorm:"foreignKey:ProgramID"` // A Program has many Cycles
}

// Cycle struct
type Cycle struct {
	gorm.Model
	Name      string
	ProgramID uint    // Foreign key for Program
	Program   Program `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Foreign key relationship with Program
	Weeks     []Week  `gorm:"foreignKey:CycleID"`                             // A Cycle has many Weeks
	Groups    []Group `gorm:"foreignKey:CycleID"`                             // A Cycle has many Groups
}

// Week struct
type Week struct {
	gorm.Model
	Name    string
	CycleID uint  // Foreign key for Cycle
	Cycle   Cycle `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Foreign key relationship with Cycle
}

// Group struct
type Group struct {
	gorm.Model
	Name    string
	CycleID uint  // Foreign key for Cycle
	Cycle   Cycle `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Foreign key relationship with Cycle
}
