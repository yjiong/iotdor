// Code generated by entc, DO NOT EDIT.

package ent

import (
	"dcmbroker/ent/device"
	"dcmbroker/ent/gateway"
	"dcmbroker/ent/group"
	"dcmbroker/ent/schema"
	"dcmbroker/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	deviceMixin := schema.Device{}.Mixin()
	deviceMixinFields0 := deviceMixin[0].Fields()
	_ = deviceMixinFields0
	deviceFields := schema.Device{}.Fields()
	_ = deviceFields
	// deviceDescCreateTime is the schema descriptor for create_time field.
	deviceDescCreateTime := deviceMixinFields0[0].Descriptor()
	// device.DefaultCreateTime holds the default value on creation for the create_time field.
	device.DefaultCreateTime = deviceDescCreateTime.Default.(func() time.Time)
	// deviceDescUpdateTime is the schema descriptor for update_time field.
	deviceDescUpdateTime := deviceMixinFields0[1].Descriptor()
	// device.DefaultUpdateTime holds the default value on creation for the update_time field.
	device.DefaultUpdateTime = deviceDescUpdateTime.Default.(func() time.Time)
	// device.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	device.UpdateDefaultUpdateTime = deviceDescUpdateTime.UpdateDefault.(func() time.Time)
	gatewayMixin := schema.Gateway{}.Mixin()
	gatewayMixinFields0 := gatewayMixin[0].Fields()
	_ = gatewayMixinFields0
	gatewayFields := schema.Gateway{}.Fields()
	_ = gatewayFields
	// gatewayDescCreateTime is the schema descriptor for create_time field.
	gatewayDescCreateTime := gatewayMixinFields0[0].Descriptor()
	// gateway.DefaultCreateTime holds the default value on creation for the create_time field.
	gateway.DefaultCreateTime = gatewayDescCreateTime.Default.(func() time.Time)
	// gatewayDescUpdateTime is the schema descriptor for update_time field.
	gatewayDescUpdateTime := gatewayMixinFields0[1].Descriptor()
	// gateway.DefaultUpdateTime holds the default value on creation for the update_time field.
	gateway.DefaultUpdateTime = gatewayDescUpdateTime.Default.(func() time.Time)
	// gateway.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	gateway.UpdateDefaultUpdateTime = gatewayDescUpdateTime.UpdateDefault.(func() time.Time)
	// gatewayDescUpInterval is the schema descriptor for upInterval field.
	gatewayDescUpInterval := gatewayFields[5].Descriptor()
	// gateway.DefaultUpInterval holds the default value on creation for the upInterval field.
	gateway.DefaultUpInterval = gatewayDescUpInterval.Default.(int)
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescName is the schema descriptor for name field.
	groupDescName := groupFields[0].Descriptor()
	// group.NameValidator is a validator for the "name" field. It is called by the builders before save.
	group.NameValidator = func() func(string) error {
		validators := groupDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields0[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields0[1].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
}
