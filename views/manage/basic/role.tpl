<div>
    <div style="text-align:center;">
        <a href="javascript:void(0);" class="uk-button" style="float:left;" id="btnPrev">上一页</a>
        <a href="javascript:void(0);" class="uk-button" id="btnAdd">创建角色</a>
        <a href="javascript:void(0);" class="uk-button" style="float:right;" id="btnNext">下一页</a>
    </div>
    <table class="uk-table uk-table-hover">
        <thead>
            <tr>
                <th>角色名</th>
                <th>权限</th>
                <th>操作</th>
            </tr>
        </thead>
        <tbody id="tbMessage">
            {{range .roleList}}
            <tr>
                <td>{{.RoleName}}</td>
                <td>{{.Avaliables}}</td>
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
        <form id="formAdd" class="uk-form" action="/manage/basic/role/add" method="post">
            <fieldset>
                <legend>创建角色</legend>
                <div class="uk-form-row">
                    <input type="text" name="rolename" class="uk-width-1-1" placeholder="角色名"/>
                </div>
                <h3>权限列表</h3>
                <input type="hidden" name="avaliables" />
                <div class="uk-form-row">
                    <p>日常功能：</p>
                    <label class="avaliableLabel" for="登陆">登陆</label><input id="登陆" type="checkbox" class="avaliable" avaliable="登陆"/>
                    <label class="avaliableLabel" for="进行签到">进行签到</label><input id="进行签到" type="checkbox" class="avaliable" avaliable="进行签到"/>
                    <label class="avaliableLabel" for="使用记录板">使用记录板</label><input id="使用记录板" type="checkbox" class="avaliable" avaliable="使用记录板"/>
                </div>
                <div class="uk-form-row">
                    <p>基础管理权限：</p>
                    <label class="avaliableLabel" for="用户管理">用户管理</label><input id="用户管理" type="checkbox" class="avaliable" avaliable="用户管理"/>
                    <label class="avaliableLabel" for="角色管理">角色管理</label><input id="角色管理" type="checkbox" class="avaliable" avaliable="角色管理"/>
                    <label class="avaliableLabel" for="查看日志">查看日志</label><input id="查看日志" type="checkbox" class="avaliable" avaliable="查看日志"/>
                    <label class="avaliableLabel" for="选项设置">选项设置</label><input id="选项设置" type="checkbox" class="avaliable" avaliable="选项设置"/>
                </div>
                <div class="uk-form-row">
                    <p>网站内容管理权限：</p>
                    <label class="avaliableLabel" for="文案修改">文案修改</label><input id="文案修改" type="checkbox" class="avaliable" avaliable="首页文案修改"/>
                    <label class="avaliableLabel" for="文章管理">文章管理</label><input id="文章管理" type="checkbox" class="avaliable" avaliable="文章管理"/>
                    <label class="avaliableLabel" for="商品管理">商品管理</label><input id="商品管理" type="checkbox" class="avaliable" avaliable="商品管理"/>
                </div>
                <div class="uk-form-row">
                    <p>微信相关权限：</p>
                    <label class="avaliableLabel" for="消息查看">消息查看</label><input id="消息查看" type="checkbox" class="avaliable" avaliable="消息查看"/>
                    <label class="avaliableLabel" for="粉丝管理">粉丝管理</label><input id="粉丝管理" type="checkbox" class="avaliable" avaliable="粉丝管理"/>
                    <label class="avaliableLabel" for="订单管理">订单管理</label><input id="订单管理" type="checkbox" class="avaliable" avaliable="订单管理"/>
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
        <form id="formModify" class="uk-form" action="/manage/basic/role/modify" method="post">
            <fieldset>
                <legend>修改角色</legend>
                <input type="hidden" name="id"/>
                <input type="hidden" name="roleid"/>
                <div class="uk-form-row">
                    <input type="text" name="rolename" class="uk-width-1-1" placeholder="角色名"/>
                </div>
                <h3>权限列表</h3>
                <input type="hidden" name="avaliables" />
                <div class="uk-form-row">
                    <p>日常功能：</p>
                    <label class="avaliableLabel" for="登陆2">登陆</label><input id="登陆2" type="checkbox" class="avaliable" avaliable="登陆"/>
                    <label class="avaliableLabel" for="进行签到2">进行签到</label><input id="进行签到2" type="checkbox" class="avaliable" avaliable="进行签到"/>
                    <label class="avaliableLabel" for="使用记录板2">使用记录板</label><input id="使用记录板2" type="checkbox" class="avaliable" avaliable="使用记录板"/>
                </div>
                <div class="uk-form-row">
                    <p>基础管理权限：</p>
                    <label class="avaliableLabel" for="用户管理2">用户管理</label><input id="用户管理2" type="checkbox" class="avaliable" avaliable="用户管理"/>
                    <label class="avaliableLabel" for="角色管理2">角色管理</label><input id="角色管理2" type="checkbox" class="avaliable" avaliable="角色管理"/>
                    <label class="avaliableLabel" for="查看日志2">查看日志</label><input id="查看日志2" type="checkbox" class="avaliable" avaliable="查看日志"/>
                    <label class="avaliableLabel" for="选项设置2">选项设置</label><input id="选项设置2" type="checkbox" class="avaliable" avaliable="选项设置"/>
                </div>
                <div class="uk-form-row">
                    <p>网站内容管理权限：</p>
                    <label class="avaliableLabel" for="文案修改2">文案修改</label><input id="文案修改2" type="checkbox" class="avaliable" avaliable="首页文案修改"/>
                    <label class="avaliableLabel" for="文章管理2">文章管理</label><input id="文章管理2" type="checkbox" class="avaliable" avaliable="文章管理"/>
                    <label class="avaliableLabel" for="商品管理2">商品管理</label><input id="商品管理2" type="checkbox" class="avaliable" avaliable="商品管理"/>
                </div>
                <div class="uk-form-row">
                    <p>微信相关权限：</p>
                    <label class="avaliableLabel" for="消息查看2">消息查看</label><input id="消息查看2" type="checkbox" class="avaliable" avaliable="消息查看"/>
                    <label class="avaliableLabel" for="粉丝管理2">粉丝管理</label><input id="粉丝管理2" type="checkbox" class="avaliable" avaliable="粉丝管理"/>
                    <label class="avaliableLabel" for="订单管理2">订单管理</label><input id="订单管理2" type="checkbox" class="avaliable" avaliable="订单管理"/>
                </div>
                <div class="uk-form-row">
                    <a href="javascript:void(0);" class="uk-button uk-button-success" id="btnModifySubmit" style='float:left;'>提交</a>
                    <a href="javascript:void(0);" class="uk-button uk-button-danger" id="btnDeleteSubmit" style='float:right;'>删除</a>
                </div>
            </fieldset>
        </form>
    </div>
</div>
