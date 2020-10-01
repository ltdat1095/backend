import Vue from "vue";
import SuiVue from 'semantic-ui-vue';
import SignUpComponent from './components/SignUp.vue';

Vue.use(SuiVue);

const vue = new Vue({
  el: "#signUpComponent",
  components: {
    SignUpComponent
  }
});

export default vue;
