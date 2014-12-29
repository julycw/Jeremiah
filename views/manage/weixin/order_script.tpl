<script>
$(function(){
	var $codeImage = $("#orderCodeImage");
	$codeImage.attr("src",$codeImage.attr("mask-src")+$codeImage.attr("orderid"));

	$(".btnShowOrderDetails").on("click",function(){
		var $this= $(this);
		var $icon = $this.find("i");
		var $detailWrapper = $this.parent().parent().parent().find(".order-details");
		if($icon.hasClass("uk-icon-chevron-down")){
			SectionAction($icon,$detailWrapper,"show");
		}else{
			SectionAction($icon,$detailWrapper,"hide");
		}
	});

	$(".btnShowDayDetail").on("click",function(){
		var $this = $(this);
		var $articles = $this.parent().parent().find("article").each(function(){
			var $this_article = $(this);
			var $icon = $this_article.find(".btnShowOrderDetails i");
			var $detailWrapper = $this_article.find(".order-details");
			SectionAction($icon,$detailWrapper,"show");
		});
	});

	$(".btnHideDayDetail").on("click",function(){
		var $this = $(this);
		var $articles = $this.parent().parent().find("article").each(function(){
			var $this_article = $(this);
			var $icon = $this_article.find(".btnShowOrderDetails i");
			var $detailWrapper = $this_article.find(".order-details");
			SectionAction($icon,$detailWrapper,"hide");
		});
	});

	$(".btn-submit-details").on("click",function(){
		var $this = $(this);
		$this.hide();
		var orderId = $this.attr("target-id");
		var detail = $this.parent().parent().find("div .order-details-input").val();
		$.post("/api/manage/weixin/fix/addFixOrderDetail",{id:orderId,detail:detail},function(data,status){
			if (status == "success"){
				if(data.Id >0){
					var $detailList = $("#order-details-list-" + orderId);
					appendListItem($detailList,data,true);
					$this.parent().parent().find("div .order-details-input").val("");
				}else{
					alert("添加失败，请填写内容");
				}
			}else{
				alert("网络错误");
			}
			$this.fadeIn();
		});
	});

	$("#btnBindUser").on("click",function(){
		var orderNumber = $("#orderNumber").text();
		var $userInfoWrapper = $("#panel-user ul");
		$userInfoWrapper.empty();
		$.post("/api/manage/weixin/fix/getOrderBindedUser",{
			orderNumber:orderNumber
		},function(data,status){
			if (data.Id > 0 ){
				$("#panel-code").hide();
				$userInfoWrapper.append("<li>用户ID: <b id='bindUserId'>"+data.UserId+"</b></li>");
				$userInfoWrapper.append("<li>微信号: 订阅号无法获取</li>");
				$userInfoWrapper.append("<li>姓名: <b>"+data.Name+"</b></li>");
				$userInfoWrapper.append("<li>电话: <b>"+data.Phone+"</b></li>");
				$userInfoWrapper.append("<li>昵称: <b>"+data.NickName+"</b></li>");
				$("#panel-info").show();
			}else{
				$("#panel-code").fadeIn();
				$("#btnSelectUser").click();
			}
		});
	});

	$("#btnSelectUser").on("click",function(){
		alert("暂无用户可选");
	});

	$("#btnAddOrder").on("click",function(){
		var orderNumber = $("#orderNumber").text();
		var userId = $("#bindUserId").text();
		if(userId.length == 0 || userId == ""){
			$("#btnSelectUser").click();
			return;
		}
		var description = $("#inputDescription").val();
		if(description.length == 0 || description == ""){
			alert("请输入描述");
			return;
		}
		$.post("/api/manage/weixin/fix/addFixOrder",{
			orderNumber:orderNumber,
			userID:userId,
			description:description
		},function(data,status){
			if(data.Id >0){
				alert("提交成功");
				document.location.href = "/manage/weixin/order";
			}else{
				alert("提交失败");
			}
		});
	});


	$("#btnOrderRefresh").on("click",function(){
		var $this = $(this);
		var $codeImage = $("#orderCodeImage");
		$codeImage.attr("src","/static/img/logo.jpg");
		$this.addClass("uk-icon-spin");
		$.post("/api/manage/weixin/fix/getOrderNumber",function(data,status){
			if (data!="" && data.length >= 8){
				$("#orderNumber").text(data);
				$codeImage.attr("src",$codeImage.attr("mask-src")+data);
				$("#panel-info").fadeOut(function(){
					$("#panel-code").fadeIn();
				});
			}
			setTimeout(function () {
				$this.removeClass("uk-icon-spin");
		    }, 1000);
		});
	});
});

function SectionAction($icon,$detailWrapper,action){
	if(action == "show"){
		$icon.removeClass("uk-icon-chevron-down");
		$icon.addClass("uk-icon-chevron-up");
		$detailWrapper.slideDown();

		var $loadingTip = $detailWrapper.find(".loading-tip");
		if($loadingTip.is(":visible")){
			var id = $detailWrapper.attr("target-id");
			$.post("/api/manage/weixin/fix/getFixOrderDetailsByOrderId",{id:id},function(data,status){
				if(status == "success"){
					var $detailList = $("#order-details-list-" + id);
					for (var i = data.length-1; i >= 0 ; i--) {
						var detail = data[i];
						appendListItem($detailList,detail);
					};
					$loadingTip.fadeOut();
				}else{
					alert("网络错误");
				}
			});
		}
	}else{
		$icon.removeClass("uk-icon-chevron-up");
		$icon.addClass("uk-icon-chevron-down");
		$detailWrapper.slideUp();
	}
}

function appendListItem($target,item,isNew){
	var htmlStr = "";
	htmlStr += "<li>";
	htmlStr += "<div style='float:right;width:97px;'>";
	if (isNew != undefined && isNew == true){
		htmlStr +="<b>New!!!</b><br/>"
	}
	htmlStr += item.CreateOn.substr(5,5) +" "+item.CreateOn.substr(11,8);
	htmlStr += "</div>";
	htmlStr += "<div style='margin-right:105px;'>";
	htmlStr += item.Detail;
	htmlStr += "</div>";
	htmlStr += "</li>";
	$target.prepend(htmlStr);
}
</script>