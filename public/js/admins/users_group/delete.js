import { getCookie, handleAjaxError, loader } from "/static/js/common/helpers.js";
import Alert from "../../components/alert.js";


// Biến lưu trữ ID của các items đang được chọn
var $listID = {};
// Mảng chứa ID để gửi lên server
var $idsArr = [];

$(document).ready(function () {
    /**
     * Start delete user group  
     */
    var isDeleteGroup = 0;

    // Bắt sự kiện click vào nút Xóa (thùng rác) trên TỪNG DÒNG của Datatable
    $(document).on("click", "#delete-item-btn", function (evt) {
        isDeleteGroup = 0;
        var tr = $(this).closest("tr");
        var $rowData = $('#user_group_table').DataTable().row(tr).data();
        var $dataID = $rowData.id;
        $idsArr = [$dataID.toString()];
        $listID = {}; // Clear existing items
        $listID[$dataID] = 1;
    });

    $("#delete_modal").on("shown.bs.modal", function (e) {
        isDeleteGroup = 0;
    });

    // Delete Roles
    $(document).on("click", "#delete_record", function (evt) {
        evt.preventDefault();
        if (isDeleteGroup == 1) {
            return;
        }

        isDeleteGroup = 1;

        var checked = document.querySelectorAll(
            '.form-check-all input[type="checkbox"]:checked'
        );
        if (checked.length > 0) {
            $idsArr.length = 0;
            checked.forEach(function (check) {
                $idsArr.push(check.value);
                $listID[check.value] = 1;
            });
        }
        loader();
        if ($idsArr.length > 0) {
            $.ajax({
                url: $("#remove-actions").data("url"),
                method: "DELETE",
                dataType: "json",
                contentType: "application/json",
                data: JSON.stringify({ ids: $idsArr }),
                enctype: $(this).attr("enctype") || "multipart/form-data",
                headers: {
                    "X-CSRF-Token": getCookie("csrf_"),
                },
                success: function (res) {

                    Alert.success(res.data.msg);
                    $("#delete_modal").modal("hide");
                    var $res = [];
                    var allRows = $('#user_group_table').DataTable().rows().data().toArray();
                    allRows.forEach(function ($allRow) {
                        if (!$listID.hasOwnProperty($allRow.id.toString())) {
                            $res.push($allRow);
                        }
                    });
                    $('#user_group_table').DataTable().rows().remove().draw();
                    $('#user_group_table').DataTable().rows.add($res).draw();
                    $("#remove-actions").hide();
                    checkAll.checked = false;

                    $idsArr.forEach(function ($row) {
                        var $detailID = $("#info_status").data("id");
                        if ($detailID == $row) {
                            var data = $('#user_group_table').DataTable().row(0).data();
                            if (typeof detailGroup !== 'undefined') detailGroup(data);
                        }
                    });
                },
                error: function (xhr) {
                    handleAjaxError(xhr);
                    $("#delete_modal").modal("hide");
                },
            });

            loader(false);
            isDeleteGroup = 1;
        }
    });
});
