package address

import (
	"strings"

	"github.com/aghape-pkg/geocode"
	"github.com/aghape-pkg/phone"
	"github.com/aghape/db/common/utils"
	"github.com/moisespsena-go/aorm"
)

type AddressPhone struct {
	phone.Phone
	AddressID string `gorm:"size:24"`
}

func (p *AddressPhone) Clean(db *aorm.DB) {
	utils.TrimStrings(&p.Note, &p.Phone.Phone)
}

type Address struct {
	aorm.AuditedModel
	Phones       []AddressPhone         `gorm:"foreignkey:AddressID"`
	ContactName  string                 `gorm:"size:255"`
	RegionID     string                 `gorm:"size:10"`
	Region       *geocode.GeoCodeRegion `gorm:"SAVE_ASSOCIATIONS:false;preload:*"`
	AddressLine1 string                 `gorm:"size:255"`
	AddressLine2 string                 `gorm:"size:255"`
}

func (Address) GetGormInlinePreloadFields() []string {
	return []string{"Region"}
}

func (e *Address) Clean(db *aorm.DB) {
	utils.TrimStrings(&e.ContactName, &e.AddressLine1, &e.AddressLine2)
}

func (a *Address) Stringify() string {
	var parts []string
	if a.ContactName != "" {
		parts = append(parts, a.ContactName)
	}

	var parts2 []string

	if a.Region != nil {
		parts2 = append(parts2, a.Region.Country.Name+", "+a.Region.Name)
	}
	if a.AddressLine1 != "" {
		parts2 = append(parts2, a.AddressLine1)
	}
	if a.AddressLine2 != "" {
		parts2 = append(parts2, a.AddressLine2)
	}

	if len(parts2) > 0 {
		parts = append(parts, strings.Join(parts2, ", "))
	}

	return strings.Join(parts, ": ")
}
