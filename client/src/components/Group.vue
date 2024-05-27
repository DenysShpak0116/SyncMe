<template>
    <div class="groupa">
      <div class="container">
        <div class="group-wrapper">
          <div class="group-head">
            <div class="group-head-img">
                <img src="../assets/logouser.jpg" alt="img">
            </div>
            <div class="group-head-info">
                <h2 class="head-name">Famous mangakas</h2>
                <p class="head-text">Hello there! Here you can find posts from your favourite mangakas, let’s enjoy the masterpieces together ^_^ </p>
                <button class="group-follow">Follow the group</button>
            </div>
          </div>
          <div class="group-authors">
            <div class="author">
                <div class="author-img">
                    <img src="../assets/logouser.jpg" alt="img">
                </div>
                <h3 class="author-name">Mikio Ikemoto</h3>
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>

<script>

export default {
  name: 'GroupC',
  data(){
        return{
          
        }
  },
  created() {
    console.log(this.$route.params.id)
  },
  methods:{
    async getGroupInfo() {
      
    },
  }
}
</script>