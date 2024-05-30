<template>
    <div class="groupa">
      <div class="container">
        <div class="group-wrapper" :style = "bgImg">
          <div class="group-head">
            <div class="group-head-img">
                <img :src="authors.group?.GroupImage" alt="img">
            </div>
            <div class="group-head-info">
                <h2 class="head-name">{{ authors.group?.Name }}</h2>
                <p class="head-text">{{ authors.group?.Description }}</p>
                <button class="group-follow">Follow the group</button>
            </div>
          </div>
          <div class="group-authors">
            <div class="author" v-for = "a in authorsARR" :key = "a?.AuthorId" :style = "bgImg1(a)">
                <div class="author-img">
                    <img :src="a.AuthorImage" alt="img">
                </div>
                <h3 class="author-name">{{ a.Name }}</h3>
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
    <PreLoader :isLoading = "load" />
  </template>

<script>
import Post from '../components/Post.vue'
import PreLoader from '../components/PreLoader.vue'

export default {
  name: 'GroupC',
  components: {
    Post,
    PreLoader
  },
  created() {
    if(this.$route.params?.id){
      this.$store.dispatch('getAuthors',this.id)
    }
  },
  data(){
        return{
          
        }
  },
  methods:{
    log(a){
      console.log(a)
    },
    bgImg1(a){
      return {
        backgroundImage: `url(${a.AuthorBackgroundImage})`,
        'background-position': 'center',
      }
    },
  },
  computed:{
    id(){
      return this.$route.params.id
    },
    bgImg(){
      return {
        backgroundImage: `url(${this.authors?.group?.GroupBackgroundImage})`,
        'background-position': 'center',
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
    },
    authorsARR(){
        return this.$store.getters.getAuthors1.authors
    },
    load(){
        return this.$store.getters.getLoad
    },
  }
  }

</script>