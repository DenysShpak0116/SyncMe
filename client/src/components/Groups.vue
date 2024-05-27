<template>
    <div class="groups">
        <div class="container">
            <h1 class="groups-head">Search for a new group</h1>
            <div class="groups-wrapper" v-for = "group in groups" :key = group.GroupId>
                {{ log(group) }}
                <div class="group">
                    <div class="group-img">
                        <img :src="group.GroupBackgroundImage" alt="img">
                    </div>
                    <h3 class="group-name">{{group.Name}}</h3>
                    <div class="group-text">{{group.Description}}</div>
                    <router-link
                        :to="{name:'group',params:{id:group.GroupId}}">
                    <button class="view-group-btn">View group</button>
                    </router-link>
                </div>
            </div>
        </div>
    </div>
  </template>

<script>
export default {
  name: 'GroupsC',
    created(){
        this.$store.dispatch('getGroups')
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
    groups(){
        return this.$store.getters.getGroups1
    }
  }
}
</script>