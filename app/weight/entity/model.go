package entity

import "awesomeSystem/app/model"

// GetAllPersonsInfo

//type AllDo struct {
//	PageNum int `json:"page_num"`
//	PageSize int `json:"page_size"`
//}

//type AllDocumentsAck struct {
//	Status bool `json:"status"`
//	Res []*entity.WeightRecord `json:"res"`
//	ErrInfo string `json:"err_info"`
//}

type Page struct {
	PageSize int64 `json:"page_size"`
	PageNum  int64 `json:"page_num"`
}

type AllDocumentsAck struct {
	Status  bool             `json:"status"`
	Res     [][]*processNode `json:"res"`
	ErrInfo string           `json:"err_info"`
}

//type AllWeighMaterialRecordPageAck struct {
//	Status  bool                 `json:"status"`
//	Res     weightMaterialRecord `json:"res"`
//	ErrInfo string               `json:"err_info"`
//}

type WeightMaterialRecord struct {
	Data  []model.WeighMaterialRecord `json:"data"`
	Total int64                       `json:"total"`
}

type CalMaterialRecord struct {
	Data  []model.CalculateMaterialRecord `json:"data"`
	Total int64                           `json:"total"`
}

type AllWeighRecordPageAck struct {
	Status  bool                 `json:"status"`
	Res     []allWeighRecordPage `json:"res"`
	ErrInfo string               `json:"err_info"`
}

type allWeighRecordPage struct {
	MaterialCode       string `json:"material_code"`
	MaterialType       string `json:"material_type"`
	MaterialName       string `json:"material_name"`
	Specifications     string `json:"specifications"`
	Supplier           string `json:"supplier"`
	Craft              string `json:"craft"`
	Texture            string `json:"texture"`
	Process            string `json:"process"`
	PurchaseStatus     string `json:"purchase_status"`
	ReceivingWarehouse string `json:"receiving_warehouse"`
}

type aa struct {
	id string `json:"_id"`
}

type AllDocumentsPageAck struct {
	Status  bool       `json:"status"`
	Res     resultData `json:"res"`
	ErrInfo string     `json:"err_info"`
}

type resultData struct {
	Data  [][]*processNode `json:"data"`
	Total int64            `json:"total"`
}

type processNode struct {
	Id                 string         `bson:"_id"`
	MaterialCode       string         `bson:"material_code"`
	MaterialType       string         `bson:"material_type"`
	MaterialName       string         `bson:"material_name"`
	Specifications     string         `bson:"specifications"`
	Supplier           string         `bson:"supplier"`
	Craft              string         `bson:"craft"`
	Texture            string         `bson:"texture"`
	Process            string         `bson:"process"`
	PurchaseStatus     string         `bson:"purchase_status"`
	ReceivingWarehouse string         `bson:"receiving_warehouse"`
	WeighStage         string         `bson:"weigh_stage"`
	RecordLog          []model.Record `bson:"record_log"`
}

type AllParameterAck struct {
	Status  bool       `json:"status"`
	Res     Parameters `json:"res"`
	ErrInfo string     `json:"err_info"`
}

type Parameters struct {
	Craft          []string `json:"crafts"`
	Texture        []string `json:"texture"`
	Process        []string `json:"process"`
	PurchaseStatus []string `json:"purchase_status"`
}

type NewWeighRecord struct {
	MaterialCode       string  `json:"material_code"`
	MaterialType       string  `json:"material_type"`
	MaterialName       string  `json:"material_name"`
	Specifications     string  `json:"specifications"`
	Supplier           string  `json:"supplier"`
	Craft              string  `json:"craft"`
	Texture            string  `json:"texture"`
	Process            string  `json:"process"`
	PurchaseStatus     string  `json:"purchase_status"`
	ReceivingWarehouse string  `json:"receiving_warehouse"`
	WeighStage         string  `json:"weigh_stage"`
	WeighNum           float64 `json:"weigh_num"`
	WeighTime          string  `json:"weigh_time"`
	BpmTaskId          int     `json:"bpm_task_id"`
	CreateTime         string  `json:"create_time"`
}
type NewWeighRecordAck struct {
	Status  bool   `json:"status"`
	Res     string `json:"res"`
	ErrInfo string `json:"err_info"`
}

