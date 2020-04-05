package c_admin

import (
	"encoding/json"
	"errors"
	"github.com/facebookgo/inject"
	"github.com/zhuxiujia/GoMybatisMall/app/com/example/app_context"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/enum"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/service"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/utils"
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/vo"
	"github.com/zhuxiujia/easy_mvc"
	"strings"
)

type MallSkuController struct {
	easy_mvc.Controller `doc:"(后台接口)商品"`

	MallSkuService           *service.MallSkuService           `inject:"MallSkuService"`
	MallSpecificationService *service.MallSpecificationService `inject:"MallSpecificationService"`
	MallCoverImageService    *service.MallCoverImageService    `inject:"MallCoverImageService"`
	MallClassService         *service.MallClassService         `inject:"MallClassService"`

	Delete  func(id string) interface{}                                                                                                                                                                                                                                 `path:"/admin/user/mall/sku/delete" arg:"id" doc:"商城商品删除" doc_arg:""`
	Page    func(title string, status *int, class_name string, min_amount *int, max_amount *int, tag1 string, tag2 string, tag3 string, order_by string, sort string, page int, size int) interface{}                                                                   `path:"/admin/user/mall/sku/page" arg:"title,status,class_name,min_amount,max_amount,tag1,tag2,tag3,order_by:create_time,sort:desc,page:0,size:5" doc:"商城商品分页" doc_arg:""`
	Add     func(title string, second_title string, content string, total_num int, shop_amount int, amount int, order_time_limit int, status int, class_name string, tag1 string, tag2 string, tag3 string, cover_images string, specs string) interface{}              `path:"/admin/user/mall/sku/add" arg:"title,second_title,content,total_num,shop_amount,amount,order_time_limit,status,class_name,tag1,tag2,tag3,cover_images,specs" doc:"商城商品添加" doc_arg:""`
	Detail  func(id string) interface{}                                                                                                                                                                                                                                 `path:"/admin/user/mall/sku/detail" arg:"id" doc:"商城商品详情" doc_arg:""`
	Shelves func(id string, status int) interface{}                                                                                                                                                                                                                     `path:"/admin/user/mall/sku/shelves" arg:"id,status" doc:"商城商品上下架" doc_arg:""`
	Update  func(id string, title string, second_title string, content string, total_num int, remain_num int, shop_amount int, amount int, order_time_limit int, status int, class_name string, tag1 string, tag2 string, tag3 string, cover_images string) interface{} `path:"/admin/user/mall/sku/update" arg:"id,title,second_title,content,total_num,remain_num,shop_amount,amount,order_time_limit,status,class_name,tag1,tag2,tag3,cover_images" doc:"商城商品更新" doc_arg:""`
}

