<script>
$(function(){
	$(".form-field").on("input propertychange",function(){
		var $this = $(this);
		var $btnsWrapper = $this.next(".wrapper-btns");
		if($this.val() != $this.attr("value-origin")){
			$btnsWrapper.css("display","inline-block");
			$this.removeClass("uk-form-success");
			$this.removeClass("uk-form-danger");
			$this.addClass("uk-form-modify");
		}else{
			$btnsWrapper.css("display","none");
			$this.removeClass("uk-form-modify");
			$this.removeClass("uk-form-success");
			$this.removeClass("uk-form-danger");
		}
	});
	$(".btn-save").on("click",function(){
		var $this = $(this);
		var $btnsWrapper = $this.parent();
		var $formField = $btnsWrapper.prev(".form-field");
		var id = $formField.attr("target-id");
		var value = $formField.val();
		$.post("/api/manage/basic/option/modifyOptionById",{
			id:id,
			value:value
		},function(data,status){
			if(status == "success"){
				if (data == "success"){
					$formField.attr("value-origin",value);
					$formField.val(value);
					$formField.removeClass("uk-form-modify");
					$formField.removeClass("uk-form-danger");
					$formField.addClass("uk-form-success");
					$btnsWrapper.fadeOut();
				}else{
					$formField.removeClass("uk-form-modify");
					$formField.removeClass("uk-form-success");
					$formField.addClass("uk-form-danger");
					alert("保存失败");
				}
			}else{
				alert("好像断网了？");
			}
		});
	});
	$(".btn-cancel").on("click",function(){
		var $this = $(this);
		var $btnsWrapper = $this.parent();
		var $formField = $btnsWrapper.prev(".form-field");
		var temp = $formField.attr("value-origin");
		$formField.val(temp);
		$btnsWrapper.fadeOut();
		$formField.removeClass("uk-form-modify");
		$formField.removeClass("uk-form-success");
		$formField.removeClass("uk-form-danger");
	});
});
</script>