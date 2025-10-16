<script lang="ts" setup>

</script>

<template>
  <div class="nav">
    <router-link to="/">首页</router-link>
    <router-link to="/home">home</router-link>
    <router-link to="/about">关于</router-link>
  </div>
  <div class="view">
    <router-view/>
  </div>

</template>