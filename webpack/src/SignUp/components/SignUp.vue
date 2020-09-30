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
      <div class="field">
        <label>First Name</label>
        <input type="text" v-model="firstName" placeholder="First Name">
      </div>
      <div class="field">
        <label>Last Name</label>
        <input type="text" v-model="lastName" placeholder="Last Name">
      </div>
      <div class="field">
          <label>Age</label>
          <input type="number" v-model.number="age" placeholder="Age">
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
export default class SignUpComponent extends Vue {
  email = "";
  password = "";
  firstName = "";
  lastName = "";
  age=0;

  async submit() {
    try {
      const { data } = await axios.post('/api/signUp', {
        email: this.email,
        password: this.password,
        firstName: this.firstName,
        lastName: this.lastName,
        age: this.age
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
