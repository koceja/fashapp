var images = [
	'https://m.media-amazon.com/images/I/71D9ljAwRLL._AC_UL320_.jpg',
	'https://m.media-amazon.com/images/I/71SOWu25AtL._AC_UL320_.jpg',
	'https://m.media-amazon.com/images/I/61Xz4E+Kp0L._AC_UL320_.jpg',
	'https://m.media-amazon.com/images/I/81iiLBqVkdL._AC_UL320_.jpg',
	'https://m.media-amazon.com/images/I/81rMweUpVoL._AC_UL320_.jpg',
	'https://m.media-amazon.com/images/I/81urAB+9qwL._AC_UL320_.jpg',
	'https://m.media-amazon.com/images/I/51z17bjs+KL._AC_UL320_.jpg',
	'https://m.media-amazon.com/images/I/518wdUguROL._AC_UL320_.jpg',
	'https://m.media-amazon.com/images/I/71dwaaVqzWL._AC_UL320_.jpg',
	'https://m.media-amazon.com/images/I/71m8F-BxjpL._AC_UL320_.jpg',
	'https://m.media-amazon.com/images/I/81R32Du+dsL._AC_UL320_.jpg',
	'https://m.media-amazon.com/images/I/81CJQ+eXoeL._AC_UL320_.jpg',
	'https://m.media-amazon.com/images/I/719k7ivfgOL._AC_UL320_.jpg',
	'https://m.media-amazon.com/images/I/81oz9owYoHL._AC_UL320_.jpg',
	'https://m.media-amazon.com/images/I/61d02VIrAGL._AC_UL320_.jpg',
	'https://m.media-amazon.com/images/I/61tOF0F0ryL._AC_UL320_.jpg',
	'https://m.media-amazon.com/images/I/71GKlkGwhwL._AC_UL320_.jpg',
	'https://m.media-amazon.com/images/I/51Hdpd4YNqL._AC_UL320_.jpg',
	'https://m.media-amazon.com/images/I/61CNc6KusOL._AC_UL320_.jpg',
	'https://m.media-amazon.com/images/I/91Lgd-4F1cL._AC_UL320_.jpg',
];

        if ($(document).height() < 2000) {
            var i;
                for (i = 0; i < 16; i++) {
                    const url = images[Math.floor(Math.random() * images.length)];

                    $("<div class='col mb-4'><div class='card h-100'><img src='" + url + "' class='card-img-top' alt='...' /><div class='card-body'>          <a href='#' class='card-link'>Like</a>        </div>      </div></div>").hide().appendTo(".row").fadeIn();

                }
        }
        $(window).scroll(function() {
            console.log($(window).scrollTop());
            console.log($(document).height() - $(window).height());
            if($(window).scrollTop() > $(document).height() - $(window).height() - 50) {
                var i;
                for (i = 0; i < 8; i++) {
                    const url = images[Math.floor(Math.random() * images.length)];

                    $("<div class='col mb-4'><div class='card h-100'><img src='" + url + "' class='card-img-top' alt='...' /><div class='card-body'>          <a href='#' class='card-link'>Like</a>        </div>      </div></div>").hide().appendTo(".row").fadeIn();

                }
   }
});