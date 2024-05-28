import { createStore } from 'vuex';

export default createStore({
  state:{
    groups: [],
    authors:[]
  },
  getters:{
    getGroups1(state){
        return state.groups
    },
    getAuthors1(state){
      return state.groups
  }
  },
  mutations: {
    setGroups(state,payload){
        state.groups = payload
    },
    setAuthors(state,payload){
      state.authors = payload
  },
  },
  actions:{
        async getGroups({commit}){
            try {
                const response = await fetch('http://localhost:3000/groups/get', {
                    method: 'GET',
                });
                let res = await response.json()
                commit('setGroups',res)
            } catch (error) {
                console.error('Ошибка при получении данных:', error);
            }
        },
        async getAuthors({commit},payload){
          try {
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
          } catch (error) {
              console.error('Ошибка при получении данных:', error);
          }
      },
  }
})
