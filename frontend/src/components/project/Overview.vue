<template>
  <v-flex>
    <v-card>
      <v-list subheader dense>
        <v-subheader>Your progress</v-subheader>
        <v-list-tile avatar v-for="item in data.project.ranking" v-bind:key="item.title" :disabled="item.disabled">
          <v-list-tile-avatar>
            <img :src="item.avatar"/>
          </v-list-tile-avatar>
          <v-list-tile-content>
            <v-list-tile-title v-html="item.name"></v-list-tile-title>
          </v-list-tile-content>
          <v-list-tile-action><strong>{{ item.points }} points</strong></v-list-tile-action>
        </v-list-tile>
      </v-list>
    </v-card>
    
    <v-list style="background-color: #FAFAFA">
      <v-subheader>
        Predefined activities
      </v-subheader>
      <v-list-tile avatar v-for="item in data.project.duties" v-bind:key="item.title">
        <v-list-tile-avatar>
          <v-icon v-bind:class="[item.iconClass]">{{ item.icon }}</v-icon>
        </v-list-tile-avatar>
        <v-list-tile-content>
          <v-list-tile-title @click.stop="openDuty(item.id)">{{ item.name }} ({{ item.reward }} pkt)</v-list-tile-title>
        </v-list-tile-content>
        <v-list-tile-action>
          <v-btn icon ripple @click.stop="confirm(data.project.ID, item.id)">
            <v-icon color="green darken-1">add</v-icon>
          </v-btn>
        </v-list-tile-action>
      </v-list-tile>
    </v-list>
    </v-card>

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
          <v-toolbar-title>Add a new activity</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-toolbar-items>
            <v-btn dark flat @click="addDuty">Save</v-btn>
            <v-menu bottom right offset-y>
            </v-menu>
          </v-toolbar-items>
        </v-toolbar>
        <v-form class="pa-3">
          <v-text-field
            color="white"
            label="Name"
            v-model="objective.name"
            required
          ></v-text-field>
          <v-text-field
            label="Reward"
            type="number"
            v-model="objective.reward"
            required
          ></v-text-field>
        </v-form>
      </v-card>
    </v-dialog>
    <v-btn fab absolute right bottom dark color="#607D8B" style="bottom: 70px;" @click.stop="dialog = true">
      <v-icon dark>add</v-icon>
    </v-btn>
  </v-flex>
</template>

<script>
import { API_URL } from "@/constants.js";

export default {
  data() {
    return {
      objective: {
        name: "",
        reward: 1,
        icon: ["house", "restaurant", "event_seat", "rowing"][Math.abs(Math.floor(Math.random()*4-0.1))]
      },
      data: this.$store,
      dialog: false
    };
  },
  methods: {
    confirm(projectId, id) {
      this.$http
        .post(API_URL + "projects/" + projectId + "/duties/" + id + "/confirm")
        .then(
          response => {
            this.$store.bus.$emit("refresh-project");
          },
          response => {}
        );
    },
    addDuty() {
      this.objective.reward = Number(this.objective.reward)
      this.$http
        .post(API_URL + "projects/" + this.data.project.ID + "/duties", this.objective)
        .then(
          response => {
            this.dialog = false;
            this.$store.bus.$emit("refresh-project");
          },
          response => {}
        );
    },
    openDuty(id) {
        this.$router.push('./duty/'+id)
    },
  }
};
</script>

<style>
.disabled {
  opacity: 0.5;
}
</style>