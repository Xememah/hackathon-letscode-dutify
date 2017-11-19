<template>
  <v-flex>
    <v-card>
      <v-list subheader dense>
        <v-subheader>Your progress</v-subheader>
        <v-list-tile avatar v-for="item in data.project.ranking" v-bind:key="item.title" :disabled="item.disabled">
          <v-list-tile-avatar>
            <img src="/static/img/avatars/2.jpg"/>
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
      <v-list-tile avatar v-for="item in data.project.duties" v-bind:key="item.title" @click="confirm">
        <v-list-tile-avatar>
          <v-icon v-bind:class="[item.iconClass]">{{ item.icon }}</v-icon>
        </v-list-tile-avatar>
        <v-list-tile-content>
          <v-list-tile-title>{{ item.name }}</v-list-tile-title>
        </v-list-tile-content>
        <v-list-tile-action>
          <v-btn icon ripple>
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
export default {
  data() {
    return {
      data: this.$store,
    };
  }
};
</script>

<style>
  .disabled {
    opacity: 0.5;
  }
</style>