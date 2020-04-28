<template>
    <div class="ui container">
        <div class="ui form">
            <div class="field">
                <label>Email</label>
                <input type="text" v-model="email" placeholder="Email">
            </div>
            <div class="field">
                <label>Password</label>
                <input type="password" v-model="password" placeholder="Password">
            </div>
            <button class="ui button" @click="submit">Submit</button>
        </div>
    </div>
</template>

<script lang="ts">
import { Vue, Component, Prop } from 'vue-property-decorator'
import SuiVue from 'semantic-ui-vue';
import axios from 'axios';

Vue.use(SuiVue);
@Component
export default class LoginComponent extends Vue {
  email = ""
  password = ""

  async submit() {
    try {
      const { data } = await axios.post('/login', {
        email: this.email,
        password: this.password,
      })
      if (data.status != 200) {
        alert(data.message)
      } else {
        alert(data.message)
      }
    } catch (error) {
      console.log(error)
    }
  }
}
</script>

<style>
</style>
