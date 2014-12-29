<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <title>贝贝电脑</title>
    <meta charset="utf-8" />
    <meta name="keywords" content="" />
    <meta name="description" content="" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <link rel="stylesheet" href="/static/uikit/css/uikit.docs.min.css">
    <script src="/static/js/jquery-1.11.0.min.js"></script>
    <script>
        jQuery.prototype.serializeObject=function(){  
        var obj=new Object();  
        $.each(this.serializeArray(),function(index,param){  
            if(!(param.name in obj)){  
                obj[param.name]=param.value;  
            }  
        });  
        return obj;  
    };  
    </script>
</head>

<body style="">
	<div class="main-container-inner" style="padding:15px 25px;">
        {{.LayoutContent}}
    </div>
</body>
</html>
