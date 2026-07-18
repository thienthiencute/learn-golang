import Alert from "../../components/alert.js"
import { handleAjaxError } from "/static/js/common/helpers.js"

$(document).ready(function () {
    // Handle delete button
    $('#user_group_table').on('click', '.delete-btn', function () {
        var id = $(this).data('id');
        $.ajax({
            url: '/api/admins/users-group/delete',
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify({ id: id }),
            success: function (response) {
                Alert.success('Xóa nhóm người dùng thành công');
                $('#user_group_table').DataTable().ajax.reload();
            },
            error: function (xhr) {
                handleAjaxError(xhr);
            }
        });
    });
});