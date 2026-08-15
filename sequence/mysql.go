package sequence

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const sqlReplaceIntoStub = `REPLACE INTO sequence (stub) VALUES('a')`

type MySQL struct {
	conn sqlx.SqlConn
}

func NewMySQL(dsn string) *MySQL {
	return &MySQL{
		conn: sqlx.NewMysql(dsn),
	}
}

func (m *MySQL) Next(ctx context.Context) (seq uint64, err error) {
	result, err := m.conn.ExecCtx(ctx, sqlReplaceIntoStub)
	if err != nil {
		logx.Errorw("conn.ExecCtx falied", logx.LogField{Key: "err", Value: err.Error()})
		return 0, err
	}

	lastid, err := result.LastInsertId()
	if err != nil {
		logx.Errorw("rest.LastInsertId falied", logx.LogField{Key: "err", Value: err.Error()})
		return 0, err
	}

	return uint64(lastid), nil

}
