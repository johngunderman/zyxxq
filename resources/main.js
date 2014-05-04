"use strict";

window.onload = function() {
    $("#content-box-div")[0].focus();
    if (!supports_html5_storage()) {
        console.log("local storage not supported");
    } else {
        // TODO: pull this out into a separate function that
        // escapes all the html (except divs because we need
        // them for newlines.
        $("#content-box-div").html(localStorage["content"]);
        $("#content-box-div").keyup(function() {
            localStorage["content"] = $("#content-box-div").html();
        });
    }
    $("#submit-button").click(function() {
        $.post("/post",
               $("#content-box-div").html(),
               function(data, textStatus, jqXHR) {
                   //TODO: give the user their edit-key.
                   //localStorage["content"] = "";
                   location.reload();
               })});
}

function supports_html5_storage() {
  try {
    return 'localStorage' in window && window['localStorage'] !== null;
  } catch (e) {
    return false;
  }
}