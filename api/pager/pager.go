package pager

import (
	"errors"
	"reflect"
	basePager "repo.wooramel.cn/mission/tools/pager"
	"xorm.io/xorm"
)

// PageModel 分页模型
type PageModel struct {
	Page         int         `json:"page"`
	PageSize     int         `json:"page_size"`
	TotalPages   int64       `json:"total_pages"`
	TotalRecords int64       `json:"total_records"`
	Records      interface{} `json:"records"`
}

// Help xorm 分页
// page: 当前页码
// pageSize: 每页条数
// rowsSlicePtr: 要求传入一个slice的指针类型 &[]model.AuthUser，否则出现错误
// rowElement: slice里面的单个元素，会将其中的非0值元素作为查询条件
// session: 为带有查询条件的session 可以携带一些特殊的查询条件 rowElement 无法做到的，如：like 等
// 如果查询总数为0则不会继续往后查询
// 会赋值传入的slice为查询结果，并返回查询总数total，或执行中出现的error错误
func Help(page, pageSize int, rowsSlicePtr interface{}, rowElement interface{}, session *xorm.Session) (total int64, err error) {
	// 保证分页数据合法性
	page, pageSize = basePager.ValPageAndPageSize(page, pageSize)

	// 保证传入的是slice类型
	sliceValue := reflect.Indirect(reflect.ValueOf(rowsSlicePtr))
	if sliceValue.Kind() != reflect.Slice {
		return 0, errors.New("需要传入一个slice类型")
	}

	// 查询具体数据
	startNum := (page - 1) * pageSize
	err = session.Limit(pageSize, startNum).Find(rowsSlicePtr, rowElement)
	if err != nil {
		return 0, err
	}
	total, err = session.Count(rowElement)
	return
}
