<template>
  <v-app light flat>
    <!-- <v-navigation-drawer
      fixed
      :mini-variant="miniVariant"
      :clipped="clipped"
      v-model="drawer"
      app
    >
    </v-navigation-drawer> -->
    <v-toolbar fixed app :clipped-left="clipped" color="white" flat>
      <v-toolbar-side-icon @click.stop="drawer = !drawer" light></v-toolbar-side-icon>
      <!--<v-btn
        icon
        light
        @click.stop="miniVariant = !miniVariant"
      >
        <v-icon v-html="miniVariant ? 'chevron_right' : 'chevron_left'"></v-icon>
      </v-btn>
       <v-btn
        icon
        light
        @click.stop="clipped = !clipped"
      >
        <v-icon>web</v-icon>
      </v-btn>
      <v-btn
        icon
        light
        @click.stop="fixed = !fixed"
      >
        <v-icon>remove</v-icon>
      </v-btn> -->
      <v-toolbar-title v-text="title"></v-toolbar-title>
      <v-spacer></v-spacer>
      <!-- <v-btn
        icon
        light
        @click.stop="rightDrawer = !rightDrawer"
      >
        <v-icon>menu</v-icon>
      </v-btn> -->
    </v-toolbar>
    <v-content>
      <v-container fluid>
        <v-slide-y-transition mode="out-in">
          <v-layout column align-center>
            <router-view></router-view>
          </v-layout>
        </v-slide-y-transition>
      </v-container>
    </v-content>
    <!-- <v-navigation-drawer
      temporary
      :right="right"
      v-model="rightDrawer"
      fixed
    >
      <v-list>
        <v-list-tile @click="right = !right">
          <v-list-tile-action>
            <v-icon light>compare_arrows</v-icon>
          </v-list-tile-action>
          <v-list-tile-title>Switch drawer (click me)</v-list-tile-title>
        </v-list-tile>
      </v-list>

    </v-navigation-drawer> -->
    <!-- <v-footer :fixed="fixed" app>
      <span>&copy; 2017</span>
    </v-footer> -->
    <v-bottom-nav
      absolute
      shift
      :value="true"
      :active.sync="accent_color"
      :color="computedColor"
    >
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
      clipped: false,
      drawer: true,
      fixed: false,
      miniVariant: false,
      right: true,
      rightDrawer: false,
      title: "Symborify",
      accent_color: 0
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
