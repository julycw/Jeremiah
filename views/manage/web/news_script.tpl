<script src="/static/codemirror/lib/codemirror.js"></script>
<script src="/static/codemirror/mode/markdown/markdown.js"></script>
<script src="/static/codemirror/addon/mode/overlay.js"></script>
<script src="/static/codemirror/mode/xml/xml.js"></script>
<script src="/static/codemirror/mode/gfm/gfm.js"></script>
<script src="/static/js/marked.min.js"></script>
<script src="/static/js/htmleditor.js"></script>
<script>
$(function(){
	var tipModal = $.UIkit.modal("#modal");
	var addModal = $.UIkit.modal("#add-modal");
	var modifyModal = $.UIkit.modal("#modify-modal");
	var viewModal = $.UIkit.modal("#view-modal");
	$("#btnAdd").on("click", function() {
        addModal.show();
    });

	$(".btn-modify").on("click", function() {
		$("#editHtmlWrapper").html("<textarea id=\"editHtml\" name=\"html\"></textarea>");
		$.post("/api/manage/web/news/getNewsById",{id:$(this).attr("target")},function(data){
			$("#formModify :input[name='id']").val(data.Id);
			$("#formModify :input[name='title']").val(data.Title);
			var textarea = $("#editHtml");
			textarea.text(data.Html);
			$.UIkit.htmleditor(textarea);
			$("#formModify :input[name='category']").val(data.Category);
			$("#formModify :input[name='state']").val(data.State);
        	modifyModal.show();
		});
    });

	$(".btn-preview").on("click", function() {
		viewModal.show();
    });

	$("#btnAddSubmit").on("click",function(){
		$("#formAdd").submit();
	});
	$("#btnModifySubmit").on("click",function(){
		$("#formModify").submit();
	});
	$("#btnDeleteSubmit").on("click",function(){
		var id = $("#formModify input[name='id']").val();
		$.post("/api/manage/web/news/deleteNewsById",{id:id},function(data){
			if(data == "success"){
				document.location.href="/manage/web/news/";
			}else{
				alert(data);
			}
		});
	});

});
</script>