package domain

type M304 struct {
	ID         int
	HouseID    int
	MacAddr    string
	DhcpFlg    bool
	IpAddr     *string
	NetMask    *string
	Defgw      *string
	Dns        *string
	VenderName string
	NodeName   *string
}

func NewM304(houseID int, macAddr string, dhcpFlg bool, ipAddr *string, netMask *string, defgw *string, dns *string, venderName string, nodeName *string) *M304 {
	return &M304{
		HouseID:    houseID,
		MacAddr:    macAddr,
		DhcpFlg:    dhcpFlg,
		IpAddr:     ipAddr,
		NetMask:    netMask,
		Defgw:      defgw,
		Dns:        dns,
		VenderName: venderName,
		NodeName:   nodeName,
	}
}

func NewM304WithID(id int, houseID int, macAddr string, dhcpFlg bool, ipAddr *string, netMask *string, defgw *string, dns *string, venderName string, nodeName *string) *M304 {
	return &M304{
		ID:         id,
		HouseID:    houseID,
		MacAddr:    macAddr,
		DhcpFlg:    dhcpFlg,
		IpAddr:     ipAddr,
		NetMask:    netMask,
		Defgw:      defgw,
		Dns:        dns,
		VenderName: venderName,
		NodeName:   nodeName,
	}
}
