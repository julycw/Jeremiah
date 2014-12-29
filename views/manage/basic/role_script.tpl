<script>
$(function(){
	var tipModal = $.UIkit.modal("#modal");
	var addModal = $.UIkit.modal("#add-modal");
	var modifyModal = $.UIkit.modal("#modify-modal");
	$("#btnAdd").on("click", function() {
        addModal.show();
    });

	$(".btn-modify").on("click", function() {
		$.post("/api/manage/basic/role/getRoleById",{id:$(this).attr("target")},function(data){
			$("#formModify :input[name='id']").val(data.Id);
			$("#formModify :input[name='roleid']").val(data.RoleId);
			$("#formModify :input[name='rolename']").val(data.RoleName);
			$("#formModify input[type='checkbox']").prop("checked",false);
			var avaliables = data.Avaliables.split(",");
			for (var i = avaliables.length - 1; i >= 0; i--) {
				$("#formModify input[avaliable='"+avaliables[i]+"']").prop("checked",true);
			};
        	modifyModal.show();
		});
    });

	$("#btnAddSubmit").on("click",function(){
		var avaliableList = "";
		$("#formAdd .avaliable").each(function(){
			if($(this).is(":checked")){
				if(avaliableList !== ""){
					avaliableList = avaliableList+",";
				}
				avaliableList = avaliableList+$(this).attr("avaliable");
			}
		});
		$("#formAdd input[name='avaliables']").val(avaliableList);
		$("#formAdd").submit();
	});
	$("#btnModifySubmit").on("click",function(){
		var avaliableList = "";
		$("#formModify .avaliable").each(function(){
			if($(this).is(":checked")){
				if(avaliableList !== ""){
					avaliableList = avaliableList+",";
				}
				avaliableList = avaliableList+$(this).attr("avaliable");
			}
		});
		$("#formModify input[name='avaliables']").val(avaliableList);
		$("#formModify").submit();
	});

	$("#btnDeleteSubmit").on("click",function(){
		var id = $("#formModify input[name='id']").val();
		$.post("/api/manage/basic/role/deleteRoleById",{id:id},function(data){
			if(data == "success"){
				document.location.href="/manage/basic/role/";
			}else{
				alert(data);
			}
		});
	});
});
</script>