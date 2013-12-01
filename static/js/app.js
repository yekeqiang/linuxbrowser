var SESSIONMESSAGE = "SESSIONMESSAGE";

function mkdir(directory) {

    var dirname = prompt("Enter dirname ", "");//将输入的内容赋给变量 name ，

    if(dirname) {
        $.ajax({
            url: "/operation",
            type: "get",
            dataType: "json",
            data: {action:"mkdir", dirname: directory + dirname},
            success: function(json){
              if(json[0] == '0') {
                  $.cookie(SESSIONMESSAGE, "create directory success");
                  window.location.href = location.href
              } else {
                  $.scojs_message(json[1], $.scojs_message.TYPE_ERROR);
              }
            }
        });
     }

    return ;
}

function createFile(directory) {

    var filename = prompt("Enter filename ", "");//将输入的内容赋给变量 name ，

    if(filename) {
        $.ajax({
            url: "/operation",
            type: "get",
            dataType: "json",
            data: {action:"create", filename : directory + filename},
            success: function(json){
              if(json[0] == '0') {
                   $.cookie(SESSIONMESSAGE, "create file success");
                  window.location.href = location.href
              } else {
                  $.scojs_message(json[1], $.scojs_message.TYPE_ERROR);
              }
            }
        });
     }

    return ;
}

function rename(directory, filename) {

    var newname = prompt("Enter newname ", "");//将输入的内容赋给变量 name ，

    if(newname) {
        $.ajax({
            url: "/operation",
            type: "get",
            dataType: "json",
            data: {action:"rename", oldname: directory + filename,  newname: directory + newname},
            success: function(json){
              if(json[0] == '0') {
                  $.cookie(SESSIONMESSAGE, "rename success");
                  window.location.href = location.href
              } else {
                  $.scojs_message(json[1], $.scojs_message.TYPE_ERROR);
              }
            }
        });
     }

    return ;
}

function fileRemove(filename) {
    if (confirm("Are you sure delete '" + filename + "' ? ")) {  
        $.ajax({
            url: "/operation",
            type: "get",
            dataType: "json",
            data: {action:"remove", filename: filename},
            success: function(json){
              if(json[0] == '0') {
                   $.cookie(SESSIONMESSAGE, "remove file success");
                  window.location.href = location.href
              } else {
                  $.scojs_message(json[1], $.scojs_message.TYPE_ERROR);
              }
            }
        });
    }
    return ;
}

function dirRemove(filename) {
    if (confirm("Are you sure delete '" + filename + "' ? Warning ! This is a directory ")) {  
        $.ajax({
            url: "/operation",
            type: "get",
            dataType: "json",
            data: {action:"remove", filename: filename},
            success: function(json){
              if(json[0] == '0') {
                  $.cookie(SESSIONMESSAGE, "remove directory success");
                  window.location.href = location.href
              } else {
                  $.scojs_message(json[1], $.scojs_message.TYPE_ERROR);
              }
            }
        });
    }
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
