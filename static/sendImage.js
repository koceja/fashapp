$(document).ready(function() {
    $(".row").on("click", "div.card-link", function() {
        const imageUrl = $(this).parent().parent().children("a").children("img").attr('src');
        const path = window.location.pathname
        const imgData = {image: imageUrl}
        console.log(path);
        var successFunc = function() {
            console.log("success!");
        }
        
        $.ajax({
            type: "POST",
            url: path,
            data: imgData,
            success: successFunc,
      });
    });
})
