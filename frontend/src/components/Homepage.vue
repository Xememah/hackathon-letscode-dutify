<template>
  <v-flex>
    <v-card>
      <v-list subheader dense>
        <v-subheader>Your progress</v-subheader>
        <v-list-tile avatar v-for="item in progress" v-bind:key="item.title" :disabled="item.disabled">
          <v-list-tile-avatar>
            <img v-bind:src="item.avatar"/>
          </v-list-tile-avatar>
          <v-list-tile-content>
            <v-list-tile-title v-html="item.title"></v-list-tile-title>
          </v-list-tile-content>
          <v-list-tile-action><strong>{{ item.points }} points</strong></v-list-tile-action>
        </v-list-tile>
      </v-list>
    
    </v-card>
    
    <!-- <v-container class="text-xs-center">
      <v-avatar style="background-color: #607D8B">
        <v-icon dark>people</v-icon>
      </v-avatar>
    </v-container> -->

    <v-list style="background-color: #FAFAFA">
      <v-subheader>
        Predefined activities
      </v-subheader>
      <v-list-tile avatar v-for="item in actions" v-bind:key="item.title" @click="">
        <v-list-tile-avatar>
          <v-icon v-bind:class="[item.iconClass]">{{ item.icon }}</v-icon>
        </v-list-tile-avatar>
        <v-list-tile-content>
          <v-list-tile-title>{{ item.title }}</v-list-tile-title>
          <v-list-tile-sub-title>{{ item.subtitle }}</v-list-tile-sub-title>
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
      progress: [
        {
          points: 72,
          title: "Weronika",
          avatar: "/static/img/avatars/1.jpg",
          disabled: true
        },
        {
          points: 51,
          title: "Marek",
          avatar: "/static/img/avatars/2.jpg",
          disabled: false        
        },
        {
          points: 47,
          title: "Gracjan",
          avatar: "/static/img/avatars/3.jpg ",
          disabled: true
        }
      ],
      actions: [
        {
          title: "Bathroom cleaning",
          points: 10,
          icon: "home"
        },
        {
          title: "Cooking dinner",
          points: 7,
          icon: "restaurant"
        },
        {
          title: "Dishwashing",
          points: 3,
          icon: "local_dining"
        },
        {
          title: "Shopping for groceries",
          points: 2,
          icon: "shopping_basket"
        }
      ]
    };
  }
};
</script>

<style>
  .disabled {
    opacity: 0.5;
  }
</style>