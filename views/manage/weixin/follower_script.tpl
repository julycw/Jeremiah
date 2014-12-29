<script>
var modal;

$(function() {
    modal = $.UIkit.modal("#modify-modal");
    getFollower(0,100);

    $("#tbFollower").on("click", ".btnModify", function() {
        var followerId = $(this).attr("follower-id");
        var url = "/api/manage/weixin/follower/getFollowerById";
        $.post(url, {
            id: followerId
        }, function(data) {
            $("#formModify input[name='id']").val(data.Id);
            $("#formModify input[name='userid']").val(data.UserId);
            $("#formModify input[name='username']").val(data.UserName);
            $("#formModify input[name='nickname']").val(data.NickName);
            $("#formModify input[name='description']").val(data.Description);
            $("#formModify input[name='name']").val(data.Name);
            $("#formModify input[name='phone']").val(data.Phone);
            $("#formModify input[name='sex']").val(data.Sex);
            $("#Name").text(data.Name);
            $("#Sex").text(data.Sex);
            $("#Phone").text(data.Phone);
            $("#UserId").text(data.UserId);
            modal.show();
        })
    });

    $("#btnUpload").on("click",function(){
        var form = $("#formModify");
        form.submit();
    });
});


function getFollower(page, size) {
    var $tbFollower = $("#tbFollower");
    var url = "/api/manage/weixin/follower/getFollowerList";
    var params = {};
    if (size != undefined && !isNaN(size) && size > 0) {
        params["size"] = size;
    }

    if (page != undefined && !isNaN(size) && size > 0) {
        params["page"] = page;
    }

    $.post(url, params, function(data) {
        for (var i = 0; i < data.length; i++) {
            $tbFollower.append("<tr><td>" + "..." + data[i].UserId.substr(data[i].UserId.length - 6, 6) + "</td><td>" + data[i].Name + "</td><td>" + data[i].NickName + "</td><td>" + data[i].Description + "</td><td><a follower-id='" + data[i].Id + "' href='javascript:void(0);' class='btnModify'>修改</a></td>");
        };
    });
}
</script>
