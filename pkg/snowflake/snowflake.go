package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"time"
)

var node *snowflake.Node //并不通过全局node产生id

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	//雪花算法由三部分组成：时间戳、机器ID、同一毫秒内产生的ID序列
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	snowflake.Epoch = st.UnixNano() / 1000000 //用于设置用户ID中时间戳的起始时间
	node, err = snowflake.NewNode(machineID)
	return
}

func GenID() int64 { //使用封装的函数生成ID
	return node.Generate().Int64()
}
