<template>
    <Toast />
    <div class="card relative">
        <Button class="absolute" icon="pi pi-trash" severity="danger"
            style="right: 0; top: -4rem; background-color: var(--error-color);"
            @click="showDialogDelete" label="Eliminar" />
        <div class="ml-5 border-circle absolute w-12rem h-12rem flex align-items-center justify-content-center" style="top: -6rem; background-color: #F0F0F0;">
            <img src="../../assets/images/project/project.png" class="w-10rem h-10rem" />
        </div>
        <div class="p-6 bg-white border-round shadow-2" style="margin-top: 7rem; padding-top: 7.5rem !important;">
            <div class="flex">
                <i class="pi pi-pencil" style="font-size: 1.5rem; color: #9594A4;"></i>
                <h2 class="text-xl m-0 ml-2" style="color: #9594A4;">Editar proyecto</h2>
            </div>
            <div class="col-12 md:col-10 mt-2">
                <div class="grid formgrid">
                    <div class="field col-12 md:col-6">
                        <label for="name">Nombre</label>
                        <InputText
                            id="name"
                            type="text"
                            v-model="projectRegister.Name"
                            placeholder="Nombre del proyecto"
                            class="w-full"
                            style="background-color: #fdfbdf; border: none;" />

                        <label for="description" class="mt-3">Descripcion</label>
                        <Textarea
                            v-model="projectRegister.Description"
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
                            v-model="projectRegister.InitialDate"
                            dateFormat="dd/mm/yy"
                            showIcon
                            iconDisplay="input"
                            inputId="start_date"
                            class="w-full"
                            :inputStyle="{ backgroundColor: '#fdfbdf', border: 'none' }" />

                        <label for="end_date" class="mt-3">Fecha de finalización</label>
                        <Calendar
                            v-model="projectRegister.FinalDate"
                            dateFormat="dd/mm/yy"
                            showIcon
                            iconDisplay="input"
                            inputId="end_date"
                            class="w-full"
                            :inputStyle="{ backgroundColor: '#fdfbdf', border: 'none' }" />

                        <label for="status" class="mt-3">Estado</label>
                        <Dropdown
                            v-model="projectRegister.Status"
                            :options="status"
                            optionLabel="name"
                            optionValue="Status"
                            placeholder="Seleciona el estado"
                            class="w-full text-base"
                            style="background-color: #fdfbdf; border: none;" />
                    </div>
                </div>
            </div>
            <div class="flex">
                <h2 class="text-xl m-0 ml-2" style="color: #9594A4;">Integrantes</h2>
            </div>
            <AutoComplete
                v-model="selectedEmployee"
                :suggestions="filteredEmployees"
                optionLabel="Name"
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
                    class="p-button-text ml-4 mr-2 font-bold" @click="showDialogEdit"/>
            </div>
        </div>
    </div>

    <!-- Dialog, eliminación -->
    <ConfirmDialog :visible="isDialogDeleteVisible" :title="'Eliminación de proyecto'"
        :content="'Se enviará una solicitud al administrador del sistema para eliminar este proyecto. El administrador revisará la solicitud y podrá aprobarla o rechazarla. Hasta que esto suceda, el proyecto seguirá activo en el sistema. Te notificaremos cuando haya una respuesta.'"
        :severity="'warn'" :cancelText="'Cancelar'" :acceptText="'Enviar solicitud'"
        :onAcceptAction="handleAcceptDelete" acceptIcon="pi pi-send" @update:visible="isDialogDeleteVisible = $event" />
    <!-- Dialog, actualización -->
    <ConfirmDialog :visible="isDialogEditVisible" :title="'Actualización de información'"
        :content="'Se enviará una solicitud al administrador del sistema para actualizar la información. Una vez revisada, el administrador podrá aprobar o rechazar la modificación. Hasta que esto suceda, los cambios no serán visibles. Te notificaremos cuando haya una respuesta.'"
        :severity="'warn'" :cancelText="'Cancelar'" :acceptText="'Enviar solicitud'" :onAcceptAction="handleAcceptEdit"
        acceptIcon="pi pi-send" @update:visible="isDialogEditVisible = $event" />
</template>

