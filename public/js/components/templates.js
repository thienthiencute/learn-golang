const BASE_URL = window.location.origin;

//path: doraemon/submissions
function actionBar(id, path) {
  return `
        <ul class="list-inline hstack gap-2 mb-0 d-flex">   
            <li class="list-inline-item item_edit" data-bs-toggle="tooltip" data-bs-trigger="hover" data-bs-placement="top" aria-label="Edit" data-bs-original-title="Edit">
                <a class="edit-item-btn" href="${BASE_URL}/${path}/edit/${id}">
                    <i class="ri-pencil-fill align-bottom"></i>
                </a>
            </li>    
            <li class="list-inline-item item_delete" data-bs-toggle="tooltip" data-bs-trigger="hover" data-bs-placement="top" aria-label="Delete" data-bs-original-title="Delete">
                <a id="delete-item-btn" href="#delete_modal" class="text-danger" data-bs-toggle="modal">
                    <i class="ri-delete-bin-5-fill align-bottom"></i>
                </a>
            </li>   
        </ul>
    `;
}

function checkBtnDatatable(id) {
  return `<div class="form-check-all">
        <input class="form-check-input" type="checkbox" name="chk_child" value=${id}>
    </div>`;
}

function statusTemplateNoDropdown(status, color, isEvent) {
  const eventClass = isEvent ? "change_status cursor-pointer" : "";
  var dataStatus;
  if (status == "active") {
    dataStatus = "inactive";
  } else {
    dataStatus = "active";
  }

  return `<span class="badge ${color} text-uppercase ${eventClass}" data-status="${dataStatus}" data-key="t-${status}-status">${status}</span>`;
}

function statusTemplateWithDropdown(otherStatuses, status, color) {
  let dropdownItems = Object.entries(otherStatuses)
    .map(([val, label]) => {
      return `<li><a class="dropdown-item change_status text-capitalize" href="javascript:void(0);" data-status="${label}" data-key="t-${label}-status">${label}</a></li>`;
    })
    .join("");

  return `
        <div class="dropdown cursor-pointer">
            <span class="badge ${color} text-uppercase dropdown-toggle" data-status="${status}" data-key="t-${status}-status" data-bs-toggle="dropdown" aria-expanded="false">${status}</span>
            <ul class="dropdown-menu">${dropdownItems}</ul>
        </div>
    `;
}

function modalNotiUpdateStatus(status, indexRow) {
  return `
        <div class="modal fade zoomIn" id="update_status_modal" tabindex="-1" aria-hidden="true">
            <div class="modal-dialog modal-dialog-centered">
                <div class="modal-content">
                    <div class="modal-header">
                        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close" id="btn-close"></button>
                    </div>
                    <div class="modal-body p-5 text-center">
                        <div class="swal2-icon swal2-warning swal2-icon-show" style="display: flex;"><div class="swal2-icon-content">!</div></div>                    <div class="mt-4 text-center">
                            <h4 class="fs-semibold" data-key="t-change-status">Change status</h4>
                            <p class="text-muted fs-14 mb-4 pt-1"><span data-key="t-noti-change-status">Are you sure you want to change the status to </span><b class="text-capitalize" data-key="t-${status}">${status}</b>?</p>
                            <div class="hstack gap-2 justify-content-center remove">
                                <input type="hidden" name="delete_id" id="delete_ids" value="">
                                <button class="btn btn-link link-success fw-medium text-decoration-none" id="deleteRecord-close" data-bs-dismiss="modal">
                                    <i class="ri-close-line me-1 align-middle"></i>
                                    <span data-key="t-cancel">Cancel</span>
                                </button>
                                <button type="submit" class="btn btn-danger" id="update_status" data-status="${status}" data-key="t-ok">Ok</button>
                                <input type="hidden" id="row_index" value="${indexRow}">
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    `;
}

function renderAudit(createdBy, createdAt, updatedBy, updatedAt) {
  const created =
    createdBy || createdAt
      ? `<strong data-key="t-created">Created</strong>: ${createdBy || ""}${
          createdBy && createdAt ? " - " : ""
        }${createdAt || ""}`
      : "";
  const updated =
    updatedBy || updatedAt
      ? `</br><strong data-key="t-updated">Updated</strong>: ${
          updatedBy || ""
        }${updatedBy && updatedAt ? " - " : ""}${updatedAt || ""}`
      : "";
  return created + updated;
}
function renderAuditUpdate(updatedBy, updatedAt) {
  if (!updatedBy && !updatedAt) return "";

  return `
    <div class="audit-info"">
      <div><strong data-key="t-updated">Updated:</strong></div>
      ${updatedBy ? `<div>${updatedBy}</div>` : ""}
      ${updatedAt ? `<div>${updatedAt}</div>` : ""}
    </div>
  `;
}

