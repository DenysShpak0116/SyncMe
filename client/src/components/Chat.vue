<template>
    <div class="chat">
        <div class="container">
            <div class="chat-wrapper">
                <div class="chat-left-block">
                    <div class="chat-search-block">
                        <div class="chat-search">
                            <input type="text" v-model="chatSearch">
                            <div class="chat-search-icon">
                                <svg width="30" height="30" viewBox="0 0 35 35" fill="none" xmlns="http://www.w3.org/2000/svg">
                                    <path d="M14.6998 2.1001C8.1209 2.1001 2.7998 7.42119 2.7998 14.0001C2.7998 20.579 8.1209 25.9001 14.6998 25.9001C17.0486 25.9001 19.2225 25.211 21.0654 24.0407L29.6623 32.6376L32.6373 29.6626L24.1498 21.197C25.6756 19.1954 26.5998 16.7153 26.5998 14.0001C26.5998 7.42119 21.2787 2.1001 14.6998 2.1001ZM14.6998 4.9001C19.7393 4.9001 23.7998 8.96064 23.7998 14.0001C23.7998 19.0396 19.7393 23.1001 14.6998 23.1001C9.66035 23.1001 5.5998 19.0396 5.5998 14.0001C5.5998 8.96064 9.66035 4.9001 14.6998 4.9001Z" fill="#794D98"/>
                                </svg>
                        </div>
                    </div>
                    </div>
                    <div class="chat-users">
                        <div class="chat-user" @click = "chatClick(el.UserId,el.UserName)" v-for="el in chats.chats" :key="el.UserId">
                            <div class="chat-user-img">
                                <img src="../assets/logouser.jpg" alt="Logo">
                            </div>
                            <div class="chat-user-info">
                                <p class="chat-user-name">
                                    {{ el.UserName }}
                                </p>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="chat-right-block">
                    <div class="chat-right-head">
                        <div class="chat-user-img">
                            <img src="../assets/logouser.jpg" alt="Logo">
                        </div>
                        <p class="chat-user-name">
                            {{ name }}
                        </p>
                        <div class="chat-right-head-more">
                            <svg xmlns="http://www.w3.org/2000/svg" width="10" height="42" viewBox="0 0 13 48" fill="none">
                            <path d="M6.5 12.5588C9.81622 12.5588 12.5 9.85685 12.5 6.52941C12.5 3.20198 9.81622 0.5 6.5 0.5C3.18378 0.5 0.5 3.20198 0.5 6.52941C0.5 9.85685 3.18378 12.5588 6.5 12.5588Z" fill="#744E8E" stroke="#0D0811"/>
                            <path d="M6.5 29.5588C9.81622 29.5588 12.5 26.8568 12.5 23.5294C12.5 20.202 9.81622 17.5 6.5 17.5C3.18378 17.5 0.5 20.202 0.5 23.5294C0.5 26.8568 3.18378 29.5588 6.5 29.5588Z" fill="#744E8E" stroke="#0D0811"/>
                            <path d="M6.5 46.5588C9.81622 46.5588 12.5 43.8568 12.5 40.5294C12.5 37.202 9.81622 34.5 6.5 34.5C3.18378 34.5 0.5 37.202 0.5 40.5294C0.5 43.8568 3.18378 46.5588 6.5 46.5588Z" fill="#744E8E" stroke="#0D0811"/>
                            </svg>
                        </div>
                    </div>
                    <div class="chat-right-main">
                        <div class="chat-right-messages" ref="scroll">
                            <div :class="getClass(e)" v-for="e in chat.chat" :key="e.MessageId">
                                <div class="chat-right-message">
                                    <p class="ctext">{{ e.Text }}</p>
                                    <p class="time">{{ e.SentAt }}</p>
                                </div>
                            </div>
                        </div>
                        <div class="chat-right-func">
                            <div class="file-block">
                                <svg width="30" height="30" viewBox="0 0 36 36" fill="none" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
                                <rect width="36" height="36" fill="url(#pattern0_92_3106)"/>
                                <defs>
                                <pattern id="pattern0_92_3106" patternContentUnits="objectBoundingBox" width="1" height="1">
                                <use xlink:href="#image0_92_3106" transform="scale(0.0104167)"/>
                                </pattern>
                                <image id="image0_92_3106" width="96" height="96" xlink:href="data:image/png;base64,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"/>
                                </defs>
                                </svg>
                            </div>
                            <div class="message-input-block">
                                <input type="text" v-model="message" placeholder="Write a message...">
                                <div class="send-svg" @click = "sendMsg()">
                                    <svg width="26" height="26" viewBox="0 0 26 26" fill="none" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
                                    <rect width="26" height="26" fill="url(#pattern0_312_1104)"/>
                                    <defs>
                                    <pattern id="pattern0_312_1104" patternContentUnits="objectBoundingBox" width="1" height="1">
                                    <use xlink:href="#image0_312_1104" transform="scale(0.0104167)"/>
                                    </pattern>
                                    <image id="image0_312_1104" width="96" height="96" xlink:href="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAGAAAABgCAYAAADimHc4AAAACXBIWXMAAAsTAAALEwEAmpwYAAAC50lEQVR4nO2dTW7UMBiGfQRgDRtaxBVgx5pbAIepVPE5cAyKxPTzlE4PgCo09oDELYDhp+xAQoM8C6gqVymJ7c+O30fKbkaK3zd5konisVIAAAAAAAAAEODVvtthci8N2W+s3Wej3YHpVvcRVgbmz97dMeTWRrvNxY3JnqCIxGyP/ED4KCITrO3XvgL+buTezDv7INe+NQFr9+XKBUBN8WHtXv9vASgiIoZWj4YWADVFYNa9v2a0/Tm6BFwjhsPkjqMUgCIENYQiCtEQikilIft7bBGMX9aXc9gtH/ccxb/mZB8abec4I4Q05Evyn2Vy91BEAoy2ix4NLc5/HkUIaOiETm9c/B6KENBQCBQhoKEQKEJAQyEYF+v8GkIRhWgoBOOMyK+hEChCQEMhUISAhkKgCAENhUAR5zja+3A9h4ZCoAgBDYVovggJDYVotghJDYVosghpDanWiyhFQ80WUZqGmiyiRA01VUTJGmqiiBo0NPkiatHQZIvo1VC67QeTmx12y7tjx+CnW21fDhu9T3bh59WpnHjHe9cLlbDxM3lmT5e3YowlRhF+YgvTajfG/kTUUPLtIOp4xhZB9oVqREOb7VFH7izFuIYW4c8C1VIBRrvvKcY1+Ewgt06xP+UqiOKe8lUpSPoibMitj5+/vRljLFHuhsits94JSemHyZ35Iy1G+DFvQ4/08rbKSd+O+wkeqlC49h9iV3kU4eeZqcLg2oP/NxD7pGcH8TAuJbXoh6dyxNemH55i8DXoh6ccfMn64RaCL1E/3ErwpemHWwu+FP1wq8FL64dbDl5SP4zgZfTDCF5GP4zgZfTDCF5GP4zgZfTDCF5GP4zgZfSD4IX0g+CF9IPghfSD4IX0g+CF9IO/rUwI/ri1dP20+li46tcOCcHLvPWM4IX0g+CF9IPgh4NFfITBMlbCYCE3YbYLd/Y4HitfJMRPNsZinsL4yWZ+HpbR9iOT/YTlbAEAAAAAAABADeYPnu7lpSKou+IAAAAASUVORK5CYII="/>
                                    </defs>
                                    </svg>
                                </div>
                            </div>
                            <div class="emoji-block">
                                    <img src="../assets/smile.png" alt="emoji">
                                </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
  </template>

