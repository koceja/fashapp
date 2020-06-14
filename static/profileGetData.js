$(document).ready(function() {
        const urls = $(".card-columns").html()
        console.log(urls)
        $(urls).appendTo(".card-columns")
})
