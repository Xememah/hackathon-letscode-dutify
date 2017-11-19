<template>
    <div>
        <v-layout>
        <v-slide-y-transition mode="out-in">
            <router-view></router-view>
        </v-slide-y-transition>
        </v-layout>
        <v-bottom-nav absolute shift :value="true" :color="computedColor">
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
        </v-bottom-nav>
    </div>
</template>
<script>
import { API_URL } from "@/constants.js";

export default {
  name: "Project",
  data() {
    return {
      project: {}
    };
  },
  mounted: function() {
    this.$store.bus.$on("refresh-project", (id)  => {
        if(!id) {
            id = this.$store.project.ID
        }
      this.fetchProject(id, ()=>{});
    });
  },
  beforeRouteUpdate(to, from, next) {
    if (to.params.projectId != to.params.projectId) {
      this.fetchProject(to.params.projectId, function() {
        next();
      });
    } else {
      next();
    }
  },
  created() {
    let params = this.$route.params;
    this.fetchProject(params.projectId, () => {});
  },
  methods: {
    fetchProject(id, callback) {
      this.$http.get(API_URL + "projects/" + id + "/").then(
        response => {
          response.body.max = 1
          for (let entry of response.body.ranking) {
            if(entry.points > response.body.max) {
                response.body.max = entry.points
            }
            for (let user of response.body.users) {
              if (user.id == entry.user) {
                entry.name = user.name;
                entry.avatar = 'static/img/avatars/'+(user.id%8+1)+'.jpg';
              }
            }
          }
          this.$store.project = response.body;
          this.project = this.$store.project;
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
      switch (this.$route.name) {
        case 'home':
          return "blue-grey";
          break;
        case 'score':
          return "teal";
          break;
        case 'profile':
          return "brown";
          break;
        deafult:
          return "black";
          break;
      }
    }
  }
};
</script>