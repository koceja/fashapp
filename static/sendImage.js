$(document).ready(function() {
    $(".row").on("click", "div.card-link", function() {
        const imageUrl = $(this).parent().parent().children("a").children("img").attr('src');
        var path = window.location.pathname
        path = path.split("/")
        const person = path[2]
        console.log(person)
        console.log(imageUrl)

        const submitData = {personId: person, 
                        image: imageUrl}
        var successFunc = function() {
            console.log("Sent image to database");
        }
        
        $.ajax({
            type: "POST",
            url: path,
            data: submitData,
            success: successFunc,
      });
    });
})
