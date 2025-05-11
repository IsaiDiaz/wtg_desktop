<template>
  <Toast />
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
      <DataTable
        v-model:selection="selectedProjects"
        :value="projects" dataKey="ID"
        :loading="isLoading"
        @row-click="goToProjectInfoPage">
        <template #empty> No se encontraron proyectos. </template>
        <template #loading> Cargando proyectos... </template>
        <Column selectionMode="multiple" headerStyle="width: 3rem"></Column>
        <Column field="Name" header="Nombre"></Column>
        <Column field="Description" header="Descripción"></Column>
        <Column field="InitialDate" header="Fecha de inicio">
          <template #body="slotProps">
            {{ formatDate(slotProps.data.InitialDate) }}
          </template>
        </Column>
        <Column field="FinalDate" header="Fecha de finalización">
          <template #body="slotProps">
            {{ formatDate(slotProps.data.FinalDate) }}
          </template>
        </Column>
        <Column header="Estado">
          <template #body="slotProps">
            <Tag :value="slotProps.data.Status ? 'Activo' : 'Inactivo'"
              :severity="slotProps.data.Status ? 'success' : 'danger'">
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
import { onMounted, ref } from 'vue';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Button from 'primevue/button';
import Tag from 'primevue/tag';
import Toast from 'primevue/toast';
import { useRouter } from 'vue-router';
import { GetAllProjects } from '../../../wailsjs/go/desktop/ProjectHandler';
import { project } from '../../../wailsjs/go/models';
import { useToast } from "primevue/usetoast";

const toast = useToast();
const isLoading = ref(true);

onMounted(async () => {
  try {
    const response = await GetAllProjects();
    projects.value = response;
  } catch (error) {
    console.error('Error fetching projects:', error);
    toast.add({ severity: 'error', summary: 'Ups! Ocurrió un error', detail: 'No se pudieron cargar los proyectos', life: 2000 });
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

const router = useRouter();

var projects = ref<project.Project[]>([]);

const selectedProjects = ref();

function goToProjectInfoPage(event: any) {
  const projectId = event.data.ID;
  console.log(projectId);
  router.push({ path: `/project/${projectId}` });
}

function goToProjectCreatePage(event: any) {
  router.push({ path: `/project/new` });
}

function formatDate(date: string) {
  return new Date(date).toLocaleDateString('es-ES');
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