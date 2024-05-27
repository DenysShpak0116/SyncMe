import { createStore } from 'vuex';

export default createStore({
  state:{
    groups: []
  },
  getters:{
    getGroups1(state){
        return state.groups
    }
  },
  mutations: {
    setGroups(state,payload){
        state.groups = payload
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
        }
  }
})
