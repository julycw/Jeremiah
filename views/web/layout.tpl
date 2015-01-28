<!DOCTYPE HTML>
<html>
<head>
	<title>{{.Title}}</title>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
    <link rel="stylesheet" type="text/css" href="/static/uikit/css/uikit.min.css">
	<link rel="stylesheet" type="text/css" href="/static/home/css/style.css"/>
	<link rel="stylesheet" type="text/css" href="/static/home/css/slider.css"/>
	<link rel="stylesheet" type="text/css" href="/static/uikit/css/uikit.docs.min.css">
    <link rel="stylesheet" type="text/css" href="/static/uikit/css/docs.css">
	<link rel="stylesheet" type="text/css" href="/static/css/site.css"/>
    {{.LayoutHeader}}
</head>
<body>
  	<div class="header">
  	 	<div class="header_top">
	        <div class="wrap">		
		 		<div class="logo">
						<a href="/"><h1><span class="black">耶利米|智能家居</span></h1></a>
					</div>	
					 <div class="menu">
					    <ul id="menu-items">
							<li><a href="/">首页</a></li>
							<li><a href="/about">关于</a></li>
							<li><a href="/gallery">展示</a></li>
							<li><a href="/services">服务</a></li>
							<li><a href="/contact">联系</a></li>
							<div class="clear"></div>
						</ul>
					 </div>						
	    		 <div class="clear"></div>
	        </div>
	    </div>
	 	{{.LayoutHead}}
	</div>
  	<div class="main">
	 	{{.LayoutContent}}
	</div>
 	<div class="copy-right">
		<div class="wrap">
		  	<p class="copy">慈溪市耶利米智能科技有限公司 <i id="copyrightsign">©</i> All Rights Reseverd | 浙ICP备1500142 | Power by  <a href="http://weibo.com/u/2703314922" target="_blank">July</a></p><!--  <p class="design"> </p> -->
	      	<div class="clear"></div>
	      	<a href="http://webscan.360.cn/index/checkwebsite/url/www.16me.cn"><img border="0" src="http://img.webscan.360.cn/status/pai/hash/6aee8b7837e3280bc6451857a2867a8a"/></a>
		</div>	
	</div>
	<a href="javascript:void(0)" id="goTop" class="uk-button" data-uk-smooth-scroll="" style="display: inline;"><i class="uk-icon-chevron-up"></i></a>

	<script src="/static/js/jquery-1.11.0.min.js"></script>
    <script src="/static/uikit/js/uikit.min.js"></script>
    <script src="/static/js/site.js"></script>
    <script src="/static/uikit/js/addons/sticky.js"></script>
    {{.LayoutFooter}}
</body>
</html>

