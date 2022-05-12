package domain

import "time"

type Shop struct {
	Id          uint64
	Title       string
	Description string
	ManagerIDs  []uint64
}

type Product struct {
	Id           uint64
	Title        string
	Description  string
	Price        uint64
	Availability bool
	// AssemblyTime is measured in minutes
	AssemblyTime uint64
	PartsAmount  uint64
	Rating       float64
	Size         string
	Category     string
	ImageLinks   []string
	VideoLink    string
	ShopId       uint64
}

type Order struct {
	Id              uint64
	UserID          uint64
	Items           map[uint64]uint64
	CreatedAt       time.Time
	TotalPrice      uint64
	PickUp          bool
	DeliveryAddress string
	PaymentMethod   string
	CallNeeded      bool
	Status          string
}
