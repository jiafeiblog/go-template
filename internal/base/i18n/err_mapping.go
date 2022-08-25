package i18n

import "template/pkg/mistake"

const (
	// ErrFailedToCreateOrder 创建订单失败
	ErrFailedToCreateOrder = 320001
	// ErrUploadImage 上传文件失败
	ErrUploadImage = 320002
	// ErrUploadImageFormat 上传文件格式错误
	ErrUploadImageFormat = 320003
	// ErrOrderShared 已经分享了
	ErrOrderShared = 320004
)

func init() {
	mistake.RegisterCodeErrorInfo(ErrFailedToCreateOrder, "Failed to create order")
	mistake.RegisterCodeErrorInfo(ErrUploadImage, "upload image failed")
	mistake.RegisterCodeErrorInfo(ErrUploadImageFormat, "upload image format failed")
	mistake.RegisterCodeErrorInfo(ErrOrderShared, "already shared")
}
