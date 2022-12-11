package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LawFirmModel = (*customLawFirmModel)(nil)

type (
	// LawFirmModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLawFirmModel.
	LawFirmModel interface {
		lawFirmModel
	}

	customLawFirmModel struct {
		*defaultLawFirmModel
	}
)

// NewLawFirmModel returns a model for the database table.
func NewLawFirmModel(conn sqlx.SqlConn) LawFirmModel {
	return &customLawFirmModel{
		defaultLawFirmModel: newLawFirmModel(conn),
	}
}
