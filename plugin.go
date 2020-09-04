package address

import (
	"github.com/ecletus-pkg/admin"
	"github.com/ecletus-pkg/geocode"
	"github.com/ecletus-pkg/phone"
	"github.com/ecletus/db"
	"github.com/ecletus/plug"
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
	db.Events(p).DBOnMigrate(func(e *db.DBEvent) error {
		return e.AutoMigrate(&Address{}, &Phone{}).Error
	})
}
