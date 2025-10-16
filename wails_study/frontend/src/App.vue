<script lang="ts" setup>
import {EventsOn} from "../wailsjs/runtime";
import {router} from "./router";

EventsOn("router", function (name) {
  router.push({name})
})


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