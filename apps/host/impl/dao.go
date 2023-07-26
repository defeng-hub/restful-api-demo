package impl

import (
	"context"
	"fmt"
	"github.com/defeng-hub/restful-api-demo/apps/host"
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
