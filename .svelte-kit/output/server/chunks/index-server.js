import { X as escape_html } from "./index.js";
import "clsx";
import { w as writable } from "./index2.js";
import axios from "axios";
const replacements = {
  translate: /* @__PURE__ */ new Map([
    [true, "yes"],
    [false, "no"]
  ])
};
function attr(name, value, is_boolean = false) {
  if (value == null || !value && is_boolean || value === "" && name === "class") return "";
  const normalized = name in replacements && replacements[name].get(value) || value;
  const assignment = is_boolean ? "" : `="${escape_html(normalized, true)}"`;
  return ` ${name}${assignment}`;
}
const agendas = writable([]);
const agendasError = writable("");
const apiBaseUrl = "/config_api/agendas";
function getAgendas() {
  agendasError.set("");
  axios.get(`${apiBaseUrl}`).then((res) => {
    agendas.set(res.data);
  }).catch((err) => {
    console.log("An error has occurred while retrieving agendas");
    console.log(err);
    if (err.response?.data?.message) {
      agendasError.set(JSON.stringify(err.response.data.message));
    } else {
      agendasError.set(JSON.stringify(err));
    }
  });
}
function getAgenda(id) {
  agendasError.set("");
  return axios.get(`${apiBaseUrl}/${id}`).then((res) => {
    return Promise.resolve(res.data);
  }).catch((err) => {
    console.log("An error has occurred while retrieving agenda");
    console.log(err);
    if (err.response?.data?.message) {
      agendasError.set(JSON.stringify(err.response.data.message));
    } else {
      agendasError.set(JSON.stringify(err));
    }
    return Promise.reject(err);
  });
}
function deleteAgenda(id) {
  agendasError.set("");
  return axios.delete(`${apiBaseUrl}/${id}`).then((res) => {
    getAgendas();
    return Promise.resolve();
  }).catch((err) => {
    console.log("An error as occurred while deleting agenda");
    console.log(err);
    if (err.response?.data?.message) {
      agendasError.set(JSON.stringify(err.response.data.message));
    } else {
      agendasError.set(JSON.stringify(err));
    }
    return Promise.reject(err);
  });
}
const SvelteMap = globalThis.Map;
export {
  SvelteMap as S,
  attr as a,
  agendas as b,
  agendasError as c,
  deleteAgenda as d,
  getAgenda as g
};
