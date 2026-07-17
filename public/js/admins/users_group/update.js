import Alert from "../../components/alert.js";

$(document).ready(function () {
    // Fetch data on load
    var groupId = $('#groupId').val();
    if (groupId) {
        $.ajax({
            url: '/api/admins/users-group/detail/' + groupId,
            type: 'GET',
            success: function (response) {
                if (response && response.data) {
                    $('#groupName').val(response.data.name);
                    $('#groupDesc').val(response.data.description);
                    $('#groupStatus').prop('checked', response.data.status === 1);
                }
            },
            error: function () {
                alert('Failed to load user group details');
            }
        });
    }

    // Handle Save button click
    $('#update_user_group_btn').on('click', function () {
        // Collect data from the form
        var formData = {
            id: parseInt($('#groupId').val(), 10),
            name: $('#groupName').val(),
            status: $('#groupStatus').is(':checked') ? 1 : 2,
            description: $('#groupDesc').val()
        };

        // Validate basic required fields
        if (!formData.name) {
            Alert.error('Name is required!');
            return;
        }

        // Send AJAX POST request to update API
        $.ajax({
            url: '/api/admins/users-group/update',
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(formData),
            success: function (response) {
                Alert.success('User group updated successfully');
                setTimeout(() => {
                    location.href = '/admins/users-group/list';
                }, 1500);
            },
            error: function (xhr, status, error) {
                var errorMessage = 'An error occurred while updating user group';
                if (xhr.responseJSON && xhr.responseJSON.error) {
                    errorMessage = xhr.responseJSON.error;
                }
                Alert.error(errorMessage);
            }
        });
    });
});
