// Example usage
function debounce(cb, wait = 400) {
  let timeout;
  return function (...args) {
    const context = this;
    clearTimeout(timeout);
    timeout = setTimeout(() => cb.apply(context, args), wait);
  };
}

const formatPercentCurrency = function (
  number,
  unit = null,
  minimumFractionDigits = 1,
  maximumFractionDigits = 2,
) {
  if (number === 0 || number === "0" || number === "") {
    return "-";
  }

  // Bước 1: Làm sạch chuỗi (nếu là string)
  if (typeof number === "string") {
    // Loại bỏ mọi ký tự không phải số, dấu trừ, hoặc dấu chấm
    number = number.replace(/[^0-9.-]+/g, "");
  }

  const num = parseFloat(number);
  if (isNaN(num)) return "-";

  // Làm tròn đến 2 chữ số thập phân
  let rounded = Math.round(num * 100) / 100;
  const isInteger = rounded % 1 === 0;

  let formatted;

  if (unit) {
    formatted = isInteger
      ? rounded.toLocaleString("en-US", { maximumFractionDigits: 0 })
      : rounded.toLocaleString("en-US", {
          minimumFractionDigits,
          maximumFractionDigits,
        });

    return formatted + " " + unit;
  }

  formatted = isInteger
    ? rounded.toLocaleString("en-US", { maximumFractionDigits: 0 })
    : rounded.toLocaleString("en-US", {
        minimumFractionDigits: 1,
        maximumFractionDigits: 2,
      });

  return formatted;
};

const setParamURL = (key, value) => {
  const url = new URL(window.location.href);
  const params = url.searchParams;
  params.set(key, value);
  window.history.replaceState(null, "", url);
};

const getParamURL = (key) => {
  const params = new URL(window.location.href).searchParams;
  const value = params.get(key);
  return value;
};

function removeParamURL(param) {
  const url = new URL(window.location.href);
  url.searchParams.delete(param);
  window.history.replaceState(null, "", url.toString());
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

// Format number currency input
function formatCurrency(value) {
  if (value === null || value === undefined || value === "") return "";

  const num = Number(value);

  if (isNaN(num)) return "";

  return new Intl.NumberFormat("en-US").format(num) + "₫";
}

// Gắn event cho tất cả input có class .number-currency-format
document.querySelectorAll(".number-currency-format").forEach((input) => {
  // Khi focus: chỉ hiển thị số thô
  input.addEventListener("focus", function () {
    let raw = this.value.replace(/\D/g, "");
    this.value = raw;
  });

  // Khi blur: định dạng lại
  input.addEventListener("blur", function () {
    let raw = this.value.replace(/\D/g, "");
    this.value = formatCurrency(raw === "" ? "" : Number(raw));
  });

  // Format sẵn ban đầu
  let initRaw = input.value.replace(/\D/g, "");
  input.value = formatCurrency(initRaw === "" ? "" : Number(initRaw));
});
