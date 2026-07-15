import Alert from "/static/js/components/alert.js";
import { getLanguage } from "/static/js/common/inits.js";
import { multiFilesOpt } from "/static/js/ecommerce/components/multi_files_opt.js";
import {
  renderAttachmentItem,
  getAttachmentType,
} from "/static/js/ecommerce/components/attachment_renderer.js";

const MAX_IMAGE_COUNT = 3;
const MAX_VIDEO_COUNT = 1;

const API_BASE_URL =
  document
    .querySelector('meta[name="api-base-url"]')
    ?.getAttribute("content") || "";
let lastName = "";
let firstName = "";
let image = "";
const pathDefault = "images/default/";

const replaceIdForHref = function (id, href) {
  let string = href.replace("/0/", `/${id}/`);

  return string.replace(/0$/, `${id}`);
};

const getCookie = function (name) {
  let cookieValue = document.cookie.replace(
    /(?:(?:^|.*;\s*)username\s*\=\s*([^;]*).*$)|^.*$/,
    "$1",
  );

  if (document.cookie && document.cookie !== "") {
    var cookies = document.cookie.split(";");
    for (var i = 0; i < cookies.length; i++) {
      var cookie = cookies[i].trim();
      if (cookie.substring(0, name.length + 1) === name + "=") {
        cookieValue = decodeURIComponent(cookie.substring(name.length + 1));
        break;
      }
    }
  }
  return cookieValue;
};

function errorAlert(msg) {
  Swal.fire({
    toast: true,
    position: "top-end",
    icon: "error",
    title: msg,
    showConfirmButton: !1,
    timer: 3500,
    showCloseButton: !0,
    timerProgressBar: true,
    didOpen: (toast) => {
      toast.onmouseenter = Swal.stopTimer;
      toast.onmouseleave = Swal.resumeTimer;
    },
  });
}

function forbiddenAlert() {
  Swal.fire({
    html:
      '<div class="mt-3"><lord-icon src="' +
      domain +
      '/static/themes/json/tdrtiskw.json" trigger="loop" colors="primary:#f06548,secondary:#f7b84b" style="width:120px;height:120px"></lord-icon><div class="mt-4 pt-2 fs-15"><h4>Oops...!</h4><p class="text-muted mx-4 mb-0">You do not have permission to perform this action</p></div></div>',
    showCancelButton: !0,
    showConfirmButton: !1,
    cancelButtonClass: "btn btn-primary w-xs mb-1",
    cancelButtonText: "Back",
    buttonsStyling: !1,
    showCloseButton: !0,
  });
}

function handleAjaxError(xhr) {
  if (xhr.status === 403) {
    forbiddenAlert();
    return;
  }
  const errResponse = JSON.parse(xhr.responseText);
  const errMessages = Array.isArray(errResponse.details.msg)
    ? errResponse.details.msg.join(" <br>")
    : errResponse.details.msg;
  errorAlert(errMessages);
}

function handleAjaxErrorWithFooter(xhr, footer) {
  if (xhr.status === 403) {
    forbiddenAlert();
    return;
  }
  const errResponse = JSON.parse(xhr.responseText);
  const errMessages = Array.isArray(errResponse.details.msg)
    ? errResponse.details.msg.join(" <br>")
    : errResponse.details.msg;
  Alert.errorWithFooter(errMessages, footer);
}

const numberFormat = function (
  field,
  delimiterSymbol = ".",
  decimalSymbol = ",",
) {
  if (typeof field !== "string") {
    throw new TypeError("Field selector must be a string.");
  }

  const cleaveNumber = document.querySelectorAll(field);
  cleaveNumber.forEach(function (field) {
    new Cleave(field, {
      numeral: true,
      numeralDecimalMark: decimalSymbol,
      delimiter: delimiterSymbol,
    });
  });
};

const numberFormatToDisplay = function (
  number,
  decimals = 0,
  thousandSeparator = ",",
  decimalSeparator = ".",
) {
  let integerPart = String(number);
  let decimalPart = "";
  if (decimals > 0) {
    const parts = number.toFixed(decimals).split(".");
    integerPart = parts[0];
    decimalPart = parts[1];
  }

  integerPart = integerPart.replace(/\B(?=(\d{3})+(?!\d))/g, thousandSeparator);
  if (decimals > 0) {
    return integerPart + decimalSeparator + decimalPart;
  }

  return integerPart;
};

