<template>
  <v-flex>
    <v-card>
    <v-list subheader dense>
      <v-subheader>Your progress</v-subheader>
      <v-list-tile avatar v-for="item in items" v-bind:key="item.title" :disabled="item.disabled">
        <v-list-tile-avatar>
          <img v-bind:src="item.avatar"/>
        </v-list-tile-avatar>
        <v-list-tile-content>
          <v-list-tile-title v-html="item.title"></v-list-tile-title>
        </v-list-tile-content>
        <v-list-tile-action><strong>{{ item.points }} points</strong></v-list-tile-action>
      </v-list-tile>
    </v-list>

    </v-card>
    <v-container>
      <v-card>
        test
      </v-card>
    </v-container>
  </v-flex>
</template>

<script>
export default {
  data() {
    return {
      items: [
        {
          points: 72,
          title: "Weronika",
          avatar: "/static/img/avatars/1.jpg",
          disabled: true
        },
        {
          points: 51,
          title: "Marek",
          avatar: "/static/img/avatars/2.jpg",
          disabled: false        
        },
        {
          points: 47,
          title: "Gracjan",
          avatar: "/static/img/avatars/3.jpg ",
          disabled: true
        }
      ]
    };
  }
};
</script>

<style>
  .disabled {
    opacity: 0.5;
  }
</style>