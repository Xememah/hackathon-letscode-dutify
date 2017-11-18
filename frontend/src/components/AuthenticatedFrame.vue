<template>
<v-navigation-drawer fixed v-model="drawer" app>
  </v-navigation-drawer>
  <v-toolbar app color="white" flat>
    <v-toolbar-side-icon @click.stop="drawer = !drawer" light></v-toolbar-side-icon>
    <v-toolbar-title v-text="title"></v-toolbar-title>
    <v-spacer></v-spacer>
  </v-toolbar>
  <v-layout>
    <v-slide-y-transition mode="out-in">
      <router-view></router-view>
    </v-slide-y-transition>
  </v-layout>
  <v-bottom-nav absolute shift :value="true" :active.sync="accent_color" :color="computedColor">
    <v-btn href="#/" dark>
      <span>Home</span>
      <v-icon>home</v-icon>
    </v-btn>
    <v-btn href="#/score" dark>
      <span>Score</span>
      <v-icon>assessment</v-icon>
    </v-btn>
    <v-btn href="#/profile" dark>
      <span>Profile</span>
      <v-icon>face</v-icon>
    </v-btn>
    <v-btn href="#/settings" dark>
      <span>Settings</span>
      <v-icon>settings</v-icon>
    </v-btn>
  </v-bottom-nav>
</template>