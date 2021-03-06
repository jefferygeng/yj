package autocode

import (
	"github.com/jefferygeng/yj/server/global"
	"github.com/jefferygeng/yj/server/model/autocode"
	autoCodeReq "github.com/jefferygeng/yj/server/model/autocode/request"
	"github.com/jefferygeng/yj/server/model/common/request"
)

type AgreementService struct {
}

// CreateAgreement 创建Agreement记录
// Author [piexlmax](https://github.com/piexlmax)
func (agreementService *AgreementService) CreateAgreement(agreement autocode.Agreement) (err error) {
	err = global.GVA_DB.Create(&agreement).Error
	return err
}

// DeleteAgreement 删除Agreement记录
// Author [piexlmax](https://github.com/piexlmax)
func (agreementService *AgreementService) DeleteAgreement(agreement autocode.Agreement) (err error) {
	err = global.GVA_DB.Delete(&agreement).Error
	return err
}

// DeleteAgreementByIds 批量删除Agreement记录
// Author [piexlmax](https://github.com/piexlmax)
func (agreementService *AgreementService) DeleteAgreementByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Agreement{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAgreement 更新Agreement记录
// Author [piexlmax](https://github.com/piexlmax)
func (agreementService *AgreementService) UpdateAgreement(agreement autocode.Agreement) (err error) {
	err = global.GVA_DB.Save(&agreement).Error
	return err
}

// GetAgreement 根据id获取Agreement记录
// Author [piexlmax](https://github.com/piexlmax)
func (agreementService *AgreementService) GetAgreement(id uint) (err error, agreement autocode.Agreement) {
	err = global.GVA_DB.Where("id = ?", id).First(&agreement).Error
	return
}

// GetAgreementInfoList 分页获取Agreement记录
// Author [piexlmax](https://github.com/piexlmax)
func (agreementService *AgreementService) GetAgreementInfoList(info autoCodeReq.AgreementSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.Agreement{})
	var agreements []autocode.Agreement
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Currency != nil {
		db = db.Where("`currency` = ?", info.Currency)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Supplier").Find(&agreements).Error
	return err, agreements, total
}
