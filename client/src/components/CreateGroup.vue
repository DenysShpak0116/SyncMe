<template>
    <div class="login">
        <div class="container">
            <div class="login-wrapper">
                <h2 class = "login-head">
                  Create Group
                </h2>
                <input type="text" class="login-input" v-model="name" placeholder="Group name">
                <input type="text" class="login-input" v-model="description" placeholder="Group description">
                <input type="text" class="login-input" v-model="logo" placeholder="Group logo">
                <input type="text" class="login-input" v-model="image" placeholder="Background image">
                <input type="text" class="login-input" v-model="author" placeholder="Author">
                <div class="login-btn">
                    <button class="login-btn-button" @click="create">
                        Create Group
                    </button>
                </div>
            </div>
        </div>
    </div>
  </template>

<script>
export default {
  name: 'CreateGroupC',
    created(){
    },
  data(){
        return{
           name:"",
           description:"",
           logo:"",
           image:"",
           author:""
        }
  },
  methods:{
    log(a){
        console.log(a)
    },
    async create(){
      let resArr = this.author.split(',')
      try {
        const res = await fetch('https://syncme-server-a6c96ce1c319.herokuapp.com/groups/add', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(
            { 
                name: this.name, 
                description: this.description,
                group_image: this.logo, 
                group_background_image: this.image,
                user_id: +localStorage.getItem('id')
            }
        )
        });
        let result = await res.json()
        for(let link of resArr){
            const res1 = await fetch('https://syncme-server-a6c96ce1c319.herokuapp.com/authors/add', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(
              { 
                  author_link: link, 
                  group_id: result.group_id,
              }
          )
          });
          
          let result1 = await res1.json()
          console.log(result1)
        }
        
        this.$router.push({name:'groups'})
      } catch (error) {
        // Handle error
      }
    }
  },
  computed:{
    
  }
}
</script>