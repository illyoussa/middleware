import * as server from '../entries/pages/_layout.server.js';

export const index = 0;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/pages/_layout.svelte.js')).default;
export { server };
export const server_id = "src/routes/+layout.server.js";
export const imports = ["_app/immutable/nodes/0.BSgB03xI.js","_app/immutable/chunks/disclose-version.wNEiCm-t.js","_app/immutable/chunks/runtime.CGf8NitW.js","_app/immutable/chunks/legacy.B129RShy.js","_app/immutable/chunks/entry.tons6ZdQ.js","_app/immutable/chunks/index.BHUNO-cV.js","_app/immutable/chunks/index-client.D2jsEopY.js"];
export const stylesheets = ["_app/immutable/assets/0.C1IPNx-7.css"];
export const fonts = [];
