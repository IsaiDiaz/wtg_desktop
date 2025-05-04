<!-- src/components/Sidebar.vue -->
<template>
  <div class="flex flex-column w-15rem min-h-screen bg-white p-3">
    <div class="py-5 px-2">
      <img style="width: 100%;" src="../assets/images/feedback-logo.png" alt="Logo">
    </div>
    <ul class="list-none p-0 m-0">
      <SidebarButton
        route="/"
        :icon="HomeLogo"
        title="Inicio"
      />
      <SidebarButton
        route="/users"
        :icon="UserLogo"
        title="Usuarios"
        enabled
      />
      <SidebarButton
        route="/attendance"
        :icon="AttendanceLogo"
        title="Asistencia"
      />
      <SidebarButton
        route="/projects"
        :icon="ProjectsLogo"
        title="Proyectos"
      />
      <SidebarButton
        route="/cards"
        :icon="CardLogo"
        title="Tarjetas"
      />
      <SidebarButton
        route="/notifications"
        :icon="NotificationLogo"
        title="Notificaciones"
      />

    </ul>
  </div>
</template>

<script setup lang="ts">
import SidebarButton from './SidebarButton.vue';
import UserLogo from '../assets/images/icons/user-icon.svg';
import HomeLogo from '../assets/images/icons/home-icon.svg';
import AttendanceLogo from '../assets/images/icons/attendance-icon.svg';
import CardLogo from '../assets/images/icons/card-icon.svg';
import NotificationLogo from '../assets/images/icons/notification-icon.svg';
import ProjectsLogo from '../assets/images/icons/projects-icon.svg';
</script>
