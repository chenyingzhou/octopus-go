package claw

import (
	"github.com/chenyingzhou/octopus-go/consts"
	"time"
)

type Config struct {
	Id           int32
	Name         string
	SourceTree   *SourceTree
	TargetType   consts.SourceType
	TargetSource string
	TargetSet    string
	Comments     string
	Deleted      bool
	CreatedTime  time.Time
	UpdatedTime  time.Time
}

type SourceTree struct {
	SourceType     consts.SourceType //数据来源类型
	DataSource     string            //数据来源
	DataSet        string            //数据集
	Priority       int32             //数据集相同时的优先级
	Subscribed     bool              //是否订阅
	IdColumn       string            //ID列名
	TimeColumn     string            //时间列名
	DeleteColumn   string            //软删除列名
	ExtraCondition string            //额外条件(仅针对MYSQL)
	Fields         []SourceField     //字段映射
	Relations      []SourceRelation  //关联关系
}

type SourceRelation struct {
	SourceTree         SourceTree    // 子节点
	SourceRelationType string        //多字段之间关系(AND/OR)
	Fields             []SourceField // 与子节点的关联字段
}

type SourceField struct {
	Column                string                       //源字段
	Target                string                       //目标字段
	SourceFieldTargetType consts.SourceFieldTargetType //源字段->目标字段 的转换类型
}
