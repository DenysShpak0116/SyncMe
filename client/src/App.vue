<template>
  <div class = "wrapper">
    <Header :logged = "isLogged" @logout = "logout"></Header>
    <main class="main">
      <router-view @reg = "regF"></router-view>
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
  created() {
    this.$store.dispatch('getUserInfo')
    setTimeout(() =>{localStorage.setItem('id',this.$store.getters?.getUserInfo1?.user?.userId)},4000)
    
  },
  data(){
        return{
            isLogged:localStorage.getItem('loginToken'),
        }
  },
  methods:{
    regF(token){
      this.isLogged = token;
    },
    logout(){
      this.isLogged = false;
    }
  }
}
</script>
