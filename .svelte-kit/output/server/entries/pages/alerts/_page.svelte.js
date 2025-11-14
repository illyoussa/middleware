import { Z as ensure_array_like, _ as store_get, X as escape_html, $ as unsubscribe_stores, T as pop, R as push } from "../../../chunks/index.js";
import { w as writable } from "../../../chunks/index2.js";
import axios from "axios";
import { D as Delete_modal } from "../../../chunks/delete-modal.js";
import { g as getAgenda, a as attr, S as SvelteMap } from "../../../chunks/index-server.js";
const alerts = writable([]);
const alertsError = writable("");
const apiBaseUrl = "/config_api/alerts";
function getAlerts() {
  alertsError.set("");
  axios.get(`${apiBaseUrl}`).then((res) => {
    alerts.set(res.data);
  }).catch((err) => {
    console.log("An error has occurred while retrieving alerts");
    console.log(err);
    if (err.response?.data?.message) {
      alertsError.set(JSON.stringify(err.response.data.message));
    } else {
      alertsError.set(JSON.stringify(err));
    }
  });
}
function deleteAlert(id) {
  alertsError.set("");
  return axios.delete(`${apiBaseUrl}/${id}`).then((res) => {
    getAlerts();
    return Promise.resolve();
  }).catch((err) => {
    console.log("An error as occurred while deleting alert");
    console.log(err);
    if (err.response?.data?.message) {
      alertsError.set(JSON.stringify(err.response.data.message));
    } else {
      alertsError.set(JSON.stringify(err));
    }
    return Promise.reject(err);
  });
}
function _page($$payload, $$props) {
  push();
  var $$store_subs;
  let newEmail = "";
  let newAgendaId = "";
  let editable = new SvelteMap();
  let alertAgenda = new SvelteMap();
  alerts.subscribe((value) => {
    value.forEach((a) => {
      if (!alertAgenda.get(a.agendaId)) {
        getAgenda(a.agendaId).then((data) => {
          alertAgenda.set(a.agendaId, data);
        }).catch(() => {
        });
      }
    });
  });
  function deleteAlertById(id) {
    console.log(`Deleting alert ${id}`);
    deleteAlert(id).then(() => {
      editable.set(id, false);
    }).catch(() => {
    });
  }
  const each_array = ensure_array_like(store_get($$store_subs ??= {}, "$alerts", alerts));
  $$payload.out += `<div class="d-flex flex-column"><h1 class="text-center color-yellow mb-lg-5">Alerts</h1> <div class="d-flex flex-column mb-5"><h5 class="text-center fw-semibold mb-2">Add an alert :</h5> <form class="d-flex justify-content-center"><div class="d-flex me-4"><label for="email" class="fs-5 align-self-center me-2">UCA Email</label> <input id="email" class="form-control-sm" placeholder="justine.bachelard@ext.uca.fr"${attr("value", newEmail)}></div> <div class="d-flex me-4"><label for="agenda" class="fs-5 align-self-center me-2">Agenda ID</label> <input id="agenda" class="form-control-sm" placeholder="6ce750e5-05b8-4c4c-8fee-d9e381dbf364 "${attr("value", newAgendaId)}></div> <input class="btn btn-yellow" type="submit" value="Submit"></form></div> `;
  if (store_get($$store_subs ??= {}, "$alertsError", alertsError)) {
    $$payload.out += "<!--[-->";
    $$payload.out += `<div class="text-center text-danger form-text fs-5">An error occurred : ${escape_html(store_get($$store_subs ??= {}, "$alertsError", alertsError))}</div>`;
  } else {
    $$payload.out += "<!--[!-->";
  }
  $$payload.out += `<!--]--> <table class="table w-50 align-self-center"><thead><tr><th class="color-yellow" scope="col">Email</th><th class="color-yellow" scope="col">Agenda</th><th class="color-yellow" scope="col">ID</th><th scope="col">Actions</th></tr></thead><tbody><!--[-->`;
  for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
    let alert = each_array[$$index];
    $$payload.out += `<tr><td>`;
    if (!editable.get(alert.id)) {
      $$payload.out += "<!--[-->";
      $$payload.out += `<p class="color-yellow fw-bold">${escape_html(alert.email)}</p>`;
    } else {
      $$payload.out += "<!--[!-->";
      $$payload.out += `<input class="form-control-sm"${attr("value", alert.email)}>`;
    }
    $$payload.out += `<!--]--></td><td>`;
    if (!editable.get(alert.id)) {
      $$payload.out += "<!--[-->";
      $$payload.out += `${escape_html(alert.agendaId)} <br> ${escape_html(alertAgenda.get(alert.agendaId)?.name ? alertAgenda.get(alert.agendaId)?.name : "Pb while getting agenda name")}`;
    } else {
      $$payload.out += "<!--[!-->";
      $$payload.out += `<input class="form-control-sm"${attr("value", alert.agendaId)}>`;
    }
    $$payload.out += `<!--]--></td><td>${escape_html(alert.id)}</td><td class="justify-content-around">`;
    if (!editable.get(alert.id)) {
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
    uniqueName: "alerts",
    deleteFunction: deleteAlertById
  });
  $$payload.out += `<!----></div>`;
  if ($$store_subs) unsubscribe_stores($$store_subs);
  pop();
}
export {
  _page as default
};
