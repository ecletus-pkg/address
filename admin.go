package address

import (
	"github.com/ecletus-pkg/geocode"
	"github.com/ecletus-pkg/phone"
	"github.com/ecletus/admin"
)

type AddressGetter interface {
	GetQorAddress() *Address
}

func AddSubResource(res *admin.Resource, value interface{}, fieldName ...string) *admin.Resource {
	cfg := &admin.Config{
		Setup: func(r *admin.Resource) {
			r.SetI18nModel(&Address{})
			PrepareResource(r)
			res.SetMeta(&admin.Meta{Name: fieldName[0], Resource: r})
		},
	}

	if len(fieldName) == 0 || fieldName[0] == "" {
		fieldName = []string{"Adresses"}
		res.Meta(&admin.Meta{
			Name:  fieldName[0],
			Label: GetResource(res.GetAdmin()).PluralLabelKey(),
		})
	} else {
		cfg.LabelKey = res.ChildrenLabelKey(fieldName[0])
	}

	return res.AddResource(&admin.SubConfig{FieldName: fieldName[0]}, value, cfg)
}

func PrepareResource(res *admin.Resource) {
	phone.AddSubResource(res, &AddressPhone{})
	geocode.InitRegionMeta(res)
	res.EditAttrs(
		"ContactName",
		&admin.Section{Rows: [][]string{{geocode.COUNTRY, geocode.REGION}}},
		&admin.Section{Rows: [][]string{
			{"AddressLine1", "AddressLine2"},
			{"AddressLine3", "AddressLine4"},
		}},
		"Cep",
		"Phones",
	)
	res.ShowAttrs(admin.META_STRING)
	res.NewAttrs(res.EditAttrs())
	res.IndexAttrs("String")
}

func GetResource(Admin *admin.Admin) *admin.Resource {
	return Admin.GetResourceByID("Address")
}
