import { createStore } from 'vuex';

export default createStore({
  state:{
    groups: [],
    authors:[],
    userInfo:[],
    load:false,
    name:localStorage.getItem("name"),
    posts:[]
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
    },
    getName(state){
      return state.name
    },
    getPosts1(state){
      return state.posts
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
    setName(state,payload){
      state.name = payload
    },
    setPosts(state,payload){
      state.posts = payload
    }
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
          let res;
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
              res = await response.json()
              commit('setAuthors',res)
          } catch (error) {
              console.error('Ошибка при получении данных:', error);
          }
          try {
            let arr = [];
            if(res.authors){
              for(let el of res.authors){
                const response = await fetch(`http://localhost:3000/authors/get/${el.AuthorId}`, {
                  method: 'GET',
                });
                let result = await response.json()
                arr.push(result)
              }
            }
            commit('setPosts',arr)
            commit('setLoad',false)
          } catch (error) {
              console.error('Ошибка при получении данных:', error);
          }
        },
        async getUserInfo({ commit }) {
          try {
              commit('setLoad',true)
              let token = localStorage.getItem("loginToken");
              const response = await fetch('http://localhost:3000/validate', {
                  method: 'POST',
                  headers: { 'Content-Type': 'application/json' },
                  body: JSON.stringify({ token: token })
              });
              let res = await response.json();
              commit('setUserInfo', res);
          } catch (error) {
              console.error('Error fetching data:', error);
          }
        },
        async getName({commit},payload){
          console.log(payload)
          commit('setName',payload)
        },
        async getPosts({commit},payload){
          commit('setLoad',true)
            const response = await fetch(`http://localhost:3000/authors/get/${payload}`, {
              method: 'GET',
            });
            let result = await response.json()   
            commit('setPosts',result)
            commit('setLoad',false)
        },
  }
})
