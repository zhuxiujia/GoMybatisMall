package c_admin

import (
	"github.com/facebookgo/inject"
	"github.com/zhuxiujia/GoMybatisMall/app/com/example/app_context"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/utils"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
	"github.com/zhuxiujia/easy_mvc"
	"net/http"
	"time"
)

type UserController struct {
	easy_mvc.Controller `doc:"用户控制器（后台）"`

	UserService *service.UserService `inject:"UserService"`

	AddressPage func(login_id string, phone string, real_name string, page int, size int) interface{} `path:"/admin/user/platform/address/page" arg:"login_id,phone,real_name,page:0,size:5" doc:"用户地址分页" doc_arg:""`

	Page func(id string, phone string, channel string, real_name string, invitation_code string, inviter_code string, TimeStart string, TimeEnd string, page int, size int) interface{} `path:"/admin/user/platform/page" arg:"id,phone,channel,real_name,invitation_code,inviter_code,time_start,time_end,page:0,size:5" doc:"用户分页"`

	PageExcel func(phone string, channel string, real_name string, invitation_code string, inviter_code string, TimeStart string, TimeEnd string, page int, size int, writer http.ResponseWriter) interface{} `path:"/admin/user/platform/page/excel" arg:"phone,channel,real_name,invitation_code,inviter_code,time_start,time_end,page:0,size:1000,w" doc:"用户分页(导出excel)"`
}

func (it *UserController) Routers() {

	it.AddressPage = func(login_id string, phone string, real_name string, page int, page_size int) interface{} {
		var arg = service.AddressPageDTO{}
		arg.Page = page
		arg.PageSize = page_size
		arg.RealName = real_name
		arg.Phone = phone
		arg.UserId = login_id
		var result, e = it.UserService.AddressPage(arg)
		if e != nil {
			return e
		}
		return vo.ResultVO{}.NewSuccess(result)
	}

	it.Page = func(id string, phone string, channel string, realname string, invitation_code string, inviter_code string, TimeStart string, TimeEnd string, page int, size int) interface{} {

		var timeAble = vo.TimeRangeable{}.New(TimeStart, TimeEnd)
		var startTime = timeAble.TimeStart()
		var endTime = timeAble.TimeEnd()
		var data, e = it.UserService.Page(service.UserPageDTO{
			Id:             id,
			Phone:          phone,
			Channel:        channel,
			Realname:       realname,
			InvitationCode: invitation_code,
			InviterCode:    inviter_code,
			TimeStart:      startTime,
			TimeEnd:        endTime,
			Pageable:       vo.Pageable{}.New(page, size),
		})
		if e != nil {
			return e
		}
		return vo.ResultVO{}.NewSuccess(data)
	}

	it.PageExcel = func(phone string, channel string, real_name string, invitation_code string, inviter_code string, TimeStart string, TimeEnd string, page int, size int, writer http.ResponseWriter) interface{} {
		var res = it.Page("", phone, channel, real_name, invitation_code, inviter_code, TimeStart, TimeEnd, page, size)
		var resVO = res.(vo.ResultVO)
		if resVO.Code != 1 {
			return resVO
		}
		var data = resVO.Data.(vo.PageVO).Content
		writer.Header().Set("Content-Type", "application/vnd.ms-excel")
		writer.Header().Set("Content-Disposition", "attachment;filename=excel_"+time.Now().String()+".xlsx")
		utils.Export(utils.ExportDTO{
			Titles: []utils.ExcelTitle{
				{
					Title: "id",
					Field: "id",
				},
				{
					Title: "手机号",
					Field: "phone",
				},
				{
					Title: "姓名",
					Field: "realname",
				},
				{
					Title: "邀请码(本人)",
					Field: "invitation_code",
				},
				{
					Title: "邀请码(受邀)",
					Field: "inviter_code",
				},
				{
					Title: "地址",
					Field: "address",
				},
				{
					Title: "注册渠道",
					Field: "channel",
				},
				{
					Title: "注册时间",
					Field: "create_time",
				},
			},
			DataArray: data,
		}, writer)
		return nil
	}

	//finish
	it.Init(it)
	app_context.Context.Provide(&inject.Object{
		Value: it,
	})
}
