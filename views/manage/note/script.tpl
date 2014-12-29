<script>
$(function() {
    $(document).on({
        "click": function() {
            addTag($.trim($(this).text()));
        },
    }, ".tag-select");

    $("#btnAddNote").on({
        "click": function() {
            $("#formNote").submit();
        },
    });

    $(".tag-list").each(function() {
        var tags = $(this).attr("tags");
        if (tags == "" || tags == undefined) {
            return;
        }

        var tag_list = tags.split(",");
        for (var i = 0; i < tag_list.length; i++) {
            if (tag_list[i] == "") continue;
            $(this).append('<span class="uk-badge tag-piece">' + tag_list[i] + '</span>&nbsp;');
        };
    });
});

function addTag(tagName) {
    var $tagInput = $("input[name='tags']");
    var t = $tagInput.val();

    var tags = t.split(",");
    var index = tags.indexOf(tagName)
    if (index != -1) {
        tags.splice(index,1);
    } else {
        tags.push(tagName);
    }

    if (t == undefined || t == "") {
        t = tagName;
    } else {
        t = t + "," + tagName;
    }

    t = "";
    for (var i = 0; i < tags.length; i++) {
    	if (tags[i] == undefined || tags[i] == "") continue;
        if (t == "") {
            t = tags[i];
        } else {
            t = t + "," + tags[i];
        }
    }

    $tagInput.val(t);
}
</script>
