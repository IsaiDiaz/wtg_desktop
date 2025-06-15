<template>
    <div class="card relative">
        <div class="ml-5 border-circle absolute w-12rem h-12rem flex align-items-center justify-content-center" style="top: -6rem; background-color: #F0F0F0;">
            <img src="../../assets/images/project/project.png" class="w-10rem h-10rem" />
        </div>
        <div class="p-6 bg-white border-round shadow-2" style="margin-top: 7rem; padding-top: 7.5rem !important;">
            <div class="flex">
                <h2 class="text-xl m-0 ml-2" style="color: #9594A4;">Crear proyecto</h2>
            </div>
            <div class="col-12 md:col-10 mt-2">
                <div class="grid formgrid">
                    <div class="field col-12 md:col-6">
                        <label for="name">Nombre</label>
                        <InputText
                            id="name"
                            type="text"
                            v-model="newProject.Name"
                            placeholder="Nombre del proyecto"
                            class="w-full"
                            style="background-color: #fdfbdf; border: none;" />

                        <label for="description" class="mt-3">Descripcion</label>
                        <Textarea 
                            v-model="newProject.Description"
                            id="description"
                            placeholder="Descripcion del proyecto" 
                            rows="5" 
                            cols="30"
                            class="w-full"
                            style="background-color: #fdfbdf; border: none; resize: none;"/>
                    </div>
                    <div class="field col-12 md:col-6">
                        <label for="start_date">Fecha de inicio</label>
                        <Calendar 
                            v-model="newProject.InitialDate"
                            dateFormat="dd/mm/yy"
                            showIcon 
                            iconDisplay="input"
                            inputId="start_date"
                            class="w-full"
                            :inputStyle="{ backgroundColor: '#fdfbdf', border: 'none' }" />
                        <label for="end_date" class="mt-3">Fecha de finalización</label>
                        <Calendar 
                            v-model="newProject.FinalDate"
                            dateFormat="dd/mm/yy"
                            showIcon 
                            iconDisplay="input"
                            inputId="end_date"
                            class="w-full"
                            :inputStyle="{ backgroundColor: '#fdfbdf', border: 'none' }" />
                        <label for="status" class="mt-3">Estado</label>
                        <Dropdown 
                            v-model="newProject.Status"
                            :options="status" 
                            optionLabel="name" 
                            optionValue="Status"
                            placeholder="Seleciona el estado"
                            class="w-full text-base" 
                            :style="{ backgroundColor: '#fdfbdf', border: 'none' }" />
                    </div>
                </div>
            </div>
            <div class="flex">
                <h2 class="text-xl m-0 ml-2" style="color: #9594A4;">Integrantes</h2>
            </div>
            <AutoComplete
                v-model="selectedEmployee"
                :suggestions="filteredEmployees"
                optionLabel="name"
                placeholder="Agregar nuevo integrante"
                @complete="searchEmployees"
                @item-select="addEmployee"
                class="col-12 md:col-5 m-0"
                :inputStyle="{ backgroundColor: '#fdfbdf', border: 'none', width: '100% !important' }"/>
            <div class="col-12 md:col-10">
                <div class="bg-gray-100 flex flex-wrap gap-2 mb-2 p-3 border-round" style="min-height: 62px;">
                    <div v-for="employee in selectedEmployees" :key="employee.ID" 
                        class="text-base bg-gray-400 text-white flex border-round align-items-center justify-content-between p-1 pl-2">
                        {{ employee.Name }}
                        <i class="pi pi-times pl-2 pr-1" style="color: white; cursor: pointer;" 
                            @click="removeEmployee(employee)"></i>
                    </div>
                </div>
            </div>

            <div class="flex justify-content-end">
                <Button label="Cancelar" icon="pi pi-times-circle" class="p-button-text"
                    severity="danger" @click="onCancel"/>
                <Button label="Guardar" icon="pi pi-save" 
                    style="background-color: #EFE627; color: white; border-radius: 5px; font-weight: 700 !important;"
                    class="p-button-text ml-4 mr-2 font-bold" @click="save"/>
            </div>
        </div>
    </div>
