package dto

// SaveDictTypeRequest 保存字典类型
type SaveDictTypeRequest struct {
	DictId   int    `json:"dictId"`
	DictName string `json:"dictName"`
	DictType string `json:"dictType"`
	Status   string `json:"status"`
	CreateBy string `json:"createBy"`
	UpdateBy string `json:"updateBy"`
	Remark   string `json:"remark"`
}

// DictTypeListRequest 字典类型列表
type DictTypeListRequest struct {
	PageRequest
	DictName  string `query:"dictName" form:"dictName"`
	DictType  string `query:"dictType" form:"dictType"`
	Status    string `query:"status" form:"status"`
	BeginTime string `query:"params[beginTime]" form:"params[beginTime]"`
	EndTime   string `query:"params[endTime]" form:"params[endTime]"`
}

// CreateDictTypeRequest 新增字典类型
type CreateDictTypeRequest struct {
	DictName string `json:"dictName"`
	DictType string `json:"dictType"`
	Status   string `json:"status"`
	Remark   string `json:"remark"`
}

// UpdateDictTypeRequest 更新字典类型
type UpdateDictTypeRequest struct {
	DictId   int    `json:"dictId"`
	DictName string `json:"dictName"`
	DictType string `json:"dictType"`
	Status   string `json:"status"`
	Remark   string `json:"remark"`
}

// SaveDictDataRequest 保存字典数据
type SaveDictDataRequest struct {
	DictCode  int    `json:"dictCode"`
	DictSort  int    `json:"dictSort"`
	DictLabel string `json:"dictLabel"`
	DictValue string `json:"dictValue"`
	DictType  string `json:"dictType"`
	CssClass  string `json:"cssClass"`
	ListClass string `json:"listClass"`
	IsDefault string `json:"isDefault"`
	Status    string `json:"status"`
	CreateBy  string `json:"createBy"`
	UpdateBy  string `json:"updateBy"`
	Remark    string `json:"remark"`
}

// DictDataListRequest 字典数据列表
type DictDataListRequest struct {
	PageRequest
	DictType  string `query:"dictType" form:"dictType"`
	DictLabel string `query:"dictLabel" form:"dictLabel"`
	Status    string `query:"status" form:"status"`
}

// CreateDictDataRequest 新增字典数据
type CreateDictDataRequest struct {
	DictSort  int    `json:"dictSort"`
	DictLabel string `json:"dictLabel"`
	DictValue string `json:"dictValue"`
	DictType  string `json:"dictType"`
	CssClass  string `json:"cssClass"`
	ListClass string `json:"listClass"`
	IsDefault string `json:"isDefault"`
	Status    string `json:"status"`
	Remark    string `json:"remark"`
}

// UpdateDictDataRequest 更新字典数据
type UpdateDictDataRequest struct {
	DictCode  int    `json:"dictCode"`
	DictSort  int    `json:"dictSort"`
	DictLabel string `json:"dictLabel"`
	DictValue string `json:"dictValue"`
	DictType  string `json:"dictType"`
	CssClass  string `json:"cssClass"`
	ListClass string `json:"listClass"`
	IsDefault string `json:"isDefault"`
	Status    string `json:"status"`
	Remark    string `json:"remark"`
}
