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
                <h2 class="text-xl m-0 ml-2" style="color: #9594A4;">{{ projectRegister.Name }}</h2>
            </div>
            <div class="col-12 md:col-10 mt-2">
                <div class="grid formgrid">
                    <div class="field col-12 md:col-6">
                        <label for="description">Descripcion</label>
                        <Textarea 
                            v-model="projectRegister.Description"
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
                        <div class="card-yellow">{{ projectRegister.InitialDate }}</div>

                        <label for="end_date" class="mt-2">Fecha de finalización</label>
                        <div class="card-yellow">{{ projectRegister.FinalDate }}</div>

                        <label for="status" class="mt-2">Estado</label>
                        <div class="card-yellow">
                            {{ projectRegister.Status ? 'Activo' : 'Inactivo' }}
                        </div>
                    </div>
                </div>
            </div>
            <div class="flex">
                <h2 class="text-xl m-0 ml-2" style="color: #9594A4;">Integrantes</h2>
            </div>
            <div class="col-12 md:col-10">
                <div class="bg-gray-100 flex flex-wrap gap-2 mb-2 p-3 border-round" style="min-height: 62px;">
                    <div v-for="employee in selectedEmployees" :key="employee.ID" 
                        class="text-base bg-gray-400 text-white flex border-round align-items-center justify-content-between p-1 pl-2 pr-2">
                        {{ employee.Name }}
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
  
<script lang="ts" setup>
    import { onMounted, ref } from 'vue';
    import Textarea from 'primevue/textarea';
    import Button from 'primevue/button';
    import { useRouter } from 'vue-router';
    import { useToast } from "primevue/usetoast";

    import { GetProjectByID } from '../../../wailsjs/go/desktop/ProjectHandler';
    import { GetEmployeeByID } from '../../../wailsjs/go/desktop/EmployeeHandler';
    import { GetProjectEmployeesByProjectID } from '../../../wailsjs/go/desktop/ProjectEmployeeHandler';
    import { project, employee } from '../../../wailsjs/go/models';
    
    const router = useRouter();
    const toast = useToast();
    const isLoading = ref(true);

    var projectRegister = ref<project.Project>({
        ID: 0,
        Name: '',
        InitialDate: undefined,
        IsCurrent: false,
        Status: false,
        Description: '',
        convertValues: function (a: any, classs: any, asMap?: boolean) {
            throw new Error('Function not implemented.');
        }
    });

    var selectedEmployees = ref<employee.Employee[]>([]);

    onMounted(async () => {
        try {
            // Get the project by ID
            var id = router.currentRoute.value.params.id;
            const response = await GetProjectByID(Number(id));
            projectRegister.value = response;
            projectRegister.value.InitialDate = new Date(projectRegister.value.InitialDate).toLocaleDateString('es-ES');
            projectRegister.value.FinalDate = new Date(projectRegister.value.FinalDate).toLocaleDateString('es-ES');
            
            // Get the employees associated with the project
            const employees = await GetProjectEmployeesByProjectID(Number(id));
            if (employees && employees.length > 0) {
                for (const emp of employees) {
                    const employeeData = await GetEmployeeByID(Number(emp.EmployeeID));
                    if (employeeData) {
                        selectedEmployees.value.push(employeeData);
                    }
                }
            } else {
                selectedEmployees.value = [];
            }
        } catch (error) {
            console.error('Error fetching projects:', error);
            toast.add({ severity: 'error', summary: 'Ups! Ocurrió un error', detail: 'No se pudieron cargar los datos del proyecto', life: 2000 });
        } finally {
            isLoading.value = false;
        }
    });

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