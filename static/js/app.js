var SESSIONMESSAGE = "SESSIONMESSAGE";

function readFile(file) {

	$.ajax({
		url: "/operation",
		type: "get",
		dataType: "json",
		data: {action : 'read' , file : file},
		success: function(json){
			if(json[0] == '0') {
				$("#readTitle").html(file);
				$("#readModal .modal-body").html("<pre>"+json[1]+"</pre>");
				$("#readModal").modal();
			} else {
				$.scojs_message(json[1], $.scojs_message.TYPE_ERROR);
			}
		}
	});
}

function mkdir(directory, param2) {

    if(param2 === undefined) {
        var input = "<input id=\"prompt\" type=\"text\" class=\"form-control\" placeholder=\"Enter dirname\" />";
        $(".modal-body").html(input);
        $("#confirmYes").attr("onclick","mkdir('" + directory + "', true);");
        $("#confirm").modal();
        return ;
    }
    app_ajax({action:'mkdir' , dirname : directory + $("#prompt").val()}, "create dir success ");
}

function createFile(directory, param2) {

    if(param2 === undefined) {
        var input = "<input id=\"prompt\" type=\"text\" class=\"form-control\" placeholder=\"Enter filename\" />";
        $(".modal-body").html(input);
        $("#confirmYes").attr("onclick","createFile('" + directory  + "', true);");
        $("#confirm").modal();
        return ;
    }
    app_ajax({action:'create' , filename : directory + $("#prompt").val()}, "create file success ");
}

function rename(directory, filename, param3) {

    if(param3 === undefined) {
        var input = "<input id=\"prompt\" type=\"text\" class=\"form-control\" placeholder=\"Enter filename\" />";
        $(".modal-body").html(input);
        $("#confirmYes").attr("onclick","rename('" + directory  + "','" + filename + "', true);");
        $("#confirm").modal();
        return ;
    }
    app_ajax({action: 'rename' , oldname: directory + filename, newname: directory + $("#prompt").val()}, "rename success ");
}

function app_ajax(data, message) {
	$.ajax({
		url: "/operation",
		type: "get",
		dataType: "json",
		data: data,
		success: function(json){
		  if(json[0] == '0') {
			  $.cookie(SESSIONMESSAGE, message);
			  window.location.href = location.href
		  } else {
			  $.scojs_message(json[1], $.scojs_message.TYPE_ERROR);
		  }
		}
	});
}

function dirRemove(filename) {
    var message = "Are you sure delete '" + filename + "' ? Warning ! This is a directory ";
    $(".modal-body").html(message);
    $("#confirmYes").attr("onclick","app_ajax({action:'remove',filename:'" + filename + "'}, 'remove directory success');");
    $("#confirm").modal();
    return ;
}

function fileRemove(filename) {
    var message = "Are you sure delete '" + filename + "' ? ";
    $(".modal-body").html(message);
    $("#confirmYes").attr("onclick","app_ajax({action:'remove', filename:'" + filename + "'}, 'remove directory success');");
    $("#confirm").modal();
    return ;
}

function back() {
    var url = window.location.href;
    if(url.substr(url.length - 1, url.length) == "/") {
      url = url.substr(0, url.length - 1);
    }
    var strs = url.split("/");
    if(strs.length > 3) {
      strs[strs.length - 1] = "";
      url = strs.join("/");
      window.location.href = url;  
    }
    return;
}

$(function() {
    var url = window.location.href;
    if(url.substr(url.length - 1, url.length) == "/") {
      url = url.substr(0, url.length - 1);
    }
    var strs = url.split("/");
    var str = "";
    for(var i = 3; i < strs.length; i++) {
        str += "/" + strs[i];
        $(".breadcrumb").append("<li><a href=\"" + str + "/\">" + strs[i] + "</a></li>");
    }
    if($.cookie(SESSIONMESSAGE)) {
      $.scojs_message($.cookie(SESSIONMESSAGE), $.scojs_message.TYPE_OK);
      $.cookie(SESSIONMESSAGE, "", { expires: -1 });
    }
});
