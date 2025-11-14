import { X as escape_html, a0 as bind_props, T as pop, R as push } from "./index.js";
import { a as attr } from "./index-server.js";
function Delete_modal($$payload, $$props) {
  push();
  let { uniqueName, deleteFunction } = $$props;
  let modalName = `deleteModal-${uniqueName}`;
  function show(id, name) {
    if (deleteModal === void 0) {
      deleteModal = new bootstrap.Modal(`#${modalName}`);
    }
    itemName = name;
    deleteModal.show();
  }
  let itemName = "";
  let deleteModal = void 0;
  $$payload.out += `<div class="modal fade modal-lg"${attr("id", modalName)} aria-hidden="true" data-bs-backdrop="static"><div class="modal-dialog"><div class="modal-content"><div class="modal-header"><h1 class="modal-title fs-5">Deletion confirmation</h1> <button type="button" class="btn-close" data-bs-dismiss="modal"></button></div> <div class="modal-body d-flex align-items-start justify-content-between flex-md-column">Are you sure you want to delete ${escape_html(itemName)} ?</div> <div class="modal-footer"><button type="button" class="btn btn-secondary rounded-0" data-bs-dismiss="modal">Cancel</button> <button type="button" class="btn btn-danger delete-button rounded-0 bg-danger">Delete</button></div></div></div></div>`;
  bind_props($$props, { show });
  pop();
}
export {
  Delete_modal as D
};
