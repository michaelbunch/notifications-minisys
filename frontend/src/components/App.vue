<template>
  <div class="w-6/12">
    <h1 class="text-5xl">Notifications</h1>
    <div>
      <div v-for="notification in notifications" class="flex mb-2 bg-gray-100 rounded">
        <div class="flex-none bg-green-600 text-white px-2 py-1 text-sm rounded">{{ notification.domain }}</div>
        <div class="flex-1 ml-2 text-sm py-1">{{ notification.subject }}</div>
        <div class="flex-1 ml-2 text-sm py-1 pr-2 text-right">
          Performed by {{ notification.actor_name }} on
          {{ formatDate(notification.createdAt) }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import { DateTime } from "luxon";
export default {
  data() {
    return {
      user_id: 1,
      notifications: [],
      pollInterval: 10000
    }
  },
  computed:{

  },
  methods: {
    async getNotifications() {
      try {
        const response = await axios.get('/api/notifications', {
          params: {
            user_id: this.user_id
          }
        });
        this.notifications = response.data;
      } catch (error) {
        console.error(error);
      }
    },
    pollNotifications () {
      this.polling = setInterval(() => {
        this.getNotifications()
      }, this.pollInterval)
    },
    formatDate(datetime) {
      return DateTime.fromSQL(datetime).toFormat('LLL d, yyyy @ t');
    }
  },
  mounted() {
    this.getNotifications();
    this.pollNotifications();
  }
}
</script>