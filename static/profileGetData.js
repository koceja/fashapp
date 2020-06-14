$(document).ready(function() {
        const urls = $(".card-columns").html()
        $(".card-columns").empty()
        if (urls == "") {
            return
        } else {
            $(urls).appendTo(".card-columns")
        }
})
