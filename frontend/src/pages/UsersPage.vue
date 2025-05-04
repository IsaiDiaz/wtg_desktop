<template>
  <div class="card">
    <div class="flex justify-content-end align-items-center mb-2 font-bold">
      <!-- Botón Añadir Usuario -->
      <Button label="Añadir Usuario" icon="pi pi-plus"
        style="background-color: #EFE627; color: white; border-radius: 5px; font-weight: 700 !important;"
        class="p-button-text mr-2 font-bold" />
      <!-- Botón Eliminar -->
      <Button icon="pi pi-trash" severity="danger" style="background-color: var(--error-color);"
        @click="showDialog"
        aria-label="Eliminar" />
    </div>
    <div class="p-3 bg-white border-round shadow-2" style="position: relative;">
      <div class="flex align-items-center justify-content-center gap-2 py-2 px-3 text-lg font-bold"
        style="position: absolute; top: -46px; left: 0px; background-color: #2A5C9F; color: white; border-top-left-radius: 20px; border-top-right-radius: 20px; ">
        <img src="../assets/images/user-icon-white.svg" width="30"></img>
        Usuarios
      </div>
      <DataTable v-model:selection="selectedUsers" :value="users" dataKey="id" @row-click="goToUserInfoPage">
        <Column selectionMode="multiple" headerStyle="width: 3rem"></Column>
        <Column field="ci" header="CI"></Column>
        <Column field="name" header="Nombre"></Column>
        <Column field="email" header="Correo electrónico"></Column>
        <Column header="Rol">
          <template #body="slotProps">
            <Tag :value="slotProps.data.role"
              :severity="slotProps.data.role === 'Administrador' ? 'warn' : slotProps.data.role === 'Auditor' ? 'Success' : 'info'">
            </Tag>
          </template>
        </Column>
        <Column field="phone" header="Teléfono"></Column>
      </DataTable>
    </div>
    <!-- Dialog, eliminación -->
    <ConfirmDialog :visible="isDialogVisible" :title="'Eliminación de usuario'"
      :content="'Se enviará una solicitud al administrador del sistema para eliminar este usuario. El administrador revisará la solicitud y podrá aprobarla o rechazarla. Hasta que esto suceda, el usuario seguirá activo en el sistema. Te notificaremos cuando haya una respuesta.'"
      :severity="'warn'" :cancelText="'Cancelar'" :acceptText="'Enviar solicitud'" :onAcceptAction="handleAccept"
      acceptIcon="pi pi-send"
      @update:visible="isDialogVisible = $event" />
    <!-- Dialog, actualización -->
    <!-- <ConfirmDialog :visible="isDialogVisible" :title="'Actualización de información'"
      :content="'Se enviará una solicitud al administrador del sistema para actualizar la información. Una vez revisada, el administrador podrá aprobar o rechazar la modificación. Hasta que esto suceda, los cambios no serán visibles. Te notificaremos cuando haya una respuesta.'"
      :severity="'warn'" :cancelText="'Cancelar'" :acceptText="'Enviar solicitud'" :onAcceptAction="handleAccept"
      acceptIcon="pi pi-send"
      @update:visible="isDialogVisible = $event" /> -->
  </div>
</template>


<script lang="ts" setup>
import ConfirmDialog from '../components/dialogs/ConfirmDialog.vue';
import { ref, reactive } from 'vue'
// import Sidebar from '../components/Sidebar.vue'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import Button from 'primevue/button';
import Tag from 'primevue/tag'
import { useRouter } from 'vue-router'

const isDialogVisible = ref(false);

const showDialog = () => {
  isDialogVisible.value = true;
};

const handleAccept = () => {
  console.log('Accepted!');

};

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
    role: 'Empleado',
    phone: '555-5678'
  },
  {
    id: 3,
    ci: '11223344',
    name: 'Carlos García',
    email: 'carlos.garcia@example.com',
    role: 'Auditor',
    phone: '555-9101'
  },
  {
    id: 4,
    ci: '44332211',
    name: 'Ana Torres',
    email: 'ana.torres@example.com',
    role: 'Empleado',
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
const selectedUsers = ref();

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