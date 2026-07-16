$(document).ready(function () {
    $('#user_group_table').DataTable({
        ajax: {
            url: '/api/admins/users-group/list',
            type: 'POST',
            dataSrc: function (response) {
                console.log(response)
                if (!response.data) {
                    return [];
                }

                // var $firstRowData = response.data[0];
                // detailUser($firstRowData);
                return response.data;
            }
        },
        columns: [
            { data: 'id' },
            { data: 'name' },
            { data: 'description' },
            { data: 'status' },
            { data: 'custom' },
            // {
            //     render: function (data, type, row) {
            //         return `
            //             <button class="btn btn-sm btn-info edit-btn" data-id="${row.id}">Edit</button>
            //             <button class="btn btn-sm btn-danger delete-btn" data-id="${row.id}">Delete</button>
            //         `;
            //     }
            // }
        ]
    });

    // Handle Save button click
    $('#save_user_group_btn').on('click', function () {
        // Collect data from the form
        var formData = {
            name: $('#groupName').val(),
            description: $('#groupDesc').val()
        };

        // Validate basic required fields
        if (!formData.name) {
            alert('Name is required!');
            return;
        }

        // Send AJAX POST request to create API
        $.ajax({
            url: '/api/admins/users-group/create',
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(formData),
            success: function (response) {
                // Hide modal and reset form
                $('#create_user_group_modal').modal('hide');
                $('#create_user_group_form')[0].reset();

                // Show success message
                if (typeof Swal !== 'undefined') {
                    Swal.fire('Success', 'User group created successfully', 'success');
                } else {
                    alert('User group created successfully');
                }

                // Reload datatable
                $('#user_group_table').DataTable().ajax.reload();
            },
            error: function (xhr, status, error) {
                var errorMessage = 'An error occurred while creating user group';
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
