<div>
    <div style="text-align:center;">
        <a href="javascript:void(0);" class="uk-button" style="float:left;" id="btnPrev">上一页</a>
        <a href="javascript:void(0);" class="uk-button" id="btnAdd">添加产品</a>
        <a href="javascript:void(0);" class="uk-button" style="float:right;" id="btnNext">下一页</a>
    </div>
    <table class="uk-table uk-table-hover">
        <thead>
            <tr>
                <th>产品标题</th>
                <th>类型</th>
                <th>简述</th>
                <th>可见</th>
                <th>推荐</th>
                <th>操作</th>
            </tr>
        </thead>
        <tbody id="tbMessage">
            {{range .productList}}
            <tr>
                <td>{{.Title}}</td>
                <td>{{productType .Type}}</td>
                <td>{{.Description}}</td>
                <td>{{if compare .Visible  1}}<i class="uk-icon-check"></i>{{else}}X{{end}}</td>
                <td>{{if compare .Hot  1}}<i class="uk-icon-check"></i>{{else}}X{{end}}</td>
                <td style="min-width:80px;">
                    <div class="uk-button-group">
                        <a class="uk-button uk-button-small btn-modify" target="{{.Id}}">修改</a>
                        {{if compare .Visible 1}}
                        <a href="/gallery/{{.Id}}" target="_blank" class="uk-button uk-button-small">查看</a>
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
        <form id="formAdd" class="uk-form" action="/manage/web/product/add" method="post" enctype="multipart/form-data">
            <fieldset>
                <legend>添加产品</legend>
                <div class="uk-form-row">
                    <input type="text" name="title" class="uk-width-1-1" placeholder="产品标题"/>
                </div>
                <div class="uk-form-row">
                    <select name="type" class="uk-width-1-1">
                        <option value="1">智能家居</option>
                    </select>
                </div>
                <div class="uk-form-row">
                    <input type="text" name="description" class="uk-width-1-1" placeholder="简述"/>
                </div>
                <div class="uk-form-row">
                    <input type="text" name="sort" class="uk-width-1-1" placeholder="排序（数字越大越前面）"/>
                </div>
                <div class="uk-form-row">
                    <label for="visible">可见</label><input id="visible" type="checkbox" name="visible" value="1" checked/>
                </div>
                <div class="uk-form-row">
                    <label for="hot">推荐</label><input id="hot" type="checkbox" name="hot" value="1"/>
                </div>
                <div class="uk-form-row">
                    <label for="hot">上传图片   </label><input type="file" name="image"/>
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
        <form id="formModify" class="uk-form" action="/manage/web/product/modify" method="post">
            <fieldset>
                <legend>修改产品</legend>
                <input type="hidden" name="id"/>
                <input type="hidden" name="image"/>
                <input type="hidden" name="images"/>
                <div class="uk-form-row">
                    <input type="text" name="title" class="uk-width-1-1" placeholder="产品标题"/>
                </div>
                <div class="uk-form-row">
                    <select name="type" class="uk-width-1-1">
                        <option value="1">智能家居</option>
                    </select>
                </div>
                <div class="uk-form-row">
                    <input type="text" name="description" class="uk-width-1-1" placeholder="简述"/>
                </div>
                <div class="uk-form-row">
                    <input type="text" name="sort" class="uk-width-1-1" placeholder="排序（数字越大越前面）"/>
                </div>
                <div class="uk-form-row">
                    <label for="visible">可见</label><input id="visible" type="checkbox" name="visible" value="1"/>
                </div>
                <div class="uk-form-row">
                    <label for="hot">推荐</label><input id="hot" type="checkbox" name="hot" value="1"/>
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