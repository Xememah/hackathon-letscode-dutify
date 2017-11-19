<template>
  <v-app light flat>
    <v-navigation-drawer fixed v-model="drawer" disable-resize-watcher disable-route-watcher app>
      <v-container>
      <img src="/static/img/logo/dark-logo.svg" alt="Dutify" style="height: 25px;">
      
      <v-list>
        <v-list-tile avatar v-for="item in this.data.projects" v-bind:key="item.ID" :id="item.ID" @click="goto(item.ID)">
          <v-list-tile-avatar>
            <v-icon v-bind:class="[item.iconClass]">{{ item.icon }}</v-icon>
          </v-list-tile-avatar>
          <v-list-tile-content>
            <v-list-tile-title>{{ item.name }}</v-list-tile-title>
            <v-list-tile-sub-title>5 people</v-list-tile-sub-title>
          </v-list-tile-content>
        </v-list-tile>

        <v-divider inset></v-divider>

        <v-list-tile avatar @click="dialog = true">
          <v-list-tile-avatar>
            <v-icon class="grey lighten-5">add</v-icon>
          </v-list-tile-avatar>
          <v-list-tile-content>
            <v-list-tile-title>Create new</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
        <v-list-tile avatar @click="joinDialog = true">
          <v-list-tile-avatar>
            <v-icon class="grey lighten-5">add</v-icon>
          </v-list-tile-avatar>
          <v-list-tile-content>
            <v-list-tile-title>Join</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
      </v-container>
    <v-dialog
      v-model="dialog"
      fullscreen
      transition="dialog-bottom-transition"
      :overlay=false scrollable>
      <v-card>
          <v-toolbar style="flex: 0 0 auto;" dark class="primary">
          <v-btn icon @click.native="dialog = false" dark>
            <v-icon>close</v-icon>
          </v-btn>
          <v-toolbar-title>Add a new project</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-toolbar-items>
            <v-btn dark flat @click="addProject">Add</v-btn>
            <v-menu bottom right offset-y>
            </v-menu>
          </v-toolbar-items>
        </v-toolbar>
        <v-form class="pa-3">
          <v-text-field
            color="white"
            label="Name"
            v-model="project.name"
            required
          ></v-text-field>
        </v-form>
      </v-card>
    </v-dialog>
    <v-dialog
      v-model="joinDialog"
      fullscreen
      transition="dialog-bottom-transition"
      :overlay=false scrollable>
      <v-card>
          <v-toolbar style="flex: 0 0 auto;" dark class="primary">
          <v-btn icon @click.native="joinDialog = false" dark>
            <v-icon>close</v-icon>
          </v-btn>
          <v-toolbar-title>Join a project</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-toolbar-items>
            <v-btn dark flat @click="joinProject">Join</v-btn>
            <v-menu bottom right offset-y>
            </v-menu>
          </v-toolbar-items>
        </v-toolbar>
        <v-form class="pa-3">
          <v-text-field
            color="white"
            label="Secret code"
            v-model="project.id"
            required
          ></v-text-field>
        </v-form>
      </v-card>
    </v-dialog>
    </v-navigation-drawer>
    
    <v-toolbar app color="white" flat>
      <v-toolbar-side-icon @click.stop="drawer = !drawer" light></v-toolbar-side-icon>
      <v-toolbar-title>Dutify {{ data.project.name ? ('- '+data.project.name) : '' }}</v-toolbar-title>
      <v-spacer></v-spacer>
    </v-toolbar>
    <router-view></router-view>
  </v-app>
</template>

<script>
import {
  API_URL
} from './constants.js'
export default {
  mounted() {
    this.fetchProjects()
  },
  methods: {
    fetchProjects: function() {
      if (this.$store.user.token()) {
        this.$http.get(API_URL + "projects/").then(
          response => {
            this.data.projects = response.body
            if(this.data.projects.length > 0) {
              if(this.$route.fullPath == '/') {
                this.goto(this.data.projects[0].ID)
              }
            }
          },
          response => {
            console.log("TODO")
          }
        );
      } else if(this.$route.fullPath != '/login') {
        this.$router.push('login')
      }
    },
    addProject() {
      this.dialog = false;
      let self = this;
      this.$http.post(API_URL + "projects/", this.project).then(
        response => {
          self.fetchProjects()
        },
        response => {
          console.log("TODO")
        }
      );
    },
    joinProject() {
      this.joinDialog = false;
      let self = this;
      this.$http.post(API_URL + "projects/"+this.project.id+"/join").then(
        response => {
          self.fetchProjects()
          goto(this.project.id)
        },
        response => {
          console.log("TODO")
        }
      );
    },
    goto: function(id) {
      this.$router.push('/project/'+id);
      this.$store.bus.$emit("refresh-project", id);
    }
  },
  data() {
    return {
      project: {
        id: 0,
        name: '',
        hidden: false,
        icon: ["house", "restaurant", "event_seat", "rowing"][Math.abs(Math.floor(Math.random()*4-0.1))]
      },
      data: this.$store,
      dialog: false,
      joinDialog: false,
      drawer: false,
      fixed: false,
      title: "Dutify",
      accent_color: 0,
    };
  }
};
</script>
