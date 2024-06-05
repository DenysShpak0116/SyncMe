<template>
  <div v-if="visible" class="notification">
    <p>{{ message }}</p>
    <button @click="hideNotification">Close</button>
  </div>
</template>

<script>
export default {
  name: 'NotificationC',
  props: {
    message: {
      type: String,
      default: ''
    },
    duration: {
      type: Number,
      default: 3000 // Продолжительность отображения уведомления в миллисекундах
    }
  },
  data() {
    return {
      visible: false
    };
  },
  methods: {
    showNotification() {
      this.visible = true;
      setTimeout(this.hideNotification, this.duration);
    },
    hideNotification() {
      this.visible = false;
    }
  },
  watch: {
    message(newMessage) {
      if (newMessage) {
        this.showNotification();
      }
    }
  }
};
</script>

<style scoped>
.notification {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  padding: 20px 40px;
  background-color: #f44336; /* Красный цвет для ошибки */
  color: white;
  border-radius: 4px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  z-index: 10000;
  display: flex;
  align-items: center;
  justify-content: center;
}
.notification button {
  background: none;
  border: none;
  color: white;
  margin-left: 20px;
  cursor: pointer;
}
</style>
