package domain

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
	Rating       float32
	Size         string
	Category     string
	ImageLinks   []string
	ShopId       uint64
}
