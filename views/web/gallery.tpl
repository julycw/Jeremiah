<div class="gallery">
	<ul class="uk-grid uk-text-center uk-grid-preserve" data-uk-grid-match data-uk-grid-margin>
		{{range .productList}}
        <li class="uk-width-1-1 uk-width-medium-1-2 uk-width-large-1-4"><div class="uk-panel"><a href="/gallery/{{.Id}}"><img src="/resource/image/{{getFileName .Image}}/360x220" alt="{{.Title}}"></a></div></li>
		{{end}}
    </ul>
</div>
