<div>
    <h3 class="tm-article-subtitle">筛选</h3>
    <form class="uk-form" id="search-form">
        <div class="uk-grid">
            <div class="uk-width-1-1 uk-width-large-1-2">
                <label for="model">型号:</label>
                <input id="model" name="model" type="text" style="width:100%;"></input>
            </div>
        </div>
        <div class="uk-grid" style="margin-top:0;">
            <div class="uk-width-1-1 uk-width-large-1-2">
                <label for="suggest-price-min">建议售价:</label>
                <div class="uk-grid">
                    <div class="uk-width-1-2" >
                        <input id="suggest-price-min" name="suggestPriceMin" style="width:100%;" type="number" value="0" min="0" step="50"></input>
                    </div>
                    <div class="uk-width-1-2" >
                        <input name="suggestPriceMax" type="number" style="width:100%;" value="15000" min="0" step="50"></input>
                    </div>
                </div>
            </div>
        </div>
        <div class="uk-grid" style="margin-top:0;">
            <div class="uk-width-1-1 uk-width-large-1-2">
                <div class="uk-overflow-container">
                    <table class="uk-table uk-table-condensed">
                        <thead>
                            <tr>
                                <th>CPU</th>
                                <th></th>
                                <th>显卡</th>
                                <th>内存</th>
                                <th>硬盘</th>
                                <th>系统</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td>
                                    <label>任意</label><input type="radio" name="cpu" value="all" checked/>
                                </td>
                                <td>
                                </td>
                                <td>
                                    <label>任意</label><input type="radio" name="graphic" value="all" checked/>
                                </td>
                                <td>
                                    <label>任意</label><input type="radio" name="memory" value="all" checked/>
                                </td>
                                <td>
                                    <label>任意</label><input type="radio" name="disk" value="all" checked/>
                                </td>
                                <td>
                                    <label>任意</label><input type="radio" name="os" value="all" checked/>
                                </td>
                            </tr>
                            <tr>
                                <td>
                                    <label>Intel Core i3</label><input type="radio" name="cpu" value="i3"/>
                                </td>
                                <td>
                                    <label>AMD E1</label><input type="radio" name="cpu" value="E1"/>
                                </td>
                                <td>
                                    <label>512M</label><input type="radio"  name="graphic" value="512M"/>
                                </td>
                                <td>
                                    <label>2G</label><input type="radio" name="memory" value="2G"/>
                                </td>
                                <td>
                                    <label>320G</label><input type="radio" name="disk" value="320G"/>
                                </td>
                                <td>
                                    <label>Win7</label><input type="radio" name="os" value="Win7"/>
                                </td>
                            </tr>
                            <tr>
                                <td>
                                    <label>Intel Core i5</label><input type="radio" name="cpu" value="i5"/>
                                </td>
                                <td>
                                    <label>AMD sempron</label><input type="radio" name="cpu" value="sempron"/>
                                </td>
                                <td>
                                    <label>1G</label><input type="radio"  name="graphic" value="1G"/>
                                </td>
                                <td>
                                    <label>4G</label><input type="radio" name="memory" value="4G"/>
                                </td>
                                <td>
                                    <label>500G</label><input type="radio" name="disk" value="500G"/>
                                </td>
                                <td>
                                    <label>Win8</label><input type="radio" name="os" value="Win8"/>
                                </td>
                            </tr>
                            <tr>
                                <td>
                                    <label>Intel Core i7</label><input type="radio" name="cpu" value="i7"/>
                                </td>
                                <td>
                                    <label>AMD Athlon x2</label><input type="radio" name="cpu" value="Athlon"/>
                                </td>
                                <td>
                                    <label>2G</label><input type="radio"  name="graphic" value="2G"/>
                                </td>
                                <td>
                                    <label>8G</label><input type="radio" name="memory" value="8G"/>
                                </td>
                                <td>
                                    <label>750G</label><input type="radio" name="disk" value="750G"/>
                                </td>
                                <td>
                                    <label>Linux</label><input type="radio" name="os" value="Linux"/>
                                </td>
                            </tr>
                            <tr>
                                <td>
                                    <label>Intel Celeron 双核</label><input type="radio" name="cpu" value="Celeron"/>
                                </td>
                                <td>
                                    <label>AMD A4</label><input type="radio" name="cpu" value="A4"/>
                                </td>
                                <td>
                                    <label>高清显卡</label><input type="radio"  name="graphic" value="高清显卡"/>
                                </td>
                                <td>
                                    <label>10G</label><input type="radio" name="memory" value="10G"/>
                                </td>
                                <td>
                                    <label>1TB</label><input type="radio" name="disk" value="1T"/>
                                </td>
                                <td>
                                </td>
                            </tr>
                            <tr>
                                <td>
                                    <label>Intel Pentium 双核</label><input type="radio" name="cpu" value="Pentium"/>
                                </td>
                                <td>
                                    <label>AMD A6</label><input type="radio" name="cpu" value="A6"/>
                                </td>
                                <td>
                                </td>
                                <td>
                                </td>
                                <td>
                                    <label>2TB</label><input type="radio" name="disk" value="2T"/>
                                </td>
                                <td>
                                </td>
                            </tr>
                            <tr>
                                <td>
                                    <label>Intel Celeron 四核</label><input type="radio" name="cpu" value="Celeron"/>
                                </td>
                                <td>
                                    <label>AMD A8</label><input type="radio" name="cpu" value="A8"/>
                                </td>
                                <td>
                                </td>
                                <td>
                                </td>
                                <td>
                                    <label>SSD</label><input type="radio" name="disk" value="SSD"/>
                                </td>
                                <td>
                                </td>
                            </tr>
                            <tr>
                                <td>
                                    <label>Intel Pentium 四核</label><input type="radio" name="cpu" value="Pentium"/>
                                </td>
                                <td>
                                </td>
                                <td>
                                </td>
                                <td>
                                </td>
                                <td>
                                </td>
                                <td>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    </form>
    <br/> 
    <button id="btnSearch" class="uk-button uk-button-small">查询</button>
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
                <th>建议售价</th>
                <th>CPU</th>
                <th>显卡</th>
                <th>内存</th>
                <th>硬盘</th>
            </tr>
        </thead>
        <tbody id="erp-list">
        </tbody>
    </table>
</div>
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