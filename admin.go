package address

import (
	"github.com/aghape-pkg/geocode"
	"github.com/aghape-pkg/phone"
	"github.com/aghape/admin"
	"github.com/aghape/admin/admincommon"
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
	admincommon.RecordInfoFields(res)
	phone.AddSubResource(res, &AddressPhone{})
	geocode.InitRegionMeta(res)
	res.ShowAttrs("ContactName", geocode.COUNTRY, geocode.REGION, "AddressLine1", "AddressLine2", "Phones")
	res.EditAttrs(res.ShowAttrs())
	res.NewAttrs(res.EditAttrs())
}

func GetResource(Admin *admin.Admin) *admin.Resource {
	return Admin.GetResourceByID("Address")
}
