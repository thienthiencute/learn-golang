import Alert from "../../components/alert.js"
import { handleAjaxError, ischeckboxcheck } from "/static/js/common/helpers.js"
import { checkBtnDatatable, statusTemplateWithDropdown } from "/static/js/components/templates.js"

var user_group_dt;
var statusArray = [];
var statusTemplates = {};
var statusColors = {
    active: "badge-soft-success",
    inactive: "badge-soft-warning",
    deleted: "badge-soft-danger",
};

const userGroupList = function () {
    const initial = function () {
        user_group_dt = $('#user_group_table').DataTable({
            ajax: {
                url: '/api/admins/roles/list',
                type: 'POST',
                dataSrc: function (response) {
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
                {
                    data: 'status',
                    render: function (data) {
                        return statusTemplates[data] || "";
                    }
                },
            ]
        });

        // Handle Save button click
        $('#save_user_group_btn').on('click', function () {
            var formData = {
                name: $('#groupName').val(),
                description: $('#groupDesc').val(),
                status: parseInt($('input[name="status"]:checked').val() || 1)
            };

            if (!formData.name) {
                Alert.error('Name is required!');
                return;
            }

            $.ajax({
                url: '/api/admins/roles/create',
                type: 'POST',
                contentType: 'application/json',
                data: JSON.stringify(formData),
                success: function (response) {
                    $('#create_user_group_modal').modal('hide');
                    $('#create_user_group_form')[0].reset();
                    Alert.success('User group created successfully');
                    user_group_dt.ajax.reload();
                },
                error: function (xhr) {
                    handleAjaxError(xhr)
                }
            });
        });

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
    };

    return {
        init: function () {
            try {
                statusArray = JSON.parse($("#status_list").val() || "[]");
            } catch (e) {
                console.error("Error parsing status data:", e);
            }
            initStatusTemplates();
            initial();
        }
    };
}();

document.addEventListener("DOMContentLoaded", function () {
    userGroupList.init();
});

function initStatusTemplates() {
    statusArray.forEach((value) => {
        const otherStatuses = statusArray.filter((item) => item !== value);
        statusTemplates[value] = statusTemplateWithDropdown(
            otherStatuses,
            value,
            statusColors[value] || "badge-soft-secondary"
        );
    });
}