function renderInfo(img, name, path, sku, countVariant) {
  const variantKey = countVariant > 1 ? "t-variants" : "t-variant";
  let skuHtml = "";
  if (sku) {
    skuHtml = `<label class="fs-12 text-muted mb-0">SKU: <span class="variant_sku fw-normal">${sku}</span></label>`;
  }
  return `
    <div class="d-flex gap-2">
      <div style="height: 80px; object-fit: cover; aspect-ratio: 3/4;">
        <img 
          src="${img}" 
          class="rounded w-100 h-100 product_image" 
          loading="lazy"
          style="cursor: zoom-in"
          data-bs-toggle="popover"
          data-bs-html="true"
          data-bs-content='<img src="${img}" class="w-100" />'
          >
      </div>

      <div class="audit-info">
        <div style="font-size: 16px;" class="mb-1">
          <h5 class="fs-16 mb-1">
              <a href="#" class="text-secondary">${name || ""}</a>
          </h5>
        </div>
        <p class="text-muted fs-12 mb-2">${path || ""}</p>
        ${skuHtml}
        <div class="d-flex align-items-center justify-content-center gap-1 badge rounded-pill border border-warning text-warning mb-0" style="width: fit-content;">
          <span class="fs-8">${countVariant}</span>
          <span data-key="${variantKey}">Biến thể</span>
        </div>
      </div>
    </div>
  `;
}

function renderOrderNumber(row, link = "/ecommerce/orders/edit/") {
  const isVerified =
    row.confirmed_status === "confirmed" || row.has_fulfillments;

  return `
    <div class="d-flex align-items-center gap-1">
      <a class="${
        row.cancelled_status == "cancelled" ? "text-danger" : ""
      }" href="${link}${row.id}">
        ${row.order_number}
      </a>

      ${
        isVerified
          ? `
            <div>
              <i class="ri-check-line me-1 align-bottom text-success fs-5 fw-bold"
              data-bs-toggle="tooltip"
              data-bs-placement="top"
              data-key="t-ready-confirm"
              title="Đã xác thực"></i>
            </div>
          `
          : ""
      }

      ${
        row.is_delivery_ready
          ? `
            <i
              class="ri-truck-line text-success fs-5"
              data-bs-toggle="tooltip"
              data-bs-placement="top"
              data-key="t-ready-delivery-info"
              title="Đủ thông tin để tạo giao hàng"
            ></i>
          `
          : `
            <i
              class="ri-truck-line text-warning fs-5"
              data-bs-toggle="tooltip"
              data-bs-placement="top"
              data-key="t-missing-delivery-info"
              title="Thiếu thông tin giao hàng"
            ></i>
          `
      }
    </div>
  `;
}

function renderNameCustomer(name) {
  return `
    <span class="order-customer cursor-pointer" title="${name}">${name}</span>
  `;
}

function renderCategoriesSync(categories, isCatePan, prdID) {
  const hasCategories = categories && categories.length > 0;

  const categoryHtml = hasCategories
    ? categories
        .map(function (category) {
          return `
            <a href="#" 
               class="fw-medium link-secondary category_item" 
               data-cate="${category.id}">
              ${category.name}
            </a>
          `;
        })
        .join(", ")
    : `<span class="text-muted">Không có danh mục</span>`;

  const checkBtn = hasCategories
    ? isCatePan
      ? `<span class="badge bg-success">✔</span>`
      : `<span class="badge bg-danger">✖</span>`
    : "";

  const actionBtn =
    !hasCategories || !isCatePan
      ? `<button class="btn btn-sm btn-outline-danger btn-sync-category"
          data-bs-toggle="modal" data-bs-target="#syncCatesPancakeModal"
          hx-post="/api/synchronization/pancake/products/sync-cates"
          hx-target="#selectCatesPancake"
          hx-swap="outerHTML"
          hx-trigger="click delay:300ms"
          hx-vals="js:{&quot;id&quot;:&quot;${prdID}&quot;}">Sync</button>`
      : "";

  return `
    <div class="d-flex flex-column">
      <div>${categoryHtml}</div>
    </div>
  `;
}

export {
  actionBar,
  checkBtnDatatable,
  statusTemplateNoDropdown,
  statusTemplateWithDropdown,
  modalNotiUpdateStatus,
  renderAudit,
  renderAuditUpdate,
  renderInfo,
  renderOrderNumber,
  renderNameCustomer,
  renderCategoriesSync,
};