func (it *MallSkuController) Routers() {

	it.Delete = func(id string) interface{} {
		var e = it.MallSkuService.Delete(id)
		if e != nil {
			return e
		}
		return vo.ResultVO{}.NewSuccess(nil)
	}

	it.Page = func(title string, status *int, class_name string, min_amount *int, max_amount *int, tag1 string, tag2 string, tag3 string, order_by string, sort string, page int, size int) interface{} {

		var data, e = it.MallSkuService.Page(service.MallSkuPageDTO{
			Title:     title,
			Status:    status,
			ClassName: class_name,
			MinAmount: min_amount,
			MaxAmount: max_amount,
			Order_by:  order_by,
			Sort:      sort,
			Tag1:      tag1,
			Tag2:      tag2,
			Tag3:      tag3,
			Pageable:  vo.Pageable{}.New(page, size),
		})
		if e != nil {
			return e
		}
		return vo.ResultVO{}.NewSuccess(data)
	}

	it.Add = func(title string, second_title string, content string, total_num int, shop_amount int, amount int, order_time_limit int, status int, class_name string, tag1 string, tag2 string, tag3 string, cover_images string, specs string) interface{} {
		var sku = model.MallSku{
			Id:             utils.CreateUUID(),
			Title:          title,
			SecondTitle:    second_title,
			Content:        content,
			TotalNum:       total_num,
			RemainNum:      total_num,
			ShopAmount:     shop_amount,
			Amount:         amount,
			OrderTimeLimit: order_time_limit,
			Status:         enum.NewMallSkuStatus(status),
			ClassName:      class_name,
			Tag1:           tag1,
			Tag2:           tag2,
			Tag3:           tag3,
		}
		var spcsArray = strings.Split(specs, ",")
		var dto = service.MallSkuAddDTO{
			MallSku: sku,
		}
		if len(spcsArray) != 0 {
			dto.Specifications = &spcsArray
		}
		var e = it.MallSkuService.Add(dto)
		if e != nil {
			return e
		}
		if cover_images != "" {
			var imgs = []string{}
			json.Unmarshal([]byte(cover_images), &imgs)
			for _, v := range imgs {
				it.MallCoverImageService.Add(model.MallCoverImage{
					Id:    utils.CreateUUID(),
					Img:   v,
					SkuId: sku.Id,
				})
			}
		}

		return vo.ResultVO{}.NewSuccess(sku)
	}

	it.Update = func(id string, title string, second_title string, content string, total_num int, remain_num int, shop_amount int, amount int, order_time_limit int, status int, class_name string, tag1 string, tag2 string, tag3 string, cover_images string) interface{} {
		sku, e := it.MallSkuService.Find(id)
		if e != nil {
			return e
		}
		sku.Title = title
		sku.SecondTitle = second_title
		sku.Content = content
		sku.TotalNum = total_num
		sku.RemainNum = remain_num
		sku.ShopAmount = shop_amount
		sku.Amount = amount
		sku.OrderTimeLimit = order_time_limit
		sku.Status = enum.NewMallSkuStatus(status)
		sku.ClassName = class_name
		sku.Tag1 = tag1
		sku.Tag2 = tag2
		sku.Tag3 = tag3
		_, e = it.MallSkuService.Update(service.MallSkuAddDTO{
			MallSku: sku.MallSku,
		})
		if e != nil {
			return e
		}
		e = it.MallCoverImageService.DeleteBySkuId(sku.Id)
		if e != nil {
			return e
		}
		var imgs = []string{}
		json.Unmarshal([]byte(cover_images), &imgs)
		for _, v := range imgs {
			it.MallCoverImageService.Add(model.MallCoverImage{
				Id:    utils.CreateUUID(),
				Img:   v,
				SkuId: sku.Id,
			})
		}
		return vo.ResultVO{}.NewSuccess(nil)
	}

	it.Detail = func(id string) interface{} {
		var data, e = it.MallSkuService.Find(id)
		if e != nil {
			return e
		}
		if data.Id != "" {
			//MallSpecificationVOs
			specs, e := it.MallSpecificationService.FindBySkuIds([]string{data.Id})
			if e != nil {
				return e
			}
			var spec = specs[data.Id]
			if spec != nil {
				data.MallSpecificationVOs = spec
			}
			//images
			images, e := it.MallCoverImageService.FindBySkuId([]string{data.Id})
			if e != nil {
				return e
			}
			var img = images[data.Id]
			if img != nil && len(*img) != 0 {
				data.MallCoverImageVOs = img
			}
			//MallClassVOs
			return vo.ResultVO{}.NewSuccess(data)
		}
		return vo.ResultVO{}.NewSuccess(nil)
	}

	it.Shelves = func(id string, status int) interface{} {
		var skuVO, e = it.MallSkuService.Find(id)
		if e != nil {
			return e
		}
		if skuVO.Id == "" {
			return errors.New("该商品不存在！")
		}
		skuVO.Status = enum.NewMallSkuStatus(status)
		return vo.ResultVO{}.NewSuccess(skuVO)
	}

	it.Init(it)
	app_context.Context.Provide(&inject.Object{
		Value: it,
	})
}
