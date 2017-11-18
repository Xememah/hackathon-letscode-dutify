<template>
  <v-flex>
    <v-list subheader dense>
      <v-subheader>In comparison</v-subheader>
      <v-list-tile avatar v-for="item in items" v-bind:key="item.title" :disabled="item.disabled">
        <v-list-tile-avatar>
          <img v-bind:src="item.avatar"/>
        </v-list-tile-avatar>
        <v-list-tile-content>
          <v-list-tile-title v-html="item.title"></v-list-tile-title>
        </v-list-tile-content>
        <v-list-tile-action>{{ item.points }} p</v-list-tile-action>
      </v-list-tile>
    </v-list>
  </v-flex>
</template>

<script>
export default {
  data() {
    return {
      items: [
        {
          points: 72,
          title: "Weronika Symbor",
          avatar: "http://i.pravatar.cc/300",
          disabled: true
        },
        {
          points: 51,
          title: "Marek Kochanowski",
          avatar: "http://i.pravatar.cc/300",
          disabled: false        },
        {
          points: 47,
          title: "Gracjan Brzoskowski",
          avatar: "http://i.pravatar.cc/300",
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