<template>
    <div class="card relative">
        <img src="../../assets/images/user/user.png" class="ml-5 w-12rem h-12rem border-circle absolute"
            style="top: -6rem;" />
        <div class="p-6 bg-white border-round shadow-2" style="margin-top: 7rem; padding-top: 7.5rem !important;">
            <div class="flex">
                <h2 class="text-xl m-0 ml-2" style="color: #9594A4;">Crear usuario</h2>
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
                            placeholder="Selecciona el rol"
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
                        <label for="lastname">Apellidos (falta el campo en BD)</label>
                        <!--<InputText
                            id="lastname"
                            type="text"
                            v-model="user.LastName"
                            placeholder="Apellidos"
                            class="w-full"
                            style="background-color: #fdfbdf; border: none;" />-->
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
                <Button label="Asignar nuevo proyecto" icon="pi pi-plus" size="small" severity="secondary"
                    variant="outlined" rounded @click="addProject" />
            </div>
            <DataTable :value="projects" dataKey="id">
                <Column field="name" header="Nombre"></Column>
                <Column field="startDate" header="Fecha de inicio"></Column>
                <Column field="endDate" header="Fecha de finalización"></Column>
                <Column header="Estado">
                    <template #body="slotProps">
                        <Tag :value="slotProps.data.status" :severity="slotProps.data.status === 'En pausa' ? 'warn' : slotProps.data.status === 'Finalizado' ?
                            'success' : slotProps.data.status === 'Cancelado' ? 'danger' : 'info'">
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
                    class="p-button-text ml-4 mr-2 font-bold" @click="createUser" />
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
    import { ref } from 'vue';
    import InputText from 'primevue/inputtext';
    import Dropdown from 'primevue/dropdown';
    import Calendar from 'primevue/calendar';
    import Button from 'primevue/button';
    import Column from 'primevue/column';
    import Tag from 'primevue/tag';
    import DataTable from 'primevue/datatable';
    import { useRouter } from 'vue-router';

    import { CreateEmployee } from '../../../wailsjs/go/desktop/EmployeeHandler';
    import { employee } from '../../../wailsjs/go/models';

    const router = useRouter();

    const user = ref<employee.Employee>({
        ID: 0,
        Name: '',
        CI: '',
        BirthDate: new Date(),
        StartDate: new Date(),
        PhotoURL: '../../assets/images/user/user.png',
        Auth: 3,
        CategoryID: 0,
        Email: '',
        Phone: '',
        convertValues: () => {}
    });

    const roles = ref([
        { name: 'Administrador', Auth: 1 },
        { name: 'Auditor', Auth: 2 },
        { name: 'Empleado', Auth: 3 }
    ]);

    const projects = ref([
        {
            id: 1,
            name: 'Proyecto Alpha',
            description: 'Desarrollo de una aplicación web',
            startDate: '2023-01-15',
            endDate: '2023-06-30',
            status: 'En curso',
        },
    ]);

    function onCancel() {
        router.push({ path: `/users` });
    }

    async function createUser() {
        try {
            await CreateEmployee(user.value);
            router.push({ path: `/users` });
        } catch (error) {
            console.error('Error creating user:', error);
        }
    }

    function addProject() {
        // Logic to add a new project
    }

    function removeProject(projectId: number) {
        // Logic to remove a project
    }
</script>

<style>
label {
    font-size: 13px;
    font-weight: 600;
    color: #9594A4;
}
</style>