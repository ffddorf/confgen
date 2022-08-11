package edgeos

import "github.com/ffddorf/confgen/netbox/models"

type InterfaceType string

const (
	InterfaceTypeBonding        InterfaceType = "bonding"         // Bonding interface name
	InterfaceTypeBridge         InterfaceType = "bridge"          // Bridge interface name
	InterfaceTypeEthernet       InterfaceType = "ethernet"        // Ethernet interface name
	InterfaceTypeInput          InterfaceType = "input"           // Input functional block (IFB) interface name
	InterfaceTypeIpv6Tunnel     InterfaceType = "ipv6-tunnel"     // IPv6 Tunnel interface
	InterfaceTypeL2tpClient     InterfaceType = "l2tp-client"     // L2TP client interface name
	InterfaceTypeL2tpv3         InterfaceType = "l2tpv3"          // L2TPv3 interface
	InterfaceTypeLoopback       InterfaceType = "loopback"        // Loopback interface name
	InterfaceTypeOpenvpn        InterfaceType = "openvpn"         // OpenVPN tunnel interface name
	InterfaceTypePptpClient     InterfaceType = "pptp-client"     // PPTP client interface name
	InterfaceTypePseudoEthernet InterfaceType = "pseudo-ethernet" // Pseudo Ethernet device name
	InterfaceTypeSwitch         InterfaceType = "switch"          // Switch interface name
	InterfaceTypeTunnel         InterfaceType = "tunnel"          // Tunnel interface
	InterfaceTypeVti            InterfaceType = "vti"             // Virtual Tunnel interface
	InterfaceTypeWirelessmodem  InterfaceType = "wirelessmodem"   // Wireless modem interface name
)

func InterfaceTypeFromNetbox(netboxType models.DcimInterfaceTypeChoices) (edgeosType InterfaceType, ok bool) {
	ok = true
	switch netboxType {
	case models.DcimInterfaceTypeChoicesA100baseTx,
		models.DcimInterfaceTypeChoicesA1000baseT,
		models.DcimInterfaceTypeChoicesA25gbaseT,
		models.DcimInterfaceTypeChoicesA5gbaseT,
		models.DcimInterfaceTypeChoicesA10gbaseT,
		models.DcimInterfaceTypeChoicesA10gbaseCx4,
		models.DcimInterfaceTypeChoicesA1000baseXGbic,
		models.DcimInterfaceTypeChoicesA1000baseXSfp,
		models.DcimInterfaceTypeChoicesA10gbaseXSfpp,
		models.DcimInterfaceTypeChoicesA10gbaseXXfp,
		models.DcimInterfaceTypeChoicesA10gbaseXXenpak,
		models.DcimInterfaceTypeChoicesA10gbaseXX2,
		models.DcimInterfaceTypeChoicesA25gbaseXSfp28,
		models.DcimInterfaceTypeChoicesA50gbaseXSfp56,
		models.DcimInterfaceTypeChoicesA40gbaseXQsfpp,
		models.DcimInterfaceTypeChoicesA50gbaseXSfp28,
		models.DcimInterfaceTypeChoicesA100gbaseXCfp,
		models.DcimInterfaceTypeChoicesA100gbaseXCfp2,
		models.DcimInterfaceTypeChoicesA200gbaseXCfp2,
		models.DcimInterfaceTypeChoicesA100gbaseXCfp4,
		models.DcimInterfaceTypeChoicesA100gbaseXCpak,
		models.DcimInterfaceTypeChoicesA100gbaseXQsfp28,
		models.DcimInterfaceTypeChoicesA200gbaseXQsfp56,
		models.DcimInterfaceTypeChoicesA400gbaseXQsfpdd,
		models.DcimInterfaceTypeChoicesA400gbaseXOsfp,
		models.DcimInterfaceTypeChoicesA1gfcSfp,
		models.DcimInterfaceTypeChoicesA2gfcSfp,
		models.DcimInterfaceTypeChoicesA4gfcSfp,
		models.DcimInterfaceTypeChoicesA8gfcSfpp,
		models.DcimInterfaceTypeChoicesA16gfcSfpp,
		models.DcimInterfaceTypeChoicesA32gfcSfp28,
		models.DcimInterfaceTypeChoicesA64gfcQsfpp,
		models.DcimInterfaceTypeChoicesA128gfcQsfp28:
		edgeosType = InterfaceTypeEthernet
	case models.DcimInterfaceTypeChoicesBridge:
		edgeosType = InterfaceTypeBridge
	case models.DcimInterfaceTypeChoicesLag:
		edgeosType = InterfaceTypeBonding
	case models.DcimInterfaceTypeChoicesVirtual:
		edgeosType = InterfaceTypeLoopback
	default:
		ok = false
	}
	return
}