</template>
  
<script lang="ts" setup>
    import { ref } from 'vue';
    import { onMounted } from 'vue';
    import { useToast } from 'primevue';
    import InputText from 'primevue/inputtext';
    import Textarea from 'primevue/textarea';
    import Dropdown from 'primevue/dropdown';
    import Calendar from 'primevue/calendar';
    import Button from 'primevue/button';
    import AutoComplete from 'primevue/autocomplete';
    import { useRouter } from 'vue-router';

    import { GetAllEmployees } from '../../../wailsjs/go/desktop/EmployeeHandler';
    import { CreateProject } from '../../../wailsjs/go/desktop/ProjectHandler';
    import { CreateProjectEmployee } from '../../../wailsjs/go/desktop/ProjectEmployeeHandler';
    import { project, employee, projectemployee } from '../../../wailsjs/go/models';

    const toast = useToast();
    const router = useRouter();

    const newProject = ref<project.Project>({
        ID: 0,
        Name: '',
        InitialDate: new Date(),
        FinalDate: new Date(),
        IsCurrent: true,
        Status: true,
        Description: '',
        convertValues: () => {}
    });

    const status = ref([
        { name: 'Activo', Status: true },
        { name: 'Inactivo', Status: false },
        //{ name: 'Finalizado', Status: '3' },
        //{ name: 'Cancelado', Status: '4' }
        //TODO : Existiran mas estados para los proyectos?
    ]);

    const isLoading = ref(true);
    var employees = ref<employee.Employee[]>([]);
    onMounted(async () => {
        try {
            const response = await GetAllEmployees();
            employees.value = response;
        } catch (error) {
            console.error('Error fetching employees:', error);
            toast.add({ severity: 'error', summary: 'Ups! Ocurrió un error', detail: 'No se pudieron cargar los empleados', life: 2000 });
        } finally {
            isLoading.value = false;
        }
    });

    const selectedEmployees = ref<employee.Employee[]>([]);
    const selectedEmployee = ref<employee.Employee | null>(null);
    const filteredEmployees = ref<employee.Employee[]>([]);

    function searchEmployees(event: any) {
        const query = event.query.toLowerCase();
        filteredEmployees.value = employees.value.filter(
            (e) =>
            e.Name.toLowerCase().includes(query) &&
            !selectedEmployees.value.some((sel) => sel.ID === e.ID)
        );
    }

    function addEmployee(event: any) {
        const employee = event.value;
        if (!selectedEmployees.value.some((e) => e.ID === employee.ID)) {
            selectedEmployees.value.push(employee);
        }
        selectedEmployee.value = null;
    }

    function removeEmployee(employeeToRemove: any) {
        selectedEmployees.value = selectedEmployees.value.filter(
            (e) => e.ID !== employeeToRemove.id
        );
    }

    function onCancel() {
        router.push({ path: `/projects` });
    }

    function save() {
        CreateProject(newProject.value)
            .then(() => {
                router.push({ path: `/projects` });
            })
            .catch((error) => {
                console.error('Error creating project:', error);
            });

        if (selectedEmployees.value.length > 0) {
            for (const employee of selectedEmployees.value) {
                const projectEmployee: projectemployee.ProjectEmployee = {
                    ID: 0,
                    ProjectID: newProject.value.ID,
                    EmployeeID: employee.ID,
                    CurrentCategoryID: 0,
                    Status: true
                };
                CreateProjectEmployee(projectEmployee)
                    .then(() => {
                        console.log(`Empleado ${employee.Name} agregado al proyecto.`);
                    })
                    .catch((error) => {
                        console.error(`Error al agregar empleado ${employee.Name} al proyecto:`, error);
                    });
            }
        }
    }
</script>
  
<style>
    label {
        font-size: 13px;
        font-weight: 600;
        color: #9594A4;
    }
</style>