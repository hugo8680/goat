package dto

// SaveConfigRequest 保存参数
type SaveConfigRequest struct {
	ConfigId    int    `json:"configId"`
	ConfigName  string `json:"configName"`
	ConfigKey   string `json:"configKey"`
	ConfigValue string `json:"configValue"`
	ConfigType  string `json:"configType"`
	CreateBy    string `json:"createBy"`
	UpdateBy    string `json:"updateBy"`
	Remark      string `json:"remark"`
}

// ConfigListRequest 参数列表
type ConfigListRequest struct {
	PageRequest
	ConfigName string `query:"configName" form:"configName"`
	ConfigKey  string `query:"configKey" form:"configKey"`
	ConfigType string `query:"configType" form:"configType"`
	BeginTime  string `query:"params[beginTime]" form:"params[beginTime]"`
	EndTime    string `query:"params[endTime]" form:"params[endTime]"`
}

// CreateConfigRequest 新增参数
type CreateConfigRequest struct {
	ConfigName  string `json:"configName"`
	ConfigKey   string `json:"configKey"`
	ConfigValue string `json:"configValue"`
	ConfigType  string `json:"configType"`
	Remark      string `json:"remark"`
}

// UpdateConfigRequest 更新参数
type UpdateConfigRequest struct {
	ConfigId    int    `json:"configId"`
	ConfigName  string `json:"configName"`
	ConfigKey   string `json:"configKey"`
	ConfigValue string `json:"configValue"`
	ConfigType  string `json:"configType"`
	Remark      string `json:"remark"`
}
