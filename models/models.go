package models

import (
	"gorm.io/gorm"
)

type Program struct {
	gorm.Model
	Name   string  `gorm:"index:idx_name_code,unique" json:"name"`
	Cycles []Cycle `gorm:"foreignKey:ProgramID" json:"cycles"`
}

type Cycle struct {
	gorm.Model
	Name      string  `gorm:"index:idx_cycle_name_program_id,unique" json:"name"`
	ProgramID uint    `gorm:"index:idx_cycle_name_program_id" json:"program_id"`
	Program   Program `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"program"`
	Weeks     []Week  `gorm:"foreignKey:CycleID" json:"weeks"`
	Groups    []Group `gorm:"foreignKey:CycleID" json:"groups"`
}

// Week struct
type Week struct {
	gorm.Model
	Name    string `gorm:"index:idx_week_name_cycle_id,unique" json:"name"`
	CycleID uint   `gorm:"index:idx_week_name_cycle_id,unique" json:"cycle_id"`
	Cycle   Cycle  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"cycle"`
}

// Group struct
type Group struct {
	gorm.Model
	Name    string `gorm:"index:idx_group_name_cycle_id,unique" json:"name"`
	CycleID uint   `gorm:"index:idx_group_name_cycle_id,unique" json:"cycle_id"`
	Cycle   Cycle  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"cycle"`
}

type ProgramsResponse struct {
	ID     uint             `json:"id"`
	Name   string           `json:"name"`
	Cycles []CyclesResponse `gorm:"foreignKey:ProgramID" json:"cycles"`
}

type CyclesResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	ProgramID uint   `json:"program_id"`
	Program   ProgramsResponse
	Weeks     []WeeksResponse  `gorm:"foreignKey:CycleID" json:"weeks"`
	Groups    []GroupsResponse `gorm:"foreignKey:CycleID" json:"groups"`
}

type WeeksResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	CycleID uint   `json:"cycle_id"`
	Cycle   CyclesResponse
}

type GroupsResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	CycleID uint   `json:"cycle_id"`
	Cycle   CyclesResponse
}
