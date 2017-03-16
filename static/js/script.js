// Copyright (c) 2017. Marcus Brummer.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

$(function() {

// Configuration
// ------------------------------------------------------------------------

var SIDEBAR_MIN = 300;
var SIDEBAR_MAX = 3600;
var MAIN_MIN = 200;

var API_ENDPOINT = "http://localhost:3000/api/";

// History
// ------------------------------------------------------------------------
History = {
    current: HOME_DIR,
    history: [],
    pushPath: function(path) {
        this.current = path;
        History.history.push(path);
    },

    popPath: function() {
        return history.pop();
    },
};

// Api
// ------------------------------------------------------------------------
var Api = {
    listFiles: function(path, onSuccess, onError) {
        $.getJSON(API_ENDPOINT + "list/" + path, onSuccess).fail(onError);
    },

};

// Ui
// ------------------------------------------------------------------------
var Ui = {
    adjustSidebar: function(width) {
        var x = width - $('#sidebar').offset().left;
        if (x > SIDEBAR_MIN && x < SIDEBAR_MAX && width < ($(window).width() - MAIN_MIN)) {  
            $('#sidebar').css("width", x);
            $('#main').css("width", $(window).width() - x - 6);
            $("#split-bar").css("margin-left", x);
        }
    },

    resetSidebar: function() {
        this.adjustSidebar($("#split-bar").css("margin-left").replace("px", ""));
    },

    setCurrentPath: function(path) {
        var html = '<a class="path-segment" href="#" data-path="/">/</a>' ;
        var currentPath = "/"
        $.each(path.split("/"), function(index, value) {
            if(value.length > 0) {
                currentPath += value + "/";
                html += '<a class="path-segment" href="#" data-path="'+currentPath+'">' + value + '</a>/'
            }
        });

        $("#current-path").html(html);
    },

    loadPath: function(path) {
        Api.listFiles(path, function(data) {
            History.pushPath(path);
            Ui.setCurrentPath(History.current);

            var html = "";
            $.each(data, function(index, file) {
                var currentPath = path + "/" + file.name; 
                if(file.type == 5) {
                    html += "<div><a href='#' class='folder-link' data-path='"+currentPath+"'>" + file.name + "</a></div>";
                } else {
                    html += "<div>" + file.name + " </div>";
                }
            });
            $("#files").html(html);
            
            Ui.resetSidebar();
        }, function() {
            console.log("Failed to fetch files");
        });
    },
};

// Sidebar resizing events
// ------------------------------------------------------------------------
$('#split-bar').mousedown(function(e) {
    e.preventDefault();
    $(document).mousemove(function(e) {
        e.preventDefault();
        Ui.adjustSidebar(e.pageX);
        console.log(e.pageX);
    })
});

$(document).mouseup(function (e) {
    $(document).unbind('mousemove');
});

$(window).resize(function(e) {
    Ui.resetSidebar();
});

// Header events
// ------------------------------------------------------------------------
$(document).on("click", ".path-segment", function(e) {
    e.preventDefault();
    var path = $(this).attr("data-path");
    Ui.loadPath(path);
});

// Folder events
// ------------------------------------------------------------------------
$(document).on("click", ".folder-link", function(e) {
    e.preventDefault();
    var path = $(this).attr("data-path");
    Ui.loadPath(path);
});

// Initial Setup
// ------------------------------------------------------------------------
Ui.adjustSidebar(350);
Ui.loadPath(HOME_DIR);

});
