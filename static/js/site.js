$(function(){
    $(window).scroll(function(){
        if ($(window).scrollTop()>200){
            $("#goTop").fadeIn(300);
        }else{
            $("#goTop").fadeOut(200);
        }
    });
    $("#goTop").hide();

    MenuHighlight();

    //彩蛋
    $("#copyrightsign").on("mouseenter",function(){
        timer = setInterval(function(){
            joke++;
            if(joke>=5){
                $("body").animate({opacity:"0.1"},1500,function(){
                    window.location.href= "/manage";
                });
            }
        },1000);
    });
    $("#copyrightsign").on("mouseout",function(){
        joke = 0;
        console.info("mouseout");
        clearInterval(timer);
    });
});

var timer = undefined;
var joke = 0;

//菜单加亮
function MenuHighlight(){
    console.info(document.URL);
    $("#menu-items li").each(function() {
        var $this = $(this);
        var $linkA = $this.find("a");
        var linkstr = $linkA.attr("href");
        var index = document.URL.indexOf(linkstr)
        if (document.URL.indexOf(linkstr) != -1) {
            $("#menu-items li").attr("class", "");
            $this.attr("class", "active");
        }
    });
};

//数组去重
Array.prototype.unique = function() {
    var n = {},
        r = []; //n为hash表，r为临时数组
    for (var i = 0; i < this.length; i++) //遍历当前数组
    {
        if (!n[this[i]]) //如果hash表中没有当前项
        {
            n[this[i]] = true; //存入hash表
            r.push(this[i]); //把当前数组的当前项push到临时数组里面
        }
    }
    return r;
}

//删除数组中下标为n的元素
Array.prototype.del = function(n) {　
    return this.splice(n,1);
}