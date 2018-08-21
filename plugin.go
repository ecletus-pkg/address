package address

import (
	"github.com/aghape-pkg/geocode"
	"github.com/aghape-pkg/phone"
	"github.com/aghape/admin/adminplugin"
	"github.com/aghape/db"
	"github.com/aghape/plug"
)

type Plugin struct {
	plug.EventDispatcher
	db.DBNames
	adminplugin.AdminNames
}

func (Plugin) After() []interface{} {
	return []interface{}{&geocode.Plugin{}, &phone.Plugin{}}
}

func (p *Plugin) OnRegister() {
	p.AdminNames.OnInitResources(p, func(e *adminplugin.AdminEvent) {
		e.Admin.AddResource(&Address{}, &adminplugin.Config{Setup: PrepareResource, Invisible: true})
	})
	db.Events(p).DBOnMigrateGorm(func(e *db.GormDBEvent) error {
		return e.DB.AutoMigrate(&Address{}, &AddressPhone{}).Error
	})
}
