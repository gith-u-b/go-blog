package e

var MsgFlags = map[int]string {
	SUCCESS : "ok",
	ERROR : "fail",
	INVALID_PARAMS : "请求参数错误",
	ERROR_EXIST_TAG : "已存在该标签名称",
	ERROR_NOT_EXIST_TAG : "该标签不存在",
	ERROR_NOT_EXIST_ARTICLE : "该文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL : "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT : "Token已超时",
	ERROR_AUTH_TOKEN : "Token生成失败",
	ERROR_AUTH : "Token错误",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL: "保存图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL: "检查图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验图片错误，图片格式或大小有问题",
	ERROR_CHECK_EXIST_ARTICLE_FAIL: "文章不存在",
	ERROR_GET_ARTICLE_FAIL: "获取文章失败",
	ERROR_GET_ARTICLES_FAIL: "获取文章列表失败",
	ERROR_UPDATE_ARTICLE_FAIL: "更新文章失败",
	ERROR_EXPORT_TAG_FAIL: "导出标签失败",
	ADD_TAG_FAIL: "添加标签失败",
	ERROR_GEN_ARTICLE_POSTER_FAIL: "生成文章海报失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
