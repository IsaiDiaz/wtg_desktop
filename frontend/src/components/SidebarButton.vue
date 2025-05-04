<template>
  <li class="sidebar-button-container">
    <RouterLink
      :to="route"
      :class="['sidebar-button flex align-items-center justify-content-start gap-2', { enabled: enabled }]"
    >
      <span class="icon text-white flex items-center justify-center pr-2 rounded">
        <div class="img-container flex items-center justify-content-center">
          <img :src=icon alt="icon" style="width: 100%;" />
        </div>
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
.sidebar-button-container {
  list-style: none;
  margin-bottom: 14px;
}

.sidebar-button:hover {
  background-color: #f7f089;
}

.sidebar-button {
  gap: 0.5rem;
  padding: 8px;
  background-color: rgba(239, 229, 39, 0.25);
  color: #2F4858;
  text-decoration: none;
  border-radius: 5px;
  transition: background-color 0.3s ease, color 0.3s ease;
  font-weight: 600;
}
.enabled {
  background-color: #EFE627;
  color: #2F4858;
}

.img-container {
  background-color: var(--primary-color);
  width: 30px;
  height: 30px;
  padding: 5px;
  border-radius: 4px;
}
</style>