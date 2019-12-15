import Vue from "vue";
import SuiVue from 'semantic-ui-vue';
import LoginComponent from "./Login/components/Login.vue";

Vue.use(SuiVue);

const v = new Vue({
  el: "#app",
  components: {
    LoginComponent
  }
});
