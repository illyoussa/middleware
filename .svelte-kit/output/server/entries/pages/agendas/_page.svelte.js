import { Z as ensure_array_like, _ as store_get, X as escape_html, $ as unsubscribe_stores, T as pop, R as push } from "../../../chunks/index.js";
import { a as attr, S as SvelteMap, b as agendas, c as agendasError, d as deleteAgenda } from "../../../chunks/index-server.js";
import { D as Delete_modal } from "../../../chunks/delete-modal.js";
function _page($$payload, $$props) {
  push();
  var $$store_subs;
  let newUcaId = "";
  let newUcaName = "";
  let editable = new SvelteMap();
  function deleteAgendaById(id) {
    console.log(`Deleting agenda ${id}`);
    deleteAgenda(id).then(() => {
      editable.set(id, false);
    }).catch(() => {
    });
  }
  const each_array = ensure_array_like(store_get($$store_subs ??= {}, "$agendas", agendas));
  $$payload.out += `<div class="d-flex flex-column"><h1 class="text-center color-yellow mb-lg-5">Agendas</h1> <div class="d-flex flex-column mb-5"><h5 class="text-center fw-semibold mb-2">Add a agenda :</h5> <form class="d-flex justify-content-center"><div class="d-flex me-4"><label for="id" class="fs-5 align-self-center me-2">UCA ID</label> <input id="id" class="form-control-sm" placeholder="56529"${attr("value", newUcaId)}></div> <div class="d-flex me-4"><label for="name" class="fs-5 align-self-center me-2">Name</label> <input id="name" class="form-control-sm" placeholder="M1 - Tutorat L2"${attr("value", newUcaName)}></div> <input class="btn btn-yellow" type="submit" value="Submit"></form></div> `;
  if (store_get($$store_subs ??= {}, "$agendasError", agendasError)) {
    $$payload.out += "<!--[-->";
    $$payload.out += `<div class="text-center text-danger form-text fs-5">An error occurred : ${escape_html(store_get($$store_subs ??= {}, "$agendasError", agendasError))}</div>`;
  } else {
    $$payload.out += "<!--[!-->";
  }
  $$payload.out += `<!--]--> <table class="table w-50 align-self-center"><thead><tr><th class="color-yellow" scope="col">Name</th><th class="color-yellow" scope="col">UCA ID</th><th class="color-yellow" scope="col">ID</th><th scope="col">Actions</th></tr></thead><tbody><!--[-->`;
  for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
    let agenda = each_array[$$index];
    $$payload.out += `<tr><td>`;
    if (!editable.get(agenda.id)) {
      $$payload.out += "<!--[-->";
      $$payload.out += `${escape_html(agenda.name)}`;
    } else {
      $$payload.out += "<!--[!-->";
      $$payload.out += `<input class="form-control-sm"${attr("value", agenda.name)}>`;
    }
    $$payload.out += `<!--]--></td><td>`;
    if (!editable.get(agenda.id)) {
      $$payload.out += "<!--[-->";
      $$payload.out += `${escape_html(agenda.ucaId)}`;
    } else {
      $$payload.out += "<!--[!-->";
      $$payload.out += `<input class="form-control-sm"${attr("value", agenda.ucaId)}>`;
    }
    $$payload.out += `<!--]--></td><td>${escape_html(agenda.id)}</td><td class="justify-content-around">`;
    if (!editable.get(agenda.id)) {
      $$payload.out += "<!--[-->";
      $$payload.out += `<button class="bg-transparent border-0"><span class="fa fa-pen-to-square text-warning"></span></button> <button class="bg-transparent border-0"><span class="fa fa-trash-can text-danger"></span></button>`;
    } else {
      $$payload.out += "<!--[!-->";
      $$payload.out += `<button class="bg-transparent border-0"><span class="fa fa-check" style="color: #0b6f33"></span></button>`;
    }
    $$payload.out += `<!--]--></td></tr>`;
  }
  $$payload.out += `<!--]--></tbody></table> `;
  Delete_modal($$payload, {
    uniqueName: "agendas",
    deleteFunction: deleteAgendaById
  });
  $$payload.out += `<!----></div>`;
  if ($$store_subs) unsubscribe_stores($$store_subs);
  pop();
}
export {
  _page as default
};
