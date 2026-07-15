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
});
