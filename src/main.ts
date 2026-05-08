import { mount } from "svelte";
import App from "./App.svelte";
import * as VIAM from "@viamrobotics/sdk";
import Cookies from "js-cookie";
import "./app.css";

let apiKeyId = "";
let apiKeySecret = "";
let host = "";
let machineId = "";

// Extract machine credentials from the cookie set by Viam's app platform
try {
  const machineCookieKey = window.location.pathname.split("/")[2];
  ({
    apiKey: { id: apiKeyId, key: apiKeySecret },
    machineId: machineId,
    hostname: host,
  } = JSON.parse(Cookies.get(machineCookieKey)!));
} catch {
  // Running locally — use env vars from .env.local
  apiKeyId = import.meta.env.VITE_API_KEY_ID ?? "";
  apiKeySecret = import.meta.env.VITE_API_KEY_SECRET ?? "";
  machineId = import.meta.env.VITE_MACHINE_ID ?? "";
  host = import.meta.env.VITE_HOST ?? "";
}

// Mount the Svelte app
const app = mount(App, {
  target: document.getElementById("app")!,
  props: {
    apiKeyId,
    apiKeySecret,
    machineId,
    host,
  },
});

export default app;
