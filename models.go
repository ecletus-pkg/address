package address

import (
	"strings"

	"github.com/ecletus-pkg/geocode"
	"github.com/ecletus-pkg/phone"
	"github.com/ecletus/db/common/utils"
	"github.com/moisespsena-go/aorm"
	"github.com/moisespsena-go/bid"
)

type Phone struct {
	phone.Phone
	AddressID bid.BID
}

func (p *Phone) Clean(db *aorm.DB) {
	utils.TrimStrings(&p.Note, &p.Phone.Number)
}

type Address struct {
	aorm.AuditedModel
	Phones       []Phone         `aorm:"fkc:{field:AddressID;cascade}"`
	ContactName  string          `aorm:"size:255"`
	RegionID     string          `aorm:"size:10"`
	Region       *geocode.Region `aorm:"save_associations:false"`
	AddressLine1 string          `aorm:"size:255"`
	AddressLine2 string          `aorm:"size:255"`
	AddressLine3 string          `aorm:"size:255"`
	AddressLine4 string          `aorm:"size:255"`
	Zip          string          `aorm:"size:32"`
}

func (Address) GetAormInlinePreloadFields() []string {
	return []string{"*", "Region"}
}

func (e *Address) Clean(db *aorm.DB) {
	utils.TrimStrings(&e.ContactName, &e.AddressLine1, &e.AddressLine2, &e.AddressLine3, &e.Zip)
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
