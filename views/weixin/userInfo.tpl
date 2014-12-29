<div class="uk-panel">
    <form id="form" class="uk-form uk-form-stacked">
        <fieldset>
            <legend>用户信息完善</legend>
            <input type="hidden" name="id" value="{{.User.Id}}"></input>
            <div class="uk-form-row">
                <label class="uk-form-label" for="form-s-it">唯一ID</label>
                <p>{{.User.UserId}}</p>
            </div>
            <div class="uk-form-row">
                <label class="uk-form-label" for="form-s-it">姓名</label>
                <div class="uk-form-controls">
                    <input type="text" id="form-s-it" name="name" placeholder="请输入姓名" value="{{.User.Name}}" class="uk-width-1-1">
                </div>
            </div>
            <div class="uk-form-row">
                <label class="uk-form-label" for="form-s-ip">电话</label>
                <div class="uk-form-controls">
                    <input type="text" id="form-s-ip" name="phone" placeholder="请输入电话" value="{{.User.Phone}}" class="uk-width-1-1">
                </div>
            </div>
            <div class="uk-form-row">
                <span class="uk-form-label">性别</span>
                <div class="uk-form-controls"> 
                    <label><input type="radio" name="sex" value="男" {{if compare .User.Sex "男"}}checked{{end}}> 男 </label>
                    &nbsp;&nbsp;&nbsp;&nbsp;
                    <label><input type="radio" name="sex" value="女" {{if compare .User.Sex "女"}}checked{{end}}> 女 </label>
                </div>
            </div>
            <div class="uk-form-row">
                <a id="btnSubmit" class="uk-button">提交</a>
            </div>
        </fieldset>
    </form>
</div>
<script>
$(function(){
    $("#btnSubmit").on("click",function(){
        var params = $("#form").serializeObject();
        $.post("/weixin/user/save",params,function(data,status){
            if(data.Result == true){
                alert("保存成功");
            }else{
                alert(data.Message);
            }
        });
    });
});
</script>