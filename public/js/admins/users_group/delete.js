import Alert from "../../components/alert.js"
import { handleAjaxError } from "/static/js/common/helpers.js"

$(document).ready(function () {
    // Handle delete button
    $('#user_group_table').on('click', '.delete-btn', function () {
        var id = $(this).data('id');
        $('#delete_ids').val(id);
        $('#deleteRecordModal').modal('show');
    });

    $('#delete_record').on('click', function() {
        var id = $('#delete_ids').val();
        if (!id) return;

        $.ajax({
            url: '/api/admins/users-group/delete',
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify({ id: parseInt(id) }),
            success: function (response) {
                $('#deleteRecordModal').modal('hide');
                Alert.success('Xóa nhóm người dùng thành công');
                $('#user_group_table').DataTable().ajax.reload();
            },
            error: function (xhr) {
                $('#deleteRecordModal').modal('hide');
                handleAjaxError(xhr);
            }
        });
    });
});