package address

import (
	"github.com/aghape-pkg/geocode"
	"github.com/aghape-pkg/phone"
	"github.com/aghape-pkg/admin"
	"github.com/aghape/db"
	"github.com/aghape/plug"
)

type Plugin struct {
	plug.EventDispatcher
	db.DBNames
	admin_plugin.AdminNames
}

func (Plugin) After() []interface{} {
	return []interface{}{&geocode.Plugin{}, &phone.Plugin{}}
}

func (p *Plugin) OnRegister() {
	admin_plugin.Events(p).InitResources(func(e *admin_plugin.AdminEvent) {
		e.Admin.AddResource(&Address{}, &admin_plugin.Config{Setup: PrepareResource, Invisible: true})
	})
	db.Events(p).DBOnMigrateGorm(func(e *db.GormDBEvent) error {
		return e.DB.AutoMigrate(&Address{}, &AddressPhone{}).Error
	})
}
