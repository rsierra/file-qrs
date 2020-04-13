$(document).ready(function () {
  console.log('La magia aquí');
  // var $view = $('<div id="qr-code">');
  // $("body").append($view);
  // $view.hide();

  $('<div id="qr-code">').appendTo('body').hide();
});

$("body").click(function (e) {
  var target_selector = ".qr-link";
  var QRVisibe = $(e.target).is(target_selector) || $(e.target).parents(target_selector).length > 0;

  // if there is click event outside IMG then close the qr-view box
  if (!QRVisibe)
    $('#qr-code').fadeOut();
});

$(document).on('click', 'a.qr-link', function (e) {
  // prevent default click behaviour
  e.preventDefault();

  // get full link url
  var href = $(this).prop('href');

  // set its location and do show
  $("#qr-code").css("top", (e.pageY) + "px").css("left", (e.pageX) + "px");
  $("#qr-code").fadeIn();

  // set qr-code content
  $("#qr-code").empty().qrcode(
    {
      text: encodeURI(decodeURI(href)),
      correctLevel: 1, // QRErrorCorrectLevel.L,
      width: 300,
      height: 300
    }
  );

  var data = {
    'key_1': 'value 1',
    'key_2': 'value_2',
  };

  var card_template = $('#card-template').html();

  for (var key in data) {
    // console.log(key+':'+data[key]);
    // CAUTION: \w "word" character (letter+digit+underscore)
    card_template = card_template.replace(/\{\{(\w+)\}\}/g /* {{key}} */, function (str, key) {
      return data[key];
    });
  } // for

  $("#card-template-content").remove();
  var $card_template = $(card_template);
  $card_template.appendTo('body');

});
