<script lang="ts" setup>
import {EventsOn} from "../wailsjs/runtime";
import {router} from "./router";
import {Quit, WindowMinimise} from "../wailsjs/runtime";

EventsOn("router", function (name) {
  router.push({name})
})


</script>

<template>
  <div class="nav" style="--wails-draggable:drag">
    <div class="left">
      <router-link to="/">首页</router-link>
      <router-link to="/home">home</router-link>
      <router-link to="/about">关于</router-link>
    </div>
    <div class="action">
      <span @click="WindowMinimise">-</span>
      <span @click="Quit">x</span>
    </div>
  </div>
  <div class="view">
    <router-view/>
  </div>

</template>


<style>
body {
  margin: 0;
}

.nav {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 40px;
  background-color: #b1b1b1;
  padding: 0 20px;

  .left {
    a {
      margin-right: 20px;
      text-decoration: none;
    }
  }

  .action {
    span {
      margin-left: 10px;
      cursor: pointer;

      &:hover {
        color: #1d1dee;
      }
    }
  }

}
</style>