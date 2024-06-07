import { createStore } from 'vuex';

export default createStore({
  state:{
    groups: [],
    authors:[],
    userInfo:[],
    load:false,
    name:localStorage.getItem("name"),
    posts:[],
    msgUsers:[],
    chat:[],
    msg:null
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
    },
    getMsgUsers1(state){
      return state.msgUsers
    },
    getChat1(state){
      return state.chat
    },
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
      if(Array.isArray(payload)){
        for(let el of payload){
          for(let post of el.posts){
            if(post.comments == null){
              post.comments = []
            }
          }
        }
      }else{
        for(let post of payload.posts){
          if(post.comments == null){
            post.comments = []
          }
        }
      }
      
      state.posts = payload
    },
    setMsgUsers(state,payload){
      state.msgUsers = payload
    },
    setChat(state,payload){
      state.chat = payload
    },
    sendMsg(state,payload){
      state.msg = payload
    },
    addComment(state,payload){
      const date = new Date();
      const year = date.getFullYear();
      const month = String(date.getMonth() + 1).padStart(2, '0');
      const day = String(date.getDate()).padStart(2, '0');
      if(Array.isArray(state.posts)){
        for(let o of state.posts){
          for(let post of o.posts){
            if(post.postId == payload.postId){
              post.comments.push({
                date:`${year}-${month}-${day}`,
                text:payload.text,
                user_name:state.name,
              })
            }
          }
        }
      }else{
        for(let post of state.posts.posts){
          if(post.postId == payload.postId){
            post.comments.push({
              date:`${year}-${month}-${day}`,
              text:payload.text,
              user_name:state.name,
            })
          }
        }
      }
    },
  },
  actions:{
        async getGroups({commit}){
            try {
                commit('setLoad',true)
                const response = await fetch('https://syncme-server-a6c96ce1c319.herokuapp.com/groups/get', {
                    method: 'GET',
                });
                let res = await response.json()
                commit('setGroups',res)
                commit('setLoad',false)
            } catch (error) {
                console.error('Помилка при отриманні даних:', error);
            }
        },
        async getAuthors({commit},payload){
          let res;
          try {
                commit('setLoad',true)
              const response = await fetch('https://syncme-server-a6c96ce1c319.herokuapp.com/authors/get', {
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
              console.error('Помилка при отриманні даних:', error);
          }
          try {
            let arr = [];
            if(res.authors){
              for(let el of res.authors){
                const response = await fetch(`https://syncme-server-a6c96ce1c319.herokuapp.com/authors/get/${el.AuthorId}`, {
                  method: 'GET',
                });
                let result = await response.json()
                arr.push(result)
              }
            }
            commit('setPosts',arr)
            commit('setLoad',false)
          } catch (error) {
              console.error('Помилка при отриманні даних:', error);
          }
        },
        async getUserInfo({ commit }) {
          try {
              commit('setLoad',true)
              let token = localStorage.getItem("loginToken");
              const response = await fetch('https://syncme-server-a6c96ce1c319.herokuapp.com/validate', {
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
            const response = await fetch(`https://syncme-server-a6c96ce1c319.herokuapp.com/authors/get/${payload}`, {
              method: 'GET',
            });
            let result = await response.json()   
            commit('setPosts',result)
            commit('setLoad',false)
        },
        async getMsgUsers({commit},payload){
            let token = localStorage.getItem("loginToken");
            const res = await fetch('https://syncme-server-a6c96ce1c319.herokuapp.com/validate', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ token: token })
            });
            let f = await res.json();
            commit('setUserInfo', f);

            const response = await fetch(`https://syncme-server-a6c96ce1c319.herokuapp.com/messages/chats/${payload}`, {
              method: 'GET',
            });
            let result = await response.json()   
            commit('setMsgUsers',result)
        },
        async getChat({commit},payload){
          const response = await fetch(`https://syncme-server-a6c96ce1c319.herokuapp.com/messages/get`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ 
              disscusser_id:payload.disscusser_id,
              current_user_id:payload.current_user_id,
             })
          });
          let result = await response.json()   
          commit('setChat',result)
      },
      async sendMsg({commit},payload){
        await fetch(`https://syncme-server-a6c96ce1c319.herokuapp.com/messages/add`, {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ 
            message_text:payload.message_text,
            sent_at:payload.sent_at,
            user_from_id:+payload.user_from_id,
            user_to_id:+payload.user_to_id,
           })
        });
        commit("sendMsg",1)
        const response = await fetch(`https://syncme-server-a6c96ce1c319.herokuapp.com/messages/get`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ 
              disscusser_id:+payload.user_to_id,
              current_user_id:+payload.user_from_id,
             })
          });
          let result = await response.json()   
          commit('setChat',result)
    },
    async addComment({commit},payload){
      await fetch(`https://syncme-server-a6c96ce1c319.herokuapp.com/authors/addcomment`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ 
          post_id:+payload.post_id,
          user_id:+payload.user_id,
          text:payload.text,
         })
      });
      commit("addComment",{
        postId:+payload.post_id,
        text:payload.text,
      })
  },
  }
})
