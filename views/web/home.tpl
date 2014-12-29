<div class="wrap">
	<div class="icon_grids">
		<div class="section group">
			<div class="grid_1_of_3 images_1_of_3">
				<div class="grid_icon"><a href="/about"><img src="/static/img/g1.png"></a></div>
			  	<h3>成功案例</h3>
			  	<div class="text-wrapper">
				  	<p>{{if .success_case}}
					{{.success_case|str2html}}
					{{else}}
					Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.
					{{end}}
					</p>
				</div>
		     	<div class="button"><span><a href="/about">更 多</a></span></div>
			</div>
			<div class="grid_1_of_3 images_1_of_3">
				<div class="grid_icon"><a href="/services"><img src="/static/img/g2.png"></a></div>
				<h3>客户定制</h3>
				<div class="text-wrapper">
				<p>{{if .custom_made}}
				{{.custom_made|str2html}}
				{{else}}
				Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.
				{{end}}</p></div>
				<div class="button"><span><a href="/services">更 多</a></span></div>
			</div>
			<div class="grid_1_of_3 images_1_of_3">
				<div class="grid_icon"><a href="/services"><img src="/static/img/g3.png"></a></div>
				<h3>售后服务</h3>
				<div class="text-wrapper">
				<p>{{if .service}}
				{{.service |str2html}}
				{{else}}
				Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.
				{{end}}</p></div>
				<div class="button"><span><a href="/services">更 多</a></span></div>
			</div>
		</div>
	</div>
    <hr class="uk-article-divider">
	<div class="section group">						
		<div class="col_1_of_3 span_1_of_3">
			<h3>欢迎！</h3>
			<p>{{if .welcome}}
		  {{.welcome|str2html}}
		  {{else}}
		  Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.
		  {{end}}</p>
	 	</div>
	   	<div class="col_1_of_3 span_1_of_3">
			<h3>我们做什么</h3>
		   	<ul style="padding-left:0px;">
		   	{{if .whatWeDoList}}
			   	{{range .whatWeDoList}}
			   	<li><a href="/news/{{.Id}}">{{.Title}}</a></li>
			   	{{end}}
		   	{{else}}
				<li>Lorem ipsum dolor sit amet qui officia</li>
				<li>Duis aute irure dolor in culpa qui</li>
				<li>Sunt in culpa qui officia vel illum qui</li>
				<li>vel illum qui dolorem eum wise man therefore</li>
				<li>The wise man therefore in culpa qui officia</li>
				<li>Sunt in culpa qui officia culpa qui officia</li>
				<li>Lorem ipsum dolor sit amet qui officia</li>
				<li>Duis aute irure dolor in culpa qui</li>
				<li>Sunt in culpa qui officia vel illum qui</li>
				<li>vel illum qui dolorem eum wise man therefore</li>
			{{end}}
	        </ul>
	      	<div class="button"><span><a href="/WhatWeDo">更多</a></span></div>
		</div>
		<div class="col_1_of_3 span_1_of_3">
		  	<div class="new-products">
				<h3>最新产品</h3>
				{{if .productList}}
				{{range .productList}}
				<p><a href="/gallery/{{.Id}}"><img src="/resource/image/{{getFileName .Image}}/160x100" alt="{{.Title}}"></a></p>
				{{end}}
				{{else}}
					<p><a href="/gallery"><img src="/static/img/product/product1.jpg" alt=""></a></p>
					<p><a href="/gallery"><img src="/static/img/product/product2.jpg" alt=""></a></p>
					<p><a href="/gallery"><img src="/static/img/product/product3.jpg" alt=""></a></p>
					<p><a href="/gallery"><img src="/static/img/product/product4.jpg" alt=""></a></p>
					<p><a href="/gallery"><img src="/static/img/product/product5.jpg" alt=""></a></p>
					<p><a href="/gallery"><img src="/static/img/product/product1.jpg" alt=""></a></p>
				{{end}}
				<div class="button"><span><a href="/gallery">查看全部</a></span></div>
	      	</div>
    	</div>
	</div>
    <hr class="uk-article-divider">
</div>

