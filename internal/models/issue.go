package models

type Status struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type Type struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Priority struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type Company struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Contact struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ServiceObject struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Agreement struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type EquipmentKind struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type EquipmentManufacturer struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type EquipmentModel struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Equipment struct {
	ID                    int                   `json:"id"`
	SerialNumber          string                `json:"serial_number"`
	InventoryNumber       string                `json:"inventory_number"`
	EquipmentKind         EquipmentKind         `json:"equipment_kind"`
	EquipmentManufacturer EquipmentManufacturer `json:"equipment_manufacturer"`
	EquipmentModel        EquipmentModel        `json:"equipment_model"`
}

type Issue struct {
	ID                int           `json:"id"`
	Title             string        `json:"title"`
	CreatedAt         string        `json:"created_at"`
	CompletedAt       string        `json:"completed_at"`
	DeadlineAt        string        `json:"deadline_at"`
	PlannedReactionAt string        `json:"planned_reaction_at"`
	ReactedAt         string        `json:"reacted_at"`
	DelayTo           *string       `json:"delay_to"`
	UpdatedAt         string        `json:"updated_at"`
	Status            Status        `json:"status"`
	Type              Type          `json:"type"`
	Priority          Priority      `json:"priority"`
	Company           Company       `json:"company"`
	Contact           Contact       `json:"contact"`
	ServiceObject     ServiceObject `json:"service_object"`
	Agreement         Agreement     `json:"agreement"`
	Equipments        []Equipment   `json:"equipments"`
}
