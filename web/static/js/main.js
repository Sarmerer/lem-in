$(document).ready(function () {
  console.log("ban");
  $("#form").submit(function (e) {
    e.preventDefault(); // avoid to execute the actual submit of the form.
    var form = $(this);
    var url = form.attr("action");
    $.ajax({
      type: "POST",
      url: "/upload",
      data: form.serialize(), // serializes the form's elements.
      success: function (data) {
        alert(data); // show response from the php script.
      },
    });
  });
});
