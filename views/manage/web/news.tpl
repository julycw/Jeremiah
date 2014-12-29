<div>
    <div style="text-align:center;">
        <a href="javascript:void(0);" class="uk-button" style="float:left;" id="btnPrev">上一页</a>
        <a href="javascript:void(0);" class="uk-button" id="btnAdd">新建新闻</a>
        <a href="javascript:void(0);" class="uk-button" style="float:right;" id="btnNext">下一页</a>
    </div>
    <table class="uk-table uk-table-hover">
        <thead>
            <tr>
                <th>分类</th>
                <th>标题</th>
                <th>状态</th>
                <th>操作</th>
            </tr>
        </thead>
        <tbody id="tbMessage">
            {{range .newsList}}
            <tr>
                <td>{{newsCategory .Category}}</td>
                <td>{{.Title}}</td>
                <td>{{newsState .State}}</td>
                <td>
                    <div class="uk-button-group">
                        <a class="uk-button uk-button-small btn-modify" target="{{.Id}}">修改</a>
                        {{if compare .State 2}}
                        <a href="/news/{{.Id}}" target="_blank" class="uk-button uk-button-small">查看</a>
                        {{else}}
                        <a class="uk-button uk-button-small btn-preview" target="{{.Id}}">预览</a>
                        {{end}}
                    </div>
                </td>
            </tr>
            {{end}}
        </tbody>
    </table>
</div>
<div id="add-modal" class="uk-modal">
    <div class="uk-modal-dialog uk-modal-dialog-large">
        <a href="" class="uk-modal-close uk-close"></a>
        <form id="formAdd" class="uk-form" action="/manage/web/news/add" method="post">
            <fieldset>
                <legend>添加新闻</legend>
                <div class="uk-form-row">
                    <input type="text" name="title" class="uk-width-1-1" placeholder="标题"/>
                </div>
                <div class="uk-form-row">
                    <lable for="state">状态</label>
                    <select name="state" class="uk-width-1-1" placeholder="状态">
                        <option value="1">准备发布</option>
                        <option value="2">已发布</option>
                        <option value="3">不可见</option>
                    </select>
                </div>
                <div class="uk-form-row">
                    <lable for="category">分类</label>
                    <select name="category" class="uk-width-1-1" placeholder="分类">
                        <option value="1">我们做什么</option>
                        <option value="2">最新热点</option>
                        <option value="3">一般新闻</option>
                    </select>
                </div>
                <div class="uk-form-row" style="min-height:200px;">
                    <textarea name="html" data-uk-htmleditor="{}"></textarea>
                </div>
                <div class="uk-form-row">
                    <a href="javascript:void(0);" class="uk-button" id="btnAddSubmit">提交</a>
                </div>
            </fieldset>
        </form>
    </div>
</div>

<div id="modify-modal" class="uk-modal">
    <div class="uk-modal-dialog uk-modal-dialog-large">
        <a href="" class="uk-modal-close uk-close"></a>
        <form id="formModify" class="uk-form" action="/manage/web/news/modify" method="post">
            <fieldset>
                <legend>修改新闻</legend>
                <input type="hidden" name="id"/>
                <div class="uk-form-row">
                    <input type="text" name="title" class="uk-width-1-1" placeholder="标题"/>
                </div>
                <div class="uk-form-row">
                    <lable for="state">状态</label>
                    <select name="state" class="uk-width-1-1" placeholder="状态">
                        <option value="1">准备发布</option>
                        <option value="2">已发布</option>
                        <option value="3">不可见</option>
                    </select>
                </div>
                <div class="uk-form-row">
                    <lable for="category">分类</label>
                    <select name="category" class="uk-width-1-1" placeholder="分类">
                        <option value="1">我们做什么</option>
                        <option value="2">最新热点</option>
                        <option value="3">一般新闻</option>
                    </select>
                </div>
                <div class="uk-form-row" id="editHtmlWrapper" style="min-height:200px;">
                    <textarea id="editHtml" name="html"></textarea>
                </div>
                <div class="uk-form-row">
                    <a href="javascript:void(0);" class="uk-button uk-button-success" id="btnModifySubmit">提交</a>
                    <a href="javascript:void(0);" class="uk-button uk-button-danger" id="btnDeleteSubmit" style='float:right;'>删除</a>
                </div>
            </fieldset>
        </form>
    </div>
</div>

<div id="view-modal" class="uk-modal">
    <div class="uk-modal-dialog">
        <a href="" class="uk-modal-close uk-close"></a>
        <h1>不好了</h1>
        <p>这个功能还没来得及做 :(</p>
    </div>
</div>