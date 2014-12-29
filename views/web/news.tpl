<div class="news-panel uk-width-1-1 uk-width-medium-4-5 uk-width-large-4-5 uk-container-center">
	<article class="detail">
		<header style="background-color:white;">
			<h1 class="uk-article-title" data-uk-sticky style="background-color:white;padding-top:10px;">
                <a href="/news" class="news-btn-list uk-icon-button uk-icon-list"></a>
				<a href="javascript:void(0);" class="news-btn-send uk-icon-button uk-icon-send"></a>
				<span style="margin-right:50px;display:block;">{{.news.Title}}</span>
			</h1>
			<p class="uk-article-meta">
				作者<span class="ago">{{.news.CreateBy}}</span>
				发表于<span class="ago"> {{dateformat .news.CreateOn "2006-01-02 15:04"}}</span>
			</p>
		</header>
    	<hr class="uk-article-divider">
		<section>
		{{str2html .news.Html}}
		</section>
    	<hr class="uk-article-divider">
    	<h4 class="tm-article-subtitle">微信评论</h4>
		<ul class="uk-comment-list">
            <li>
                <article class="uk-comment">
                    <header class="uk-comment-header">
                        <img class="uk-comment-avatar" src="/static/img/avator.jpg" alt="">
                        <h4 class="uk-comment-title">张三</h4>
                        <div class="uk-comment-meta">于 2014年12月24日 在微信中评论</div>
                    </header>
                    <div class="uk-comment-body">
                        <p>巴拉巴拉。。。</p>
                    </div>
                </article>
                <ul>
                    <li>
                        <article class="uk-comment">
                            <header class="uk-comment-header">
                                <img class="uk-comment-avatar" src="/static/img/avator.jpg" alt="">
                                <h4 class="uk-comment-title">李四</h4>
                        		<div class="uk-comment-meta">于 2014年12月24日 在微信中评论</div>
                            </header>
                            <div class="uk-comment-body">
                                <p>巴拉巴拉。。。巴拉巴拉。。。</p>
                            </div>
                        </article>
                    </li>
                    <li>
                        <article class="uk-comment">
                            <header class="uk-comment-header">
                                <img class="uk-comment-avatar" src="/static/img/avator.jpg" alt="">
                                <h4 class="uk-comment-title">李四</h4>
                        		<div class="uk-comment-meta">于 2014年12月24日 在微信中评论</div>
                            </header>
                            <div class="uk-comment-body">
                                <p>巴拉巴拉。。。巴拉巴拉。。。</p>
                            </div>
                        </article>
                    </li>
                </ul>
            </li>
            <li>
                <article class="uk-comment">
                    <header class="uk-comment-header">
                        <img class="uk-comment-avatar" src="/static/img/avator.jpg" alt="">
                        <h4 class="uk-comment-title">王五</h4>
                		<div class="uk-comment-meta">于 2014年12月24日 在微信中评论</div>
                    </header>
                    <div class="uk-comment-body">
                        <p>巴拉巴拉。。。巴拉巴拉。。。</p>
                    </div>
                </article>
            </li>
        </ul>
	</div>
</div>

				