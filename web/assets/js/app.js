// Show QR code

$(document).on('click', 'a.qr-link', function (e) {
  console.log(e);
  // prevent default click behaviour
  e.preventDefault();

  // get full link url
  var href = $(this).prop('href');

  // set its location and do show
  $("#qr-code").fadeIn();

  // set qr-code content
  $("#qr-code .content").empty().qrcode(
    {
      text: encodeURI(decodeURI(href)),
      correctLevel: 1, // QRErrorCorrectLevel.L,
      width: 300,
      height: 300
    }
  );
});

// Hide QR code

$(document).on('click', 'body', function (e) {
  var target_selector = ".qr-link";
  var QRVisibe = $(e.target).is(target_selector) || $(e.target).parents(target_selector).length

  // if there is click event outside IMG then close the qr-view box
  if (!QRVisibe)
    $('#qr-code').fadeOut();
});
