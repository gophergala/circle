$(function(){
    $('.dir').dblclick(function(){
        var path = $(this).data('path');
        $.post("/", {path: path}).done(function(data){
            $('body').html(data)
        })
    })
})