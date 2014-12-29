<div style="text-align:center;padding-top:60px;">
	{{if .success}}
	<p style="font-family: cursive;font-size: 28px;font-weight: bold;color: orange;">配对成功</p>
	<p style="font-family: cursive;font-size: 28px;font-weight: bold;color: orange;">{{.orderNumber}}</p>
	<p style="font-family: cursive;font-size: 28px;font-weight: bold;color: orange;">订单查询号</p>
	<p style="font-family: cursive;font-size: 28px;font-weight: bold;color: orange;">{{.orderShortNumber}}</p>
	{{else}}
	<p style="font-family: cursive;font-size: 28px;font-weight: bold;">未发现匹配信息</p>
	<br/>
	<p style="font-family: cursive;font-size: 28px;font-weight: bold;color: orange;"> <a href="javascript:history.back();" class="uk-button">返回</a></p>
	{{end}}
</div>
