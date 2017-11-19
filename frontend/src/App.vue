<template>
  <v-app light flat>
    <v-navigation-drawer fixed v-model="drawer" stateless app>
      <v-container>
      <img src="/static/img/logo/dark-logo.svg" alt="Dutify" style="height: 25px;">
      </v-container>
    </v-navigation-drawer>
    
    <v-toolbar app color="white" flat>
      <v-toolbar-side-icon @click.stop="drawer = !drawer" light></v-toolbar-side-icon>
      <v-toolbar-title v-text="title"></v-toolbar-title>
      <v-spacer></v-spacer>
    </v-toolbar>
    <router-view></router-view>
  </v-app>
</template>

<script>
export default {
  data() {
    return {
      drawer: false,
      fixed: false,
      title: "Dutify",
    };
  },
};
</script>
