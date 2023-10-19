import { createApp } from "vue";
import App from "./App.vue";
import "bootstrap/dist/css/bootstrap.css";
import "bootstrap-vue/dist/bootstrap-vue.css";
import VueGoogleMaps from "@fawmi/vue-google-maps";
import "./styles/global.css";
import "@fortawesome/fontawesome-free/css/all.css";
import "@fortawesome/fontawesome-free/js/all.js";

const apiKey = process.env.VUE_APP_GOOGLE_MAP_API_KEY;
const app = createApp(App);
app
  .use(VueGoogleMaps, {
    load: {
      key: apiKey,
    },
  })
  .mount("#app");
