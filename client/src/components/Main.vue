<template>
    <div class="post">
      <div class="container">
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
  name: 'MainC',
  components: {
    Post,
    PreLoader
  },
  data(){
        return{
          logged:localStorage.getItem('loginToken'),
          comment:""
        }
  },
  created() {
    this.$store.dispatch('getPosts',9)
  },
  methods:{
    async validateToken() {
      try {
        const response = await fetch('https://syncme-server-a6c96ce1c319.herokuapp.com/validate', {
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
  },
  computed:{
    posts(){
        return this.$store.getters.getPosts1
    },
    load(){
        return this.$store.getters.getLoad
    },
  }
}
</script>