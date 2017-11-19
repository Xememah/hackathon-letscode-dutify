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
          <v-list-tile-title>{{ item.name }} ({{ item.reward }} pkt)</v-list-tile-title>
        </v-list-tile-content>
        <v-list-tile-action>
          <v-btn icon ripple @click="confirm(data.project.ID, item.id)">
            <v-icon color="green darken-1">add</v-icon>
          </v-btn>
        </v-list-tile-action>
      </v-list-tile>
    </v-list>
    </v-card>
    <v-btn fab fixed right bottom dark color="#607D8B" style="margin-bottom: 60px;">
      <v-icon dark>add</v-icon>
    </v-btn>
  </v-flex>
</template>

<script>
import { API_URL } from "@/constants.js";

export default {
  data() {
    return {
      data: this.$store
    };
  },
  methods: {
    confirm(projectId, id) {
      this.$http.post(API_URL + "projects/" + projectId + "/duties/"+id+"/confirm").then(
        response => {
            this.$store.bus.$emit('refresh-project')
            this.$router.push('./score')
        },
        response => {
        }
      );
    }
  }
};
</script>

<style>
.disabled {
  opacity: 0.5;
}
</style>