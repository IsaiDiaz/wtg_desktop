<template>
    <div class="card relative">
        <Button label="Editar" icon="pi pi-pencil"
            style="right: 0; top: -4rem; background-color: #EFE627; color: white; border-radius: 5px; font-weight: 700 !important;"
            class="absolute p-button-text mr-2 font-bold" @click="goToProjectEditPage" />
        <div class="ml-5 border-circle absolute w-12rem h-12rem flex align-items-center justify-content-center" style="top: -6rem; background-color: #F0F0F0;">
            <img src="../../assets/images/project/project.png" class="w-10rem h-10rem" />
        </div>
        <div class="p-6 bg-white border-round shadow-2" style="margin-top: 7rem; padding-top: 7.5rem !important;">
            <div class="flex">
                <h2 class="text-xl m-0 ml-2" style="color: #9594A4;">{{ project.Name }}</h2>
            </div>
            <div class="col-12 md:col-10 mt-2">
                <div class="grid formgrid">
                    <div class="field col-12 md:col-6">
                        <label for="description">Descripcion</label>
                        <Textarea 
                            v-model="project.Description"
                            id="description"
                            placeholder="Descripcion del proyecto" 
                            rows="8" 
                            cols="30"
                            class="w-full"
                            disabled
                            style="background-color: #fdfbdf; border: none; resize: none; color: black; padding: 10px 15px;"/>
                    </div>
                    <div class="field col-12 md:col-6">
                        <label for="start_date">Fecha de inicio</label>
                        <div class="card-yellow">{{ project.InitialDate }}</div>

                        <label for="end_date" class="mt-2">Fecha de finalización</label>
                        <div class="card-yellow">{{ project.FinalDate }}</div>

                        <label for="status" class="mt-2">Estado</label>
                        <div class="card-yellow">{{ project.Status }}</div>
                    </div>
                </div>
            </div>
            <div class="flex">
                <h2 class="text-xl m-0 ml-2" style="color: #9594A4;">Integrantes</h2>
            </div>
            <div class="col-12 md:col-10">
                <div class="bg-gray-100 flex flex-wrap gap-2 mb-2 p-3 border-round" style="min-height: 62px;">
                    <div v-for="employee in selectedEmployees" :key="employee.id" 
                        class="text-base bg-gray-400 text-white flex border-round align-items-center justify-content-between p-1 pl-2 pr-2">
                        {{ employee.name }}
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
  
<script lang="ts" setup>
    import ConfirmDialog from '../../components/dialogs/ConfirmDialog.vue';
    import { onMounted, ref } from 'vue';
    import InputText from 'primevue/inputtext';
    import Textarea from 'primevue/textarea';
    import Dropdown from 'primevue/dropdown';
    import Calendar from 'primevue/calendar';
    import Button from 'primevue/button';
    import { useRouter } from 'vue-router';
    import { GetProjectByID } from '../../../wailsjs/go/desktop/ProjectHandler';
    import { project } from '../../../wailsjs/go/models';
    import { useToast } from "primevue/usetoast";

    const router = useRouter();
    const toast = useToast();
    const isLoading = ref(true);

    var project = ref<project.Project>({});
    /*var project = ref({
        id: 1,
        name: 'Proyecto Alpha',
        description: 'Proyecto de desarrollo de software para la gestión de proyectos y tareas.',
        start_date: '2020-01-01',
        end_date: '2026-01-01',
        status: 'Activo',
        statusId: '1',
        members: [
            { id: 1, name: 'Juan Perez Lopez' },
            { id: 2, name: 'Micaela Gordillo Alcocer' },
            { id: 3, name: 'Luis Fernández' },
            { id: 4, name: 'María Chávez' },
            { id: 5, name: 'Carlos López' }
        ],
    });*/

    onMounted(async () => {
        try {
            const projectId = router.currentRoute.value.params.id;
            const response = await GetProjectByID(parseInt(projectId));
            project.value = response;
            project.value.InitialDate = new Date(project.value.InitialDate).toLocaleDateString('es-ES');
            project.value.FinalDate = new Date(project.value.FinalDate).toLocaleDateString('es-ES');
        } catch (error) {
            console.error('Error fetching projects:', error);
            toast.add({ severity: 'error', summary: 'Ups! Ocurrió un error', detail: 'No se pudieron cargar los datos del proyecto', life: 2000 });
        } finally {
            isLoading.value = false;
        }
    });

    const status = ref([
        { name: 'Activo', statusId: '1' },
        { name: 'Inactivo', statusId: '2' },
        { name: 'Finalizado', statusId: '3' },
        { name: 'Cancelado', statusId: '4' }
    ]);

    const selectedEmployees = ref(project.value.members);

    function onCancel() {
        router.push({ path: `/projects` });
    }

    function goToProjectEditPage() {
        const projectId = router.currentRoute.value.params.id;
        router.push({ path: `/project/${projectId}/edit` });
    }
</script>
  
<style>
    label {
        font-size: 13px;
        font-weight: 600;
        color: #9594A4;
    }
    .card-yellow {
        font-size: 16px;
        background-color: #fdfbdf;
        border-radius: 10px;
        padding: 10px 15px;
        margin-top: 0;
    }
</style>