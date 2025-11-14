export const manifest = (() => {
function __memo(fn) {
	let value;
	return () => value ??= (value = fn());
}

return {
	appDir: "_app",
	appPath: "_app",
	assets: new Set(["favicon.png","fontawesome/css/brands.css","fontawesome/css/fontawesome.css","fontawesome/css/regular.css","fontawesome/css/solid.css","fontawesome/webfonts/fa-brands-400.ttf","fontawesome/webfonts/fa-brands-400.woff2","fontawesome/webfonts/fa-regular-400.ttf","fontawesome/webfonts/fa-regular-400.woff2","fontawesome/webfonts/fa-solid-900.ttf","fontawesome/webfonts/fa-solid-900.woff2","fontawesome/webfonts/fa-v4compatibility.ttf","fontawesome/webfonts/fa-v4compatibility.woff2","js/bootstrap.min.js","style/bootstrap.css","style/bootstrap_pulse.css","style/custom.css"]),
	mimeTypes: {".png":"image/png",".css":"text/css",".ttf":"font/ttf",".woff2":"font/woff2",".js":"text/javascript"},
	_: {
		client: {"start":"_app/immutable/entry/start.BCLYsi8y.js","app":"_app/immutable/entry/app.C4LdfXA8.js","imports":["_app/immutable/entry/start.BCLYsi8y.js","_app/immutable/chunks/entry.tons6ZdQ.js","_app/immutable/chunks/runtime.CGf8NitW.js","_app/immutable/chunks/index.BHUNO-cV.js","_app/immutable/chunks/index-client.D2jsEopY.js","_app/immutable/entry/app.C4LdfXA8.js","_app/immutable/chunks/runtime.CGf8NitW.js","_app/immutable/chunks/render.7UwKtW8V.js","_app/immutable/chunks/disclose-version.wNEiCm-t.js","_app/immutable/chunks/store.DSIwpa3B.js","_app/immutable/chunks/index-client.D2jsEopY.js","_app/immutable/chunks/props.DZC5C3sh.js"],"stylesheets":[],"fonts":[],"uses_env_dynamic_public":false},
		nodes: [
			__memo(() => import('./nodes/0.js')),
			__memo(() => import('./nodes/1.js'))
		],
		routes: [
			
		],
		matchers: async () => {
			
			return {  };
		},
		server_assets: {}
	}
}
})();
