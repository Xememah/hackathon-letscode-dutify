<template>
  <v-flex>
    <v-card>
      <v-card-text>
        <v-form v-model="valid">
          <!-- <v-text-field
            color="white"
            label="Name"
            v-model="credentials.name"
            required
          ></v-text-field> -->
          <v-text-field
            color="white"
            label="E-mail"
            v-model="credentials.email"
            :rules="emailRules"
            required
          ></v-text-field>
          <v-text-field
            label="Password"
            type="password"
            v-model="credentials.password"
            required
          ></v-text-field>
          <v-btn color="primary" v-on:click="login">Log in</v-btn>
        </v-form>
      </v-card-text>
    </v-card>
  </v-flex>
</template>

<script>
import { API_URL } from "@/constants.js";
export default {
  data() {
    return {
      valid: false,
      credentials: {
        email: "",
        password: "",
        name: "John Smith"
      },
      emailRules: [
        v => !!v || "E-mail is required",
        v =>
          /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(v) ||
          "E-mail must be valid"
      ]
    };
  },
  methods: {
    login() {
      let self = this
      this.$http
        .post(API_URL + "accounts/login/", this.credentials)
        .then(response => {
          window.localStorage.setItem("token", response.data.token);
          self.$router.push('/homepage');
        })
        .catch(errors => {
          console.log(errors);
        });
    }
  }
};
</script>