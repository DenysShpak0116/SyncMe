<template>
    <div class="container">
        <div class="notifications-container">
            <h2>Notifications</h2>
            <div class="notification" v-for="el of notifications" :key="el.NotificationId">
                <p class="notification-text">{{el.Text}}</p>
                <p class="notification-date">{{el.Date.slice(0,10)}}</p>
            </div>
        </div>
    </div>
  </template>

<script>
export default {
  name: 'NotificationsC',
  async created(){
    const response = await fetch(`https://syncme-server-a6c96ce1c319.herokuapp.com/notifications/${this.uInfo}`, {
      method: 'GET',
    });
    let result = await response.json() 
    this.notifications = result
  },
  data(){
        return{
            notifications:[]
        }
  },
  methods:{
    
  },
  computed:{
    uInfo(){
      return this.$store.getters?.getUserInfo1?.user.userId
    },
  }
    
}
</script>