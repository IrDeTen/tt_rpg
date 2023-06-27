package models

type CampaignSetting struct {
	ID          uint
	Name        string
	Description string
	Author      User
	System      TTSystem
	Public      bool
}

type Campaign struct {
	ID          uint
	Name        string
	Description string
	Setting     CampaignSetting
	GameMaster  User
	Players     []User
	Public      bool
}

type TTSystem interface {
	Create()
	Update()
	Delete()
}

type ttSystem struct {
	ID           uint
	Name         string
	Desctription string
	Version      string
	Public       bool
	Creators     []User
	Editors      []User
}
