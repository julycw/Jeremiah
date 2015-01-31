<script>
	$(function(){
		//设置样式
		$("#scanForm div").addClass("uk-form-row");
		$("#scanForm div label").addClass("uk-form-label");
		$("#scanForm div input").addClass("uk-width-1-1").wrap("<div class=\"uk-form-controls\"></div>");
		//数据库ID不允许修改
		$("#idStr").attr("readonly","true");
		//事件绑定
		$("#btnSubmit").on("click",function(){
			// var computerData = $("#scanForm").serializeObject();
			var computerData = serializeForm("scanForm");
			$.post("/api/manage/web/erp/add",computerData,function(data){
				console.info(data);
			});
		});
		$("#btnClose").on("click",function(){
			window.opener=null;
			window.open('','_self');
			window.close();
		});
	});
</script>