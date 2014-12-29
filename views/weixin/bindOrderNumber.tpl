<div class="uk-panel">
	<form action="/weixin/bindOrderNumber" method="post" class="uk-form uk-form-stacked">
		<fieldset>
            <legend>订单配对</legend>
            <input type="hidden" name="userID" value="{{.userID}}" />
			<div class="uk-form-row">
                <label class="uk-form-label" for="form-s-it">订单配对号</label>
				<input id="form-s-it" type="text" name="order_suffix" value="" placeholder="请输入订单配对号" class="uk-width-1-1" />
			</div>
			<div class="uk-form-row">
				<input type="submit" class="uk-button" value="提交"/>
			</div>
        </fieldset>
	</form>
</div>
