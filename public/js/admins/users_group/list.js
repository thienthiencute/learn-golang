import Alert from "../../components/alert.js"
import { handleAjaxError, ischeckboxcheck } from "/static/js/common/helpers.js"
import { checkBtnDatatable, statusTemplateWithDropdown, modalNotiUpdateStatus } from "/static/js/components/templates.js"

var user_group_dt;
var statusArray = [];
var statusTemplates = {};
var statusColors = {
    active: "badge-soft-success",
    inactive: "badge-soft-warning",
    deleted: "badge-soft-danger",
};

var listData = {};

const userGroupList = function () {
    const initial = function () {
        user_group_dt = $('#user_group_table').DataTable({
            serverSide: true,
            processing: true,
            ajax: {
                url: '/api/admins/roles/list',
                type: 'POST',
                dataType: "json",
                data : function(d) {
                    console.log(listData);
                    
                    d = { ...d, ...listData };
                    return d;
                },
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
                { 
                    data: null,
                    render: function (data, type, row, meta) {
                        return meta.row + 1;
                    }
                },
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

        $(document).on('input', '#search_box', function() {
            console.log("========================");
            console.log($("#search_box").val());
            listData.search_name = $("#search_box").val();
            user_group_dt.draw();

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
            console.log(123);
            
            
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

        // Mở Modal
        $(document).on('click', '.change_status', function (e) {
            e.preventDefault();
            let tr = $(this).closest('tr');
            if (tr.hasClass('child')) tr = tr.prev('.parent');
            let indexRow = user_group_dt.row(tr).index();
            let newStatusString = $(this).data('status');
            processModalNotiUpdateStatus(newStatusString, indexRow);
        });

        $(document).on("click", "#update_status", function () {
            var $status = $(this).data("status");
            var $rowIndex = $("#row_index").val();
            var $rowData = user_group_dt.row($rowIndex).data();
            if ($rowData.id == "") {
                $("#update_status_modal").modal("hide");
                Alert.error("Please select the row you want to change the status for");
                return;
            }


            var statusInt = 1;
            if ($status === 'inactive') statusInt = 2;
            if ($status === 'deleted') statusInt = 3;

            $.ajax({
                url: "/api/admins/roles/update-status",
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
                    user_group_dt.row($rowIndex).data($rowData);
                    user_group_dt.draw(false);
                    // getLanguage();
                },
                error: function (xhr) {
                    handleAjaxError(xhr);
                },
            });

            $("#update_status_modal").modal("hide");
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

function processModalNotiUpdateStatus(statusName, rowIndex) {
    var str = modalNotiUpdateStatus(statusName, rowIndex);
    // getLanguage();
    var $modifiedStr = $(str);
    $("#update_status_html").html($modifiedStr);
    $("#update_status_modal").modal("show");
}