# OCR 云码

OCR 采用云码提供的 api

开发文档: <https://www.yunmaiocr.com/doc>

## OCR 类型

云码对不同的 OCR 类型提供了相应的服务

服务类型需在 `type` 字段指定

1. 数英汉字类型

    通用数英1-4位 10110

    通用数英5-8位 10111

    通用数英9~11位 10112

    通用数英12位及以上 10113

    通用数英1~6位plus 10103

    定制-数英5位~qcs 9001

    定制-纯数字4位 193

2. 中文类型

    通用中文字符1~2位 10114

    通用中文字符 3~5位 10115

    通用中文字符6~8位 10116

    通用中文字符9位及以上 10117

    定制-XX西游苦行中文字符 10107

3. 计算类型

    通用数字计算题 50100

    通用中文计算题 50101

    定制-计算题 cni 452

## 其他

详细文档: <https://www.jfbym.com/demo/>
