package logic

// cmd name define
const (
	ProtocolVer     = "Iotd"
	SysHelp         = "help"
	InitSysSet      = "set /sys/info"
	InitSysGet      = "get /sys/info"
	GetHistory      = "get /dev/history"
	ListDevSupport  = "list /dev/supported"
	DelDevItem      = "del /dev/item"
	SetSysCommif    = "set /sys/commif"
	ListSysCommif   = "list /sys/commif"
	ListDevItems    = "list /dev/items"
	SetSysTime      = "set /sys/time"
	SetSysInterval  = "set /sys/interval"
	UpdateSysDrive  = "update /sys/drive"
	AddSysSchedule  = "add /sys/schedule"
	DelSysSchedule  = "del /sys/schedule"
	ListSysSchedule = "list /sys/schedule"
	SetSysExternal  = "set /sys/external"
	AutoUpdata      = "push /dev/vars"
)
