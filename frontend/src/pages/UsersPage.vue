<template>
  <div class="card">
    <div class="flex justify-content-end align-items-center mb-4">
      <!-- Botón Añadir Usuario -->
      <Button label="Añadir Usuario" icon="pi pi-plus"
        style="background-color: #EFE627; color: #2F4858; border-radius: 5px;" class="p-button-text mr-2" />
      <!-- Botón Eliminar -->
      <Button icon="pi pi-trash" severity="danger" aria-label="Eliminar" />
    </div>
    <div class="p-3 bg-white border-round shadow-2"
      style="position: relative;">
      <div class="flex align-items-center justify-content-center py-3 px-5 text-xl font-bold"
        style="position: absolute; top: -60px; left: 0px; background-color: #2A5C9F; color: white; border-top-left-radius: 20px; border-top-right-radius: 20px; ">
        <i class="pi pi-users mr-1 text-2xl"></i>
        Usuarios
      </div>
      <DataTable :value="users" tableStyle="min-width: 50rem" @row-click="goToUserInfoPage">
        <Column field="ci" header="CI"></Column>
        <Column field="name" header="Nombre"></Column>
        <Column field="email" header="Correo electrónico"></Column>
        <Column field="role" header="Rol"></Column>
        <Column field="phone" header="Teléfono"></Column>
      </DataTable>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, reactive } from 'vue'
// import Sidebar from '../components/Sidebar.vue'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import Button from 'primevue/button';
import { useRouter } from 'vue-router'

const router = useRouter()

const users = ref([
  {
    id: 1,
    ci: '12345678',
    name: 'Juan Pérez',
    email: 'juan.perez@example.com',
    role: 'Administrador',
    phone: '555-1234'
  },
  {
    id: 2,
    ci: '87654321',
    name: 'María López',
    email: 'maria.lopez@example.com',
    role: 'Usuario',
    phone: '555-5678'
  },
  {
    id: 3,
    ci: '11223344',
    name: 'Carlos García',
    email: 'carlos.garcia@example.com',
    role: 'Moderador',
    phone: '555-9101'
  },
  {
    id: 4,
    ci: '44332211',
    name: 'Ana Torres',
    email: 'ana.torres@example.com',
    role: 'Usuario',
    phone: '555-1213'
  },
  {
    id: 5,
    ci: '55667788',
    name: 'Luis Martínez',
    email: 'luis.martinez@example.com',
    role: 'Administrador',
    phone: '555-1415'
  }
]);

function goToUserInfoPage(event: any) {
  const clientId = event.data.id;
  router.push({ path: `/users/${clientId}` });
}
</script>

<style>
.p-datatable-thead>tr>th {
  background-color: #F5F4F9;
  color: #9594A4;
  font-weight: bold;
  font-size: 16px;
}

.p-datatable-thead>tr:first-child>th:first-child {
  border-top-left-radius: 30px;
}

.p-datatable-thead>tr:first-child>th:last-child {
  border-top-right-radius: 30px;
}
</style>