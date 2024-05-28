<template>
    <div class="groupa">
      <div class="container">
        <div class="group-wrapper" :style = "bgImg">
          <div class="group-head">
            <div class="group-head-img">
                <img :src="groups.GroupImage" alt="img">
            </div>
            <div class="group-head-info">
                <h2 class="head-name">{{ groups.Name }}</h2>
                <p class="head-text">{{ groups.Description }}</p>
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
        <div class="posts-wrapper">
          <Post></Post>
          <Post></Post>
          <Post></Post>
        </div>
      </div>
    </div>
  </template>

<script>
import Post from '../components/Post.vue'

export default {
  name: 'GroupC',
  components: {
    Post
  },
  created() {
    this.$store.dispatch('getAuthors',this.$route.params.id)
  },
  data(){
        return{
          
        }
  },
  methods:{
    log(a){
      console.log(a)
    }
  },
  computed:{
    id(){
      return this.$route.params.id
    },
    bgImg(){
      return {
        backgroundImage: `url(${this.groups.GroupBackgroundImage})`,
        'background-position': 'center'
      }
    },
    groups(){
        let arr=this.$store.getters.getGroups1
        for(let a of arr){
          if(a.GroupId == this.$route.params.id){
            return a
          }
        }
        return this.$store.getters.getGroups1
    },
    authors(){
        return this.$store.getters.getAuthors1
    }
  }
  }

</script>