// NewCalRecord 新增理算
type NewCalRecord struct {
	MaterialCode   string  `json:"material_code"`
	MaterialType   string  `json:"material_type"`
	MaterialName   string  `json:"material_name"`
	Specifications string  `json:"specifications"`
	Craft          string  `json:"craft"`
	Texture        string  `json:"texture"`
	Process        string  `json:"process"`
	PurchaseStatus string  `json:"purchase_status"`
	CalStage       string  `json:"cal_stage"`
	CalNum         float64 `json:"cal_num"`
	CalTime        string  `json:"cal_time"`
	BpmTaskId      int     `json:"bpm_task_id"`
	CreateTime     string  `json:"create_time"`
	Validate       bool    `json:"validate"`
}
type NewCalRecordAck struct {
	Status  bool   `json:"status"`
	Res     string `json:"res"`
	ErrInfo string `json:"err_info"`
}

type NewRecord struct {
	MaterialCode       string  `json:"material_code"`
	MaterialType       string  `json:"material_type"`
	MaterialName       string  `json:"material_name"`
	Specifications     string  `json:"specifications"`
	Supplier           string  `json:"supplier"`
	Craft              string  `json:"craft"`
	Texture            string  `json:"texture"`
	Process            string  `json:"process"`
	PurchaseStatus     string  `json:"purchase_status"`
	ReceivingWarehouse string  `json:"receiving_warehouse"`
	WeighStage         string  `json:"weigh_stage"`
	CalPerson          string  `json:"cal_person"`
	CalWeight          float64 `json:"cal_weight"`
	CalTime            string  `json:"cal_time"`
}

type NewRecordAck struct {
	Status  bool   `json:"status"`
	Res     string `json:"res"`
	ErrInfo string `json:"err_info"`
}

type UpdateFlowRecord struct {
	FlowProcess []model.FlowProcessStage `bson:"flow_process"`
}

type MaterialCode struct {
	MaterialCode string `json:"material_code" form:"material_code" binding:"required"`
}

type WeightMaterialCodeAck struct {
	Status  bool    `json:"status"`
	Res     float64 `json:"res"`
	ErrInfo string  `json:"err_info"`
}

type NewCraft struct {
	Name string `json:"name"`
}

type UpCraft struct {
	//Id       string `json:"id"`
	ClientId string `json:"client_id"`
	Name     string `json:"name"`
}

type NewTexture struct {
	Name string `json:"name"`
}
type UpTexture struct {
	//Id       string `json:"id"`
	ClientId string `json:"client_id"`
	Name     string `json:"name"`
}

type NewProcess struct {
	Name string `json:"name"`
}
type UpProcess struct {
	//Id       string `json:"id"`
	ClientId string `json:"client_id"`
	Name     string `json:"name"`
}

type NewPurchaseStatus struct {
	Name string `json:"name"`
}
type UpPurchaseStatus struct {
	//Id       string `json:"id"`
	ClientId string `json:"client_id"`
	Name     string `json:"name"`
}

type CraftId struct {
	Id string `uri:"id"`
}

type TextureId struct {
	Id string `json:"id"`
}

type ProcessId struct {
	Id string `json:"id"`
}

type PurchaseStatusId struct {
	Id string `json:"id"`
}

//type PurchaseStatus struct {
//	Name string `bson:"name"`
//}

type NewParameterAck struct {
	Status  bool   `json:"status"`
	Res     string `json:"res"`
	ErrInfo string `json:"err_info"`
}

type AllCraftAck struct {
	Status  bool          `json:"status"`
	Res     []model.Craft `json:"res"`
	ErrInfo string        `json:"err_info"`
}

type AllTextureAck struct {
	Status  bool            `json:"status"`
	Res     []model.Texture `json:"res"`
	ErrInfo string          `json:"err_info"`
}

type AllProcessAck struct {
	Status  bool            `json:"status"`
	Res     []model.Process `json:"res"`
	ErrInfo string          `json:"err_info"`
}

