const Alert = {
    error: function (message) {
        Swal.fire({
            toast: true,
            position: "top-end",
            icon: "error",
            title: message,
            showConfirmButton: !1,
            timer: 3500,
            showCloseButton: !0,
            timerProgressBar: true,
            didOpen: (toast) => {
                toast.onmouseenter = Swal.stopTimer;
                toast.onmouseleave = Swal.resumeTimer;
            },
        });
    },
    errorWithFooter: function (message, footer) {
        Swal.fire({
            icon: "error",
            position: "top-end",
            title: message,
            showCloseButton: !0,
            footer: footer,
          });
    },
    success: function (message) {
        Swal.fire({
            position: "top-end",
            icon: "success",
            title: message,
            showConfirmButton: !1,
            timer: 1500,
            showCloseButton: !0,
        });
    },
    warning: function (message) {
        Swal.fire({
            position: "top-end",
            icon: "warning",
            title: message,
            showConfirmButton: !1,
            timer: 1500,
            showCloseButton: !0,
        });
    },
    successCenter: function (message = "Success!") {
        Swal.fire({
            position: "center",
            icon: "success",
            title: message,
            showConfirmButton: false,
            timer: 1000,
            showCloseButton: true,
        });
    },
    successNoTimer: function (message = "Success!") {
        Swal.fire({
            position: "top-end",
            icon: "success",
            title: message,
            showConfirmButton: true,
            showCloseButton: true,
        });
    },
    confirm: function (
        target,
        callback = {},
        title = "Are you sure?",
        confirmButtonText = "Yes",
        cancelButtonText = "No",
        onCancel = {}
    ) {
        Swal.fire({
            title,
            icon: "warning",
            showCancelButton: true,
            reverseButtons: true,
            confirmButtonColor: "#3085d6",
            cancelButtonColor: "#d33",
            cancelButtonText,
            confirmButtonText,
            reverseButtons: true,
        }).then((result) => {
            if (result.isConfirmed) {
                callback(target);
            } else {
                onCancel(target);
            }
        });
    },
};

export default Alert;
