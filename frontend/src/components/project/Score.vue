<template>
  <v-flex>
  <v-card>
    <v-card-text>
      <v-layout row class="text-xs-center">
        <v-flex v-for="item in this.data.project.ranking.slice(0,3)" v-bind:key="item.user" style="transform: rotate(180deg);">
          <div class="bar-chart" :style="calcHeight(item.points)">{{ item.points }} p</div>
        </v-flex>
      </v-layout>
      
      <v-layout row class="text-xs-center">
        <v-flex v-for="item in this.data.project.ranking.slice(0,3)" v-bind:key="item.user">
          <v-avatar>
            <img :src="item.avatar" alt="Avatar">
          </v-avatar>
        </v-flex>
      </v-layout>

      <v-layout row class="text-xs-center">
        <v-flex v-for="item in this.data.project.ranking.slice(0,3)" v-bind:key="item.user">
          {{ item.name }}
        </v-flex>
      </v-layout>
    </v-card-text>
  </v-card>
      <v-list subheader dense>
        <v-subheader>Top</v-subheader>
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
  </v-flex>
</template>

<script>
export default {
  data() {
    return {
      data: this.$store,
    }
  },
  methods: {
    calcHeight: function(points) {
      let percentage = 100;
      if(this.data.project.max > 0 && points > 0) {
        if(this.data.project.max >= points) {
          percentage = (points / this.data.project.max)*100
        }

        return "height: " + percentage + "px"
      } else {
        return "height: 0px"
      }
    }
  }
}
</script>

<style>
.bar-chart {
  background-color: #009688;
  width: 45px;
  margin: 20px auto;
  color: #fff;
  transform: rotate(180deg);
}
</style>