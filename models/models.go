package models

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
	//"reflect"
)

type Network struct {
	gorm.Model
	Name    string
	Subnets []Subnet `gorm:"type:bytes;serializer:json"`
}

type Subnet struct {
	gorm.Model
	Name     string  `gorm:"type:varchar(255)" json:"name"`
	ParentID *uint64 `gorm:"index"`
	Nodes    []Node
	Links    []Link
}

// Define a custom type for JSONB
type JSONB []map[string]string

// These interfaces are used to handle the conversion between Go data structures and PostgreSQL's JSONB type
// Implement the Scanner interface
func (j *JSONB) Scan(value interface{}) error {
	return json.Unmarshal([]byte(value.([]uint8)), j)
}

// Implement the Valuer interface
func (j JSONB) Value() (driver.Value, error) {
	return json.Marshal(j)
}

type Ltp struct {
	gorm.Model
	Name string `gorm:"type:varchar(255)" json:"name"`
	//Group       bool
	Description string `gorm:"type:varchar(255)" json:"label"`
	Busy        bool
	Hwinfo      JSONB   `gorm:"type:jsonb" json:"hwinfo"`
	NodeID      *uint64 `gorm:"index"`
	Node        Node
}

type Node struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:varchar(255)"`
	Label       string `gorm:"type:varchar(255)"`
	Posx        uint64
	Posy        uint64
	SubnetID    *uint64 `gorm:"index"`
	Subnet      Subnet
	Ltps        []Ltp
}

type Link struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(255)"`
	Status      string  `gorm:"type:varchar(255)"`
	Description string  `gorm:"type:varchar(255)"`
	Label       string  `gorm:"type:varchar(255)"`
	SrcLtpId    *uint64 `gorm:"index"`
	DestLtpId   *uint64 `gorm:"index"`
	SrcNodeId   *uint64 `gorm:"index"`
	DestNodeId  *uint64 `gorm:"index"`
	SubnetId    *uint64 `gorm:"index"`
}