const numberFormatToDisplayNumberPerCurrency = function (number, unit = null) {
  if (number === 0 || number === "0" || number === "") {
    return "-";
  }

  let formatter = new Intl.NumberFormat("en-US", {
    minimumFractionDigits: 1,
    maximumFractionDigits: 2,
  });

  let formatted = formatter.format(number);
  if (unit === "%") return formatted + "%";

  if (unit === "đ") return "₫" + formatted;
  return formatted;
};

const displayNumberOrHyphen = function (number) {
  if (number === 0 || number === "0" || number === "") {
    return "-";
  }

  return numberFormatToDisplay(number);
};

const formDataToJson = function (formData) {
  const formObject = {};
  formData.forEach((value, key) => {
    const keys = key.match(/[^\[\]]+/g);
    keys.reduce((accumulator, currentValue, index) => {
      if (index === keys.length - 1) {
        if (
          accumulator[currentValue] !== undefined &&
          Array.isArray(accumulator[currentValue])
        ) {
          accumulator[currentValue].push(value);
        } else if (accumulator[currentValue] !== undefined) {
          accumulator[currentValue] = [accumulator[currentValue], value];
        } else if (key.includes("[]")) {
          accumulator[currentValue] = [value];
        } else {
          accumulator[currentValue] = value;
        }
      } else {
        if (!accumulator[currentValue]) {
          accumulator[currentValue] = {};
        }

        return accumulator[currentValue];
      }

      return accumulator;
    }, formObject);
  });

  return formObject;
};

const convertStringToIntArray = (input) => {
  return input.map(function (value) {
    return parseInt(value);
  });
};

const convertToIdValueArray = (input) => {
  return input.map(function (value) {
    return { id: value };
  });
};

function toName(str) {
  str = str.replace(/\s+/g, " ");
  str = str.trim();
  return str;
}

