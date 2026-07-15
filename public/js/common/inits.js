function getLanguage() {
  var language = localStorage.getItem("language");
  var request = new XMLHttpRequest();
  request.open("GET", "/static/themes/lang/" + language + ".json");
  request.onreadystatechange = function () {
    if (this.readyState === 4 && this.status === 200) {
      var data = JSON.parse(this.responseText);
      Object.keys(data).forEach(function (key) {
        var elements = document.querySelectorAll("[data-key='" + key + "']");
        Array.from(elements).forEach(function (elem) {
          if (elem.hasAttribute("data-bs-original-title")) {
            elem.setAttribute("data-bs-original-title", data[key]);
            return;
          }

          if (["INPUT", "TEXTAREA"].includes(elem.tagName)) {
            elem.setAttribute("placeholder", data[key]);
            return;
          }

          for (let i = 0; i < elem.childNodes.length; i++) {
            const node = elem.childNodes[i];
            if (node.nodeType === Node.TEXT_NODE) {
              node.textContent = data[key];
              break; // only change language first node text
            }
          }
        });
      });
    }
  };
  request.send();
}
function updatePagePath(slugId, pagePathId) {
  const slugInput = document.getElementById(slugId);
  const pagePath = document.getElementById(pagePathId);

  const pattern = pagePath.getAttribute("data-pattern");
  const slugValue = slugInput.value.trim();
  if (slugValue == "") {
    return (pagePath.textContent = pattern);
  }
  pagePath.textContent = pattern.replace("{slug}", slugValue);
}

function updateSlugIdPagePath(slugId, pagePathId, id) {
  const slugInput = document.getElementById(slugId);
  const pagePath = document.getElementById(pagePathId);
  if (!slugInput || !pagePath) return;

  const pattern = pagePath.getAttribute("data-pattern") || "";
  const slugValue = slugInput.value.trim();

  if (slugValue) {
    let result = pattern.replace("{slug}", slugValue);
    if (id) {
      result = result.replace("{id}", id);
    }
    pagePath.textContent = result;
  } else {
    pagePath.textContent = pattern;
  }
}

function updateURLParam(key, value) {
    const url = new URL(window.location);  
    if (value === "all" || value === undefined || value === null || value === ""|| (key =="page" && value == 1)) {
        url.searchParams.delete(key);
    } else {
        url.searchParams.set(key, encodeURIComponent(value));
    }

    window.history.pushState({}, "", url);
}

export { getLanguage, updatePagePath, updateSlugIdPagePath, updateURLParam };
