<template>
  <v-flex>
    <v-card>
      <v-card-text class="text-xs-center">
      <v-avatar size="100px">
        <img :src="this.getAvatar()">
      </v-avatar>

      <p class="text-xs-center">
        <h2>{{ data.user.getUser().name }}</h2>
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
      data: this.$store,
    }
  },
  methods: {
    getAvatar() {
      return '/static/img/avatars/'+(Number(this.data.user.getUser().id) % 8 +1)+'.jpg'
    }
  }
}
</script>
