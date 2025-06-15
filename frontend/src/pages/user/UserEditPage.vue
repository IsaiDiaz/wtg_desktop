<template>
    <div class="card relative">
        <Button class="absolute" icon="pi pi-trash" severity="danger" 
            style="right: 0; top: -4rem; background-color: var(--error-color);" 
            @click="showDialogDelete" label="Eliminar" />
        <img :src="user.PhotoURL" class="ml-5 w-12rem h-12rem border-circle absolute" style="top: -6rem;"/>
        <div class="p-6 bg-white border-round shadow-2" style="margin-top: 7rem; padding-top: 7.5rem !important;">
            <div class="flex">
                <i class="pi pi-pencil" style="font-size: 1.5rem; color: #9594A4;"></i>
                <h2 class="text-xl m-0 ml-2" style="color: #9594A4;">Editar usuario</h2>
            </div>
            <div class="col-12 md:col-10 mt-2">
                <div class="grid formgrid mb-3">
                    <div class="field col-12 md:col-6">
                        <label for="ci">Carnet de identidad</label>
                        <InputText
                            id="ci"
                            type="text"
                            v-model="user.CI"
                            placeholder="Carnet de identidad"
                            class="w-full"
                            style="background-color: #fdfbdf; border: none;" />
                    </div>
                    <div class="field col-12 md:col-6">
                        <label for="rol">Rol</label>
                        <Dropdown 
                            v-model="user.Auth" 
                            :options="roles" 
                            optionLabel="name" 
                            optionValue="Auth"
                            placeholder="Seleciona el rol" 
                            class="w-full text-base" 
                            style="background-color: #fdfbdf; border: none;" />    
                    </div>
                </div>

                <div class="grid formgrid mb-3">
                    <div class="field col-12 md:col-6">
                        <label for="name">Nombre</label>
                        <InputText
                            id="name"
                            type="text"
                            v-model="user.Name"
                            placeholder="Nombre"
                            class="w-full"
                            style="background-color: #fdfbdf; border: none;" />
                    </div>
                    <div class="field col-12 md:col-6">
                        <label for="lastname">Apellidos</label>
                        <InputText
                            id="lastname"
                            type="text"
                            v-model="user.Name"
                            placeholder="Apellidos"
                            class="w-full"
                            style="background-color: #fdfbdf; border: none;" />
                    </div>
                </div>

                <div class="grid formgrid mb-3">
                    <div class="field col-12 md:col-6">
                        <label for="phone">Teléfono</label>
                        <InputText
                            id="phone"
                            type="text"
                            v-model="user.Phone"
                            placeholder="Teléfono"
                            class="w-full"
                            style="background-color: #fdfbdf; border: none;" />
                    </div>
                    <div class="field col-12 md:col-6">
                        <label for="email">Correo electrónico</label>
                        <InputText
                            id="email"
                            type="text"
                            v-model="user.Email"
                            placeholder="Correo electrónico"
                            class="w-full"
                            style="background-color: #fdfbdf; border: none;" />
                    </div>
                </div>

                <div class="grid formgrid mb-3">
                    <div class="field col-12 md:col-6">
                        <label for="birth_date">Fecha de nacimiento</label>
                        <Calendar 
                            v-model="user.BirthDate"
                            dateFormat="dd/mm/yy"
                            showIcon 
                            iconDisplay="input"
                            inputId="birth_date"
                            class="w-full"
                            :inputStyle="{ backgroundColor: '#fdfbdf', border: 'none' }" />
                    </div>
                    <div class="field col-12 md:col-6">
                        <label for="start_date">Fecha de ingreso</label>
                        <Calendar 
                            v-model="user.StartDate"
                            dateFormat="dd/mm/yy"
                            showIcon 
                            iconDisplay="input"
                            inputId="birth_date"
                            class="w-full"
                            :inputStyle="{ backgroundColor: '#fdfbdf', border: 'none' }" />
                    </div>
                </div>
            </div>
            <div class="flex justify-content-between align-items-center my-3">
                <h2 class="text-xl m-0 mr-3" style="color: #9594A4;">Proyectos asociados</h2>
            </div>
            <DataTable :value="projects" dataKey="id">
                <Column field="name" header="Nombre"></Column>
                <Column field="startDate" header="Fecha de inicio"></Column>
                <Column field="endDate" header="Fecha de finalización"></Column>
                <Column header="Estado">
                    <template #body="slotProps">
                        <Tag :value="slotProps.data.status" :severity="slotProps.data.status ? 'success': 'warn'" class="w-full text-center">
                            {{ slotProps.data.status ? 'Activo' : 'Inactivo' }}
                        </Tag>
                    </template>
                </Column>
                <Column header="">
                    <!-- Botón, desvincular del proyecto -->
                    <template #body="slotProps">
                        <Button label="Desvincular del proyecto" icon="pi pi-link" size="small" severity="secondary"
                            @click="removeProject(slotProps.data.id)" />
                    </template>
                </Column>
            </DataTable>
            <div class="flex justify-content-end mt-4">
                <Button label="Cancelar" icon="pi pi-times-circle" class="p-button-text" severity="danger"
                    @click="onCancel" />
                <Button label="Guardar" icon="pi pi-save"
                    style="background-color: #EFE627; color: white; border-radius: 5px; font-weight: 700 !important;"
                    class="p-button-text ml-4 mr-2 font-bold" @click="showDialogEdit" />
            </div>
        </div>
    </div>

    <!-- Dialog, eliminación -->
    <ConfirmDialog :visible="isDialogDeleteVisible" :title="'Eliminación de usuario'"
        :content="'Se enviará una solicitud al administrador del sistema para eliminar este usuario. El administrador revisará la solicitud y podrá aprobarla o rechazarla. Hasta que esto suceda, el usuario seguirá activo en el sistema. Te notificaremos cuando haya una respuesta.'"
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
    import InputText from 'primevue/inputtext';
    import Dropdown from 'primevue/dropdown';
    import Calendar from 'primevue/calendar';
    import DataTable from 'primevue/datatable';
    import Button from 'primevue/button';
    import Column from 'primevue/column';
    import Tag from 'primevue/tag';
    import { useToast } from 'primevue';
    import { useRouter } from 'vue-router';
    
    import { GetEmployeeByID, UpdateEmployee, DeleteEmployee } from '../../../wailsjs/go/desktop/EmployeeHandler';
    import { GetProjectEmployeesByEmployeeID, DeleteProjectEmployee } from '../../../wailsjs/go/desktop/ProjectEmployeeHandler';
    import { GetProjectByID } from '../../../wailsjs/go/desktop/ProjectHandler';
    import { employee, project, projectemployee } from '../../../wailsjs/go/models';

    const toast = useToast();
    const router = useRouter();
    const isLoading = ref(true);

    var user = ref<employee.Employee>({
        ID: 0,
        Name: '',
        CI: '',
        Email: '',
        Phone: '',
        PhotoURL: '',
        BirthDate: new Date(),
        StartDate: new Date(),
        Auth: 1,
        CategoryID: 1,
        convertValues: (a: any) => a 
    });
    var projectsEmployee = ref<projectemployee.ProjectEmployee[]>([]);
    var projects = ref<project.Project[]>([]);

    onMounted(async () => {
        try {
            // Get the user by ID from the route parameters
            const userId = router.currentRoute.value.params.id;
            console.log('userId', userId);
            const response = await GetEmployeeByID(Number(userId));
            user.value = response;
            user.value.BirthDate = new Date(user.value.BirthDate);
            user.value.StartDate = new Date(user.value.StartDate);
            if(user.value.PhotoURL == '') {
                user.value.PhotoURL = '../../assets/images/user/user.png';
            }

            // Get projects associated with the user
            const projectsEmp = await GetProjectEmployeesByEmployeeID(Number(userId));
            projectsEmployee.value = projectsEmp;
            if (projectsEmployee && projectsEmployee.value.length > 0) {
                for (const project of projectsEmployee.value) {
                    const projectDetails = await GetProjectByID(project.ID);
                    if (projectDetails) {
                        projects.value.push(projectDetails);
                    }
                }
            } else {
                projects.value = [];
            }
        } catch (error) {
            console.error('Error fetching user:', error);
            toast.add({ severity: 'error', summary: 'Ups! Ocurrió un error', detail: 'No se pudo cargar la información del usuario', life: 2000 });
        } finally {
            isLoading.value = false;
        }
    });

    const roles = ref([ // TODO : obtener los roles desde la API
        { name: 'Administrador', Auth: 1 },
        { name: 'Auditor', Auth: 2 },
        { name: 'Empleado', Auth: 3 }
    ]);

    const isDialogDeleteVisible = ref(false);
    const showDialogDelete = () => {
        isDialogDeleteVisible.value = true;
    };
    const handleAcceptDelete = async() => {
        isDialogDeleteVisible.value = false;
        try {
            const userId = router.currentRoute.value.params.id;
            console.log('ID del empleado a eliminar:', userId);
            await DeleteEmployee(Number(userId));
            toast.add({ severity: 'success', summary: 'Empleado eliminado', detail: 'El empleado ha sido eliminado correctamente.', life: 2000 });
            setTimeout(() => {
                router.push({ path: `/users` });
            }, 2000);
        } catch (error) {
            console.error('Error deleting user:', error);
            toast.add({ severity: 'error', summary: 'Ups! Ocurrió un error', detail: 'No se pudo eliminar al empleado', life: 2000 });
        }
    };

    const isDialogEditVisible = ref(false);
    const showDialogEdit = () => {
        isDialogEditVisible.value = true;
    };
    const handleAcceptEdit = async() => {
        isDialogEditVisible.value = false;
        try {
            const userId = router.currentRoute.value.params.id;
            console.log('ID del empleado a editar:', userId);
            await UpdateEmployee(user.value);
            toast.add({ severity: 'success', summary: 'Información actualizada', detail: 'La información del empleado ha sido actualizada correctamente.', life: 2000 });
            setTimeout(() => {
                router.push({ path: `/users` });
            }, 2000);
        } catch (error) {
            console.error('Error deleting project:', error);
            toast.add({ severity: 'error', summary: 'Ups! Ocurrió un error', detail: 'No se pudo actualizar la información del empleado', life: 2000 });
        }
    };

    function onCancel() {
        const clientId = router.currentRoute.value.params.id;
        router.push({ path: `/users/${clientId}` });
    }

    async function removeProject(projectId: number) {
        // Obtener de projectsEmployee
        const projectEmployeeID = projectsEmployee.value.find(pe => pe.ID === projectId)?.ID;
        if (!projectEmployeeID) {
            toast.add({ severity: 'error', summary: 'Error', detail: 'No se encontró el proyecto asociado.', life: 2000 });
            return;
        }
        try {
            await DeleteProjectEmployee(projectEmployeeID);
            projects.value = projects.value.filter(p => p.ID !== projectId);
            projectsEmployee.value = projectsEmployee.value.filter(pe => pe.ID !== projectEmployeeID);
            toast.add({severity: 'success', summary: 'Proyecto desvinculado', detail: 'El proyecto ha sido desvinculado correctamente.', life: 2000 });
        } catch (error) {
            console.error('Error removing project:', error);
            toast.add({ severity: 'error', summary: 'Ups! Ocurrió un error', detail: 'No se pudo desvincular el proyecto', life: 2000 });
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