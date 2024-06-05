<template>
    <div v-if="isLoading" class="preloader">
      <div class="spinner">
        <p>SyncMe</p>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      isLoading: {
        type: Boolean,
        default: true
      }
    }
  }
  </script>
  
  <style scoped>
  .preloader {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: rgba(255, 255, 255, 0.8);
    z-index: 9999;
  }
  
  .spinner {
    display: flex;
    justify-content: center;
    align-items: center;
    border: 16px solid #f3f3f3;
    border-top: 16px solid #3498db;
    border-radius: 50%;
    width: 150px;
    height: 150px;
    animation: spin 2s linear infinite;
  }
  
  .spinner p {
    position: absolute;
    font-size: 24px;
    color: #3498db;
  }
  
  @keyframes spin {
    0% { transform: rotate(0deg); }
    100% { transform: rotate(360deg); }
  }
  </style>
  