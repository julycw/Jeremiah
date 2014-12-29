<div>
    <ul class="uk-subnav uk-subnav-pill" data-uk-switcher="{connect:'#subnav-pill-content'}">
        <li class=""><a href="#">浏览</a>
        </li>
        <li class=""><a href="#">新增</a>
        </li>
    </ul>
    <ul id="subnav-pill-content" class="uk-switcher">
        <li class="">
            {{range .dayFixOrderList}}
            <div class="uk-panel uk-panel-box">
                <div style="float:right">
                    <a href="javascript:void(0);" class="btnShowDayDetail"><i class="uk-icon-folder-open-o"></i> 全部展开</a> /
                    <a href="javascript:void(0);" class="btnHideDayDetail"><i class="uk-icon-folder-o"></i> 全部缩进</a>
                </div>
                <h3 class="uk-panel-title">{{.Day}}</h3>
                {{range .FixOrderList}}
                <article class="uk-comment">
                    <header class="uk-comment-header">
                        <div style="float:left;">
                            <span class="uk-text-primary uk-text-bold">{{.UserId}}</span>的订单
                        </div>
                        <div style="float:right;">
                            <a href="javascript:void(0);" class="btnShowOrderDetails"><i class="uk-icon-chevron-down uk-icon-small"></i>查看明细</a>
                        </div>
                        <div style="clear:both;">
                        </div>
                        <hr class="uk-article-divider" style="margin-top:3px;margin-bottom:5px;"/>
                        <h4 class="uk-comment-title">{{.Description}}</h4>
                        <div class="uk-comment-meta">提交人:
                            <span class="uk-text-primary uk-text-bold">{{.CreateBy}}</span> 提交时间:<b>{{dateformat .CreateOn "15:04:05"}}</b>.
                        </div>
                    </header>
                    <section class="uk-comment-body order-details" style="display:none;" target-id="{{.Id}}">
                        <div class="loading-tip" style="text-align:center;"><i class="uk-icon-spinner uk-icon-spin uk-icon-small"></i>
                        </div>
                        <ul class="uk-list" style="margin-top:0;">
                            <li>
                                <form class="uk-form">
                                    <div style="float:right;width:70px;">
                                        <a href="javascript:void(0);" class="uk-button btn-submit-details" style="min-width:70px;" target-id="{{.Id}}">添加</a>
                                    </div>
                                    <div style="margin-right:80px;">
                                        <input type="text" placeholder="输入新的内容" class="order-details-input" style="width:100%;">
                                    </div>
                                </form>
                            </li>
                        </ul>
                        <ul id="order-details-list-{{.Id}}" class="uk-list uk-list-striped order-details-list" style="margin-top:0;">
                            
                        </ul>
                    </section>
                    <br/>
                </article>
                {{end}}
            </div>
            <br/>{{end}}
        </li>
        <li class="">
            <div class="uk-panel uk-panel-box">
                <form id="formNote" class="uk-form">
                    <div class="uk-grid" data-uk-grid-margin >
                        <div class="uk-width-1-1 uk-width-medium-1-4">
                            <div class="uk-panel">
                                <div class="uk-grid">
                                    <div class="uk-width-1-1 uk-width-medium-1-1 uk-text-center">
                                        <p>订单号: <b id="orderNumber">{{.orderNumber}}</b> <a id="btnOrderRefresh" href="javascript:void(0);"><i class="uk-icon-refresh"></i></a></p>
                                        <div class="uk-panel" id="panel-code">
                                            <span style="display:inline-block;min-width:200px;min-height:200px;">
                                            <img id="orderCodeImage" src="/static/img/logo.jpg" mask-src="http://qr.liantu.com/api.php?text=http://{{.SiteIp}}/api/manage/weixin/fix/scanFixOrder?order=" orderid="{{.orderNumber}}"/>
                                            </span>
                                            <p>
                                                <a href="javascript:void(0);" id="btnBindUser" class="uk-button">绑定用户</a>
                                                <a href="javascript:void(0);" id="btnSelectUser" class="uk-button">手动选择</a>
                                            </p>
                                        </div>
                                        <div class="uk-panel uk-text-left" id="panel-user">
                                            <ul class="uk-list uk-list-striped"></ul>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <div class="uk-width-1-1 uk-width-medium-3-4">
                            <div class="uk-panel">
                                <textarea id="inputDescription" name="content" class="uk-width-1-1" rows="12" placeholder="描述"></textarea>
                            <p>1. 打开(刷新)此页面，获取预设的订单号</p>
                            <p>2. 用户打开微信的对话界面，进入绑定订单界面，输入订单<b>识别号</b></p>
                            <p>3. 点击此页面的“绑定用户”按钮，成功后二维码变为用户相关信息</p>
                            <p>4. 填写描述并提交</p>
                            <p>PS: 订单号末尾4位为订单<b>识别号</b></p>
                            </div>
                        </div>
                    </div>
                    <div class="uk-grid" data-uk-grid-margin>
                        <div class="uk-form-row uk-width-1-1">
                            <a href="javascript:void(0);" id="btnAddOrder" class="uk-button">提交</a>
                            <span>绑定用户后才可以提交</span>
                        </div>
                    </div>
                </form>
            </div>
        </li>
    </ul>
</div>
