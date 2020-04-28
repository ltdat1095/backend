import Vue from "vue";
import SuiVue from 'semantic-ui-vue';
import LoginComponent from './components/Login.vue';

Vue.use(SuiVue);

const vue = new Vue({
  el: "#app",
  components: {
    LoginComponent
  }
});

export default vue;
