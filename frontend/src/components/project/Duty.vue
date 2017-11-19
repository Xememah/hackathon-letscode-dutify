<template>
  <v-flex>
    <v-card>
      <v-card-text class="text-xs-center">
      <v-avatar size="100px">
          <v-icon>{{ duty.icon }}</v-icon>
      </v-avatar>

      <p class="text-xs-center">
        <h2>{{ duty.name }}</h2>
      </p>
      </v-card-text>
    </v-card>
    
    <v-list style="background-color: #FAFAFA">
      <v-subheader>
          Duty History
      </v-subheader>
      <v-list-tile avatar v-for="item in this.duty.confirmations" v-bind:key="item.ID">
        <v-list-tile-avatar>
          <v-avatar>
            <img :src="getAvatar(item.UserID)" alt="Avatar">
          </v-avatar>
        </v-list-tile-avatar>
        <v-list-tile-content>
          <v-list-tile-title>{{ item.user }}</v-list-tile-title>
        </v-list-tile-content>
        <v-list-tile-action>
        </v-list-tile-action>
      </v-list-tile>
    </v-list>
    </v-card>

  </v-flex>
</template>

<script>
import { API_URL } from "@/constants.js";

export default {
  data() {
    return {
        duty: {
            name: '',
            confirmations: [],
        }
    };
  },
  created() {
      this.fetch(this.$route.params.projectId, this.$route.params.dutyId)
  },
  methods: {
    fetch(projectId, id) {
      this.$http
        .get(API_URL + "projects/" + projectId + "/duties/" + id)
        .then(
          response => {
              if(response.body.confirmations) {
                for(let confirmation of response.body.confirmations) {
                    for(let user of this.$store.project.users) {
                        if(user.id == confirmation.UserID) {
                            confirmation.user = user.name
                        }
                    }
                }
              }
              this.duty = response.body
          },
          response => {}
        );
    },
    getAvatar(id) {
      return '/static/img/avatars/'+(id % 8 +1)+'.jpg'
    }
  }
};
</script>

<style>
.disabled {
  opacity: 0.5;
}
</style>