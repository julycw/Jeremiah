<div class="product-panel uk-width-1-1 uk-width-medium-4-5 uk-width-large-4-5 uk-container-center">
	<div class="uk-grid">
        <div class="uk-width-1-1">
        	<div class="uk-panel">
        		<article class="uk-article">
        			<div class="product-header" data-uk-sticky="{top:5}">
                        <a href="/gallery" class="product-btn-list uk-icon-button uk-icon-list"></a>
        				<a href="javascript:void(0);" class="product-btn-send uk-icon-button uk-icon-send"></a>
					    <h1 class="uk-article-title" style="margin-left:20px;">{{.product.Title}}</h1>
					    <p class="uk-article-lead" style="margin-left:20px;">{{.product.Description}}</p>
        			</div>
    				<hr class="uk-article-divider">
				    <section class="product-info">
				    	<div class="product-html">
            			{{.product.Html | str2html}}
            			</div>
            		</section>
				</article>
        	</div>
        </div>
    </div>
    <hr class="uk-article-divider">
    <div class="uk-panel uk-container-center uk-text-center"><img src="{{.product.Image}}" alt="{{.data.Title}}"/></div>
</div>