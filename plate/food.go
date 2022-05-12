package plate

import "github.com/chenyingzhou/octopus-go/consts"

type Food struct {
	SourceType consts.SourceType
	DataSource string
	DataSet    string
	Ids        []int32
	Rows       []map[string]string
	Conditions []string
}
