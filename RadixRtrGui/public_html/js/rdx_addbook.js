var init_add_book_window = function () {
//show add dialog box
    $('#rdx_addBook').click(function () {
        $('#rdxAddBookDbox').show();
    });
//close add dialog box
    $('#rdxAddBookDboxClose').click(function () {
        $('#rdxAddBookDbox').hide();
    });
    //add book clicked
    $('#rdxAb_addButton').click(function () {

        var rdx_data_packet = {
            Id: RADIX_REQUEST_ADD_BOOK,
            Title: $('#rdxAb_title').val(),
            Authur: $('#rdxAb_auther').val(),
            Publisher: $('#rdxAb_publisher').val(),
            Popub: $('#rdxAb_pofpublish').val(),
        };


        $.ajax({
            type: "POST",
            contentType: 'application/json; charset=utf-8',
            dataType: 'json',
            url: "/login",
            data: JSON.stringify(rdx_data_packet),
            success: function (result) {
                console.log(result);
                $('#rdxAddBookDbox').hide();
            },
            error: function (xhr, status, error) {
                // error handling
                alert(error);
            }
        });
    });
};

