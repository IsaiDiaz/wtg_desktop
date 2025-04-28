<template>
  <li class="mb-3">
    <RouterLink
      :to="route"
      :class="['sidebar-button flex items-center gap-2', { enabled: enabled }]"
    >
      <span class="icon text-white flex items-center justify-center pr-2 rounded">
        <img :src=icon alt="icon" style="width: 20px;" />
      </span>
      {{ title }}
    </RouterLink>
  </li>
</template>

<script>
export default {
  props: {
    route: {
      type: String,
      required: true
    },
    icon: {
      type: String,
      required: true
    },
    title: {
      type: String,
      required: true
    },
    enabled: {
      type: Boolean,
      default: false
    }
  }
};
</script>

<style scoped>
.icon {
  font-size: 14px;
}
.sidebar-button:hover {
  background-color: #f7f089;
}

.sidebar-button {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 10px 20px;
  background-color: rgba(239, 229, 39, 0.25);
  color: #2F4858;
  text-decoration: none;
  border-radius: 5px;
  transition: background-color 0.3s ease, color 0.3s ease;
}
.enabled {
  background-color: #EFE627;
  color: #2F4858;
}
</style>