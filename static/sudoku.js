/*global $*/
$(function() {
    "use strict";

    $('button').click(function(e) {
        e.preventDefault();

        /* Assemble the board into a data structure. */
        var board = [];
        $.each($('[data-x][data-y]'), function(idx, elem) {
            var $elem = $(elem);
            var x = $elem.data('x');
            var y = $elem.data('y');
            var val = $elem.val();

            if (val === '') {
                val = null;
            } else {
                val = parseInt(val, 10);
            }

            board.push(val);
        });

        /* Make the request. */
        $.post('solve', {
            board: JSON.stringify(board)
        }, function(json) {
            /* On success, fill out the empty spots and mark them. */
            $.each(json, function(i, elem) {
                var x = i % 9;
                var y = Math.floor(i / 9);
                var input = $('[data-x=' + x + '][data-y=' + y + ']');
                if (input.val() === '' && elem !== undefined) {
                    input.val(elem);
                    input.addClass('non-bold');
                }
                console.log("I: " + i + " X: " + x + " Y:" + y + " :: " + elem);
            });
        });
    });
});