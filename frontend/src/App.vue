<template>
  <v-app light flat>
    <v-navigation-drawer fixed v-model="drawer" app>
      <v-container class="text-xs-center">
      <img src="/static/img/logo/dark-logo.svg" alt="Dutify" style="height: 25px;">
      
      <v-list>
        <v-list-tile avatar v-for="item in projects" v-bind:key="item.title" @click="">
          <v-list-tile-avatar>
            <v-icon v-bind:class="[item.iconClass]">{{ item.icon }}</v-icon>
          </v-list-tile-avatar>
          <v-list-tile-content>
            <v-list-tile-title>{{ item.title }}</v-list-tile-title>
            <v-list-tile-sub-title>{{ item.subtitle }}</v-list-tile-sub-title>
          </v-list-tile-content>
        </v-list-tile>

        <v-divider inset></v-divider>

        <v-list-tile avatar @click="">
          <v-list-tile-avatar>
            <v-icon class="grey lighten-5">add</v-icon>
          </v-list-tile-avatar>
          <v-list-tile-content>
            <v-list-tile-title>Create / join a new group</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
      </v-container>
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
  </v-app>
</template>

<script>
export default {
  data() {
    return {
      drawer: false,
      fixed: false,
      title: "Dutify",
      accent_color: 0,
      projects: [
        {
          title: "Me and my girlfriend",
          icon: "whatshot",
          subtitle: "2 people"
        },
        {
          title: "Flatmates",
          icon: "people",
          subtitle: "5 people"
        }
      ]
    }
  },
  computed: {
    computedColor () {
      switch (this.accent_color) {
        case 0:
          return 'blue-grey'
        break
        case 1:
          return 'teal'
        break
        case 2:
          return 'brown'
        break
        case 3:
          return 'black'
        break
      }
    }
  }
}
</script>
