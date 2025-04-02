package repository

import (
	"context"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/utils"
)

type M304Repository struct {
	queries *mysqlc.Queries
}

func NewM304Repository(queries *mysqlc.Queries) *M304Repository {
	return &M304Repository{
		queries: queries,
	}
}

func (mr M304Repository) CreateM304(newM304 domain.M304) (int, error) {
	ctx := context.Background()

	arg := mysqlc.CreateM304Params{
		HouseID:    int32(newM304.HouseID),
		MacAddr:    newM304.MacAddr,
		DhcpFlg:    newM304.DhcpFlg,
		IpAddr:     utils.PointerToNullString(newM304.IpAddr),
		NetMask:    utils.PointerToNullString(newM304.NetMask),
		Defgw:      utils.PointerToNullString(newM304.Defgw),
		Dns:        utils.PointerToNullString(newM304.Dns),
		VenderName: newM304.VenderName,
		NodeName:   utils.PointerToNullString(newM304.NodeName),
	}

	id, err := mr.queries.CreateM304(ctx, arg)
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (mr M304Repository) GetAllM304s() (*[]domain.M304, error) {
	ctx := context.Background()

	m304sRow, err := mr.queries.GetAllM304s(ctx)
	if err != nil {
		return nil, err
	}

	m304s := make([]domain.M304, len(m304sRow))
	for i, m304 := range m304sRow {
		m304s[i] = *domain.NewM304WithID(int(m304.ID), m304.UecsID.String, int(m304.HouseID), m304.MacAddr, m304.DhcpFlg, &m304.IpAddr.String, &m304.NetMask.String, &m304.Defgw.String, &m304.Dns.String, m304.VenderName, &m304.NodeName.String)
	}

	return &m304s, nil
}

func (mr M304Repository) GetM304FromID(ID int) (*domain.M304, error) {
	ctx := context.Background()

	m304Row, err := mr.queries.GetM304FromID(ctx, int32(ID))
	if err != nil {
		return nil, err
	}

	m304 := domain.NewM304WithID(
		int(m304Row.ID),
		m304Row.UecsID.String,
		int(m304Row.HouseID),
		m304Row.MacAddr,
		m304Row.DhcpFlg,
		utils.NullStringToPointer(m304Row.IpAddr),
		utils.NullStringToPointer(m304Row.NetMask),
		utils.NullStringToPointer(m304Row.Defgw),
		utils.NullStringToPointer(m304Row.Dns),
		m304Row.VenderName,
		utils.NullStringToPointer(m304Row.NodeName),
	)

	return m304, nil
}
