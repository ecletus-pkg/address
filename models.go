package address

import (
	"strings"

	"github.com/ecletus-pkg/geocode"
	"github.com/ecletus-pkg/phone"
	"github.com/ecletus/db/common/utils"
	"github.com/moisespsena-go/aorm"
)

type AddressPhone struct {
	phone.Phone
	AddressID string `gorm:"size:24"`
}

func (p *AddressPhone) Clean(db *aorm.DB) {
	utils.TrimStrings(&p.Note, &p.Phone.Number)
}

type Address struct {
	aorm.AuditedModel
	Phones       []AddressPhone         `gorm:"foreignkey:AddressID"`
	ContactName  string                 `gorm:"size:255"`
	RegionID     string                 `gorm:"size:10"`
	Region       *geocode.GeoCodeRegion `gorm:"SAVE_ASSOCIATIONS:false"`
	AddressLine1 string                 `gorm:"size:255"`
	AddressLine2 string                 `gorm:"size:255"`
	AddressLine3 string                 `gorm:"size:255"`
	AddressLine4 string                 `gorm:"size:255"`
	Cep          string                 `gorm:"size:32"`
}

func (Address) GetGormInlinePreloadFields() []string {
	return []string{"*", "Region"}
}

func (e *Address) Clean(db *aorm.DB) {
	utils.TrimStrings(&e.ContactName, &e.AddressLine1, &e.AddressLine2, &e.AddressLine3, &e.Cep)
}

func (a *Address) String() string {
	var parts []string

	if a.Region != nil {
		parts = append(parts, a.Region.Country.Name+", "+a.Region.Name)
	}
	if a.AddressLine1 != "" {
		parts = append(parts, a.AddressLine1)
	}
	if a.AddressLine2 != "" {
		parts = append(parts, a.AddressLine2)
	}
	if a.AddressLine3 != "" {
		parts = append(parts, a.AddressLine3)
	}
	if a.AddressLine4 != "" {
		parts = append(parts, a.AddressLine4)
	}

	if a.ContactName != "" {
		parts = append(parts, "("+a.ContactName+")")
	}

	return strings.Join(parts, ", ")
}
