package address

import (
	"github.com/aghape-pkg/geocode"
	"github.com/aghape-pkg/phone"
	"github.com/aghape/admin"
	"github.com/aghape/admin/admincommon"
	"github.com/aghape/aghape"
)

type AddressGetter interface {
	GetQorAddress() *Address
}

func AddSubResource(res *admin.Resource, value interface{}, fieldName ...string) *admin.Resource {
	if len(fieldName) == 0 || fieldName[0] == "" {
		fieldName = []string{"Adresses"}
	}
	return res.AddResource(&admin.SubConfig{FieldName: fieldName[0]}, value, &admin.Config{
		Setup: func(r *admin.Resource) {
			r.SetI18nModel(&Address{})
			PrepareResource(r)
			res.SetMeta(&admin.Meta{Name: fieldName[0], Resource: r})
		},
	})
}

func PrepareResource(res *admin.Resource) {
	Admin := res.GetAdmin()
	admincommon.RecordInfoFields(res)
	phone.AddSubResource(res, &AddressPhone{})
	countryRes := geocode.GetCountryResource(Admin)
	countryMeta := res.SetMeta(&admin.Meta{
		Name: "Country",
		Config: &admin.SelectOneConfig{
			Layout:             admin.BASIC_LAYOUT_HTML_WITH_ICON,
			RemoteDataResource: admin.NewDataResource(countryRes),
		},
		FormattedValuer: func(record interface{}, context *qor.Context) interface{} {
			if record != nil {
				var adr *Address
				if r, ok := record.(*Address); ok {
					adr = r
				} else if r, ok := record.(AddressGetter); ok {
					adr = r.GetQorAddress()
				}
				if adr != nil && adr.Region != nil {
					if adr.Region.Country != nil {
						return adr.Region.Country.Name
					}
				}
			}
			return ""
		},
		Valuer: func(record interface{}, context *qor.Context) interface{} {
			if record != nil {
				var adr *Address
				if r, ok := record.(*Address); ok {
					adr = r
				} else if r, ok := record.(AddressGetter); ok {
					adr = r.GetQorAddress()
				}
				if adr != nil && adr.Region != nil {
					return adr.Region.CountryID
				}
			}
			return ""
		},
	})

	regionRes := geocode.GetRegionsResource(Admin)
	res.SetMeta(&admin.Meta{
		Name:       "Region",
		Resource:   regionRes,
		Dependency: []interface{}{&admin.DependencyParent{countryMeta}},
	})

	res.ShowAttrs("ContactName", "Country", "Region", "AddressLine1", "AddressLine2", "Phones")
	res.EditAttrs(res.ShowAttrs())
	res.NewAttrs(res.EditAttrs())
}

func GetResource(Admin *admin.Admin) *admin.Resource {
	return Admin.GetResourceByID("Address")
}
