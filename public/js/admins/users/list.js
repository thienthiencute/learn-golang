import Alert from "../../components/alert.js"
import { handleAjaxError, ischeckboxcheck } from "/static/js/common/helpers.js"
import { checkBtnDatatable, statusTemplateWithDropdown, modalNotiUpdateStatus } from "/static/js/components/templates.js"

var user_dt;
var statusArray = [];
var statusTemplates = {};
var statusColors = {
  active: "badge-soft-success",
  inactive: "badge-soft-warning",
  deleted: "badge-soft-danger",
};

const userList = function () {
    const initial = function () {
        user_dt = $('#user_table').DataTable({
            ajax: {
                url: '/api/admins/users/list',
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
                { data: 'full_name_html' },
                { data: 'role_html' },
                {
                    data: 'status',
                    render: function (data) {
                        return statusTemplates[data] || "";
                    }
                },
                { data: 'audit_info_html' }
            ]
        });

        // Delete action
        $('#remove-actions').on('click', function (e) {
            e.preventDefault();
            const checkedIds = [];
            $('.form-check-input:checked').not('#checkAll').each(function () {
                checkedIds.push($(this).val());
            });

            if (checkedIds.length === 0) {
                return;
            }

            const url = $(this).data('url');

            $('#delete-record').off('click').on('click', function () {
                $.ajax({
                    url: url,
                    type: 'POST',
                    contentType: 'application/json',
                    data: JSON.stringify({ ids: checkedIds.map(id => parseInt(id, 10)) }),
                    success: function (response) {
                        Alert.success('Users deleted successfully');
                        $('#delete_modal').modal('hide');
                        user_dt.ajax.reload();
                        $('#checkAll').prop('checked', false);
                    },
                    error: function (xhr) {
                        handleAjaxError(xhr);
                        $('#delete_modal').modal('hide');
                    }
                });
            });
        });

        $('#checkAll').on('change', function () {
            $('.form-check-input').prop('checked', $(this).prop('checked'));
            ischeckboxcheck();
        });

        $("#user_table").on('click', function (evt) {
            ischeckboxcheck();
        });

        // Mở Modal
        $(document).on('click', '.change_status', function (e) {
            e.preventDefault();
            let tr = $(this).closest('tr');
            if (tr.hasClass('child')) tr = tr.prev('.parent');
            let indexRow = user_dt.row(tr).index();
            let newStatusString = $(this).data('status');
            processModalNotiUpdateStatus(newStatusString, indexRow);
        });

        $(document).on("click", "#update_status", function () {
            var $status = $(this).data("status");
            var $rowIndex = $("#row_index").val();
            var $rowData = user_dt.row($rowIndex).data();
            if ($rowData.id == "") {
                $("#update_status_modal").modal("hide");
                Alert.error("Please select the row you want to change the status for");
                return;
            }

            var statusInt = 1;
            if ($status === 'inactive') statusInt = 2;
            if ($status === 'deleted') statusInt = 3;

            $.ajax({
                url: "/api/admins/users/update-status",
                method: "PATCH",
                dataType: "json",
                contentType: "application/json",
                data: JSON.stringify({ 
                    id: parseInt($rowData.id), 
                    status: statusInt 
                }),
                success: function (res) {
                    Alert.success("Thành công");
                    $rowData.status = $status;
                    user_dt.row($rowIndex).data($rowData);
                    user_dt.draw(false);
                },
                error: function (xhr) {
                    handleAjaxError(xhr);
                },
            });

            $("#update_status_modal").modal("hide");
        });
    };

    const initStatusTemplates = function () {
        statusArray.forEach((value) => {
            const otherStatuses = statusArray.filter((item) => item !== value);
            statusTemplates[value] = statusTemplateWithDropdown(
                otherStatuses,
                value,
                statusColors[value] || "badge-soft-secondary"
            );
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

document.addEventListener("DOMContentLoaded", function (event) {
    userList.init();
});

function processModalNotiUpdateStatus(statusName, rowIndex) {
    var str = modalNotiUpdateStatus(statusName, rowIndex);
    var $modifiedStr = $(str);
    $("#update_status_html").html($modifiedStr);
    $("#update_status_modal").modal("show");
}