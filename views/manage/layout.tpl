<!DOCTYPE html>
<html lang="zh-CN">

<head>
    <title>耶利米 | 智能家居 后台管理系统 version 0.3</title>
    <meta charset="utf-8" />
    <meta name="keywords" content="" />
    <meta name="description" content="" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <link rel="stylesheet" href="/static/uikit/css/uikit.docs.min.css">
    <link rel="stylesheet" type="text/css" href="/static/uikit/css/docs.css">
    <link rel="stylesheet" type="text/css" href="/static/css/site.css">
    {{.LayoutHeader}}
</head>

<body style="background:#302E27;">
    <div class="uk-container uk-container-center uk-margin-large-bottom" style="background:white;padding:0px;padding-bottom:20px;min-height:768px;">
        <nav class="uk-navbar uk-margin-large-bottom">
            <a class="uk-navbar-brand uk-hidden-small" href="/manage"><i class="uk-icon-home uk-icon-small"></i> 耶利米 | 智能家居
            </a>
            <ul class="uk-navbar-nav uk-hidden-small" style="float:right;">
                <li>
                    <a href="/" target="_blank"><i class="uk-icon-globe uk-icon-small"></i> 访问网站</a>
                </li>
                <li class="uk-parent" data-uk-dropdown>
                    <a href="#">网站内容管理 <i class="uk-icon-caret-down"></i></a>
                    <div class="uk-dropdown uk-dropdown-navbar">
                        <ul class="uk-nav uk-nav-navbar">
                            <li>
                                <a href="/manage/web/news">文章列表</a>
                            </li>
                            <li>
                                <a href="/manage/web/product">产品管理</a>
                            </li>
                            <li>
                                <a href="/manage/basic/option/text">首页内容</a>
                            </li>
                        </ul>
                    </div>
                </li>
                <li class="uk-parent" data-uk-dropdown>
                    <a href="#">微信功能 <i class="uk-icon-caret-down"></i></a>
                    <div class="uk-dropdown uk-dropdown-navbar">
                        <ul class="uk-nav uk-nav-navbar">
                            <li>
                                <a href="/manage/basic/option/weixin">选项设置</a>
                            </li>
                            <li>
                                <a href="/manage/weixin/menu">菜单管理</a>
                            </li>
                            <li>
                                <a href="/manage/weixin/order">订单管理</a>
                            </li>
                            <li>
                                <a href="/manage/weixin/message">消息管理</a>
                            </li>
                            <li>
                                <a href="/manage/weixin/follower">粉丝管理</a>
                            </li>
                        </ul>
                    </div>
                </li>
                <li>
                    <a href="/manage/note"><i class="uk-icon-pencil uk-icon-small"></i> 记录板</a>
                </li>
                <li class="uk-parent" data-uk-dropdown>
                    <a href="#">基础信息管理 <i class="uk-icon-caret-down"></i></a>
                    <div class="uk-dropdown uk-dropdown-navbar">
                        <ul class="uk-nav uk-nav-navbar">
                            <li>
                                <a href="/manage/basic/option/other">选项设置</a>
                            </li>
                            <li>
                                <a href="/manage/basic/user">用户管理</a>
                            </li>
                            <li>
                                <a href="/manage/basic/role">角色管理</a>
                            </li>
                            <li>
                                <a href="/manage/basic/log">查看日志</a>
                            </li>
                        </ul>
                    </div>
                </li>
                <li>
                    <a href="/manage/logout"><i class="uk-icon-sign-out uk-icon-small"></i> 注销</a>
                </li>
            </ul>
            <a href="/manage/logout" style="position:absolute;right:10px;top:10px;"><i class="uk-icon-sign-out uk-icon-small"></i></a>
            <a href="#offcanvas" class="uk-navbar-toggle uk-visible-small" data-uk-offcanvas=""></a>
            <div class="uk-navbar-brand uk-navbar-center uk-visible-small" style="max-width:75%;"><a href="/manage" class="no-style">耶利米 | 智能家居</a></div>
        </nav>
        <div class="main-container-inner" style="padding:0 25px;">
            {{.LayoutContent}}
        </div>
        <div id="offcanvas" class="uk-offcanvas">

            <div class="uk-offcanvas-bar">

                <ul class="uk-nav uk-nav-offcanvas uk-nav-parent-icon" data-uk-nav="{multiple:true}">
                    <li class="uk-nav-header">网站</li>
                    <li>
                        <a href="/" target="_blank"><i class="uk-icon-globe"></i> 访问网站</a>
                    </li>
                    <li class="uk-nav-header">后台管理</li>
                    <li>
                        <a href="/manage"><i class="uk-icon-home"></i> 后台主页面</a>
                    </li>
                    <li class="uk-nav-header">细分功能</li>
                    <li class="uk-parent"><a href="#"><i class="uk-icon-wrench"></i> 网站内容管理</a>
                        <ul class="uk-nav-sub">
                            <li>
                                <a href="/manage/web/news">文章列表</a>
                            </li>
                            <li>
                                <a href="/manage/web/product">产品管理</a>
                            </li>
                            <li>
                                <a href="/manage/basic/option/text">首页内容</a>
                            </li>
                        </ul>
                    </li>
                    <li class="uk-parent"><a href="#"><i class="uk-icon-weixin"></i> 微信功能</a>
                        <ul class="uk-nav-sub">
                            <li>
                                <a href="/manage/basic/option/weixin">选项设置</a>
                            </li>
                            <li>
                                <a href="/manage/weixin/menu">菜单管理</a>
                            </li>
                            <li>
                                <a href="/manage/weixin/order">订单管理</a>
                            </li>
                            <li>
                                <a href="/manage/weixin/message">消息管理</a>
                            </li>
                            <li>
                                <a href="/manage/weixin/follower">粉丝管理</a>
                            </li>
                        </ul>
                    </li>
                    <li><a href="/manage/note"><i class="uk-icon-pencil"></i> 记事版</a>
                       
                    </li>
                    <li class="uk-parent"><a href="#"><i class="uk-icon-th-large"></i> 基础管理</a>
                        <ul class="uk-nav-sub">
                            <li>
                                <a href="/manage/basic/option">选项设置</a>
                            </li>
                            <li>
                                <a href="/manage/basic/user">用户管理</a>
                            </li>
                            <li>
                                <a href="/manage/basic/role">角色管理</a>
                            </li>
                            <li>
                                <a href="/manage/basic/log">查看日志</a>
                            </li>
                        </ul>
                    </li>
                </ul>

            </div>

        </div>

    </div>
    <script src="/static/js/jquery-1.11.0.min.js"></script>
    <script src="/static/uikit/js/uikit.min.js"></script>
    <script src="/static/js/site.js"></script>
    <script src="/static/js/utils.js"></script>
    {{.LayoutFooter}}
</body>

</html>
