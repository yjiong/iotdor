// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/yjiong/iotdor/ent/gateway"
	"github.com/yjiong/iotdor/ent/group"
)

// Gateway is the model entity for the Gateway schema.
type Gateway struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Gwid holds the value of the "gwid" field.
	Gwid string `json:"gwid,omitempty"`
	// Svrid holds the value of the "svrid" field.
	Svrid string `json:"svrid,omitempty"`
	// Broker holds the value of the "broker" field.
	Broker string `json:"broker,omitempty"`
	// InstallationLocation holds the value of the "installation_location" field.
	InstallationLocation string `json:"installation_location,omitempty"`
	// Stat holds the value of the "stat" field.
	Stat string `json:"stat,omitempty"`
	// DeleteFlag holds the value of the "delete_flag" field.
	DeleteFlag bool `json:"delete_flag,omitempty"`
	// UpInterval holds the value of the "up_interval" field.
	UpInterval int `json:"up_interval,omitempty"`
	// Version holds the value of the "version" field.
	Version string `json:"version,omitempty"`
	// Summary holds the value of the "summary" field.
	Summary string `json:"summary,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GatewayQuery when eager-loading is set.
	Edges          GatewayEdges `json:"edges"`
	group_gateways *int
}

// GatewayEdges holds the relations/edges for other nodes in the graph.
type GatewayEdges struct {
	// Devices holds the value of the devices edge.
	Devices []*Device `json:"devices,omitempty"`
	// Group holds the value of the group edge.
	Group *Group `json:"group,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// DevicesOrErr returns the Devices value or an error if the edge
// was not loaded in eager-loading.
func (e GatewayEdges) DevicesOrErr() ([]*Device, error) {
	if e.loadedTypes[0] {
		return e.Devices, nil
	}
	return nil, &NotLoadedError{edge: "devices"}
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GatewayEdges) GroupOrErr() (*Group, error) {
	if e.loadedTypes[1] {
		if e.Group == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: group.Label}
		}
		return e.Group, nil
	}
	return nil, &NotLoadedError{edge: "group"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Gateway) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case gateway.FieldDeleteFlag:
			values[i] = new(sql.NullBool)
		case gateway.FieldID, gateway.FieldUpInterval:
			values[i] = new(sql.NullInt64)
		case gateway.FieldGwid, gateway.FieldSvrid, gateway.FieldBroker, gateway.FieldInstallationLocation, gateway.FieldStat, gateway.FieldVersion, gateway.FieldSummary:
			values[i] = new(sql.NullString)
		case gateway.FieldCreateTime, gateway.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case gateway.ForeignKeys[0]: // group_gateways
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Gateway", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Gateway fields.
func (ga *Gateway) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case gateway.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ga.ID = int(value.Int64)
		case gateway.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				ga.CreateTime = value.Time
			}
		case gateway.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				ga.UpdateTime = value.Time
			}
		case gateway.FieldGwid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gwid", values[i])
			} else if value.Valid {
				ga.Gwid = value.String
			}
		case gateway.FieldSvrid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field svrid", values[i])
			} else if value.Valid {
				ga.Svrid = value.String
			}
		case gateway.FieldBroker:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field broker", values[i])
			} else if value.Valid {
				ga.Broker = value.String
			}
		case gateway.FieldInstallationLocation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field installation_location", values[i])
			} else if value.Valid {
				ga.InstallationLocation = value.String
			}
		case gateway.FieldStat:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field stat", values[i])
			} else if value.Valid {
				ga.Stat = value.String
			}
		case gateway.FieldDeleteFlag:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field delete_flag", values[i])
			} else if value.Valid {
				ga.DeleteFlag = value.Bool
			}
		case gateway.FieldUpInterval:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field up_interval", values[i])
			} else if value.Valid {
				ga.UpInterval = int(value.Int64)
			}
		case gateway.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				ga.Version = value.String
			}
		case gateway.FieldSummary:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field summary", values[i])
			} else if value.Valid {
				ga.Summary = value.String
			}
		case gateway.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field group_gateways", value)
			} else if value.Valid {
				ga.group_gateways = new(int)
				*ga.group_gateways = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryDevices queries the "devices" edge of the Gateway entity.
func (ga *Gateway) QueryDevices() *DeviceQuery {
	return (&GatewayClient{config: ga.config}).QueryDevices(ga)
}

// QueryGroup queries the "group" edge of the Gateway entity.
func (ga *Gateway) QueryGroup() *GroupQuery {
	return (&GatewayClient{config: ga.config}).QueryGroup(ga)
}

// Update returns a builder for updating this Gateway.
// Note that you need to call Gateway.Unwrap() before calling this method if this Gateway
// was returned from a transaction, and the transaction was committed or rolled back.
func (ga *Gateway) Update() *GatewayUpdateOne {
	return (&GatewayClient{config: ga.config}).UpdateOne(ga)
}

// Unwrap unwraps the Gateway entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ga *Gateway) Unwrap() *Gateway {
	_tx, ok := ga.config.driver.(*txDriver)
	if !ok {
		panic("ent: Gateway is not a transactional entity")
	}
	ga.config.driver = _tx.drv
	return ga
}

// String implements the fmt.Stringer.
func (ga *Gateway) String() string {
	var builder strings.Builder
	builder.WriteString("Gateway(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ga.ID))
	builder.WriteString("create_time=")
	builder.WriteString(ga.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(ga.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("gwid=")
	builder.WriteString(ga.Gwid)
	builder.WriteString(", ")
	builder.WriteString("svrid=")
	builder.WriteString(ga.Svrid)
	builder.WriteString(", ")
	builder.WriteString("broker=")
	builder.WriteString(ga.Broker)
	builder.WriteString(", ")
	builder.WriteString("installation_location=")
	builder.WriteString(ga.InstallationLocation)
	builder.WriteString(", ")
	builder.WriteString("stat=")
	builder.WriteString(ga.Stat)
	builder.WriteString(", ")
	builder.WriteString("delete_flag=")
	builder.WriteString(fmt.Sprintf("%v", ga.DeleteFlag))
	builder.WriteString(", ")
	builder.WriteString("up_interval=")
	builder.WriteString(fmt.Sprintf("%v", ga.UpInterval))
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(ga.Version)
	builder.WriteString(", ")
	builder.WriteString("summary=")
	builder.WriteString(ga.Summary)
	builder.WriteByte(')')
	return builder.String()
}

// Gateways is a parsable slice of Gateway.
type Gateways []*Gateway

func (ga Gateways) config(cfg config) {
	for _i := range ga {
		ga[_i].config = cfg
	}
}
