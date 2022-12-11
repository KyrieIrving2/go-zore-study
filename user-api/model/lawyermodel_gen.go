// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	lawyerFieldNames          = builder.RawFieldNames(&Lawyer{})
	lawyerRows                = strings.Join(lawyerFieldNames, ",")
	lawyerRowsExpectAutoSet   = strings.Join(stringx.Remove(lawyerFieldNames, "`id`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`"), ",")
	lawyerRowsWithPlaceHolder = strings.Join(stringx.Remove(lawyerFieldNames, "`id`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`"), "=?,") + "=?"
)

type (
	lawyerModel interface {
		Insert(ctx context.Context, data *Lawyer) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Lawyer, error)
		FindOneByEmail(ctx context.Context, email string) (*Lawyer, error)
		FindOneByLicenseNo(ctx context.Context, licenseNo string) (*Lawyer, error)
		FindOneByPhone(ctx context.Context, phone string) (*Lawyer, error)
		Update(ctx context.Context, data *Lawyer) error
		Delete(ctx context.Context, id int64) error
	}

	defaultLawyerModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Lawyer struct {
		Id                         int64           `db:"id"`                            // 主键
		LawFirmTeamId              int64           `db:"law_firm_team_id"`              // 机构团队id
		AccountId                  int64           `db:"account_id"`                    // 账号编号
		ImageAttachmentId          int64           `db:"image_attachment_id"`           // 照片
		State                      sql.NullInt64   `db:"state"`                         // 状态
		Name                       string          `db:"name"`                          // 姓名
		EnglishName                sql.NullString  `db:"english_name"`                  // 英文名
		LicenseNo                  string          `db:"license_no"`                    // 执业证号
		LicenseYear                sql.NullFloat64 `db:"license_year"`                  // 执业年限
		WorkingYear                sql.NullFloat64 `db:"working_year"`                  // 工作年限
		Gender                     string          `db:"gender"`                        // 性别
		Phone                      string          `db:"phone"`                         // 手机号
		Email                      string          `db:"email"`                         // 邮箱
		GraduateFrom               sql.NullString  `db:"graduate_from"`                 // 毕业院校
		EducationBackgroundId      sql.NullInt64   `db:"education_background_id"`       // 学历
		Birthday                   sql.NullTime    `db:"birthday"`                      // 生日
		LawFirmId                  sql.NullInt64   `db:"law_firm_id"`                   // 从业律所
		TeamScale                  int64           `db:"team_scale"`                    // 团队规模
		ProfessionalBackground     sql.NullString  `db:"professional_background"`       // 专业背景
		IsInternetServedExperience int64           `db:"is_internet_served_experience"` // 互联网服务经历
		CertificateType            sql.NullInt64   `db:"certificateType"`               // 证件类型
		Certificate                sql.NullString  `db:"certificate"`                   // 证件号
		TencentInviter             sql.NullString  `db:"tencent_inviter"`               // 邀请人
		CreatedBy                  string          `db:"created_by"`                    // 创建人
		CreatedAtUtc               time.Time       `db:"created_at_utc"`                // 创建时间
		UpdatedBy                  sql.NullString  `db:"updated_by"`                    // 更新人
		UpdatedAtUtc               sql.NullTime    `db:"updated_at_utc"`                // 更新时间
		IsInLibrary                sql.NullInt64   `db:"is_in_library"`                 // 是否系统库内供应商
	}
)

func newLawyerModel(conn sqlx.SqlConn) *defaultLawyerModel {
	return &defaultLawyerModel{
		conn:  conn,
		table: "`lawyer`",
	}
}

func (m *defaultLawyerModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultLawyerModel) FindOne(ctx context.Context, id int64) (*Lawyer, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", lawyerRows, m.table)
	var resp Lawyer
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultLawyerModel) FindOneByEmail(ctx context.Context, email string) (*Lawyer, error) {
	var resp Lawyer
	query := fmt.Sprintf("select %s from %s where `email` = ? limit 1", lawyerRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, email)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultLawyerModel) FindOneByLicenseNo(ctx context.Context, licenseNo string) (*Lawyer, error) {
	var resp Lawyer
	query := fmt.Sprintf("select %s from %s where `license_no` = ? limit 1", lawyerRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, licenseNo)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultLawyerModel) FindOneByPhone(ctx context.Context, phone string) (*Lawyer, error) {
	var resp Lawyer
	query := fmt.Sprintf("select %s from %s where `phone` = ? limit 1", lawyerRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, phone)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultLawyerModel) Insert(ctx context.Context, data *Lawyer) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, lawyerRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.LawFirmTeamId, data.AccountId, data.ImageAttachmentId, data.State, data.Name, data.EnglishName, data.LicenseNo, data.LicenseYear, data.WorkingYear, data.Gender, data.Phone, data.Email, data.GraduateFrom, data.EducationBackgroundId, data.Birthday, data.LawFirmId, data.TeamScale, data.ProfessionalBackground, data.IsInternetServedExperience, data.CertificateType, data.Certificate, data.TencentInviter, data.CreatedBy, data.CreatedAtUtc, data.UpdatedBy, data.UpdatedAtUtc, data.IsInLibrary)
	return ret, err
}

func (m *defaultLawyerModel) Update(ctx context.Context, newData *Lawyer) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, lawyerRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.LawFirmTeamId, newData.AccountId, newData.ImageAttachmentId, newData.State, newData.Name, newData.EnglishName, newData.LicenseNo, newData.LicenseYear, newData.WorkingYear, newData.Gender, newData.Phone, newData.Email, newData.GraduateFrom, newData.EducationBackgroundId, newData.Birthday, newData.LawFirmId, newData.TeamScale, newData.ProfessionalBackground, newData.IsInternetServedExperience, newData.CertificateType, newData.Certificate, newData.TencentInviter, newData.CreatedBy, newData.CreatedAtUtc, newData.UpdatedBy, newData.UpdatedAtUtc, newData.IsInLibrary, newData.Id)
	return err
}

func (m *defaultLawyerModel) tableName() string {
	return m.table
}
