package c_app

import (
	"github.com/zhuxiujia/GoMybatisMall/app/com/example/app_context"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"

	"github.com/facebookgo/inject"
	"github.com/zhuxiujia/easy_mvc"
)

type MallSkuApi struct {
	easy_mvc.Controller `doc:"商城商品API"`

	MallClassService *service.MallClassService `inject:"MallClassService"`
	MallSkuService   *service.MallSkuService   `inject:"MallSkuService"`

	ClassList func(name string, page int, size int) interface{} `path:"/api/mall/class/list" arg:"name,page:0,size:20" doc:"商品分类列表,logo_img分类logo图片url" doc_arg:""`
	Page      func(name string,
		status *int,
		class_name string,
		min_amount *int,
		max_amount *int,
		tag1 string,
		tag2 string,
		tag3 string,
		sort string,
		order_by string,
		page int,
		size int) interface{} `path:"/api/mall/sku/list" arg:"name,status:1,class_name,min_amount,max_amount,tag1,tag2,tag3,sort,order_by,page:0,size:20" doc:"商品列表,class_name是分类名称，mall_specification_vo商品规格，mall_cover_image_vos是商品缩略图，second_title副标题，shop_amount市场价，amount零售价，content商品介绍html" doc_arg:"sort:排序字段例如amount,order_by:升序传asc降序传desc,class_name:分类名称,status:1上架2下架,tag1:筛选传'热销',tag2:筛选传'新品',tag3:筛选传'精选'"`
	Detail func(id string) interface{} `path:"/api/mall/sku/detail" arg:"id" doc:"商品详情" doc_arg:""`
}

func (it *MallSkuApi) Routers() {

	it.ClassList = func(name string, page int, size int) interface{} {
		data, e := it.MallClassService.Page(service.MallClassPageDTO{
			Name:     name,
			Pageable: vo.Pageable{}.New(page, size),
		})
		if e != nil {
			return e
		}
		return vo.ResultVO{}.NewSuccess(data)
	}

	it.Page = func(title string, status *int, class_name string, min_amount *int, max_amount *int, tag1 string, tag2 string, tag3 string, sort string,
		order_by string, page int, size int) interface{} {
		if sort == "" {
			sort = "desc"
		}
		if order_by == "" {
			order_by = "create_time"
		}
		data, e := it.MallSkuService.Page(service.MallSkuPageDTO{
			Title:     title,
			Status:    status,
			ClassName: class_name,
			MinAmount: min_amount,
			MaxAmount: max_amount,
			Tag1:      tag1,
			Tag2:      tag2,
			Tag3:      tag3,
			Sort:      sort,
			Order_by:  order_by,
			Pageable:  vo.Pageable{}.New(page, size),
		})
		if e != nil {
			return e
		}
		return vo.ResultVO{}.NewSuccess(data)
	}
	it.Detail = func(id string) interface{} {
		data, e := it.MallSkuService.Find(id)
		if e != nil {
			return e
		}
		if data.Id == "" {
			return vo.ResultVO{}.NewSuccess(nil)
		} else {
			return vo.ResultVO{}.NewSuccess(data)
		}
	}

	it.Init(it)
	app_context.Context.Provide(&inject.Object{
		Value: it,
	})
}
