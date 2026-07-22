import { getCookie, handleAjaxError, loader } from "/static/js/common/helpers.js";
import Alert from "../../components/alert.js";

var $listID = {};
var $idsArr = [];

$(document).ready(function () {
    var isDeleteUser = 0;

    $(document).on("click", "#delete-item-btn", function (evt) {
        isDeleteUser = 0;
        var tr = $(this).closest("tr");
        if (tr.hasClass('child')) tr = tr.prev('.parent');
        var $rowData = $('#user_table').DataTable().row(tr).data();
        var $dataID = $rowData.id;
        $idsArr = [$dataID.toString()];
        $listID = {};
        $listID[$dataID] = 1;
    });

    $("#delete_modal").on("shown.bs.modal", function (e) {
        isDeleteUser = 0;
    });

    $(document).on("click", "#delete_record", function (evt) {
        evt.preventDefault();
        if (isDeleteUser == 1) {
            return;
        }

        isDeleteUser = 1;

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
                url: "/api/admins/users/delete",
                method: "DELETE",
                dataType: "json",
                contentType: "application/json",
                data: JSON.stringify({ ids: $idsArr.map(id => parseInt(id, 10)) }),
                headers: {
                    "X-CSRF-Token": getCookie("csrf_"),
                },
                success: function (res) {
                    Alert.success('Deleted successfully');
                    $("#delete_modal").modal("hide");
                    var $res = [];
                    var allRows = $('#user_table').DataTable().rows().data().toArray();
                    allRows.forEach(function ($allRow) {
                        if (!$listID.hasOwnProperty($allRow.id.toString())) {
                            $res.push($allRow);
                        }
                    });
                    $('#user_table').DataTable().rows().remove().draw();
                    $('#user_table').DataTable().rows.add($res).draw();
                    $("#remove-actions").hide();
                    
                    var checkAll = document.getElementById("checkAll");
                    if (checkAll) {
                        checkAll.checked = false;
                    }

                    $idsArr.forEach(function ($row) {
                        var $detailID = $("#info_status").data("id");
                        if ($detailID == $row) {
                            var data = $('#user_table').DataTable().row(0).data();
                            if (typeof detailUser !== 'undefined') detailUser(data);
                        }
                    });
                },
                error: function (xhr) {
                    handleAjaxError(xhr);
                    $("#delete_modal").modal("hide");
                },
            });

            loader(false);
            isDeleteUser = 1;
        } else {
            Alert.error('Please select at least one item to delete.');
            loader(false);
            isDeleteUser = 0;
            $("#delete_modal").modal("hide");
        }
    });
});
