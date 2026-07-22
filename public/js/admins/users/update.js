import Alert from "../../components/alert.js";
import { handleAjaxError } from "/static/js/common/helpers.js"

$(document).ready(function () {
    $('#update_user_form').on('submit', function (e) {
        e.preventDefault();
        
        var formData = {
            id: parseInt($('#userId').val(), 10),
            email: $('#email').val(),
            first_name: $('#firstName').val(),
            last_name: $('#lastName').val(),
            status: parseInt($('input[name="status"]:checked').val(), 10) || 2,
        };

        if (!formData.email || !formData.first_name || !formData.last_name) {
            Alert.error('Email, Tên và Họ không được để trống!');
            return;
        }

        $.ajax({
            url: '/api/admins/users/update',
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(formData),
            success: function (response) {
                Alert.success('User updated successfully');
                setTimeout(() => {
                    location.href = '/admins/users/list';
                }, 1500);
            },
            error: function (xhr) {
                handleAjaxError(xhr)
            }
        });
    });
});
