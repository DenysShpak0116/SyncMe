<template>
  <div class="groupa">
    <div class="container">
      <div class="group-wrapper" :style = "bgImg(uInfo)">
        <div class="group-head">
          <div class="group-head-img">
              <img :src = "uInfo?.logo" alt="img">
          </div>
          <div class="group-head-info">
              <h2 class="head-name">{{ uInfo?.username }}</h2>
              <p class="head-text"></p>
              <p class="head-text">Sex: {{ uInfo?.sex }}</p>
              <p class="head-text">Country:  {{ uInfo?.country }}</p>
          </div>
        </div>
      </div>
      <div class="posts-wrapper">
        <Post :author="posts" :info="el" v-for="el in posts.posts" :key = "el.PostId"></Post>
      </div>
    </div>
  </div>
  <PreLoader :isLoading = "load" />
</template>

<script>
import PreLoader from '../components/PreLoader.vue'
import Post from '../components/Post.vue'

export default {
name: 'GroupC',
components: {
  Post,
  PreLoader
},
created() {
  this.$store.dispatch('getUserInfo')
  this.$store.dispatch('getPosts',24)
},
data(){
      return{
        
      }
},
methods:{
  log(a){
    console.log(a)
  },
  bgImg(a){
    return {
      backgroundImage: `url(${a?.bgImage})`,
      'background-position': 'center',
    }
  },
},
computed:{
  id(){
    return this.$route.params.id
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
  uInfo(){
    return this.$store.getters?.getUserInfo1?.user
  },
  load(){
        return this.$store.getters.getLoad
    },
    posts(){
        return this.$store.getters.getPosts1
    },
}
}

</script>