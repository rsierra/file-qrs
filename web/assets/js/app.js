$(document).ready(function () {
  var $view = $("<div id='qr-code'><div class='popup'><div class='content'></div></div></div>");
  $("body").append($view);
  $view.hide();
});

$("body").click(function (e) {
  var target_selector = ".qr-link";
  var QRVisibe = $(e.target).is(target_selector) || $(e.target).parents(target_selector).length > 0

  // if there is click event outside IMG then close the qr-view box
  if (!QRVisibe)
    $('#qr-code').fadeOut();
});

$("a.qr-link").click(function (e) {
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
