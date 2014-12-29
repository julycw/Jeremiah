<div>
    {{range .dayFixOrderList}}
    <div class="uk-panel uk-panel-box">
        <h3 class="uk-panel-title">{{.Day}}</h3>
        {{range .FixOrderList}}
        <article class="uk-comment">
            <header class="uk-comment-header">
                <div style="float:left;">
                    <span class="uk-text-primary uk-text-bold">我的订单</span>
                </div>
                <div style="clear:both;">
                </div>
                <hr class="uk-article-divider" style="margin-top:3px;margin-bottom:5px;"/>
                <h4 class="uk-comment-title">{{.Description}}</h4>
                <div class="uk-comment-meta">
                    订单生成时间:<b>{{dateformat .CreateOn "15:04:05"}}</b>.
                </div>
            </header>
            <section class="uk-comment-body order-details">
                <ul class="uk-list uk-list-line" style="margin-top:0;">
                	{{range .OrderDetails}}
                	<li>
                		<div style='float:right;width:97px;'>
                			{{dateformat .CreateOn "01-02 15:04:05"}}
                		</div>
                		<div style='margin-right:105px;'>
                			{{.Detail}}
                		</div>
                	</li>
                	{{end}}
                </ul>
            </section>
            <br/>
        </article>
        {{end}}
    </div>
    <br/>
    {{end}}
</div>

