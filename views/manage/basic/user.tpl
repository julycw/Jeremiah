<div>
    <div style="text-align:center;">
        <a href="javascript:void(0);" class="uk-button" style="float:left;" id="btnPrev">上一页</a>
        <a href="javascript:void(0);" class="uk-button" id="btnAdd">添加用户</a>
        <a href="javascript:void(0);" class="uk-button" style="float:right;" id="btnNext">下一页</a>
    </div>
    <table class="uk-table uk-table-hover">
        <thead>
            <tr>
                <th>用户名</th>
                <th>昵称</th>
                <th>角色</th>
                <th>操作</th>
            </tr>
        </thead>
        <tbody id="tbMessage">
            {{range .userList}}
            <tr>
                <td>{{.UserName}}</td>
                <td>{{.NickName}}</td>
                <td>{{roleName .RoleId}}</td>
                <td>
                    <a href="javascript:void(0);" class="uk-button uk-button-small btn-modify" target="{{.Id}}">修改</a>
                </td>
            </tr>
            {{end}}
        </tbody>
    </table>
</div>

<div id="modal" class="uk-modal">
    <div class="uk-modal-dialog">
        <a href="" class="uk-modal-close uk-close"></a>
        <h1>不好了</h1>
        <p>这个功能还没来得及做 :(</p>
    </div>
</div>

<div id="add-modal" class="uk-modal">
    <div class="uk-modal-dialog">
        <a href="" class="uk-modal-close uk-close"></a>
        <form id="formAdd" class="uk-form" action="/manage/basic/user/add" method="post">
            <fieldset>
                <legend>添加用户</legend>
                <div class="uk-form-row">
                    <input type="text" name="username" class="uk-width-1-1" placeholder="用户名"/>
                </div>
                <div class="uk-form-row">
                    <input type="text" name="nickname" class="uk-width-1-1" placeholder="昵称"/>
                </div>
                <div class="uk-form-row">
                    <input type="text" name="password" class="uk-width-1-1" placeholder="密码"/>
                </div>
                <div class="uk-form-row">
                    <label for="roleid">角色：</p>
                    <select id="roleid" name="roleid" class="uk-width-1-1">
                        {{range .roleList}}
                        <option value='{{.RoleId}}'>{{.RoleName}}</option>
                        {{end}}
                    </select>
                </div>
                <div class="uk-form-row">
                    <a href="javascript:void(0);" class="uk-button" id="btnAddSubmit">提交</a>
                </div>
            </fieldset>
        </form>
    </div>
</div>

<div id="modify-modal" class="uk-modal">
    <div class="uk-modal-dialog">
        <a href="" class="uk-modal-close uk-close"></a>
        <form id="formModify" class="uk-form" action="/manage/basic/user/modify" method="post">
            <fieldset>
                <legend>修改用户</legend>
                <input type="hidden" name="id"/>
                <div class="uk-form-row">
                    <label for="username">用户名：</label>
                    <input type="text" name="username" class="uk-width-1-1" placeholder="用户名"/>
                </div>
                <div class="uk-form-row">
                    <label for="nickname">昵称：</label>
                    <input type="text" name="nickname" class="uk-width-1-1" placeholder="昵称"/>
                </div>
                <div class="uk-form-row">
                    <label for="password">密码：</label>
                    <input type="text" name="password" class="uk-width-1-1" placeholder="密码 - 不修改请留空"/>
                </div>
                <div class="uk-form-row">
                    <label for="roleid">角色：</label>
                    <select id="roleid" name="roleid" class="uk-width-1-1">
                        {{range .roleList}}
                        <option value='{{.RoleId}}'>{{.RoleName}}</option>
                        {{end}}
                    </select>
                </div>
                <div class="uk-form-row">
                    <a href="javascript:void(0);" class="uk-button uk-button-success" id="btnModifySubmit">提交</a>
                    <a href="javascript:void(0);" class="uk-button uk-button-danger" id="btnDeleteSubmit" style='float:right;'>删除</a>
                </div>
            </fieldset>
        </form>
    </div>
</div>
