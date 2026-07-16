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
                    $('#groupStatus').val(response.data.status);
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
            status: parseInt($('#groupStatus').val(), 10),
            description: $('#groupDesc').val()
        };

        // Validate basic required fields
        if (!formData.name) {
            alert('Name is required!');
            return;
        }

        // Send AJAX POST request to update API
        $.ajax({
            url: '/api/admins/users-group/update',
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(formData),
            success: function (response) {
                // Show success message
                if (typeof Swal !== 'undefined') {
                    Swal.fire({
                        title: 'Success',
                        text: 'User group updated successfully',
                        icon: 'success',
                        confirmButtonText: 'OK'
                    }).then((result) => {
                        if (result.isConfirmed) {
                            window.location.href = '/admins/users-group/list';
                        }
                    });
                } else {
                    alert('User group updated successfully');
                    window.location.href = '/admins/users-group/list';
                }
            },
            error: function (xhr, status, error) {
                var errorMessage = 'An error occurred while updating user group';
                if (xhr.responseJSON && xhr.responseJSON.error) {
                    errorMessage = xhr.responseJSON.error;
                }
                if (typeof Swal !== 'undefined') {
                    Swal.fire('Error', errorMessage, 'error');
                } else {
                    alert(errorMessage);
                }
            }
        });
    });
});
