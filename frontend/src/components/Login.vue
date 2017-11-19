<template>
  <v-flex>
    <v-card>
      <v-card-text>
        <v-form v-model="valid">
          <v-text-field
            color="white"
            label="E-mail"
            v-model="email"
            :rules="emailRules"
            required
          ></v-text-field>
          <v-text-field
            label="Password"
            v-model="password"
            required
          ></v-text-field>
          <v-btn color="primary">Log in</v-btn>
        </v-form>
      </v-card-text>
    </v-card>
  </v-flex>
</template>

<script>
  export default {
    data () {
      return {
        valid: false,
        email: '',
        emailRules: [
          (v) => !!v || 'E-mail is required',
          (v) => /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(v) || 'E-mail must be valid'
        ]
      }
    }
  }
</script>