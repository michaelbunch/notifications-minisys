<template>
  <div>
    <h1 class="text-5xl">Notifications</h1>
    <div>
      <div v-for="notification in notifications" class="flex mb-2">
        <div class="flex-none bg-green-600 text-white px-2 text-sm rounded">{{ notification.domain }}</div>
        <div class="flex-1 ml-2">{{ notification.subject }}</div>
        <div class="flex-1 ml-2">Performed by {{ notification.actor_name }}</div>
        <div class="flex-1 ml-2">{{ notification.createdAt }}</div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
export default {
  data() {
    return {
      user_id: 1,
      notifications: []
    }
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
    }
  },
  mounted() {
    this.getNotifications();
  }
}
</script>