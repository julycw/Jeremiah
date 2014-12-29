<div>
    <table class="uk-table uk-table-hover">
        <caption>粉丝列表</caption>
        <thead>
            <tr>
                <th>ID</th>
                <th>姓名</th>
                <th>备注名</th>
                <th>备注</th>
                <th>操作</th>
            </tr>
        </thead>
        <tbody id="tbFollower">
        </tbody>
    </table>
</div>

<div id="modify-modal" class="uk-modal">
    <div class="uk-modal-dialog">
        <a href="" class="uk-modal-close uk-close"></a>
        <form id="formModify" class="uk-form" action="/manage/weixin/follower/modify" method="post">
            <fieldset>
                <legend>修改信息</legend>
                <input type="hidden" name="id"/>
                <input type="hidden" name="userid"/>
                <input type="hidden" name="username"/>
                <input type="hidden" name="name"/>
                <input type="hidden" name="phone"/>
                <input type="hidden" name="sex"/>
                <div class="uk-form-row">
                    <p>ID: <b id="UserId"></b></p>
                </div>
                <div class="uk-form-row">
                    <p>姓名: <b id="Name"></b></p>
                </div>
                <div class="uk-form-row">
                    <p>性别: <b id="Sex"></b></p>
                </div>
                <div class="uk-form-row">
                    <p>电话: <b id="Phone"></b></p>
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
