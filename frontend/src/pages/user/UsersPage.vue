<template>
  <Toast/>
  <div class="card">
    <div class="flex justify-content-end align-items-center mb-2 font-bold">
      <!-- Botón Añadir Usuario -->
      <Button label="Añadir Usuario" icon="pi pi-plus" size="small"
        style="background-color: #EFE627; color: white; border-radius: 5px; font-weight: 700 !important;"
        class="p-button-text mr-2 font-bold"
        @click="goToUserCreatePage"/>
      <!-- Botón Eliminar -->
      <Button icon="pi pi-trash" severity="danger" style="background-color: var(--error-color);"
        @click="showDialog"
        size="small"
        aria-label="Eliminar" />
    </div>
    <div class="p-3 bg-white border-round shadow-2" style="position: relative;">
      <div class="flex align-items-center justify-content-center gap-2 py-2 px-3 text-lg font-bold"
        style="position: absolute; top: -46px; left: 0px; background-color: #2A5C9F; color: white; border-top-left-radius: 20px; border-top-right-radius: 20px; ">
        <img src="../../assets/images/user-icon-white.svg" width="30"></img>
        Usuarios
      </div>
      <DataTable
        v-model:selection="selectedUsers"
        :value="users" dataKey="ID"
        @row-click="goToUserInfoPage">
        <template #empty> No se encontraron usuarios. </template>
        <template #loading> Cargando usuarios... </template>
        <Column selectionMode="multiple" headerStyle="width: 3rem"></Column>
        <Column field="CI" header="CI"></Column>
        <Column field="Name" header="Nombre"></Column>
        <Column field="Email" header="Correo electrónico"></Column>
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
import ConfirmDialog from '../../components/dialogs/ConfirmDialog.vue';
import { ref, reactive } from 'vue'
// import Sidebar from '../components/Sidebar.vue'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import Button from 'primevue/button';
import Tag from 'primevue/tag'
import Toast from 'primevue/toast';
import { onMounted } from 'vue';
import { useRouter } from 'vue-router'
import { GetAllEmployees } from '../../../wailsjs/go/desktop/EmployeeHandler';
import { employee } from '../../../wailsjs/go/models';
import { useToast } from 'primevue';

const toast = useToast();
const isLoading = ref(true);

onMounted(async () => {
  try {
    const response = await GetAllEmployees();
    users.value = response;
  } catch (error) {
    console.error('Error fetching employees:', error);
    toast.add({ severity: 'error', summary: 'Ups! Ocurrió un error', detail: 'No se pudieron cargar los usuarios', life: 2000 });
  } finally {
    isLoading.value = false;
  }
});

const isDialogVisible = ref(false);
const showDialog = () => {
  isDialogVisible.value = true;
};

const handleAccept = () => {
  console.log('Accepted!');

};
const router = useRouter()
var users = ref<employee.Employee[]>([]);
const selectedUsers = ref();

function goToUserInfoPage(event: any) {
  const userId = event.data.ID;
  router.push({ path: `/users/${userId}` });
}

function goToUserCreatePage(event: any) {
  router.push({ path: `/users/new` });
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