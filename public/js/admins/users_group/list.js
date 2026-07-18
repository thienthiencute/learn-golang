import Alert from "../../components/alert.js"
import {handleAjaxError} from "/static/js/common/helpers.js"

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
            description: $('#groupDesc').val(),
            status: $('#groupStatus').is(':checked') ? 1 : 2
        };

        // Validate basic required fields
        if (!formData.name) {
            Alert.error('Name is required!');
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
                Alert.success('User group created successfully');
                $('#user_group_table').DataTable().ajax.reload();
            },
            error: function (xhr) {
                handleAjaxError(xhr)
            }
        });
    });
});
