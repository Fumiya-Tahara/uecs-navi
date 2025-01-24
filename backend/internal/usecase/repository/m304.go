package repository

import (
	"context"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
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
		IpAddr:     PointerToNullString(newM304.IpAddr),
		NetMask:    PointerToNullString(newM304.NetMask),
		Defgw:      PointerToNullString(newM304.Defgw),
		Dns:        PointerToNullString(newM304.Dns),
		VenderName: newM304.VenderName,
		NodeName:   PointerToNullString(newM304.NodeName),
	}

	id, err := mr.queries.CreateM304(ctx, arg)
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (mr M304Repository) GetM304FromID(ID int) (*domain.M304, error) {
	ctx := context.Background()

	m304Row, err := mr.queries.GetM304FromID(ctx, int32(ID))
	if err != nil {
		return nil, err
	}

	m304 := domain.NewM304(
		int(m304Row.HouseID),
		m304Row.MacAddr,
		m304Row.DhcpFlg,
		NullStringToPointer(m304Row.IpAddr),
		NullStringToPointer(m304Row.NetMask),
		NullStringToPointer(m304Row.Defgw),
		NullStringToPointer(m304Row.Dns),
		m304Row.VenderName,
		NullStringToPointer(m304Row.NodeName),
	)

	return m304, nil
}
