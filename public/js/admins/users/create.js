import Alert from "../../components/alert.js"
import { handleAjaxError } from "/static/js/common/helpers.js"

document.addEventListener("DOMContentLoaded", function () {
    $('#save_user_btn').on('click', function () {
        var formData = {
            email: $('#userEmail').val(),
            first_name: $('#userFirstName').val(),
            last_name: $('#userLastName').val(),
            role_id: parseInt($('#userRole').val()),
            status: parseInt($('input[name="status"]:checked').val() || 1)
        };

        if (!formData.email || !formData.first_name || !formData.last_name) {
            Alert.error('Email, Tên và Họ là bắt buộc!');
            return;
        }

        $.ajax({
            url: '/api/admins/users/create',
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(formData),
            success: function (response) {
                $('#create_user_modal').modal('hide');
                $('#create_user_form')[0].reset();
                Alert.success('Thêm người dùng thành công');
                setTimeout(() => {
                    location.reload();
                }, 1500);
            },
            error: function (xhr) {
                handleAjaxError(xhr)
            }
        });
    });
});
