package address

import (
	"github.com/ecletus-pkg/geocode"
	"github.com/ecletus-pkg/phone"
	"github.com/ecletus/admin"
	"github.com/pkg/errors"
)

type AddressGetter interface {
	GetQorAddress() *Address
}

func AddSubResource(setup func(res *admin.Resource), res *admin.Resource, value interface{}, fieldName ...string) error {
	return res.GetAdmin().OnResourcesAdded(func(e *admin.ResourceEvent) error {
		cfg := &admin.Config{
			Setup: func(r *admin.Resource) {
				r.SetI18nModel(&Address{})
				PrepareResource(r)
				res.SetMeta(&admin.Meta{Name: fieldName[0], Resource: r})
				if setup != nil {
					setup(r)
				}
			},
		}

		if len(fieldName) == 0 || fieldName[0] == "" {
			fieldName = []string{"Adresses"}
			res.Meta(&admin.Meta{
				Name:  fieldName[0],
				Label: e.Resource.PluralLabelKey(),
			})
		} else {
			cfg.LabelKey = res.ChildrenLabelKey(fieldName[0])
		}

		res.AddResource(&admin.SubConfig{FieldName: fieldName[0]}, value, cfg)
		return nil
	}, ResourceID)
}

func PrepareResource(res *admin.Resource) {
	if err := phone.AddSubResource(nil, res, &Phone{}); err != nil {
		panic(errors.Wrap(err, "add address phone subresource"))
	}
	if err := geocode.InitRegionMeta(nil, res); err != nil {
		panic(errors.Wrap(err, "add address region"))
	}
	res.EditAttrs(
		"ContactName",
		&admin.Section{Rows: [][]string{{geocode.COUNTRY, geocode.REGION}}},
		&admin.Section{Rows: [][]string{
			{"AddressLine1", "AddressLine2"},
			{"AddressLine3", "AddressLine4"},
		}},
		"Zip",
		"Phones",
	)
	res.ShowAttrs(admin.META_STRINGIFY)
	res.NewAttrs(res.EditAttrs())
	res.IndexAttrs(admin.META_STRINGIFY)
}

const ResourceID = "Address"
