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
  created() {
    // this.validateToken();
  },
  methods:{
    async validateToken() {
      try {
        const response = await fetch('http://localhost:3000/validate', {
          method: 'POST',
          credentials: 'include',
        });

        if (!response.ok) {
          throw new Error('Failed to validate token');
        }

        const data = await response.json();
        this.user = data.user;
      } catch (error) {
        console.error('Error:', error);
      }
    },
  }
}
</script>