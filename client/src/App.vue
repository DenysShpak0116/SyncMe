<template>
  <div class = "wrapper">
    <Header :logged = "isLogged" @logout = "logout"></Header>
    <main class="main">
      <router-view @reg = "reg"></router-view>
    </main>
    <Footer></Footer>
  </div>
</template>

<script>
import Header from './components/Header.vue'
import Footer from './components/Footer.vue'
export default {
  name: 'App',
  components: {
    Header,
    Footer,
  },
  data(){
        return{
            isLogged:localStorage.getItem('loginToken'),
        }
  },
  methods:{
    reg(token){
      this.isLogged = token;
    },
    logout(){
      this.isLogged = false;
    }
  }
}
</script>
