<!-- src/components/Sidebar.vue -->
<template>
  <div class="flex flex-column w-15rem min-h-screen bg-white p-3">
    <div class="py-5 px-2">
      <img style="width: 100%;" src="../assets/images/feedback-logo.png" alt="Logo">
    </div>
    <ul class="list-none p-0 m-0">
      <SidebarButton
        route="/"
        :icon="UserLogo"
        title="Inicio"
      />
      <SidebarButton
        route="/users"
        :icon="UserLogo"
        title="Usuarios"
        enabled
      />
    </ul>
  </div>
</template>

<script setup lang="ts">
import SidebarButton from './SidebarButton.vue';
import UserLogo from '../assets/images/user-logo.png';

</script>
