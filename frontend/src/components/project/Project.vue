<template>
    <div>
        <v-layout>
        <v-slide-y-transition mode="out-in">
            <router-view></router-view>
        </v-slide-y-transition>
        </v-layout>
        <v-bottom-nav absolute shift :value="true" :active.sync="accent_color" :color="computedColor">
        <v-btn to="./home" dark>
            <span>Home</span>
            <v-icon>home</v-icon>
        </v-btn>
        <v-btn to="./score" dark>
            <span>Score</span>
            <v-icon>assessment</v-icon>
        </v-btn>
        <v-btn to="./profile" dark>
            <span>Profile</span>
            <v-icon>face</v-icon>
        </v-btn>
        <v-btn to="./settings" dark>
            <span>Settings</span>
            <v-icon>settings</v-icon>
        </v-btn>
        </v-bottom-nav>
    </div>
</template>
<script>
import { API_URL } from "@/constants.js";

export default {
  name: "Project",
  data() {
    return {
      accent_color: 0,
      project: {}
    };
  },
  beforeRouteUpdate(to, from, next) {
    if (to.params.projectId != to.params.projectId) {
        this.fetchProject(to.params.projectId, function() {
            next();
        })
    } else {
        next();
    }
  },
  created() {
    let params = this.$route.params;
    this.$http.get(API_URL + "projects/" + params.projectId + "/").then(
      response => {
        this.project = response;
      },
      response => {
        console.log("Should not happen");
      }
    );
  },
  methods: {
    fetchProject(id, callback) {
      this.$http.get(API_URL + "projects/" + params.projectId + "/").then(
        response => {
          this.project = response;
          callback();
        },
        response => {
          callback();
        }
      );
    }
  },
  computed: {
    computedColor() {
      switch (this.accent_color) {
        case 0:
          return "blue-grey";
          break;
        case 1:
          return "teal";
          break;
        case 2:
          return "brown";
          break;
        case 3:
          return "black";
          break;
      }
    }
  }
};
</script>