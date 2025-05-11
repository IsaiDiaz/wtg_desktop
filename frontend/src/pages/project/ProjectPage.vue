<template>
  <div class="card">
    <div class="flex justify-content-end align-items-center mb-2 font-bold">
      <!-- Botón Añadir proyecto -->
      <Button label="Añadir proyecto" icon="pi pi-plus" size="small"
        style="background-color: #EFE627; color: white; border-radius: 5px; font-weight: 700 !important;"
        class="p-button-text mr-2 font-bold" 
        @click="goToProjectCreatePage"/>
      <!-- Botón Eliminar -->
      <Button icon="pi pi-trash" severity="danger" style="background-color: var(--error-color);"
        size="small"
        @click="showDialog"
        aria-label="Eliminar" />
    </div>
    <div class="p-3 bg-white border-round shadow-2" style="position: relative;">
      <div class="flex align-items-center justify-content-center gap-2 py-2 px-3 text-lg font-bold"
        style="position: absolute; top: -46px; left: 0px; background-color: #2A5C9F; color: white; border-top-left-radius: 20px; border-top-right-radius: 20px; ">
        <img src="../../assets/images/user-icon-white.svg" width="30"></img>
        Proyectos
      </div>
      <DataTable v-model:selection="selectedProjects" :value="projects" dataKey="id" @row-click="goToProjectInfoPage">
        <Column selectionMode="multiple" headerStyle="width: 3rem"></Column>
        <Column field="name" header="Nombre"></Column>
        <Column field="description" header="Descripción"></Column>
        <Column field="startDate" header="Fecha de inicio"></Column>
        <Column field="endDate" header="Fecha de finalización"></Column>
        <Column header="Estado">
          <template #body="slotProps">
            <Tag :value="slotProps.data.status"
              :severity="slotProps.data.status === 'En pausa' ? 'warn' : slotProps.data.status === 'Finalizado' ?
              'success' : slotProps.data.status === 'Cancelado' ? 'danger' : 'info'">
            </Tag>
          </template>
        </Column>
      </DataTable>
    </div>
    <!-- Dialog, eliminación -->
    <ConfirmDialog :visible="isDialogVisible" :title="'Eliminación de proyecto'"
      :content="'Se enviará una solicitud al administrador del sistema para eliminar este proyecto. El administrador revisará la solicitud y podrá aprobarla o rechazarla. Hasta que esto suceda, el proyecto seguirá activo en el sistema. Te notificaremos cuando haya una respuesta.'"
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
import { ref } from 'vue';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Button from 'primevue/button';
import Tag from 'primevue/tag';
import { useRouter } from 'vue-router';

const isDialogVisible = ref(false);

const showDialog = () => {
  isDialogVisible.value = true;
};

const handleAccept = () => {
  console.log('Accepted!');
};

const router = useRouter();

const projects = ref([
  {
    id: 1,
    name: 'Proyecto Alpha',
    description: 'Desarrollo de una aplicación web',
    startDate: '2023-01-15',
    endDate: '2023-06-30',
    status: 'En curso',
  },
  {
    id: 2,
    name: 'Proyecto Beta',
    description: 'Implementación de un sistema ERP',
    startDate: '2022-09-01',
    endDate: '2023-03-15',
    status: 'Finalizado',
  },
  {
    id: 3,
    name: 'Proyecto Gamma',
    description: 'Migración de base de datos',
    startDate: '2023-02-01',
    endDate: '2023-08-01',
    status: 'En pausa',
  },
  {
    id: 4,
    name: 'Proyecto Delta',
    description: 'Auditoría de seguridad',
    startDate: '2023-04-01',
    endDate: '2023-09-30',
    status: 'En curso',
  },
  {
    id: 5,
    name: 'Proyecto Epsilon',
    description: 'Capacitación interna',
    startDate: '2023-05-01',
    endDate: '2023-07-15',
    status: 'Cancelado',
  },
]);

const selectedProjects = ref();

function goToProjectInfoPage(event: any) {
  const projectId = event.data.id;
  router.push({ path: `/project/${projectId}` });
}

function goToProjectCreatePage(event: any) {
  router.push({ path: `/project/new` });
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