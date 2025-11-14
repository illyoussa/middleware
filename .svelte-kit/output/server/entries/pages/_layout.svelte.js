import { V as slot } from "../../chunks/index.js";
import "clsx";
import "../../chunks/client.js";
function Navbar($$payload) {
  $$payload.out += `<nav class="navbar bg-secondary" data-bs-theme="dark"><div class="container-fluid ms-4"><a class="navbar-brand color-yellow fw-bold" href="/">Timetable alerter</a></div></nav>`;
}
function Sidebar($$payload) {
  $$payload.out += `<nav class="navbar bg-secondary align-content-start" style="width: fit-content"><ul class="d-flex flex-column navbar-nav text-center"><li class="nav-item mb-4 clickable-item svelte-1h9gjgr"><a class="clickable-item svelte-1h9gjgr" href="/agendas"><span class="h1 fas fa-calendar-plus"></span></a></li> <li class="nav-item mb-4 clickable-item svelte-1h9gjgr"><a href="/alerts" class="svelte-1h9gjgr"><span class="h1 fas fa-bell"></span></a></li> <li class="nav-item mb-4 clickable-item svelte-1h9gjgr"><a class="clickable-item svelte-1h9gjgr" href="/events"><span class="h1 fas fa-calendar-days"></span></a></li></ul></nav>`;
}
function _layout($$payload, $$props) {
  $$payload.out += `<div class="d-flex flex-column">`;
  Navbar($$payload);
  $$payload.out += `<!----> <div class="d-flex flex-row" style="min-height: 95vh">`;
  Sidebar($$payload);
  $$payload.out += `<!----> <div class="mx-auto" style="margin-top: 4rem; padding: 2rem;width: 90%"><!---->`;
  slot($$payload, $$props, "default", {});
  $$payload.out += `<!----></div></div></div>`;
}
export {
  _layout as default
};
