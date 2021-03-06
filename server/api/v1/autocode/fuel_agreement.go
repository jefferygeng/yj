package autocode

import (
	"github.com/gin-gonic/gin"
	"github.com/jefferygeng/yj/server/global"
	"github.com/jefferygeng/yj/server/model/autocode"
	autocodeReq "github.com/jefferygeng/yj/server/model/autocode/request"
	"github.com/jefferygeng/yj/server/model/common/request"
	"github.com/jefferygeng/yj/server/model/common/response"
	"github.com/jefferygeng/yj/server/service"
	"go.uber.org/zap"
)

type AgreementApi struct {
}

var agreementService = service.ServiceGroupApp.AutoCodeServiceGroup.AgreementService


// CreateAgreement 创建Agreement
// @Tags 航油系统
// @Summary 创建Agreement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Agreement true "创建Agreement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /agreement/createAgreement [post]
func (agreementApi *AgreementApi) CreateAgreement(c *gin.Context) {
	var agreement autocode.Agreement
	_ = c.ShouldBindJSON(&agreement)
	if err := agreementService.CreateAgreement(agreement); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAgreement 删除Agreement
// @Tags 航油系统
// @Summary 删除Agreement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Agreement true "删除Agreement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /agreement/deleteAgreement [delete]
func (agreementApi *AgreementApi) DeleteAgreement(c *gin.Context) {
	var agreement autocode.Agreement
	_ = c.ShouldBindJSON(&agreement)
	if err := agreementService.DeleteAgreement(agreement); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAgreementByIds 批量删除Agreement
// @Tags 航油系统
// @Summary 批量删除Agreement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Agreement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /agreement/deleteAgreementByIds [delete]
func (agreementApi *AgreementApi) DeleteAgreementByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := agreementService.DeleteAgreementByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAgreement 更新Agreement
// @Tags 航油系统
// @Summary 更新Agreement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Agreement true "更新Agreement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /agreement/updateAgreement [put]
func (agreementApi *AgreementApi) UpdateAgreement(c *gin.Context) {
	var agreement autocode.Agreement
	_ = c.ShouldBindJSON(&agreement)
	if err := agreementService.UpdateAgreement(agreement); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAgreement 用id查询Agreement
// @Tags 航油系统
// @Summary 用id查询Agreement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Agreement true "用id查询Agreement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /agreement/findAgreement [get]
func (agreementApi *AgreementApi) FindAgreement(c *gin.Context) {
	var agreement autocode.Agreement
	_ = c.ShouldBindQuery(&agreement)
	if err, reagreement := agreementService.GetAgreement(agreement.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reagreement": reagreement}, c)
	}
}

// GetAgreementList 分页获取Agreement列表
// @Tags 航油系统
// @Summary 分页获取Agreement列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.AgreementSearch true "分页获取Agreement列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /agreement/getAgreementList [get]
func (agreementApi *AgreementApi) GetAgreementList(c *gin.Context) {
	var pageInfo autocodeReq.AgreementSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := agreementService.GetAgreementInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
