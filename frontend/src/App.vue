<script lang="ts" setup>
import Sidebar  from './components/Sidebar.vue'
</script>

<template>
  <div class="flex">
    <Sidebar />
    <div class="flex-1 ml-15rem p-4">
      <router-view />
    </div>
  </div>
  <!-- <router-view /> -->
  <!-- <HelloWorld/> -->
</template>

<style>
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
