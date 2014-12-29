<!DOCTYPE html>
<html lang="zh-cn" dir="ltr" class="uk-height-1-1">

    <head>
        <meta charset="utf-8">
        <title>耶利米 | 智能家居 后台管理系统 version 0.3</title>
    	<meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <link rel="shortcut icon" href="images/favicon.ico" type="image/x-icon">
        <link rel="apple-touch-icon-precomposed" href="images/apple-touch-icon.png">
        <link rel="stylesheet" href="/static/uikit/css/uikit.docs.min.css">
        <script src="/static/js/jquery-1.11.0.min.js"></script>
        <script src="/static/uikit/js/uikit.min.js"></script>
    </head>

    <body class="uk-height-1-1">

        <div class="uk-vertical-align uk-text-center uk-height-1-1" style="background-color:#FFEA99;">
            <div class="uk-vertical-align-middle" style="width: 250px;">

                <img class="uk-margin-bottom" width="140" height="140" style="border-radius:70px;" src="/static/img/logo.jpg" alt="">

                <form id="formLogin" class="uk-panel uk-panel-box uk-form" action="/manage/login" method="post">
                    <input type="hidden" name="redirect" value="{{.redirect}}"/>
                    <div class="uk-form-row">
                        <input class="uk-width-1-1 uk-form-large" type="text" name="username" placeholder="用户名">
                    </div>
                    <div class="uk-form-row">
                        <input class="uk-width-1-1 uk-form-large" type="password" name="password" placeholder="密码">
                    </div>
                    <div class="uk-form-row">
                        <a id="btnLogin" class="uk-width-1-1 uk-button uk-button-primary uk-button-large" href="#">登陆</a>
                    </div>
                    <div class="uk-form-row uk-text-small">
                        <label class="uk-float-left"><input type="checkbox"> 记住密码</label>
                        <!-- <a class="uk-float-right uk-link uk-link-muted" href="#">Forgot Password?</a> -->
                    </div>
                </div>

            </div>
        </div>
        <script type="text/javascript">
        $(function(){
        	$("#btnLogin").on("click",function(){
        		$("#formLogin").submit();
        	});
        });
        </script>
    </body>
</html>