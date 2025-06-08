package selection

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Form struct {
	LocationID    int `json:"location_id"`
	WHospital     int `json:"w_hospital"`
	WSport        int `json:"w_sport"`
	WShop         int `json:"w_shop"`
	WKindergarten int `json:"w_kindergarten"`
	WBusStop      int `json:"w_bus_stop"`
	WSchool       int `json:"w_school"`
}

func (m *Form) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, m)
	case string:
		return json.Unmarshal([]byte(v), m)
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
}

type SelectionCreate struct {
	Name    string
	Comment string
	Form    Form
}

type Selection struct {
	Id        uuid.UUID
	Name      string    `db:"name"`
	Comment   string    `db:"comment"`
	Form      Form      `db:"form"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
