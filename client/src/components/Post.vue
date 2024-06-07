<template>
    <div class="post-e">
        <div class="post-head">
            <div class = "post-head-left-block">
              <div class="post-logo-img">
                <img :src="author?.AuthorImage" alt="Logo">
              </div>
              <p class="post-head-name">
                {{ author?.Name }} from {{ author?.SocialMedia  }}
              </p>
            </div>
            <div class = "post-head-right-block">
              <div class="post-head-emotions">
                <p>{{ info.emotionalAnalysis.emotionalState }} %</p>
                <p>{{ info.emotionalAnalysis.emotionalIcon }}</p>
              </div>
              <p class="post-head-date">
                {{ info?.date?.slice(0,10) }}
              </p>
              <div class="post-head-delete">
                <svg width="25" height="25" viewBox="0 0 39 36" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd" clip-rule="evenodd" d="M0.594991 1.56343L1.191 3.12789L7.50016 11.3307L13.8093 19.5346L12.3152 21.3387L10.8211 23.1429L6.03054 28.6539L1.23897 34.165V35.0825V36H4.22209H7.20419L11.9998 30.3429L16.7955 24.6857L17.4364 24.7083L18.0783 24.7299L22.0841 30.3655L26.0898 36H32.5449H39V35.6019V35.2049L31.856 25.7863L24.712 16.3677V15.5643V14.761L30.3252 8.22857L35.9383 1.69611V0.847543V0H32.9715H30.0047L25.7326 5.14286L21.4605 10.2857H20.7522H20.045L16.2352 5.14286L12.4254 0H6.21322H0L0.594991 1.56343ZM20.4073 18.0658L29.8802 30.8571H28.9903H28.1003L18.7519 18.5637L9.40352 6.27017V5.70651V5.14286L10.1689 5.20869L10.9344 5.27451L20.4073 18.0658Z" fill="#5E3B76"/>
                </svg>
              </div>
              <div class="post-head-more">
                <svg width="5" height="20" viewBox="0 0 8 34" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M8 4C8 6.20914 6.20914 8 4 8C1.79086 8 0 6.20914 0 4C0 1.79086 1.79086 0 4 0C6.20914 0 8 1.79086 8 4Z" fill="#744E8E"/>
                <path d="M8 17C8 19.2091 6.20914 21 4 21C1.79086 21 0 19.2091 0 17C0 14.7909 1.79086 13 4 13C6.20914 13 8 14.7909 8 17Z" fill="#744E8E"/>
                <path d="M8 30C8 32.2091 6.20914 34 4 34C1.79086 34 0 32.2091 0 30C0 27.7909 1.79086 26 4 26C6.20914 26 8 27.7909 8 30Z" fill="#744E8E"/>
                </svg>
              </div>
            </div>
          </div>
          <div class="post-img">
              <img v-if="info?.photos" :src="getImg(info?.photos)" alt="img">
          </div>
          <div class="post-text">
              <p>{{ info?.textContent }}</p>
          </div>
          <div class="post-functions">
            <div class="post-like">
              <svg width="40" height="40" viewBox="0 0 50 50" fill="none" xmlns="http://www.w3.org/2000/svg">
              <rect x="20" y="50" width="50" height="10" rx="5" transform="rotate(-90 20 50)" fill="#5E3B76"/>
              <rect y="20" width="50" height="10" rx="5" fill="#5E3B76"/>
              </svg>
              <p class="post-like-count">{{ info?.countOfLikes }}</p>
            </div>
            <div class="post-comment">
              <input type="text" v-model="comment" placeholder="Write a comment...">
              <div class="post-comment-send" @click = "addComment(info?.postId)">
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
            <div class="post-repost">
              <svg width="40" height="40" viewBox="0 0 55 55" fill="none" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
              <rect width="55" height="55" fill="url(#pattern0_312_1105)"/>
              <defs>
              <pattern id="pattern0_312_1105" patternContentUnits="objectBoundingBox" width="1" height="1">
              <use xlink:href="#image0_312_1105" transform="scale(0.015625)"/>
              </pattern>
              <image id="image0_312_1105" width="64" height="64" xlink:href="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAAAAXNSR0IArs4c6QAABW5JREFUeF7tm21sU1UYx5/nrO3efAnDMIMa0QgYPxg/+BJ0QtfOSYIaCal+MAajc7TDuaAxwXbDS9YOjAmIc2uZSgQ1BhZfEhOZg7UdokQN3/iA06jEaIAAi2TQbW3PY+62lr203bk9d+1ebj+1yfP2//U5597n9BZhgb9wgesHA4DRAQucgLEEFngD6LcJblm1q2y4JFLNCW7PB1SGcMZytbh794nXLmnJr8sS2GTbuZoR70SEJVqS621LBOc5Msfe4NZjorGlAdQ5lOvoovk0At4imnQm7QjoH1wcvbu9UxkQySMNwGXfUYVER0SS5coGiaraQp4ekXzyAKwtTyODr0SS5cwGoaa9x/2RSD5pAC9YlZuLmelPBFYkkjAXNjGOL3eE3/xQJJc0ADWJy+arR4A9AKhLPJHCM9nkHIBajNPabEXGngOiZbIChPwRHkZgJals8wJAqGgdjVyVvt8Q8S4DQAoCRgfkehPUsbOFQxlLwNgD8rwJ1tbuNXd0bIoK96zOhnlZAg7HluLFF8uaGbCNRHQjAZ4EpIZAsPFnnfVNGy7nAJ6vfqe0NDr0LUO2enx1RDxKHJ4N9DbmdB7IKYB04hMg8gEhZwDqH1JuiJeauwBwVaa+1BuC4jhkOXu+bx0z4Z1A1LeExw4rYSWWqEELgAa7rzwa549yRCpkeHxPj+dcIk7G4UVUvN6dUFftvY1HsYsh3jMO+snBguGn9h1R/h0ZwAQvgy57y0bgvBWRXa/6EdBlBHK2Bxs/Vz+nBaBVvJ4Q6mze7wBY9dSOo77BgmilCkEEgNPmfYUBvpdiSh2KcbivI+w+nRJAtuL1gKAerg4VRy6kG6050a/DpqitMGbuzTQMMRYvSiN+pEwCaPIH3d4pAGTFy0Kosb59q4XF/86036gQEHBR2kNYouOA8Eim8wkk2NUWcr8+AYBe4uUgELpsLarA5dPeDMgYEGxoD7m/TALQW7wMBKet2YaEXYjMLKMxvS/vbg961gIgjQEgdNlbDiPh4zORcOQSiVCh5Y7Ruca7Hhkc1B8CnSi4El3b+pNyOXkVcNqbn2BU8M1MiE/E5MSPBUKNa7Tk0B/CRPFJAK5K73ZEti19cURA8AMgVqSyUX+RIaB+hrgyXQwiivhD7lK17fIDYar4awCsvs3I8P3UhRERQEOcs4iJ0QepAdDvFgYVQxyCk25ekuYEdMYf9GR1YOqq2rkOOf8CAAq1wLtmm1p8EsCLjylLi+KWUwCwaGICIuJQ7w972mqtO2oyAfCHPMtH45hDALhicqGckxIIe7ZnJwAg++WQXnwSgPpmc6XPzgEOIOLSsSL7CelVf4/nU/WzCADVToVQGDPtR2RVo3FIfe0rv2lFndL5zHC2AFQ/7RAyi58AYETkk0pJwYDlAWBojiP+0nF063+JgkUBJOxddt+9wPkyYPFT/h7lDxnh433FIUwvfgqATEVqBaCX4FRxpt8TxMTPWQBq4ekhiIuf0wBG9gSb90EG+C4R3Q/A+oHF918ou/RWZ+fuiGgHCv+YOZuWwGRxilUxjT8sERU/5ztAi9B0tvOiA2RAGABE6c3mPUBUQyo7owNE6RkdkGkYAn4VCH4UhSllh/gXcf5ZINwUlooz5qzLEtCjEG0xRkd0f9DTqs1vqvUcBaAea/PBCI/d8XFYOSsDQRhAnb3lJSAQevZOpiAtvsRhvT/s/lqLz2RbYQDqeQEhHpVJprcvh7g9EGwKysQVBjAbH4oeMFlWftL9xpWcAFCTqI/FI8UOMWTlMkllfTnxc8i5w9+77XvZWMIdkEg0V/8YIT0MyZKerf6aO2C2Csm2LgNAtuTmi5/RAfPlm8xWx4LvgP8BEXw2bid5SWAAAAAASUVORK5CYII="/>
              </defs>
              </svg>
            </div>
          </div>
          <div class="post-comments">
            <div :class="comentStyle" v-for="comment in displayedComments" :key="comment.CommentId">
              <div class="post-comment-img-block">
                <div class="post-comment-img">
                  <img src="../assets/logouser.jpg" alt="Logo">
                </div>
              </div>
              <div class="post-comment-main">
                <div class="post-comment-head">
                  <p class="post-comment-name">
                    {{comment.user_name}}
                  </p>
                  <p class="post-comment-date">
                    {{ comment.date.slice(0,10) }}
                  </p>
                </div>
                <div class="post-comment-text">
                  <p class="commentText">
                    {{comment.text}}
                  </p>
                </div>
                <div class="post-comment-functions">
                  <div class="post-comment-like">
                    <svg width="25" height="25" viewBox="0 0 50 50" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <rect x="20" y="50" width="50" height="10" rx="5" transform="rotate(-90 20 50)" fill="#5E3B76"/>
                    <rect y="20" width="50" height="10" rx="5" fill="#5E3B76"/>
                    </svg>
                    <p class="comment-likes">0</p>
                  </div>
                  <p class="post-comment-reply">
                    Reply
                  </p>
                </div>
              </div>
              <div class="post-comment-more">
                <svg width="5" height="20" viewBox="0 0 8 34" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M8 4C8 6.20914 6.20914 8 4 8C1.79086 8 0 6.20914 0 4C0 1.79086 1.79086 0 4 0C6.20914 0 8 1.79086 8 4Z" fill="#744E8E"/>
                <path d="M8 17C8 19.2091 6.20914 21 4 21C1.79086 21 0 19.2091 0 17C0 14.7909 1.79086 13 4 13C6.20914 13 8 14.7909 8 17Z" fill="#744E8E"/>
                <path d="M8 30C8 32.2091 6.20914 34 4 34C1.79086 34 0 32.2091 0 30C0 27.7909 1.79086 26 4 26C6.20914 26 8 27.7909 8 30Z" fill="#744E8E"/>
                </svg>
              </div>
            </div>
            <button :class="btnStyle" @click = "more">
              See more ▽
            </button>
          </div>
    </div>
  </template>

<script>
export default {
  name: 'PostC',
  data(){
        return{
          comment:"",
          comentStyle:['post-comment-block'],
          initialCommentsToShow: 1
        }
  },
  props: {
    info:{
      type: Object
    },
    author:{
      type: Object
    }
  },
  methods:{
    more(e){
      this.initialCommentsToShow += 3
      if(this.initialCommentsToShow >= this.info.comments.length){
        e.target.style.display = 'none'
      }
    },
    getImg(arr){
      if(arr == null){
        return null;
      }else{
        return arr[0].url;
      }
    },
    addComment(pId){
      this.$store.dispatch('addComment',{
        post_id:+pId,
        user_id:+this.uInfo.userId,
        text:this.comment,
      })
      this.comment = ""
      this.initialCommentsToShow += 1
    }
  },
  computed:{
    displayedComments() {
      let tempArr = [...this.info.comments]
      return tempArr.reverse().slice(0, this.initialCommentsToShow);
    },
    uInfo(){
      return this.$store.getters?.getUserInfo1?.user
    },
    btnStyle(){
      console.log(this.info.comments.length)
      if(this.info.comments.length == 0){
        return ['nonev','see-more']
      }else{
        return ['see-more']
      }
    }
  }
    
}
</script>