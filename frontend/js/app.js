$("#btn_search").click(function() {
  $.ajax({
    type: "GET",
    url: "http://localhost:8080/api/persons?q=" + $("#txt_search").val() + "/",
    dataType: "json",
    success: function (result, status, xhr) {
      var $body = $("#tbl_persons tbody");
      $body.empty();
      $.each( result, function( index, value ){
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
      alert("Result: " + status + " " + error + " " + xhr.status + " " + xhr.statusText)
    }
  });
});