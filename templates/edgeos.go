package templates

import (
	"strings"

	"github.com/ffddorf/confgen/interop/edgeos"
	"github.com/ffddorf/confgen/netbox"
	"github.com/ffddorf/confgen/netbox/models"
)

func edgeosConfigFromMap(in map[string]interface{}) string {
	out := new(strings.Builder)
	if err := edgeos.ConfigFromMap(out, in, 0); err != nil {
		panic(err)
	}
	return out.String()
}

type ChildInterface = models.DeviceDeviceDeviceTypeInterfacesInterfaceTypeChild_interfacesInterfaceType

type edgeosInterfaceDef struct {
	netbox.Interface
	VIFs []ChildInterface
}

// Structures interfaces coming from netbox
// to be more compatible with the config
// structure of EdgeOS.
func edgeosPrepareInterfaces(interfaces []netbox.Interface) map[edgeos.InterfaceType][]edgeosInterfaceDef {
	groups := make(map[edgeos.InterfaceType][]edgeosInterfaceDef)
	for _, iface := range interfaces {
		// find interface type, skip if not known
		ifType, ok := edgeos.InterfaceTypeFromNetbox(iface.Type)
		if !ok {
			continue
		}

		// skip interfaces that are children
		if iface.Parent.Id != "" {
			continue
		}

		outIface := edgeosInterfaceDef{Interface: iface}

		// add child interfaces
		if len(iface.Child_interfaces) > 0 {
			outIface.VIFs = iface.Child_interfaces
		}

		// add to list in map
		if _, ok := groups[ifType]; !ok {
			groups[ifType] = make([]edgeosInterfaceDef, 0, 1)
		}
		groups[ifType] = append(groups[ifType], outIface)
	}

	return groups
}
