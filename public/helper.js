$(function(){
    $('.dir').dblclick(function(){
        var path = $(this).data('path');
        $.post("/", {path: path}).done(function(data){
            $('body').html(data)
        })
    })
    
    $('.dir').click(function(){
        var path = $(this).data('path');
        $('div#pathfield').html(path);
    })
    
    $('#sortbutton').click(function(){
        var mode = $('button#mode').text();
        console.log(mode)
        var path = $('#pathfield').text();
        if (mode == "Classic"){
            classic(path);
        } else {
            var pattern = $('input#mode').val();
            custom(path, pattern);
        }
       
    })
    
    $('button#mode').click(function(){
        $('#modefield').toggle();
        $(this).text($(this).text() == "Classic" ? "Custom" : "Classic");
    });
    
    $('button#mode').click();
})


function classic(path){
    $.post("/sort", {path: path}).done(function(data){
            $('body').html(data)
    })
}

function custom(path, pattern){
    $.post("/regsort", {path: path, pattern: pattern}).done(function(data){
            $('body').html(data)
    })
}
    