// replace title or slug
function toSlug(str) {
  str = str.toLowerCase();
  str = str.replace(/--\|-/g, "");
  str = str.replace(/^\-+|\-+$/g, "");
  str = str.replace(/à|á|ạ|ả|ã|â|ầ|ấ|ậ|ẩ|ẫ|ă|ằ|ắ|ặ|ẳ|ẵ/g, "a");
  str = str.replace(/è|é|ẹ|ẻ|ẽ|ê|ề|ế|ệ|ể|ễ/g, "e");
  str = str.replace(/ì|í|ị|ỉ|ĩ/g, "i");
  str = str.replace(/ò|ó|ọ|ỏ|õ|ô|ồ|ố|ộ|ổ|ỗ|ơ|ờ|ớ|ợ|ở|ỡ/g, "o");
  str = str.replace(/ù|ú|ụ|ủ|ũ|ư|ừ|ứ|ự|ử|ữ/g, "u");
  str = str.replace(/ỳ|ý|ỵ|ỷ|ỹ/g, "y");
  str = str.replace(/đ/g, "d");
  str = str.replace(
    /\`|\~|\!|\@|\#|\||\$|\%|\^|\&|\*|\(|\)|\+|\=|\,|\.|\/|\?|\>|\<|\'|\"|\:|\;|_/gi,
    "-",
  );
  str = str.replace(/[^a-zA-Z0-9 ]/g, "");
  str = str.replace(/-+/g, "-");
  str = str.replace(/^\-+|\-+$/g, "");
  str = str.replace(/\s+/g, "-");
  str = "@" + str + "@";
  str = str.replace(/\@\-|\-\@|\@/gi, "");
  return str;
}

function toSlugSimple(str) {
  str = str.toLowerCase();
  str = str.replace(/à|á|ạ|ả|ã|â|ầ|ấ|ậ|ẩ|ẫ|ă|ằ|ắ|ặ|ẳ|ẵ/g, "a");
  str = str.replace(/è|é|ẹ|ẻ|ẽ|ê|ề|ế|ệ|ể|ễ/g, "e");
  str = str.replace(/ì|í|ị|ỉ|ĩ/g, "i");
  str = str.replace(/ò|ó|ọ|ỏ|õ|ô|ồ|ố|ộ|ổ|ỗ|ơ|ờ|ớ|ợ|ở|ỡ/g, "o");
  str = str.replace(/ù|ú|ụ|ủ|ũ|ư|ừ|ứ|ự|ử|ữ/g, "u");
  str = str.replace(/ỳ|ý|ỵ|ỷ|ỹ/g, "y");
  str = str.replace(/đ/g, "d");
  str = str.replace(
    /\`|\~|\!|\@|\#|\||\$|\%|\^|\&|\*|\(|\)|\+|\=|\,|\.|\/|\?|\>|\<|\'|\"|\:|\;|_/gi,
    "-",
  );
  str = str.replace(/[^a-zA-Z0-9\-]/g, "");
  str = str.replace(/-+/g, "-");
  str = str.replace(/\s+/g, "-");
  str = "@" + str + "@";
  str = str.replace(/\@\-|\-\@|\@/gi, "");

  return str;
}

function toSlugUppercase(str) {
  str = toSlug(str);
  str = str.toUpperCase();

  return str;
}

function loader(isBool = true) {
  var preloader = document.getElementById("preloader");

  if (isBool) {
    preloader.style.opacity = "0.4";
    preloader.style.visibility = "visible";
  } else {
    preloader.style.opacity = "0";
    preloader.style.visibility = "hidden";
  }
}

function showToolTip() {
  var tooltipTriggerList = [].slice.call(
    document.querySelectorAll('[data-bs-toggle="tooltip"]'),
  );
  tooltipTriggerList.map(function (tooltipTriggerEl) {
    var tooltip =
      bootstrap.Tooltip.getInstance(tooltipTriggerEl) ||
      new bootstrap.Tooltip(tooltipTriggerEl);

    tooltipTriggerEl.addEventListener("focusout", function () {
      tooltip.hide();
    });
    return tooltip;
  });
}

function showPopover() {
  var popoverTriggerList = [].slice.call(
    document.querySelectorAll('[data-bs-toggle="popover"]'),
  );
  popoverTriggerList.map(function (popoverTriggerEl) {
    var popover =
      bootstrap.Popover.getInstance(popoverTriggerEl) ||
      new bootstrap.Popover(popoverTriggerEl);

    popoverTriggerEl.addEventListener("focusout", function () {
      popover.hide();
    });

    return popover;
  });

  // Click outside to close
  document.addEventListener("click", function (e) {
    popoverTriggerList.forEach(function (popoverTriggerEl) {
      var popover = bootstrap.Popover.getInstance(popoverTriggerEl);

      if (popover) {
        const popoverElement = document.querySelector(".popover");

        const isClickInsideTrigger = popoverTriggerEl.contains(e.target);
        const isClickInsidePopover =
          popoverElement && popoverElement.contains(e.target);

        if (!isClickInsideTrigger && !isClickInsidePopover) {
          popover.hide();
        }
      }
    });
  });
}

function ischeckboxcheck() {
  Array.from(document.getElementsByName("chk_child")).forEach(function (a) {
    a.addEventListener("change", function (e) {
      1 == a.checked
        ? e.target.closest("tr").classList.add("table-active")
        : e.target.closest("tr").classList.remove("table-active");
      var t = document.querySelectorAll('[name="chk_child"]:checked').length;
      (e.target.closest("tr").classList.contains("table-active"),
        (document.getElementById("remove-actions").style.display =
          0 < t ? "block" : "none"));
    });
  });
  var checkboxesCount = document.querySelectorAll(
    '.form-check-all input[type="checkbox"]',
  ).length;
  var checkedCount = document.querySelectorAll(
    '.form-check-all input[type="checkbox"]:checked',
  ).length;
  checkboxesCount > 0 && checkboxesCount == checkedCount
    ? (checkAll.checked = true)
    : (checkAll.checked = false);
  document.getElementById("remove-actions").style.display =
    checkedCount > 0 ? "block" : "none";
}

function formatISODateTime(time) {
  const datePart = time.split("T")[0];
  const [year, month, day] = datePart.split("-");

  return `${day}-${month}-${year}`;
}

function getDateFromDateTime(dateTime) {
  if (dateTime === "" || dateTime.trim() === "") {
    return "";
  }

  const dateObj = new Date(dateTime);
  const day = String(dateObj.getDate()).padStart(2, "0");
  const month = String(dateObj.getMonth() + 1).padStart(2, "0");
  const year = dateObj.getFullYear();

  return `${day}-${month}-${year}`;
}

function saveAuth(user) {
  lastName = user.last_name;
  firstName = user.first_name;
  image = user.image;

  localStorage.setItem("first_name", firstName);
  localStorage.setItem("last_name", lastName);
  localStorage.setItem("image", image);
}

function showInfo() {
  lastName = localStorage.getItem("last_name") || "Super";
  firstName = localStorage.getItem("first_name") || "Admin";
  image = localStorage.getItem("image");
  $(".user-name-text").text(lastName + " " + firstName);

  if (image) {
    $(".header-profile-user").attr("src", API_BASE_URL + image);
  }
}

function selectFileWithCKFinder(chooseEle, showEle) {
  CKFinder.modal({
    chooseFiles: true,
    width: 800,
    height: 600,
    onInit: function (finder) {
      finder.on("files:choose", function (evt) {
        var file = evt.data.files.first();
        var url = file.getUrl();
        url = url.replace(pathDefault, "");
        chooseEle.val(url);
        showEle.attr("src", url);
      });
      finder.on("file:choose:resizedImage", function (evt) {
        var output = chooseEle;
        output.val(evt.data.resizedUrl);
      });
    },
  });
}

function selectMultiFilesWithCKFinder(positionEle, prefix_name = "") {
  CKFinder.modal({
    chooseFiles: true,
    width: 800,
    height: 600,
    onInit: function (finder) {
      finder.on("files:choose", function (evt) {
        var files = evt.data.files;
        files.each(function (file, index) {
          finder
            .request("command:send", {
              name: "ImageInfo",
              folder: file.get("folder"),
              params: { fileName: file.get("name") },
            })
            .done(function (response) {
              if (response.error) {
                return;
              }
              getLanguage();
              var url = file.getUrl();
              url = url.replace(pathDefault, "");
              positionEle.append(
                multiFilesOpt(
                  url,
                  file.attributes.name,
                  parseInt(file.attributes.size),
                  response.height,
                  response.width,
                  url,
                  "",
                  prefix_name,
                ),
              );
              if (positionEle.find(".btn_toggle_main_img").length === 1) {
                const lastItem = positionEle.find(".item_img").last();
                const mainBtn = lastItem.find(".btn_toggle_main_img");
                const isMainInput = lastItem.find(
                  `input[name="is_main${prefix_name ? "_" + prefix_name : ""}"]`,
                );

                mainBtn
                  .removeClass("btn-outline-primary")
                  .addClass("btn-primary");
                mainBtn.attr("data-is-main", "true");
                mainBtn
                  .find("span[data-key='t-main-image']")
                  .text("Main Image");

                isMainInput.val("true");
              }
            });
        });
      });
      finder.on("file:choose:resizedImage", function (evt) {
        var output = chooseEle;
        output.val(evt.data.resizedUrl);
      });
    },
  });
}

function selectMultiAttachmentWithCKFinder(positionEle) {
  CKFinder.modal({
    chooseFiles: true,
    width: 800,
    height: 600,
    onInit: function (finder) {
      finder.on("files:choose", function (evt) {
        var files = evt.data.files;

        files.each(function (file) {
          const name = file.get("name");
          const type = getAttachmentType(name);
          const size = parseInt(file.get("size"), 10);
          var url = file.getUrl().replace(pathDefault, "");

          const count = getCurrentAttachmentCount(positionEle);

          if (type === "image" && count.image >= MAX_IMAGE_COUNT) {
            Alert.error("Chỉ được chọn tối đa 3 ảnh");
            return;
          }

          if (type === "video") {
            if (count.video >= MAX_VIDEO_COUNT) {
              Alert.error("Chỉ được chọn 1 video");
              return;
            }
          }

          if (type === "image") {
            const img = new Image();
            img.onload = function () {
              positionEle.append(
                renderAttachmentItem({
                  url,
                  name,
                  size,
                  width: img.naturalWidth,
                  height: img.naturalHeight,
                }),
              );
            };
            img.src = url;
            return;
          }

          positionEle.append(
            renderAttachmentItem({
              url,
              name,
              size,
              width: 0,
              height: 0,
            }),
          );
        });
      });
    },
  });
}

function getCurrentAttachmentCount(container) {
  return {
    image: container.find('.item_media[data-type="image"]').length,
    video: container.find('.item_media[data-type="video"]').length,
  };
}

//max ordering for
function getMaxOrder(container, orderingClass) {
  var inputOrdering = container.find(orderingClass);
  var max = 0;
  inputOrdering.each(function (index, input) {
    var value = parseInt($(input).val());
    if (value > max) {
      max = value;
    }
  });
  return max + 1;
}

// Handle Enter key to trigger button clicks inside modals (e.g., edit, delete)
function handleEnterForModals(modalButtonList) {
  $(document).on("keydown", function (e) {
    if (e.key !== "Enter") return;

    const $target = $(e.target);

    // If the currently focused element is inside a <form>, skip handling
    // Let the browser handle Enter behavior for form submission
    if ($target.closest("form").length > 0) {
      return;
    }

    // Prevent default Enter behavior (e.g., submitting an implicit form)
    e.preventDefault();

    // Loop through all modals defined in the list
    for (const { modal, button } of modalButtonList) {
      if ($(modal).hasClass("show")) {
        $(button).click();
        break;
      }
    }
  });
}

// document.body.addEventListener("htmx:responseError", function (evt) {
//   const xhr = evt.detail.xhr;

//   if (xhr.status === 403) {
//     forbiddenAlert();
//     return;
//   }

//   try {
//     const errResponse = JSON.parse(xhr.responseText);
//     let errMessages = "";

//     if (Array.isArray(errResponse.details?.msg)) {
//       errMessages = errResponse.details.msg.join(" <br>");
//     } else if (typeof errResponse.details?.msg === "string") {
//       errMessages = errResponse.details.msg;
//     } else {
//       errMessages = errResponse.message || "An error occurred";
//     }

//     Alert.error(errMessages);
//   } catch (e) {
//     Alert.error(xhr.responseText || "An error occurred");
//   }
// });
// document.body.addEventListener("htmx:responseError", function(evt) {
//     const xhr = evt.detail.xhr;

//     let responseData;
//     try {
//         responseData = JSON.parse(xhr.responseText);
//     } catch {
//         responseData = { targetId: null, html: xhr.responseText };
//     }

//     if (responseData.targetId) {
//         const el = document.getElementById(responseData.targetId);
//         if (el) {
//             switch(responseData.swapMode) {
//                 case "outerHTML":
//                     el.outerHTML = responseData.html || "";
//                     break;
//                 case "innerHTML":
//                     el.innerHTML = responseData.html || "";
//                     break;
//                 case "beforeend":
//                     el.insertAdjacentHTML("beforeend", responseData.html || "");
//                     break;
//                 case "afterbegin":
//                     el.insertAdjacentHTML("afterbegin", responseData.html || "");
//                     break;
//                 case "displayNone":
//                   $(el).addClass('d-none');
//                   el.style.display = "none"
//                   break;
//                 default:
//                   el.innerHTML = responseData.html || "";
//             }
//         }
//     }

//     if (responseData.message) {
//         Alert.error(responseData.message);
//     }

//     if (responseData.triggerJS) {
//         document.dispatchEvent(new CustomEvent(responseData.triggerJS, { detail: responseData }));
//     }
// });
document.body.addEventListener("htmx:responseError", function (evt) {
  const xhr = evt.detail.xhr;
  let data;

  try {
    data = JSON.parse(xhr.responseText);
  } catch {
    return Alert.error(xhr.responseText || "Lỗi không xác định");
  }

  const msg = data?.details?.msg || data?.message;
  if (msg) Alert.error(msg);

  const actions = data?.details?.hxActions || [];
  actions.forEach((action) => {
    const el = document.getElementById(action.targetId);
    if (!el) return;
    switch (action.swapMode) {
      case "outerHTML":
        el.outerHTML = action.html || "";
        break;
      case "innerHTML":
        el.innerHTML = action.html || "";
        break;
      case "displayNone":
        $(el).addClass("d-none");
        el.style.display = "none";
        break;
      case "beforeend":
        el.insertAdjacentHTML("beforeend", action.html || "");
        break;
      case "afterbegin":
        el.insertAdjacentHTML("afterbegin", action.html || "");
        break;
    }

    if (action.value !== undefined) {
      el.value = action.value;
    }

    if (action.triggerJS) {
      document.dispatchEvent(
        new CustomEvent(action.triggerJS, { detail: action }),
      );
    }
  });
});

// Gắn CSRF token cho mọi request HTMX
document.body.addEventListener("htmx:configRequest", function (event) {
  const token = getCookie("csrf_");
  if (token) {
    event.detail.headers["X-CSRF-Token"] = token;
  }
});

/**
 * Đóng modal và reset form / dữ liệu về mặc định
 * modalSelector - The modal's ID string or the modal HTMLElement itself
 * defaultData - Optional object containing default values for inputs, If null, the form will be reset using `form.reset()
 */
function closeModalAndReset(modalSelector, defaultData = null) {
  const modalEl =
    typeof modalSelector === "string"
      ? document.getElementById(modalSelector)
      : modalSelector;
  if (!modalEl) return;

  if (defaultData) {
    Object.keys(defaultData).forEach((key) => {
      const el = modalEl.querySelector(`[name="${key}"]`);
      if (!el) return;

      if (el.type === "checkbox" || el.type === "radio") {
        el.checked = defaultData[key];
      } else {
        el.value = defaultData[key];
      }
    });
  } else {
    // Nếu không có defaultData → reset form mặc định
    const form = modalEl.querySelector("form");
    if (form) form.reset();
  }

  // Đóng modal
  const modalInstance = bootstrap.Modal.getInstance(modalEl);
  if (modalInstance) modalInstance.hide();
}

function formatVND(amount) {
  return new Intl.NumberFormat("vi-VN").format(amount) + " đ";
}


function initDateRangePicker(endOffset = { days: 1 }) {
  document.querySelectorAll(".date_range_picker").forEach(function (wrapper) {

    const startInput = wrapper.querySelector(".start_date");
    const endInput = wrapper.querySelector(".end_date");
    const toggle = wrapper.querySelector(".has_end_date");

    if (!startInput || !endInput || !toggle) return;

    const formatDate = (date) => {
      return (
        date.getFullYear() + "-" +
        String(date.getMonth() + 1).padStart(2, "0") + "-" +
        String(date.getDate()).padStart(2, "0") + " " +
        String(date.getHours()).padStart(2, "0") + ":" +
        String(date.getMinutes()).padStart(2, "0")
      );
    };

    const endPicker = flatpickr(endInput, {
      enableTime: true,
      dateFormat: "Y-m-d H:i",
      time_24hr: true,
    });

    const startPicker = flatpickr(startInput, {
      enableTime: true,
      dateFormat: "Y-m-d H:i",
      time_24hr: true,
      onChange: function (selectedDates) {
        if (!toggle.checked) return;

        if (selectedDates[0]) {
          endPicker.set("minDate", selectedDates[0]);
        }
      },
    });

    function handleToggle() {
      if (toggle.checked) {
        endInput.disabled = false;
        endInput.removeAttribute("readonly");
      } else {
        endInput.disabled = true;
        endInput.value = "";
        endPicker.clear();
      }
    }

    function addOffset(date, offset) {
      const result = new Date(date);

      if (offset.minutes) {
        result.setMinutes(result.getMinutes() + offset.minutes);
      }

      if (offset.hours) {
        result.setHours(result.getHours() + offset.hours);
      }

      if (offset.days) {
        result.setDate(result.getDate() + offset.days);
      }

      if (offset.months) {
        result.setMonth(result.getMonth() + offset.months);
      }

      if (offset.years) {
        result.setFullYear(result.getFullYear() + offset.years);
      }

      return result;
    }

    function setDefaultDates() {
      const now = new Date();
      const endDate = addOffset(now, endOffset);

      if (!startInput.value) {
        startPicker.setDate(now, false);
      }

      if (toggle.checked && !endInput.value) {
        endPicker.setDate(endDate, false);
      }
    }

    toggle.addEventListener("change", handleToggle);

    endInput.addEventListener("change", function () {
      if (!toggle.checked) return;

      const start = startPicker.selectedDates[0];
      const end = endPicker.selectedDates[0];

      if (start && end && end < start) {
        endPicker.clear();
      }
    });

    handleToggle();
    setDefaultDates();
  });
}

export {
  replaceIdForHref,
  getCookie,
  handleAjaxError,
  handleAjaxErrorWithFooter,
  numberFormat,
  numberFormatToDisplay,
  displayNumberOrHyphen,
  formDataToJson,
  convertStringToIntArray,
  convertToIdValueArray,
  toSlugUppercase,
  toName,
  toSlug,
  loader,
  showToolTip,
  showPopover,
  ischeckboxcheck,
  formatISODateTime,
  getDateFromDateTime,
  numberFormatToDisplayNumberPerCurrency,
  saveAuth,
  showInfo,
  selectFileWithCKFinder,
  selectMultiFilesWithCKFinder,
  selectMultiAttachmentWithCKFinder,
  getMaxOrder,
  handleEnterForModals,
  closeModalAndReset,
  formatVND,
  initDateRangePicker,
};
