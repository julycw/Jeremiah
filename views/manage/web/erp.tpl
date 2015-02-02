
{{if .List}}
<div>
<h3 class="tm-article-subtitle">筛选</h3>
<form class="uk-form" method="get">
	<div data-uk-margin>
		<div class="uk-width-1-1 uk-width-large-1-1">
			<label for="suggest-price-min" class="uk-width-1-1 uk-width-large-1-2">售价(￥):</label>
	        <input id="suggest-price-min" type="number" class="uk-form-small uk-form-width-mini" value="1000" min="0" step="50"> ~ 
	        <input id="suggest-price-max" type="number" class="uk-form-small uk-form-width-mini" value="8000" max="10000" step="50">
		</div>
	</div>
	<div data-uk-margin>
    </div>
	<div data-uk-margin>
        <label for="model" class="uk-width-min-1-1">型号:</label> <input type="text" id="model" class="uk-form-small">
    </div>
	<div data-uk-margin>
        <label for="graphics-model">显卡:</label> <input type="text" id="graphics-model" class="uk-form-small">
    </div>
	<div data-uk-margin>
		<label for="cpu-model">CPU:</label> 
        <select id="cpu-model" class="uk-form-small">
            <option selected="selected">全部</option>
            <option value="i3">i3</option>
            <option value="i5">i5</option>
            <option value="i7">i7</option>
        </select>
        <label for="memory-size">内存:</label> 
        <select id="memory-size" class="uk-form-small">
            <option selected="selected">全部</option>
            <option value="2GB">2GB</option>
            <option value="4GB">4GB</option>
            <option value="8GB">8GB</option>
            <option value="16GB">16GB</option>
        </select>
        <label for="disk-size">硬盘:</label> 
        <select id="disk-size" class="uk-form-small">
            <option selected="selected">全部</option>
            <option value="500GB">500GB</option>
            <option value="1TB">1TB</option>
        </select>
	</div>
	<br/> 
	<button href="/manage/web/erp" class="uk-button uk-button-small">查询</button>
</form>
</div>
<h3 class="tm-article-subtitle">查询结果</h3>
<div class="uk-overflow-container">
    <table class="uk-table uk-table-hover uk-table-striped uk-table-condensed">
        <thead>
            <tr>
            	<th></th>
                <th>名称</th>
                <th>京东价</th>
                <th>官方价</th>
                <th>建议价</th>
                <th>CPU</th>
                <th>显卡</th>
                <th>内存</th>
                <th>硬盘</th>
            </tr>
        </thead>
        <tbody>
			{{range .List}}
            <tr>
            	<td class="btn-erp-detail" itemid="{{.IDStr}}">+</td>
                <td>{{.Name}}</td>
                <td>￥{{.JDPrice}}</td>
                <td>￥{{.Price}}</td>
                <td>￥{{.SuggestPrice}}</td>
                <td>{{.CPUModel}}</td>
                <td>{{.GraphicsModel}}</td>
                <td>{{.MemorySize}}</td>
                <td>{{.DiskSize}}</td>
            </tr>
			{{end}}
        </tbody>
    </table>
</div>
{{end}}
<br/> 
<a href="/manage/web/erp/scan" class="uk-button uk-button-small">点击录入</a>
当前共录入<strong>{{.Count}}</strong>条记录

<div id="view-modal" class="uk-modal">
    <div class="uk-modal-dialog">
        <a href="" class="uk-modal-close uk-close"></a>
        <h1>不好了</h1>
        <p>这个功能还没来得及做 :(</p>
    </div>
</div>