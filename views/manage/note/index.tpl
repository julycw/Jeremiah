<div>
    <ul class="uk-subnav uk-subnav-pill" data-uk-switcher="{connect:'#subnav-pill-content'}">
        <li class=""><a href="#">浏览</a>
        </li>
        <li class=""><a href="#">新增</a>
        </li>
    </ul>
    <ul id="subnav-pill-content" class="uk-switcher">
        <li class="">
            {{range .dayNoteLists}}
            <div class="uk-panel uk-panel-box">
                <h3 class="uk-panel-title">{{.Day}}</h3>
                {{range .NoteList}}
                <article class="uk-comment">
                    <header class="uk-comment-header">
                        <div class="tag-list" tags="{{.Tags}}">
                        </div>
                        <h4 class="uk-comment-title">{{str2html .Content}}</h4>
                        <div class="uk-comment-meta">由
                            <span class="uk-text-primary uk-text-bold">{{nickName .UserName}}</span> 于 <b>{{dateformat .CreateOn "15:04:05"}} 发表</b>.</div>
                    </header>
                </article>
                {{end}}
            </div>
            <br/>{{end}}
        </li>
        <li class="">
            <div class="uk-panel uk-panel-box">
                <form id="formNote" class="uk-form" method="post" action="/manage/note/add">
                    <div class="uk-form-row">
                        <input name="tags" type="text" value="" placeholder="分类标签" readonly="readonly">
                        <br/>
                        <a href="#" class="uk-badge tag-piece tag-select">维修</a>
                        <a href="#" class="uk-badge tag-piece tag-select">售后</a>
                        <a href="#" class="uk-badge tag-piece tag-select">通知</a>
                        <a href="#" class="uk-badge tag-piece tag-select">重要</a>
                        <a href="#" class="uk-badge tag-piece tag-select">其他</a>
                    </div>

                    <div class="uk-form-row">
                        <div class="uk-form-controls">
                            <textarea id="form-s-t" name="content" cols="30" rows="5" placeholder="内容"></textarea>
                        </div>
                    </div>

                    <div class="uk-form-row">
                        <a id="btnAddNote" class="uk-button">提交</a>
                    </div>
                </form>
            </div>
        </li>
    </ul>
</div>
