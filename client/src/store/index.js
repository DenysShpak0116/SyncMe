import { createStore } from 'vuex';

export default createStore({
  state:{
    groups: [],
    authors:[],
    userInfo:[],
    load:false
  },
  getters:{
    getGroups1(state){
        return state.groups
    },
    getAuthors1(state){
      return state.authors
    },
    getUserInfo1(state){
      return state.userInfo
    },
    getLoad(state){
      return state.load
    }
  },
  mutations: {
    setGroups(state,payload){
        state.groups = payload
    },
    setAuthors(state,payload){
      state.authors = payload
    },
    setUserInfo(state,payload){
      state.userInfo = payload
    },
    setLoad(state,payload){
      state.load = payload
    },
  },
  actions:{
        async getGroups({commit}){
            try {
                commit('setLoad',true)
                const response = await fetch('http://localhost:3000/groups/get', {
                    method: 'GET',
                });
                let res = await response.json()
                commit('setGroups',res)
                commit('setLoad',false)
            } catch (error) {
                console.error('Ошибка при получении данных:', error);
            }
        },
        async getAuthors({commit},payload){
          try {
                commit('setLoad',true)
              const response = await fetch('http://localhost:3000/authors/get', {
                  method: 'POST',
                  headers: { 'Content-Type': 'application/json' },
                  body: JSON.stringify(
                    { 
                      group_id: +payload,
                    })
              });
              let res = await response.json()
              commit('setAuthors',res)
              commit('setLoad',false)
          } catch (error) {
              console.error('Ошибка при получении данных:', error);
          }
        },
        async getUserInfo({commit}){
          try {
              let token = localStorage.getItem("loginToken")
              const response = await fetch('http://localhost:3000/validate', {
                  method: 'POST',
                  headers: { 'Content-Type': 'application/json' },
                  body: JSON.stringify(
                    { 
                      token
                    })
              });
              let res = await response.json()
              commit('setUserInfo',res)
          } catch (error) {
              console.error('Ошибка при получении данных:', error);
          }
        },
  }
})
