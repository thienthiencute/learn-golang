import Alert from "../../components/alert.js";
import { handleAjaxError } from "/static/js/common/helpers.js"

$(document).ready(function () {
    // Handle Save button click
    $('#update_user_group_btn').on('click', function () {
        // Collect data from the form
        var formData = {
            id: parseInt($('#groupId').val(), 10),
            name: $('#groupName').val(),
            status: parseInt($('input[name="status"]:checked').val(), 10) || 2,
            description: $('#groupDesc').val()
        };

        // Validate basic required fields
        if (!formData.name) {
            Alert.error('Name is required!');
            return;
        }

        // Send AJAX POST request to update API
        $.ajax({
            url: '/api/admins/roles/update',
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(formData),
            success: function (response) {
                Alert.success('User group updated successfully');
                setTimeout(() => {
                    location.href = '/admins/roles/list';
                }, 1500);
            },
            error: function (xhr) {
                handleAjaxError(xhr)
            }
        });
    });
});
