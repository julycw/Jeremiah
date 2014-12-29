<div>
    <div id="search-bar"></div>
    {{range .OptionSections}}
    <h4 class="tm-article-subtitle">{{.SectionName}}</h4>
        <form class="uk-form uk-width-medium-1-1">
            {{range .Options}}
            <fieldset>
                <legend>{{.Title}}</legend>
                <div class="uk-form-row">
                    <textarea class="uk-width-3-4 form-field" value="{{.Value}}" value-origin="{{.Value}}" placeholder="{{.Description}}" target-id="{{.Id}}" cols="30" rows="5" placeholder="Textarea">{{.Value}}</textarea>
                    <span class="wrapper-btns" style="display:none;width:55px;">
                    <a href="javascript:void(0);" class="uk-button btn-save uk-width-1-1" style="margin-bottom:5px;">保存</a>
                    <a href="javascript:void(0);" class="uk-button btn-cancel uk-width-1-1">取消</a>
                    </span>
                </div>
            </fieldset>
            {{end}}
        </form>
        {{if .HasNext}}
        <hr class="uk-article-divider">
        {{end}}
    {{end}}
</div>

<div id="modify-modal" class="uk-modal">
    <div class="uk-modal-dialog">
        <a href="" class="uk-modal-close uk-close"></a>
        <form id="formModify" class="uk-form" action="/manage/weixin/follower/modify" method="post">
            <fieldset>
                <legend>修改信息</legend>
                <input type="hidden" name="id"/>
                <div class="uk-form-row">
                    <input type="text" name="userid" class="uk-width-1-1" readonly/>
                </div>
                <div class="uk-form-row">
                    <label id="lbUserName"></label>
                </div>
                <div class="uk-form-row">
                    <input type="text" name="nickname" class="uk-width-1-1" placeholder="昵称"/>
                </div>
                <div class="uk-form-row">
                    <input type="text" name="description" class="uk-width-1-1" placeholder="备注"/>
                </div>
                <div class="uk-form-row">
                    <a href="javascript:void(0);" class="uk-button" id="btnUpload">提交</a>
                </div>
            </fieldset>
        </form>
    </div>
</div>
