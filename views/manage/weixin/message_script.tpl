<script>
var modal;

$(function() {
    modal = $.UIkit.modal("#modal");
    getMessage(10);
    $("#btnMore").on("click", function() {
        var lastId = $("#tbMessage tr:last").attr("msgid");
        getMessage(15, lastId);
    });
});

function getMessage(size, lastId) {
    var $tbMessage = $("#tbMessage");
    var url = "/api/manage/weixin/message/getMessageList";
    var params = {};
    if (size != undefined && !isNaN(size) && size > 0) {
        params["size"] = size;
    }

    if (lastId != undefined) {
        params["lastID"] = lastId;
    }

    $.post(url, params, function(data) {
        if (data.length == 0) {
            modal.show();
            return;
        }
        for (var i = 0; i < data.length; i++) {
            $tbMessage.append("<tr msgid='" + data[i].Id + "'><td>" + data[i].User.UserName + "</td><td>" + data[i].User.NickName + "</td><td>" + data[i].Content + "</td><td>" + data[i].CreateOnFmt + "</td></tr>");
        };
    });
}
</script>
