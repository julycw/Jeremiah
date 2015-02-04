<script>
	var tipModal = $.UIkit.modal("#view-modal");
	var $erpList = $("#erp-list");
	var $btnSearch = $("#btnSearch");

	$(function(){
		eventBind();
		requestData();
	});

	function eventBind(){
		$($erpList).on("click",".btn-erp-detail",function(){
			tipModal.show();
		});
		$btnSearch.on("click",function(){
			requestData();
		});
	}

	function requestData(){
		var selector = serializeForm("search-form");
		$.get("/api/manage/web/erp",selector,function(data){
			$erpList.empty();
			for (var i = 0; i < data.length; i++) {
				var item = data[i];
				$erpList.append("<tr><td class=\"btn-erp-detail\" itemid=\""+item.IDStr+"\">+</td><td>"+item.Model+"</td><td>￥"+item.JDPrice+"</td><td>￥"+item.Price+"</td><td>￥"+item.SuggestPrice+"</td><td>"+item.CPUModel+"</td><td>"+item.GraphicsModel+"</td><td>"+item.MemorySize+"</td><td>"+item.DiskSize+"</td></tr>");
			};
		});
	}
</script>