<script>
import { mapState } from 'vuex';
export default {
  name: 'ChatC',
  created() {
    this.$store.dispatch('getMsgUsers',this.uInfo.userId)
  },
  data(){
        return{
            chatSearch:"",
            message:"",
            left:"msg-left",
            rigtht:"msg-right",
            name:"User",
            timeId:null
        }
  },
  watch: {
    messages(newMessages, oldMessages) {
      if (newMessages.length !== oldMessages.length) {
        this.$nextTick(() => {
            this.scrollToElement()
        });
      }
    }
  },
  unmounted(){
    clearInterval(this.timeId);
  },
  methods:{
    chatClick(e,a){
        clearInterval(this.timeId);
        localStorage.setItem("discusser",e)
        this.$store.dispatch('getChat',{
            disscusser_id:e,
            current_user_id:this.uInfo.userId,
        }) 
        this.name = a
        this.timeId = setInterval(() => {
            console.log("sadsadasdas")
            this.$store.dispatch('getChat',{
            disscusser_id:e,
            current_user_id:this.uInfo.userId,
            }) 
        }, 5000);
    },
    sendMsg(){
        if(this.message != ""){
            this.$store.dispatch('sendMsg',{
                message_text:this.message,
                sent_at:new Date(),
                user_from_id:this.uInfo.userId,
                user_to_id:localStorage.getItem("discusser")
            }) 
            this.message = ""
            setTimeout(() =>{this.scrollToElement()},2000)
        }
    },
    getClass(e){
        if(e.UserFromId == localStorage.getItem("discusser")){
            return ['chat-message-block','msg-left']
        }else{
            return ['chat-message-block','msg-right']
        }
    },
    scrollToElement() {
        const el = this.$refs.scroll;
        if (el) {
            el.scrollTop = el.scrollHeight;
        }
    }
  },
  computed:{
    ...mapState({
      messages: state => state.chat
    }),
  uInfo(){
    return this.$store.getters?.getUserInfo1?.user
  },
  chats(){
    return this.$store.getters?.getMsgUsers1
  },
  chat(){
    return this.$store.getters?.getChat1
  },
}
}
</script>