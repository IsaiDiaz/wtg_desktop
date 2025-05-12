<template>
    <Toast/>
    <div class="card relative">
        <Button label="Editar" icon="pi pi-pencil"
            style="right: 0; top: -4rem; background-color: #EFE627; color: white; border-radius: 5px; font-weight: 700 !important;"
            class="absolute p-button-text mr-2 font-bold" @click="goToUserEditPage" />
        <img :src="user?.PhotoURL" class="ml-5 w-12rem h-12rem border-circle absolute" style="top: -6rem;" />
        <div class="p-6 bg-white border-round shadow-2" style="margin-top: 7rem; padding-top: 7.5rem !important;">
            <div class="flex flex-wrap w-full">
                <i class="pi pi-info-circle mr-3" style="font-size: 1.5rem; color: #9594A4;"></i>
                <h2 class="text-xl m-0" style="color: #9594A4;">
                    Información de usuario
                </h2>
            </div>
            <div class="col-12 md:col-10 mt-2">
                <div class="grid formgrid mb-3">
                    <div class="field col-12 md:col-6">
                        <label>Carnet de identidad</label>
                        <div class="card-yellow">{{ user?.CI }}</div>
                    </div>
                    <div class="field col-12 md:col-6">
                        <label>Rol</label>
                        <div class="card-yellow">{{ user?.Auth }}</div>
                    </div>
                </div>
                <div class="grid formgrid mb-3">
                    <div class="field col-12 md:col-6">
                        <label>Nombres</label>
                        <div class="card-yellow">{{ user?.Name }}</div>
                    </div>
                    <div class="field col-12 md:col-6">
                        <label>Apellidos</label>
                        <div class="card-yellow">{{ user?.Name }}</div>
                    </div>
                </div>

                <div class="grid formgrid mb-3">
                    <div class="field col-12 md:col-6">
                        <label>Teléfono</label>
                        <div class="card-yellow">{{ user?.CI }}</div>
                    </div>
                    <div class="field col-12 md:col-6">
                        <label>Correo electrónico</label>
                        <div class="card-yellow">{{ user?.Email }}</div>
                    </div>
                </div>

                <div class="grid formgrid mb-3">
                    <div class="field col-12 md:col-6">
                        <label>Fecha de nacimiento</label>
                        <div class="card-yellow">{{ formatDate(user?.BirthDate) }}</div>
                    </div>
                    <div class="field col-12 md:col-6">
                        <label>Fecha de ingreso</label>
                        <div class="card-yellow">{{ formatDate(user?.StartDate) }}</div>
                    </div>
                </div>
            </div>
            <div class="flex justify-content-between align-items-center my-3">
                <h2 class="text-xl m-0 mr-3" style="color: #9594A4;">Proyectos asociados</h2>
                <Button label="Asignar nuevo proyecto" icon="pi pi-plus" size="small" severity="secondary"
                    variant="outlined" rounded @click="goToUserEditPage" />
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
        </div>
    </div>

    <!--Dialog change password-->
    <Dialog v-model:visible="visible" modal :style="{ width: '35rem' }" :closable="false">
        <template #header>
            <div class="inline-flex items-center justify-center gap-2">
                <span class="font-bold whitespace-nowrap">Cambiar contraseña</span>
            </div>
        </template>

        <span class="text-surface-500 dark:text-surface-400 block mb-3">
            Por favor, completa la siguiente información
        </span>

        <div class="mt-3">
            <label for="currentPassword" class="font-medium block mb-2">Contraseña actual</label>
            <Password
                id="currentPassword"
                toggleMask
                :feedback="false"
                autocomplete="current-password"
                class="w-full"
                inputClass="w-full"
                inputStyle="background-color: #fdfbdf; border: none;"
            />
        </div>
        <div class="mt-3">
            <label for="newPassword" class="font-medium block mb-2">Contraseña nueva</label>
            <Password
                id="newPassword"
                toggleMask
                :feedback="false"
                autocomplete="current-password"
                class="w-full"
                inputClass="w-full"
                inputStyle="background-color: #fdfbdf; border: none;"
            />
        </div>

        <template #footer>
            <div class="flex justify-content-between w-full">
                <Button icon="pi pi-times-circle" class="p-button-text" label="Cancelar" text severity="danger"
                    @click="visible = false" autofocus />
                <Button label="Cambiar" icon="pi pi-angle-right" iconPos="right" class="p-button-primary"
                    @click="visible = false" autofocus />
            </div>
        </template>
    </Dialog>
</template>

<script lang="ts" setup>
import { ref, reactive, onMounted } from 'vue';
import { computed } from 'vue';
import Dialog from 'primevue/dialog';
import Button from 'primevue/button';
import Password from 'primevue/password';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Tag from 'primevue/tag';
import Toast from 'primevue/toast';
import { useRouter } from 'vue-router';
import { employee } from '../../../wailsjs/go/models';
import { GetEmployeeByID } from '../../../wailsjs/go/desktop/EmployeeHandler';
import { useToast } from 'primevue';

const visible = ref(false);
const router = useRouter();
const isLoading = ref(false);
const toast = useToast();
var user = ref<employee.Employee>();

onMounted(async () => {
    try {
        const userId = router.currentRoute.value.params.id;
        const response = await GetEmployeeByID(Number(userId));
        user.value = response;
    } catch (error) {
        console.error('Error fetching user:', error);
        toast.add({ severity: 'error', summary: 'Ups! Ocurrió un error', detail: 'No se pudo cargar la información del usuario', life: 2000 });
    } finally {
        isLoading.value = false;
    }
});

// const user = ref({
//     id: 1,
//     profileImage: 'https://image.tmdb.org/t/p/w235_and_h235_face/xKs4zD0ze9aw3KtLZdzFxLYmVAu.jpg',
//     ci: '12345678',
//     name: 'Juan Luis Pérez Paredes',
//     email: 'juan.perez@example.com',
//     role: 'Administrador',
//     roleId: '1',
//     phone: '555-1234',
//     birth_date: '1990-01-01',
//     start_date: '2020-01-01',
// });

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

function goToUserEditPage() {
    const userId = router.currentRoute.value.params.id;
    console.log(`ID del usuario: ${userId}`);
    router.push({ path: `/users/${userId}/edit` });
}

function removeProject(projectId: number) {
    console.log(`Desvinculando el proyecto con ID: ${projectId}`);
}

function formatDate(date: string) {
    return new Date(date).toLocaleDateString('es-ES');
}


</script>

<style>
label {
    font-size: 13px;
    font-weight: 600;
    color: #9594A4;
}

.card-yellow {
    font-size: 14px;
    background-color: #fdfbdf;
    border-radius: 10px;
    padding: 10px 15px;
    margin-top: 5px;
}

.btn-password {
    background-color: #ffffff;
    border: none;
    cursor: pointer;
    font-size: 14px;
    color: #9594A4;
    padding: 10px 5px;
}
</style>