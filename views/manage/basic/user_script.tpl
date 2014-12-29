<script>
$(function(){
	var tipModal = $.UIkit.modal("#modal");
	var addModal = $.UIkit.modal("#add-modal");
	var modifyModal = $.UIkit.modal("#modify-modal");
	$("#btnAdd").on("click", function() {
        addModal.show();
    });

	$(".btn-modify").on("click", function() {
		$.post("/api/manage/basic/user/getUserById",{id:$(this).attr("target")},function(data){
			$("#formModify :input[name='id']").val(data.Id);
			$("#formModify :input[name='username']").val(data.UserName);
			$("#formModify :input[name='nickname']").val(data.NickName);
			$("#formModify :input[name='roleid']").val(data.RoleId);
        	modifyModal.show();
		});
    });

	$("#btnAddSubmit").on("click",function(){
		$("#formAdd").submit();
	});
	$("#btnModifySubmit").on("click",function(){
		$("#formModify").submit();
	});
	$("#btnDeleteSubmit").on("click",function(){
		var id = $("#formModify input[name='id']").val();
		$.post("/api/manage/basic/user/deleteUserById",{id:id},function(data){
			if(data == "success"){
				document.location.href="/manage/basic/user/";
			}else{
				alert(data);
			}
		});
	});
});
</script>