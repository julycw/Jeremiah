<div>
    <h4 class="tm-article-subtitle">从网页扫描数据</h4>
    <form class="uk-form">
		<input id="scanUrl" name="url" type="text" value="{{.url}}" placeholder="http://" class="uk-margin-small-top uk-width-1-1 uk-width-medium-1-2 uk-width-large-1-2">
		<button class="uk-button uk-button-success uk-margin-small-top"><strong>Go!</strong></button>
	</form>

    <h4 class="tm-article-subtitle">商品信息录入</h4>
    {{if .Error}}
    <div class="uk-alert uk-alert-danger">错误:{{.Error}}</div>
	{{end}}
    {{if .Warning}}
    <div class="uk-alert uk-alert-warning">警告:{{.Warning}}</div>
	{{end}}
	<form id="scanForm" class="uk-form uk-form-horizontal">
		{{.Form | renderform}}
	</form>
	<hr class="uk-article-divider">
	<button id="btnSubmit" class="uk-button uk-button-primary">提交</button>
	<button id="btnClose" class="uk-button uk-button-danger">关闭</button>
</div>