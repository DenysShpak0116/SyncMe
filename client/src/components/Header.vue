<script>
export default {
  name: 'HeaderC',
  data() {
    return {
        username: null,
        headerSearch:"",
        loggedIn:this.logged
    };
  },
  props: {
    logged: {
            default: localStorage.getItem('loginToken')
        }
    },
  methods: {
    logout() {
        localStorage.removeItem("loginToken");
        this.$emit('logout')
        this.$router.push({name:'login'})
    }
  },
}
    
</script>
<template>
    <header class = "header">
        <div class="container">
            <div class="header-wrapper">
                <router-link
                    :to="{name:'main'}">
                    <div class="header-logo">
                        <img src="../assets/logo.png" alt="Logo" class="header-logo-img">
                    </div>
                </router-link>
                <div class="header-search">
                    <input class = "header-input" type="text" v-model="headerSearch">
                    <div class="header-search-icon">
                        <svg class = "search-svg" viewBox="0 0 35 35" fill="none" xmlns="http://www.w3.org/2000/svg">
                            <path d="M14.6998 2.1001C8.1209 2.1001 2.7998 7.42119 2.7998 14.0001C2.7998 20.579 8.1209 25.9001 14.6998 25.9001C17.0486 25.9001 19.2225 25.211 21.0654 24.0407L29.6623 32.6376L32.6373 29.6626L24.1498 21.197C25.6756 19.1954 26.5998 16.7153 26.5998 14.0001C26.5998 7.42119 21.2787 2.1001 14.6998 2.1001ZM14.6998 4.9001C19.7393 4.9001 23.7998 8.96064 23.7998 14.0001C23.7998 19.0396 19.7393 23.1001 14.6998 23.1001C9.66035 23.1001 5.5998 19.0396 5.5998 14.0001C5.5998 8.96064 9.66035 4.9001 14.6998 4.9001Z" fill="#794D98"/>
                        </svg>
                    </div>
                </div>
                <div class="header-menu">
                    <router-link
                        :to="{name:'groups'}">
                        <div class="header-login header-login1">
                            <svg class="header-login-svg" viewBox="0 0 61 61" fill="none" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
                            <rect x="0.132935" y="6.20288" width="55" height="55" transform="rotate(-6.3363 0.132935 6.20288)" fill="url(#pattern0_305_691)"/>
                            <defs>
                            <pattern id="pattern0_305_691" patternContentUnits="objectBoundingBox" width="1" height="1">
                            <use xlink:href="#image0_305_691" transform="scale(0.015625)"/>
                            </pattern>
                            <image id="image0_305_691" width="64" height="64" xlink:href="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAAACXBIWXMAAAsTAAALEwEAmpwYAAADoElEQVR4nO2aTYgdVRCFr4ooURBRCAZU1BDfq+pEcBRcRB99q2dMdPxL1+k3EzTqwigKLoxiIEI0unEhUUE3GhHc+AfiwqAu3PgTEMEYjRDxP4KgohiMikEjt/NmEcHB5D3n9IQ+cJfV3fX16W7u6QqhVatWrWZR3q0uM8FuU+yfsyXYHbvlytAE2Vw3P1hR/OvQBBmh+ZkVmiBrAaB1gLWPANp3gLUvQQzzSXs3qm8y9VtNcV8U32KKrVH9w6j44Yj+CkTB3SGEo2Y7R+/M64/PtX9O7JYX54JpE9xhgs2hCbLh7vyLMw1G8XVR/JEEJFdcF7XKJ7JVnfFl154QmiwbxgHd8qp0jCh4bFZQ6j+bYqcJXo/qT5vggSi4LRdcaVpdWJw7tWhj2Hj0vANQCM4YAPhs6JeoYJ8Jvonq20z9pah4tBCsT27KBbHQfrcnOLExAKL6n2Nja49Nz7+p/zaHX5E9Ufzj9PhZp7+EBsAU36b63pLpU+ew+X/cBLzPAyB4L9XnWp3HAmCCX5kOeDnVFxkupwFQfEkDEAWP1wAENxEd8DYPgGLDoP5eogOeZzrghrpe8ATRAZt5ADKMD+q3sgDkijuZDpCBAz6gARBM0wAUYzgp1c+22/u/V5GhRwEQxX+Z2QSZ4i8WgJ5gMQeA+q66dumqs1nNpzU5NrmAAsDE30i1aX9PAyD+49DNJx3myZ9JtVF9igUgpU2BBwAP1gDE19EAiL/KA5Dh9hqA4iEigC08AFqVNQDBszQA6ptoAIpOdVENQPEm0QE30wBMaP/0GoDgc54DqisoACIvCjv4OjrV+RQA1oAorAaQTS3kAJAmRGH+x8hidDvMKKwerWHdfcVXI2k+aT5GYVHwThiV7JAvwO9pQBT2Ag1ALn5jA6Kwh3kAtJw4UOevEB1wFw9At9S6Tnw7C0AhWE2bE1y+dPXJqS4KvmMBSL/eRwYgdsuVaWjxP51YsDfVrFi84jhmFDaSn6LDaHwZzqI1r9hPH7ooBMtpAAQ/BbZM0Sc64CN2/6EecuI54DV2/4EZhZniKXb/gRqFCe5n9x+oUViGW9j9B2YUlkbrmuCAT1gAiqy8gN1/MPU19VwfAcClnWtOC02Qda4+Jf0bTBuTmdnfKP5cFH/L1L+Iit9HDkCwD8AxYb4oZlMLD2SH5aQJ1tYhiuDJtJWOgh2m/v0hvQAFO8KRpt6/TJKnT+1BbhL/NO9Wl7Cvt1WrVmFe6W8UnoaEQyf8NQAAAABJRU5ErkJggg=="/>
                            </defs>
                            </svg>
                        </div>
                    </router-link>
                    <router-link
                        v-if="!this.logged"
                        :to="{name:'login'}">
                        <div class="header-login">
                            <svg class="header-login-svg" viewBox="0 0 168 64" fill="none" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
                                <rect x="104" width="64" height="64" fill="url(#pattern0_1_79)"/>
                                <path d="M3.47319 45V18.8182H6.64364V42.1875H18.8141V45H3.47319ZM31.2914 45.4091C29.5186 45.4091 27.9632 44.9872 26.6252 44.1435C25.2956 43.2997 24.2559 42.1193 23.5059 40.6023C22.7644 39.0852 22.3936 37.3125 22.3936 35.2841C22.3936 33.2386 22.7644 31.4531 23.5059 29.9276C24.2559 28.402 25.2956 27.2173 26.6252 26.3736C27.9632 25.5298 29.5186 25.108 31.2914 25.108C33.0641 25.108 34.6152 25.5298 35.9448 26.3736C37.2828 27.2173 38.3226 28.402 39.0641 29.9276C39.8141 31.4531 40.1891 33.2386 40.1891 35.2841C40.1891 37.3125 39.8141 39.0852 39.0641 40.6023C38.3226 42.1193 37.2828 43.2997 35.9448 44.1435C34.6152 44.9872 33.0641 45.4091 31.2914 45.4091ZM31.2914 42.6989C32.638 42.6989 33.7459 42.3537 34.6152 41.6634C35.4846 40.973 36.128 40.0653 36.5456 38.9403C36.9632 37.8153 37.1721 36.5966 37.1721 35.2841C37.1721 33.9716 36.9632 32.7486 36.5456 31.6151C36.128 30.4815 35.4846 29.5653 34.6152 28.8665C33.7459 28.1676 32.638 27.8182 31.2914 27.8182C29.9448 27.8182 28.8368 28.1676 27.9675 28.8665C27.0982 29.5653 26.4547 30.4815 26.0371 31.6151C25.6195 32.7486 25.4107 33.9716 25.4107 35.2841C25.4107 36.5966 25.6195 37.8153 26.0371 38.9403C26.4547 40.0653 27.0982 40.973 27.9675 41.6634C28.8368 42.3537 29.9448 42.6989 31.2914 42.6989ZM52.7207 52.7727C51.2633 52.7727 50.0105 52.5852 48.9622 52.2102C47.9139 51.8438 47.0403 51.358 46.3414 50.7528C45.6511 50.1563 45.1014 49.517 44.6923 48.8352L47.0957 47.1477C47.3684 47.5057 47.7136 47.9148 48.1312 48.375C48.5488 48.8438 49.1199 49.2486 49.8443 49.5895C50.5772 49.9389 51.536 50.1136 52.7207 50.1136C54.3059 50.1136 55.6142 49.7301 56.6454 48.9631C57.6767 48.196 58.1923 46.9943 58.1923 45.358V41.3693H57.9366C57.715 41.7273 57.3997 42.1705 56.9906 42.6989C56.59 43.2188 56.0105 43.6832 55.252 44.0923C54.502 44.4929 53.4877 44.6932 52.2093 44.6932C50.6241 44.6932 49.2008 44.3182 47.9395 43.5682C46.6866 42.8182 45.6937 41.7273 44.9608 40.2955C44.2363 38.8636 43.8741 37.125 43.8741 35.0795C43.8741 33.0682 44.2278 31.3168 44.9352 29.8253C45.6426 28.3253 46.627 27.1662 47.8883 26.348C49.1497 25.5213 50.6071 25.108 52.2605 25.108C53.5389 25.108 54.5531 25.321 55.3031 25.7472C56.0616 26.1648 56.6412 26.642 57.0417 27.179C57.4508 27.7074 57.7662 28.142 57.9877 28.483H58.2946V25.3636H61.2093V45.5625C61.2093 47.25 60.8258 48.6222 60.0588 49.679C59.3002 50.7443 58.2775 51.5241 56.9906 52.0185C55.7122 52.5213 54.2889 52.7727 52.7207 52.7727ZM52.6184 41.983C53.8287 41.983 54.8514 41.706 55.6866 41.152C56.5218 40.598 57.1568 39.8011 57.5914 38.7614C58.0261 37.7216 58.2434 36.4773 58.2434 35.0284C58.2434 33.6136 58.0304 32.3651 57.6042 31.2827C57.1781 30.2003 56.5474 29.3523 55.7122 28.7386C54.877 28.125 53.8457 27.8182 52.6184 27.8182C51.34 27.8182 50.2747 28.142 49.4224 28.7898C48.5787 29.4375 47.9437 30.3068 47.5176 31.3977C47.1 32.4886 46.8912 33.6989 46.8912 35.0284C46.8912 36.392 47.1042 37.598 47.5304 38.6463C47.965 39.6861 48.6042 40.5043 49.448 41.1009C50.3002 41.6889 51.3571 41.983 52.6184 41.983ZM76.8571 45V25.3636H79.8741V45H76.8571ZM78.3912 22.0909C77.8031 22.0909 77.296 21.8906 76.8699 21.4901C76.4522 21.0895 76.2434 20.608 76.2434 20.0455C76.2434 19.483 76.4522 19.0014 76.8699 18.6009C77.296 18.2003 77.8031 18 78.3912 18C78.9792 18 79.4821 18.2003 79.8997 18.6009C80.3258 19.0014 80.5389 19.483 80.5389 20.0455C80.5389 20.608 80.3258 21.0895 79.8997 21.4901C79.4821 21.8906 78.9792 22.0909 78.3912 22.0909ZM88.4171 33.1875V45H85.4V25.3636H88.3148V28.4318H88.5705C89.0307 27.4347 89.7296 26.6335 90.6671 26.0284C91.6046 25.4148 92.8148 25.108 94.2978 25.108C95.6273 25.108 96.7907 25.3807 97.7878 25.9261C98.785 26.4631 99.5605 27.2812 100.115 28.3807C100.669 29.4716 100.945 30.8523 100.945 32.5227V45H97.9284V32.7273C97.9284 31.1847 97.5279 29.983 96.7267 29.1222C95.9256 28.2528 94.8262 27.8182 93.4284 27.8182C92.4654 27.8182 91.6046 28.027 90.8461 28.4446C90.0961 28.8622 89.5037 29.4716 89.0691 30.2727C88.6344 31.0739 88.4171 32.0455 88.4171 33.1875Z" fill="#503365"/>
                                <defs>
                                <pattern id="pattern0_1_79" patternContentUnits="objectBoundingBox" width="1" height="1">
                                <use xlink:href="#image0_1_79" transform="scale(0.015625)"/>
                                </pattern>
                                <image id="image0_1_79" width="64" height="64" xlink:href="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAAACXBIWXMAAAsTAAALEwEAmpwYAAACMElEQVR4nO2aO28UQQzHp6ABOr4ACCGRtS8pCAlHgZazL1SRkLL2paGjjoJoEIhPgHh8BApEAUrB41uk5tFDzSMR0YEiWDSrRELkIpJj72Zuxz/J5Yzk/3g8M/Y4ZxiGYRiGYdTHwsy14wQKnWmdrcP8XIuzi8dc7BDoPIG+JtTvjFrWadWcIK8Ye3MuRhjlLqP+qtvxvULIT0K942KCQW+O2vEBEbHqYiA/u3xqFCF/AAF+8PTS6RhW/9G4nf9DhAcxJL73oQRg1LcxCLAVLAJAt0L77wKufmWh/XcmAFoElLYF0HJAaUkQ/yWC9BnkCaM+Z9DttE4B0O3uVK+9O4ay4hKjbiYjAIGs7xnX0ouEspGEAAz6Mc/zI3+P9VFBoF+bLwBW7/iHg8ZfzooL/yOCCw3X8Hqjqd45Qv00kQIQiB7GroCeGDRPt1WcZ5DPEydA7TVFX+hIVQAPgTwzATDRCKBJ2QLUKmYIVA6bABuRBBn18bBnNoHe22/lGfRL9MdgB5SGdn6fO8BEXYQI9fpwKy8fGnEVzkHPDNMEadRjiLPi6k4v4OB9wNSfw1yZ9An1KYGspVcQwdFZaP+dCYAWAaVtAbQcUCabBCloe1y+xVDAeBdwC7wJ7b/z1d6AAtxP9pMUo/Q72dJJFwOEcmPsAoCsuJhgkNv+E+NYPkqC3nIx0vVlLdSXVSd4BCFPqC98E8XFTrutR7u4nNX1WdrP5ecM7ZdhGIZhGIZrEr8B/tC8WIeuLboAAAAASUVORK5CYII="/>
                                </defs>
                            </svg>
                        </div>
                    </router-link>
                    <router-link
                        v-if="this.logged"
                        :to="{name:'chat'}">
                        <div class="header-login header-login1">
                            <svg class="header-login-svg" viewBox="0 0 60 60" fill="none" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
                            <rect width="60" height="60" fill="url(#pattern0_305_694)"/>
                            <defs>
                            <pattern id="pattern0_305_694" patternContentUnits="objectBoundingBox" width="1" height="1">
                            <use xlink:href="#image0_305_694" transform="scale(0.0104167)"/>
                            </pattern>
                            <image id="image0_305_694" width="96" height="96" xlink:href="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAGAAAABgCAYAAADimHc4AAAACXBIWXMAAAsTAAALEwEAmpwYAAAD40lEQVR4nO2cO2hUURCGbxQUBO1UENQm5HEmkYCgYGXOJLJFFDVnkoBgaatiYZnWRo0SFVtLLVKkVbFUCCiotS9iYWdUbCIrY16ixtzX7tyz+38wEJbsvWf+/5yZvQ9OkgAAAAAAAAAAAAAAAAAAAPzBUL/0exduMMkr78JXJqm3c/glDV6xk6nhnrG+hk2YWmdtq3dy21P4YZ00VzRUG+/kljjZUrr47MJj6wQ5nnhUqgnswp0KJFWPK8J0eTUfZaee1QDvZHGwd5QKG7DccCswoyS68BSuFTaASV5bJ8KRhqfwsgwDvlgnwvHGQhkGWCdRjzlgALWNAeGNd/LEO/lknTQ3KDQ3zVFzrZAB4TuTjCdJ0rF6xUzhgtY/a8G4vFjQnDS3ZVk6PIWJ5dyNDXBy8V/fO0bje5nCTAXEqxeLMKO5/CtH78IlcwOGD5za9b/v+96xk96F9/ZCSqbQMevY/5tb38RucwOOdJ/YvtEx9H/0bqFeHVZfeFnUsabOy9oAprHRtMca7JeD7GTOWmReL5zM6RjT5uNdEHsDnHzgnvGutMcTkc2e5HzFmvSCjknHllqXnvEuT2He3oClZfvNO7mcJYGh7ok97OSetfjeyeyQk31pxz2ZTG5iJ+fSTqCkKNmWcHhxtHf0cJbjexo7ziRvmy48hXktIVnGOuxkgEmeZTlPUpQcif3wJHdrnWd2pD3HyMGRbUzhSjOa9Mr40jTZMsaXFCV/ovKRKZxt9AzjbGN67p0cyrpCPcm7vOdMilI4aSezg72n92etsZ7C59KEz9ujSB4UPXdSFFMBXPEm3egmG4UBq+HCizwlgHM06WY12bgMaEKT9k1ustEZ0Mgm7Q2abLQGlNmkff4me7/R+SVFafQA+bf77VkEXLvdvf7t4vVvhTTveUUSiQH1vCWkStcZ0RvAOZt0la60ozeA14yYz9qkLZpsyxrAOS+kmt1kW94AzvgrpxG3MtreAE7ZpC2abFusAP77me312oDsXBmj/q2fVfGZc8sZwKvx652cp0ux8fs5MKACYjBWQPtF0rolSKIIGEAwwHwWMlaAvRCMEmQvBhsEegDBAPNZyFgB9kIwSpC9GGwQ6AEEA8xnIWMF2AvBKEH2YrBBoAcQDDCfhYwVYC8EowTZi8EGgR5A8RuADZson/j6blIZBmDLMrLcsszJlPUy5kjDk1wtbIBux1vFF564XbatVHQ7XuuEOLLwJDeTstBteHU7XuukOJLwJA9L3z96yYQwjXIk6wvvZFFnfuni/47WNd0RVjs8fqKKCv9FtVBNSqv5AAAAAAAAAAAAAAAAAABIWomf5azDzv30uVwAAAAASUVORK5CYII="/>
                            </defs>
                            </svg>
                        </div>
                    </router-link>
                    <router-link
                        v-if="this.logged"
                        :to="{name:'main'}">
                        <div class="header-login header-login1">
                            <svg  class="header-login-svg" viewBox="0 0 56 55" fill="none" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
                            <rect x="0.120117" width="55" height="55" fill="url(#pattern0_305_692)"/>
                            <defs>
                            <pattern id="pattern0_305_692" patternContentUnits="objectBoundingBox" width="1" height="1">
                            <use xlink:href="#image0_305_692" transform="scale(0.0104167)"/>
                            </pattern>
                            <image id="image0_305_692" width="96" height="96" xlink:href="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAGAAAABgCAYAAADimHc4AAAACXBIWXMAAAsTAAALEwEAmpwYAAADwUlEQVR4nO2cSWtUQRDH++ICggqKCioixDhTNZMc5gMM6ZpIFMQlXZmIghcNCiqIevEruOPJqzcV8epyVFFR3EHEr6DBJEZxy0hrEAlZ5s1idb9XP/jDQGaYfv9/V3W/vHnPGEVRFEVRFEVRlMDp6+hbQAU+QsgPLbhPXv41gTvs/yY9vlRju3k1gXtOyLXpZJGf+fdIjzO9Mx9mNv/fELQS2gD9aTuzmv83BOBD7RhDpiHkR3UHgO6B9HhTByGP1RuAf6/0eFMH1W/+b0mPN3WQBiBHuXPX8qQB+M8IDjl+KiVeUkE+aMHds8A/kgbgP2OR79oCH+jr2L1Y+niiwRYGVxLy6ckz3ESmz7YoE7hTvV07VkgfX7CUSkPzepCPJ9ztJNWoBXfMf5f08QZFbxevJ+D7bTR+ant6XMnv3CB93EFgcWBrm2f9jNXQkx/YYrKMLfCBRhbYFlbCjwrwfpNF/A6HkCekzP9HE5kLweb7N0vOfJpaCeh+Ur5/m8kCPUXeKNTza3NoNPULMwPPt8hPAzC7Nq2An6R6i2qRT4qbjHPqhEkj/lKhBR4PwODaHBorA68yaYOQLwZgbq0eWeQzJk1QbvsyQvclmgCAx8vd25aatEDAR6VNpaQCd9ikBb+7EDcUE+uhSc1vecI4460lPTlLxWJsC7xH2kxqPIRBEzsW3AVpI6lRAZ8zsWPR3Ym2AsDdNLFDyO+iDQDdWxM7FvhjvAHwBxM7BO5bxAF8NbEjbSI1KRM70gaSBiBvImkFyBtJ2oLilIkdaQNJA5A3kbQC5I0kbUFxysSOtIGkAcibSFoB8kaStqA4ZWJH2kDSAORNJK0AeSNJW1CcMrEjbSBpAPImklaAvJGkLShOmdiRNpA0AHkTSStA3kjKWgsqr9u70CJfkjaQmhXw5d6uPYtMdPcCA78QNw9bIwv8plLkoonofoAQb8SuNVkJny26fSbolhPzvQAYcUuiXLXTAr8UNwf/j34fa67aaULAP1+BwA1Lm0L/vRLccBm4I4Tf/18XNwPFdE3a/6hvwKAWVEEI/2oYFTcCZeQnn7T/htDdyG4A7kogJ13ZW4Qt8vtycccaEwKbsLqWkK9adCPpN96N+JkfjPmKoiiKEuJ5hkU3osk0H8DrJgJ4pQE0iUV3tuFtI7hTGkCzAQBDQ485Bv6+qbAzpwG0ANvAhR5fOWp+iyiVhuYR8O0Es/9Wqh9BLBjCed9aZms7fuar+e1eE5DP+B3O5MX/scnXpytYzbfzuxVFURRFURRFUUwL+AVohXwHoKOqrAAAAABJRU5ErkJggg=="/>
                            </defs>
                            </svg>
                        </div>
                    </router-link>
                    <button v-if="this.logged" @click="logout">logout from {{ this.logged }}</button>
                </div>
                
            </div>
        </div>
    </header>
  </template> 
