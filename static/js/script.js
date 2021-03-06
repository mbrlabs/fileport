// Copyright (c) 2017. Marcus Brummer.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distriuted on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

$(function() {

// Configuration
// ------------------------------------------------------------------------

var SIDEBAR_MIN = 300;
var SIDEBAR_MAX = 3600;
var MAIN_MIN = 200;

var API_ENDPOINT = "http://" + IP + ":3000/api/";

// FileSystem
// ------------------------------------------------------------------------
Fs = {
    files: [],
};


// FileUtils
// ------------------------------------------------------------------------
FileUtils = {
    isImage: function(file) {
        return file.type == 1;
    },

    isVideo: function(file) {
        return file.type == 2;
    },

    isAudio: function(file) {
        return file.type == 3;
    },

    isText: function(file) {
        return file.type == 4;
    },

    isPDF: function(file) {
        return file.type == 5;
    },

    isFolder: function(file) {
        return file.type == 6;
    },

    getIcon: function(file) {
        if(this.isImage(file)) {
            return '<i class="icon-img fa fa-picture-o" aria-hidden="true"></i>';
        } else if(this.isVideo(file)) {
            return '<i class="icon-vid fa fa-video-camera" aria-hidden="true"></i>';
        } else if(this.isAudio(file)) {
            return '<i class="icon-audio fa fa-music" aria-hidden="true"></i>';
        } else if(this.isText(file)) {
            return '<i class="icon-txt fa fa-font" aria-hidden="true"></i>';
        } else if(this.isPDF(file)) {
            return '<i class="icon-pdf fa fa-file-pdf-o" aria-hidden="true"></i>';
        } else if(this.isFolder(file)) {
            return '<i class="icon-dir fa fa-folder" aria-hidden="true"></i>';
        }

        return '<i class="fa fa-file" aria-hidden="true"></i>';
    },
};

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
        $.getJSON(API_ENDPOINT + "list/" + path, function(data) {
            Fs.files = data;
            onSuccess(Fs.files);
        }).fail(onError);
    },

};

// Ui
// ------------------------------------------------------------------------
var Ui = {
    videoPlayer: plyr.setup('#video-player')[0],
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

    setFiles: function(files) {
        var html = "<table><thead><tr><th>NAME</th><th>SIZE</th></tr></thead><tbody>";
        $.each(files, function(index, file) {
            var currentPath = History.current + "/" + file.name; 

            var size = filesize(file.size);
            var type = FileUtils.getIcon(file) + "&nbsp;&nbsp;&nbsp;&nbsp;";
            
            var escapedName = file.name.replace("\"", "&#34;");
            escapedName = escapedName.replace("'", "&#39;");
            var link = encodeURI(API_ENDPOINT + "get/" + History.current + "/" + escapedName); 
            
            if(FileUtils.isFolder(file)) {
                html += "<tr class='folder-link' data-path='"+currentPath+"' >";
            } else if(FileUtils.isImage(file)) {
                html += "<tr data-fancybox='gallery' href='"+link+"' >";
            } else if(FileUtils.isVideo(file)) {
                html += "<tr class='video-link' video-url='"+link+"' >";
            } else {
               // name += "<a target='_blank' href='" + link + "'>" + file.name + "</a>";
                html += "<tr class='other-link' other-url='"+link+"' >";
            }

            html += "<td>" + type + file.name + "</td>";
            html += "<td>" + size + "</td>";
            html += "</tr>";
        });
        html += "</tbody></table>";

        $("#files").html(html);
        Ui.resetSidebar();
    },

    loadPath: function(path) {
        Api.listFiles(path, function(data) {
            History.pushPath(path);
            Ui.setCurrentPath(History.current);
            Ui.setFiles(data);
        }, function() {
            console.log("Failed to fetch files");
        });
    },

    showVideo: function(url) {
        Ui.videoPlayer.source({
          type:       'video',
          sources: [{src: url}]});

        $.fancybox.open({
            src  : '#video-player-modal',
            type : 'inline',
            focus : false,
            afterClose: function() {
                Ui.videoPlayer.stop();
            },
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

$("#search-input").keyup(function() {
    var term = $(this).val(); 

    if(term.length >= 2) {
        var options = {
            shouldSort: true,
            threshold: 0.6,
            location: 0,
            distance: 100,
            maxPatternLength: 32,
            minMatchCharLength: 1,
            keys: ["name"]
        };
        var fuse = new Fuse(Fs.files, options);
        Ui.setFiles(fuse.search(term));
    } else {
        Ui.setFiles(Fs.files);
    }
});

// File events
// ------------------------------------------------------------------------
$(document).on("click", ".folder-link", function(e) {
    e.preventDefault();
    var path = $(this).attr("data-path");
    Ui.loadPath(path);
});

$(document).on("click", ".video-link", function(e) {
    e.preventDefault();
    Ui.showVideo($(this).attr("video-url"));
});

$(document).on("click", ".other-link", function(e) {
    e.preventDefault();
   // Ui.showVideo($(this).attr("other-url"));
   window.location = $(this).attr("other-url");
});

// Initial Setup
// ------------------------------------------------------------------------
Ui.adjustSidebar(350);
Ui.loadPath(HOME_DIR);

});
