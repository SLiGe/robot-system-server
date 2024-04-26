package server

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetsac9cf9c7100f4adad9e81985f5ecea2d0c5e0aeb = "<!DOCTYPE html>\n{{ define \"lq/yllq.html\" }}\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n    <title>月老灵签</title>\n    <link rel=\"stylesheet\" href=\"/css/lq-common.css\">\n</head>\n<body>\n<main>\n    <div class=\"title\">\n        <h1>月老灵签第{{.data.Id}}签解签</h1>\n    </div>\n    <div class=\"content\">\n        <p><strong>{{.data.Title}}</strong></p>\n        <p><strong>签诗：</strong>{{replace .data.QianShi .newLineChar \"<br/>\" }}</p>\n        <p>\n            <img src=\"/file/img/lq-images/{{.data.LqImg}}\" width=\"200\" height=\"403\" alt=\"{{.data.Title}}\">\n        </p>\n        <p><strong>解签：</strong></p>\n        <p>{{replace .data.JieQian .newLineChar \"<br/>\"}}</p>\n        {{ if ne .data.Zhu \"\" }} <p>\n        <strong>注：</strong>{{ replace .data.Zhu .newLineChar \"<br/>\" }}\n    </p>\n        {{end}}\n        {{ if ne .data.Bhqs \"\" }}\n        <div>\n            <p><strong>白话浅释：</strong></p>\n            <p>{{ replace .data.Bhqs .newLineChar \"<br/>\" }}</p>\n        </div>\n        {{end}}\n    </div>\n</main>\n</body>\n</html>\n{{ end }}"
var _Assets0d4ea967719eea12e7d7378729b41a83d6a5cb78 = "<!DOCTYPE html>\n{{ define \"lq/cslq.html\" }}\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n    <title>财神灵签</title>\n    <link rel=\"stylesheet\" href=\"/css/lq-common.css\">\n</head>\n<body>\n<main>\n    <div class=\"title\">\n        <h1>财神灵签第{{.data.Id}}签解签</h1>\n    </div>\n    <div class=\"content\">\n        <p><strong>{{.title}}</strong></p>\n        <p>{{replace .data.Desc .newLineChar \"<br />\" }}</p>\n        <p><strong>诗曰：</strong>{{ replace .data.ShiYue .newLineChar \"<br />\" }}</p>\n        <p>\n            <img src=\"/file/img/lq-images/{{.data.LqImg}}\" width=\"200\" height=\"403\" alt=\"{{.data.Id}}\">\n        </p>\n        <p><strong>米力仙注：</strong></p>\n        <p>{{replace .data.Mlxz .newLineChar \"<br />\"}}</p>\n        <p><strong>吉凶：</strong>{{replace .data.JiXiong .newLineChar \"<br />\"}}</p>\n    </div>\n</main>\n</body>\n</html>\n{{ end }}"
var _Assetsa6cc76994691a85ebac5f36b0ed94cdbf652cde2 = "<!DOCTYPE html>\n{{ define \"lq/gylq.html\" }}\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n    <title>观音灵签</title>\n    <link rel=\"stylesheet\" href=\"/css/lq-common.css\">\n</head>\n<body>\n<main>\n    {{$newLineChar := .newLineChar}}\n    {{with .data}}\n    <div class=\"title\">\n        <h1>观音灵签第{{.Id}}签解签</h1>\n    </div>\n    <div class=\"content\">\n        <p><strong>{{.Title}}</strong></p>\n        <p><strong>诗曰：</strong>{{replace .ShiYue $newLineChar \"<br/>\" }}</p>\n        <p><strong>诗意：</strong>{{replace .ShiYi $newLineChar \"<br/>\" }}</p>\n        <p><strong>解曰：</strong>{{replace .JieYue $newLineChar \"<br/>\" }}</p>\n        <p><strong>详解：</strong>{{replace .XiangJie $newLineChar \"<br/>\" }}</p>\n        <p><strong>仙机：</strong>{{replace .XianJi $newLineChar \"<br/>\" }}</p>\n        <p>\n            <img src=\"/file/img/lq-images/{{.LqImg}}\" width=\"200\" height=\"403\" alt=\"{{.Title}}\">\n        </p>\n        <p><strong>观音灵签第：{{.Id}}签：整体解签</strong></p>\n        <p>{{replace .Ztjq $newLineChar \"<br/>\" }}</p>\n        <p><strong>本签精髓：</strong>{{replace .Bqjs $newLineChar \"<br/>\" }}</p>\n        <p><strong>凡事做事：</strong>{{replace .Fszs $newLineChar \"<br/>\" }}</p>\n        <p><strong>爱情婚姻：</strong>{{replace .Aqhy $newLineChar \"<br/>\" }}</p>\n\n        {{ if ne .Gzqz \"\" }}<p>\n        <strong>工作求职 创业事业：</strong>{{ replace .Gzqz $newLineChar \"<br/>\" }}\n    </p>{{end}}\n        {{ if ne .Ksjs \"\" }}<p>\n        <strong>考试竞赛 升迁竞选：</strong>{{ replace .Ksjs $newLineChar \"<br/>\" }}\n    </p>{{end}}\n        {{ if ne .Tzlc \"\" }}<p>\n        <strong>投资理财：</strong>{{ replace .Tzlc $newLineChar \"<br/>\" }}\n    </p>{{end}}\n        {{ if ne .Jssy \"\" }}<p>\n        <strong>经商生意：</strong>{{ replace .Jssy $newLineChar \"<br/>\" }}\n    </p>{{end}}\n        {{ if ne .Fdjy \"\" }}<p>\n        <strong>房地交易：</strong>{{ replace .Fdjy $newLineChar \"<br/>\" }}\n    </p>{{end}}\n\n        {{ if ne .Zbjk \"\" }}<p>\n        <strong>治病健康：</strong>{{ replace .Zbjk $newLineChar \"<br/>\" }}\n    </p>{{end}}\n        {{ if ne .Zhbg \"\" }}<p>\n        <strong>转换变更：</strong>{{ replace .Zhbg $newLineChar \"<br/>\" }}\n    </p>{{end}}\n        {{ if ne .Qyqz \"\" }}<p>\n        <strong>求孕求子：</strong>{{ replace .Qyqz $newLineChar \"<br/>\" }}\n    </p>{{end}}\n        {{ if ne .Gsss \"\" }}<p>\n        <strong>官司诉讼：</strong>{{ replace .Gsss $newLineChar \"<br/>\" }}\n    </p>{{end}}\n        {{ if ne .Xrxw \"\" }}<p>\n        <strong>寻人寻物：</strong>{{ replace .Xrxw $newLineChar \"<br/>\" }}\n    </p>{{end}}\n        {{ if ne .Yxcg \"\" }}<p>\n        <strong>远行出国：</strong>{{ replace .Yxcg $newLineChar \"<br/>\" }}\n    </p>{{end}}\n        {{if gt (len .QsgsList) 0}}\n        <div>\n            {{ range .QsgsList }}\n            <p><strong>{{ .Title }}</strong></p>\n            <div>\n                {{ range $key,$value := (split .Content $newLineChar) }}\n                <p>{{ $value }}</p>\n                {{ end }}\n            </div>\n            {{ end }}\n        </div>\n        {{ end }}\n    </div>\n    {{end}}\n</main>\n</body>\n</html>\n{{ end }}"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"web"}, "/web": []string{"templates"}, "/web/templates": []string{}, "/web/templates/lq": []string{"cslq.html", "gylq.html", "yllq.html"}}, map[string]*assets.File{
	"/web/templates/lq/cslq.html": &assets.File{
		Path:     "/web/templates/lq/cslq.html",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1714114086, 1714114086733819900),
		Data:     []byte(_Assets0d4ea967719eea12e7d7378729b41a83d6a5cb78),
	}, "/web/templates/lq/gylq.html": &assets.File{
		Path:     "/web/templates/lq/gylq.html",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1714114253, 1714114253356907200),
		Data:     []byte(_Assetsa6cc76994691a85ebac5f36b0ed94cdbf652cde2),
	}, "/web/templates/lq/yllq.html": &assets.File{
		Path:     "/web/templates/lq/yllq.html",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1714114253, 1714114253349608200),
		Data:     []byte(_Assetsac9cf9c7100f4adad9e81985f5ecea2d0c5e0aeb),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1714115460, 1714115460233913500),
		Data:     nil,
	}, "/web": &assets.File{
		Path:     "/web",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1714115222, 1714115222595877000),
		Data:     nil,
	}, "/web/templates": &assets.File{
		Path:     "/web/templates",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1714115222, 1714115222601072900),
		Data:     nil,
	}, "/web/templates/lq": &assets.File{
		Path:     "/web/templates/lq",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1714115222, 1714115222616135500),
		Data:     nil,
	}}, "")
