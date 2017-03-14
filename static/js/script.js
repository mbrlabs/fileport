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
    }
};


// Initial Setup
// ------------------------------------------------------------------------
Ui.adjustSidebar(350);

// Sidebar resizing events
// ------------------------------------------------------------------------
$('#split-bar').mousedown(function (e) {
    e.preventDefault();
    $(document).mousemove(function (e) {
        e.preventDefault();
        Ui.adjustSidebar(e.pageX);
    })
});
$(document).mouseup(function (e) {
    $(document).unbind('mousemove');
});

});
