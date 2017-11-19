<template>
  <v-flex>
  <v-card>
    <v-card-text>
      <v-layout row class="text-xs-center">
        <v-flex v-for="item in featured" style="transform: rotate(180deg);">
          <div class="bar-chart" :style="calcHeight(item.percentage)">{{ item.percentage }}</div>
        </v-flex>
      </v-layout>
      
      <v-layout row class="text-xs-center">
        <v-flex v-for="item in featured">
          <v-avatar>
            <img :src="item.avatar" alt="Avatar">
          </v-avatar>
        </v-flex>
      </v-layout>

      <v-layout row class="text-xs-center">
        <v-flex v-for="item in featured">
          {{ item.name }}
        </v-flex>
      </v-layout>
    </v-card-text>
  </v-card>
  </v-flex>
</template>

<script>
export default {
  data() {
    return {
      featured: [
        {
          percentage: 86,
          name: "Bożena",
          avatar: "/static/img/avatars/6.jpg"
        },
        {
          percentage: 64,
          name: "Aneta",
          avatar: "/static/img/avatars/8.jpg"
        },
        {
          percentage: 35,
          name: "Janek",
          avatar: "/static/img/avatars/4.jpg"
        }
      ]
    }
  },
  methods: {
    calcHeight: function(percentage) {
      return "height: " + percentage + "px"
    }
  }
}
</script>

<style>
.bar-chart {
  background-color: #009688;
  width: 35px;
  margin: 20px auto;
  color: #fff;
  transform: rotate(180deg);
}
</style>