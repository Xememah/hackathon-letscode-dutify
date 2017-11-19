<template>
  <v-flex>
    <v-card>
      <v-card-text class="text-xs-center">
      <v-avatar size="100px">
        <img :src="data.avatar">
      </v-avatar>

      <p class="text-xs-center">
        <h2>{{ data.name }}</h2>
      </p>
      </v-card-text>
    </v-card>
    <v-card flat>
      
    </v-card>
  </v-flex>
</template>

<script>
export default {
  data () {
    return {
      data: {
        name: "Weronika",
        avatar: "/static/img/avatars/1.jpg"
      }
    }
  }
}
</script>
