<template>
    <div class="post">
      <div class="container">
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
  name: 'MainC',
  components: {
    Post
  },
  data(){
        return{
          logged:localStorage.getItem('loginToken'),
          comment:""
        }
  },
  methods:{
    
  }
}
</script>