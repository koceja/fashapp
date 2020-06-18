    // $(".card-link").click(function() {
    //     $(this).css("color", "red");
    // });

    
   $(".row").on("click", "div.card-link", function() {
    $(this).css("color", "red");
    $(this).parent().animate({marginBottom: "1.35rem", padding: "1.15rem"}, 200);
    $(this).parent().animate({marginBottom: "1.25rem", padding: "1.25rem"}, 200);

});