<script lang="ts" setup>
    import ConfirmDialog from '../../components/dialogs/ConfirmDialog.vue';
    import { onMounted, ref } from 'vue';
    import { useRouter } from 'vue-router';
    import { useToast } from 'primevue/usetoast';
    import InputText from 'primevue/inputtext';
    import Textarea from 'primevue/textarea';
    import Dropdown from 'primevue/dropdown';
    import Calendar from 'primevue/calendar';
    import Button from 'primevue/button';
    import Toast from 'primevue/toast';
    import AutoComplete from 'primevue/autocomplete';

    import { GetProjectByID, UpdateProject } from '../../../wailsjs/go/desktop/ProjectHandler';
    import { GetProjectEmployeesByProjectID } from '../../../wailsjs/go/desktop/ProjectEmployeeHandler';
    import { GetAllEmployees, GetEmployeeByID } from '../../../wailsjs/go/desktop/EmployeeHandler';
    import { DeleteProject } from '../../../wailsjs/go/desktop/ProjectHandler';
    import { project, employee } from '../../../wailsjs/go/models';

    const router = useRouter();
    const toast = useToast();
    const isLoading = ref(true);

    var projectRegister = ref<project.Project>({
        ID: 0,
        Name: '',
        InitialDate: new Date().toLocaleDateString('es-ES'),
        FinalDate: new Date().toLocaleDateString('es-ES'),
        IsCurrent: false,
        Status: false,
        Description: '',
        convertValues: function (a: any, classs: any, asMap?: boolean) {
            throw new Error('Function not implemented.');
        }
    });

    var employees = ref<employee.Employee[]>([]);
    var selectedEmployees = ref<employee.Employee[]>([]);

    onMounted(async () => {
        try {
            // Get the project by ID
            let projectId = router.currentRoute.value.params.id;
            if (Array.isArray(projectId)) {
                projectId = projectId[0];
            }
            const response = await GetProjectByID(parseInt(projectId));
            projectRegister.value = response;
            projectRegister.value.InitialDate = new Date(projectRegister.value.InitialDate);
            projectRegister.value.FinalDate = new Date(projectRegister.value.FinalDate);
            
            // Get the employees associated with the project
            const projectEmployees = await GetProjectEmployeesByProjectID(parseInt(projectId));
            if (projectEmployees && projectEmployees.length > 0) {
                for (const emp of projectEmployees) {
                    const employeeData = await GetEmployeeByID(emp.EmployeeID);
                    if (employeeData) {
                        selectedEmployees.value.push(employeeData);
                    }
                }
            } else {
                selectedEmployees.value = [];
            }

            // Get all employees
            const employeeResponse = await GetAllEmployees();
            employees.value = employeeResponse;

        } catch (error) {
            console.error('Error fetching projects:', error);
            toast.add({ severity: 'error', summary: 'Ups! Ocurrió un error', detail: 'No se pudieron cargar los datos del proyecto', life: 2000 });
        } finally {
            isLoading.value = false;
        }
    });

    const status = ref([
        { name: 'Activo', Status: true },
        { name: 'Inactivo', Status: false },
        //{ name: 'Finalizado', Status: '3' },
        //{ name: 'Cancelado', Status: '4' }
        //TODO : Existiran mas estados para los proyectos?
    ]);

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
            (e) => e.ID !== employeeToRemove.ID
        );
    }

    const isDialogDeleteVisible = ref(false);
    const showDialogDelete = () => {
        isDialogDeleteVisible.value = true;
    };

    const handleAcceptDelete = async() => {
        //TODO: En este caso, se eliminará el proyecto, cambiar más adelante
        isDialogDeleteVisible.value = false;
        try {
            const projectId = router.currentRoute.value.params.id;
            console.log('ID del proyecto a eliminar:', projectId);
            await DeleteProject(Number(projectId));
            toast.add({ severity: 'success', summary: 'Proyecto eliminado', detail: 'El proyecto ha sido eliminado correctamente.', life: 2000 });
            setTimeout(() => {
                router.push({ path: `/projects` });
            }, 2000);

        } catch (error) {
            console.error('Error deleting project:', error);
            toast.add({ severity: 'error', summary: 'Ups! Ocurrió un error', detail: 'No se pudo eliminar el proyecto', life: 2000 });
        }
    };

    const isDialogEditVisible = ref(false);
    const showDialogEdit = () => {
        isDialogEditVisible.value = true;
    };
    const handleAcceptEdit = async() => {
        isDialogEditVisible.value = false;
        try {
            const projectId = router.currentRoute.value.params.id;
            console.log('ID del proyecto a editar:', projectId);
            await UpdateProject(projectRegister.value);
            // TODO : Funcion para actualizar los empleados del proyecto
            toast.add({ severity: 'success', summary: 'Proyecto editado', detail: 'El proyecto ha sido editado correctamente.', life: 2000 });
            setTimeout(() => {
                router.push({ path: `/project/${projectId}` });
            }, 2000);
        } catch (error) {
            console.error('Error editing project:', error);
            toast.add({ severity: 'error', summary: 'Ups! Ocurrió un error', detail: 'No se pudo editar el proyecto', life: 2000 });
        }
    };

    function onCancel() {
        const projectId = router.currentRoute.value.params.id;
        router.push({ path: `/project/${projectId}` });
    }

    function formatDate(fechaStr: string) {
        const [dia, mes, año] = fechaStr.split("/");
        return `${año}-${mes.padStart(2, '0')}-${dia.padStart(2, '0')}T00:00:00Z`;
    }
</script>

<style>
    label {
        font-size: 13px;
        font-weight: 600;
        color: #9594A4;
    }
</style>