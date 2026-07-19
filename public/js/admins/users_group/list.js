import Alert from "../../components/alert.js"
import { handleAjaxError, ischeckboxcheck } from "/static/js/common/helpers.js"
import { checkBtnDatatable } from "/static/js/components/templates.js"

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
                return response.data;
            }
        },
        columns: [
            {
                render: function (data, type, row) {
                    return checkBtnDatatable(row.id)
                }
            },
            { data: 'custom' },
            { data: 'id' },
            { data: 'name' },
            { data: 'description' },
            { data: 'status' },


        ]
    });

    // Handle Save button click
    $('#save_user_group_btn').on('click', function () {
        // Collect data from the form
        var formData = {
            name: $('#groupName').val(),
            description: $('#groupDesc').val(),
            status: parseInt($('input[name="status"]:checked').val() || 1)
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

    // checkAll
    var checkAll = document.getElementById("checkAll");

    if (checkAll) {
        checkAll.onclick = function () {
            var checkboxes = document.querySelectorAll('.form-check-all input[type="checkbox"]');
            var checkedCount = document.querySelectorAll('.form-check-all input[type="checkbox"]:checked').length;

            for (var i = 0; i < checkboxes.length; i++) {
                checkboxes[i].checked = this.checked;

                if (checkboxes[i].checked) {
                    checkboxes[i].closest("tr").classList.add("table-active");
                } else {
                    checkboxes[i].closest("tr").classList.remove("table-active");
                }
            }
            document.getElementById("remove-actions").style.display = checkedCount > 0 ? "none" : "block";
        };
    }

    $("#user_group_table").on('click', function (evt) {
        ischeckboxcheck();

    });
});
