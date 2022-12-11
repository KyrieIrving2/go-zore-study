package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LawyerModel = (*customLawyerModel)(nil)

type (
	// LawyerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLawyerModel.
	LawyerModel interface {
		lawyerModel
	}

	customLawyerModel struct {
		*defaultLawyerModel
	}
)

// NewLawyerModel returns a model for the database table.
func NewLawyerModel(conn sqlx.SqlConn) LawyerModel {
	return &customLawyerModel{
		defaultLawyerModel: newLawyerModel(conn),
	}
}
