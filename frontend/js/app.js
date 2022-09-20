$("#btn_search").click(function () {
  var url = "http://localhost:8080/api/persons";
  var search = $("#txt_search").val();
  if (search != "") {
    url += "?q=" + encodeURIComponent(search);
  }

  $.ajax({
    type: "GET",
    url: url,
    dataType: "json",
    success: function (result, status, xhr) {
      var $body = $("#tbl_persons tbody");
      $body.empty();
      $.each(result, function (index, value) {
        var rows = $("<tr data-id='" + value["id"] + "'>");
        rows.append("<td>" + value["name"] + "</td>");
        rows.append("<td>" + value["email"] + "</td>");
        rows.append("<td>" + value["role"] + "</td>");
        rows.append("<td>" + value["created_at"] + "</td>");
        rows.append("<td>" + value["updated_at"] + "</td>");
        rows.append("</tr>");
        $body.append(rows);
      });
    },
    error: function (xhr, status, error) {
      alert("Error!");
    }
  });
});