type AllPurchaseStatusAck struct {
	Status  bool                   `json:"status"`
	Res     []model.PurchaseStatus `json:"res"`
	ErrInfo string                 `json:"err_info"`
}

type WeighMultiCondition struct {
	MaterialCode       string `json:"material_code" binding:"required"`
	Supplier           string `json:"supplier"`
	Craft              string `json:"craft"`
	Texture            string `json:"texture"`
	Process            string `json:"process"`
	PurchaseStatus     string `json:"purchase_status"`
	ReceivingWarehouse string `json:"receiving_warehouse"`
	Validate           string `json:"validate"`
	PageSize           int    `json:"page_size" binding:"required"`
	PageNum            int    `json:"page_num" binding:"required"`
}

type CalMultiCondition struct {
	MaterialCode       string `json:"material_code" binding:"required"`
	Craft              string `json:"craft"`
	Texture            string `json:"texture"`
	Process            string `json:"process"`
	PurchaseStatus     string `json:"purchase_status"`
	ReceivingWarehouse string `json:"receiving_warehouse"`
	Validate           string `json:"validate"`
	PageSize           int    `json:"page_size" binding:"required"`
	PageNum            int    `json:"page_num" binding:"required"`
}

type MultiConditionAck struct {
	Status  bool   `json:"status"`
	Res     string `json:"res"`
	ErrInfo string `json:"err_info"`
}

type OutPut struct {
	MaterialCode string `json:"material_code"`
	Supplier     string `json:"supplier"`
	Craft        string `json:"craft"`
}

type ResultSlice []OutPut

//type SearchMaterial struct {
//	Code string `json:"code"`
//	Page Page `json:"page"`
//}

type SearchMaterial struct {
	Code  string `json:"code"`
	Start int64  `json:"start"`
	Limit int64  `json:"limit"`
}

type MaterialAck struct {
	Status  bool         `json:"status"`
	Res     materialData `json:"res"`
	ErrInfo string       `json:"err_info"`
}

type materialData struct {
	Data  []model.Material `json:"data"`
	Total int              `json:"total"`
}

type AddMaterial struct {
	InnerCode     string `json:"inner_code"`
	Name          string `json:"name"`
	Code          string `json:"code"`
	Specification string `json:"specification"`
	Status        bool   `json:"status"`
}

type AddMaterialAck struct {
	Status  bool   `json:"status"`
	Res     string `json:"res"`
	ErrInfo string `json:"err_info"`
}

type SearchSupplier struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Start int64  `json:"start"`
	Limit int64  `json:"limit"`
}

type SupplierAck struct {
	Status  bool         `json:"status"`
	Res     supplierData `json:"res"`
	ErrInfo string       `json:"err_info"`
}

type supplier struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Type string `json:"type"`
}

type supplierData struct {
	Data  []supplier `json:"data"`
	Total int        `json:"total"`
}

type supplierType struct {
	Status  bool   `json:"status"`
	Res     string `json:"res"`
	ErrInfo string `json:"err_info"`
}

type SearchWarehouse struct {
	//Code string `json:"code"`
	Name string `json:"name"`
}

type WarehouseAck struct {
	Status  bool          `json:"status"`
	Res     warehouseData `json:"res"`
	ErrInfo string        `json:"err_info"`
}

type warehouse struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type warehouseData struct {
	Data  []warehouse `json:"data"`
	Total int         `json:"total"`
}

//// SrmWarehouseAck SRM 中返回仓库格式
//type SrmWarehouseAck struct {
//	Msg []*category `json:"msg"`
//	Code string `json:"code"`
//}
//type category struct {
//	subStorages []*subStorage `json:"subStorages"`
//	categoryCode string `json:"categoryCode"`
//	categoryName string	`json:"categoryName"`
//}
//type subStorage struct {
//	subCode string `json:"subCode"`
//	subName string `json:"subName"`
//}

type Policy struct {
	Subject string `json:"subject"`
	Object  string `json:"object"`
	Action  string `json:"action"`
}
