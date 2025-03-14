package testdata

import "github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"

func GetTestM304s() []domain.M304 {
	ID1 := 1
	houseID1 := 1
	macAddr1 := "00:11:22:33:44:55"
	dhcpFlg1 := true
	ip1 := "192.168.1.10"
	netmask1 := "255.255.255.0"
	defgw1 := "192.168.1.1"
	dns1 := "8.8.8.8"
	venderName1 := "VendorA"
	nodeName1 := "NodeA"

	ID2 := 2
	houseID2 := 2
	macAddr2 := "00:11:22:33:44:66"
	dhcpFlg2 := false
	ip2 := "192.168.1.20"
	netmask2 := "255.255.255.0"
	defgw2 := "192.168.1.1"
	dns2 := "8.8.4.4"
	venderName2 := "VendorB"
	nodeName2 := "NodeB"

	ID3 := 3
	houseID3 := 3
	macAddr3 := "00:11:22:33:44:77"
	dhcpFlg3 := true
	ip3 := "192.168.1.30"
	netmask3 := "255.255.255.128"
	defgw3 := "192.168.1.129"
	dns3 := "1.1.1.1"
	venderName3 := "VendorC"
	nodeName3 := "NodeC"

	m304s := []domain.M304{
		*domain.NewM304WithID(ID1, houseID1, macAddr1, dhcpFlg1, &ip1, &netmask1, &defgw1, &dns1, venderName1, &nodeName1),
		*domain.NewM304WithID(ID2, houseID2, macAddr2, dhcpFlg2, &ip2, &netmask2, &defgw2, &dns2, venderName2, &nodeName2),
		*domain.NewM304WithID(ID3, houseID3, macAddr3, dhcpFlg3, &ip3, &netmask3, &defgw3, &dns3, venderName3, &nodeName3),
	}

	return m304s
}
