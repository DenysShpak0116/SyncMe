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
          <div class="author" v-for = "a in authorsARR" :key = "a?.AuthorId">
              <div class="author-img">
                  <img :src="a.AuthorImage" alt="img">
              </div>
              <h3 class="author-name">{{ a.Name }}</h3>
          </div>
        </div>
      </div>
      <div class="posts-wrapper">
        <Post></Post>
        <Post></Post>{{ log(uInfo) }}
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
  this.$store.dispatch('getUserInfo')
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
  uInfo(){
    return this.$store.getters.getUserInfo1
  }
}
}

</script>