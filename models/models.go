package models

import (
	"gorm.io/gorm"
)

type Program struct {
	gorm.Model
	Name   string  `gorm:"index:idx_name_code,unique"`
	Cycles []Cycle `gorm:"foreignKey:ProgramID"`
}

type Cycle struct {
	gorm.Model
	Name      string  `gorm:"index:idx_cycle_name_program_id,unique"`
	ProgramID uint    `gorm:"index:idx_cycle_name_program_id,unique"`
	Program   Program `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Weeks     []Week  `gorm:"foreignKey:CycleID"`
	Groups    []Group `gorm:"foreignKey:CycleID"`
}

// Week struct
type Week struct {
	gorm.Model
	Name    string `gorm:"index:idx_week_name_cycle_id,unique"`
	CycleID uint   `gorm:"index:idx_week_name_cycle_id,unique"`
	Cycle   Cycle  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Group struct
type Group struct {
	gorm.Model
	Name    string `gorm:"index:idx_group_name_cycle_id,unique"`
	CycleID uint   `gorm:"index:idx_group_name_cycle_id,unique"`
	Cycle   Cycle  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
