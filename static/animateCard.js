$(".card").mouseenter(function () {
    $(this).stop();
    $(this).animate({ marginTop: "-5px", marginBottom: "5px" }, "fast");
  });

  $(".card").mouseleave(function () {
    $(this).stop();
    $(this).animate({ marginTop: "0px", marginBottom: "0px"}, "fast");
  });