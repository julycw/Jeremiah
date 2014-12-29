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
		$.post("/api/manage/web/product/getProductById",{id:$(this).attr("target")},function(data){
			console.info(data);
			$("#formModify :input[name='id']").val(data.Id);
			$("#formModify :input[name='image']").val(data.Image);
			$("#formModify :input[name='images']").val(data.Images);
			$("#formModify :input[name='title']").val(data.Title);
			$("#formModify :input[name='description']").val(data.Description);
			var textarea = $("#editHtml");
			textarea.text(data.Html);
			$.UIkit.htmleditor(textarea);
			$("#formModify :input[name='type']").val(data.Type);
			$("#formModify :input[name='sort']").val(data.Sort);
			if(data.Hot=="1"){
				$("#formModify :input[name='hot']").prop("checked",true);
			}else{
				$("#formModify :input[name='hot']").prop("checked",false);
			}
			if(data.Visible=="1"){
				$("#formModify :input[name='visible']").prop("checked",true);
			}else{
				$("#formModify :input[name='visible']").prop("checked",false);
			}
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
		$.post("/api/manage/web/product/deleteProductById",{id:id},function(data){
			if(data == "success"){
				document.location.href="/manage/web/product/";
			}else{
				alert(data);
			}
		});
	});
});
</script>