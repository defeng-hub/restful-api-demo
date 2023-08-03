package impl

import (
	"context"
	"database/sql"
	"fmt"
	"restful-api-demo/apps/host"
)

//下水道代码
//完成模型和数据库直接的转换

// host 入库
func (s *MysqlServiceImpl) save(ctx context.Context, ins *host.Host) error {
	var (
		err error
	)
	// 默认值填充
	ins.InjectDefault()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		s.l.Errorf("start tx error, %s", err)
		return fmt.Errorf("开启事务异常")
	}

	//处理事务, 无报错就提交,无报错就回滚
	defer func() {
		if err != nil {
			tx.Rollback()

		} else {
			tx.Commit()
		}
	}()

	// 插入Resource
	rstmt, err := tx.Prepare(insertResourceSQL)
	if err != nil {
		return err
	}
	defer rstmt.Close()

	_, err = rstmt.Exec(
		ins.Id, ins.Vendor, ins.Region, ins.CreateAt, ins.ExpireAt,
		ins.Type, ins.Name, ins.Description, ins.Status,
		ins.UpdateAt, ins.SyncAt, ins.Account, ins.PublicIP, ins.PrivateIP,
	)
	if err != nil {
		return err
	}

	//插入Describe
	prepare, err := tx.Prepare(insertDescribeSQL)
	if err != nil {
		return err
	}
	defer prepare.Close()
	_, err = prepare.Exec(
		ins.Id, ins.CPU, ins.Memory, ins.GPUAmount, ins.GPUSpec,
		ins.OSType, ins.OSName, ins.SerialNumber,
	)
	if err != nil {
		return err
	}
	return nil
}

func (i *MysqlServiceImpl) update(ctx context.Context, ins *host.Host) error {
	var (
		err error
	)

	// 开启一个事务 tx
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// 通过Defer处理事务提交方式
	// 1. 无报错，则Commit 事务
	// 2. 有报错, 则Rollback 事务
	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				i.l.Error("rollback error, %s", err)
			}
		} else {
			if err := tx.Commit(); err != nil {
				i.l.Error("commit error, %s", err)
			}
		}
	}()

	var (
		resStmt, hostStmt *sql.Stmt
	)

	// 更新 Resource表
	resStmt, err = tx.PrepareContext(ctx, updateResourceSQL)
	if err != nil {
		return err
	}
	_, err = resStmt.ExecContext(ctx, ins.Vendor, ins.Region, ins.ExpireAt, ins.Name, ins.Description, ins.Id)
	if err != nil {
		return err
	}

	// 更新 Host表
	hostStmt, err = tx.PrepareContext(ctx, updateHostSQL)
	if err != nil {
		return err
	}
	_, err = hostStmt.ExecContext(ctx, ins.CPU, ins.Memory, ins.Id)
	if err != nil {
		return err
	}

	return nil
}
