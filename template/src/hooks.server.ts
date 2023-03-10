import { minify } from "html-minifier";
import { building, browser } from "$app/environment";
import type { Handle } from "@sveltejs/kit";
import { guard } from "$lib/store/auth";
import { goto } from "$app/navigation";

const minification_options = {
  collapseBooleanAttributes: true,
  collapseWhitespace: true,
  conservativeCollapse: true,
  decodeEntities: true,
  html5: true,
  ignoreCustomComments: [/^#/],
  minifyCSS: true,
  minifyJS: true,
  removeAttributeQuotes: true,
  removeComments: false, // some hydration code needs comments, so leave them in
  removeOptionalTags: true,
  removeRedundantAttributes: true,
  removeScriptTypeAttributes: true,
  removeStyleLinkTypeAttributes: true,
  sortAttributes: true,
  sortClassName: true,
};

let auth: boolean = false;

guard.subscribe((value) => {
  auth = value;
});

export const handle: Handle = async ({ event, resolve }) => {
  let page = "";

  if (browser) {
    if (!event.url.pathname.startsWith("/auth")) {
      goto("/auth/login")
    }
  }
  return resolve(event, {
    transformPageChunk: ({ html, done }) => {
      page += html;
      if (done) {
        return building ? minify(page, minification_options) : page;
      }
    },
  });
};
