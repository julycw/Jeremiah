<div class="slider">						    								     
	<div class="fluid_container">
		<div class="camera_wrap camera_azure_skin" id="camera_wrap_1" style="margin-bottom:15px">
			{{if .hotProductList}}
			{{range .hotProductList}}
			<div data-link="/gallery/{{.Id}}" data-thumb="/resource/image/{{getFileName .Image}}/160x100" data-src="{{.Image}}">   </div>
			{{end}}
			{{else}}
		    <div data-thumb="/static/img/product/thumbnails/slider-4.jpg" data-src="/static/img/product/slider-4.jpg">   </div>
		    <div data-thumb="/static/img/product/thumbnails/slider-3.jpg" data-src="/static/img/slider-3.jpg">  </div>
		   <div data-thumb="/static/img/product/product/thumbnails/slider-1.jpg" data-src="/static/img/slider-1.jpg">   </div>
		    <div data-thumb="/static/img/product/thumbnails/slider-2.jpg" data-src="/static/img/product/slider-2.jpg">  </div>
		    {{end}}
		</div>
	</div>
	<div class="clear"></div>					       
</div>