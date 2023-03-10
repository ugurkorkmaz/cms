import { writable } from "svelte/store";

export interface Route {
  path: string
  icon?: string
  auth: boolean
}
export const routes = writable<Route[]>([
  {
    path: "/auth/magicway",
    icon: "home",
    auth: false,
  },
  {
    path: "/auth/login",
    icon: "key",
    auth: false,
  },
  {
    path: "/auth/register",
    icon: "person",
    auth: false,
  },
  {
    path: "/",
    icon: "home",
    auth: true,
  },
]);

export let theme = writable("lemonade");
