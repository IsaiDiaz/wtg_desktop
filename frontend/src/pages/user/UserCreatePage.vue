<template>
    <div class="card relative">
        <img src="../../assets/images/user/user.png" class="ml-5 w-12rem h-12rem border-circle absolute" style="top: -6rem;"/>
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
                            v-model="user.ci"
                            placeholder="Carnet de identidad"
                            class="w-full"
                            style="background-color: #fdfbdf; border: none;" />
                    </div>
                    <div class="field col-12 md:col-6">
                        <label for="rol">Rol</label>
                        <Dropdown 
                            v-model="user.roleId" 
                            :options="roles" 
                            optionLabel="name" 
                            optionValue="roleId"
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
                            v-model="user.name"
                            placeholder="Nombre"
                            class="w-full"
                            style="background-color: #fdfbdf; border: none;" />
                    </div>
                    <div class="field col-12 md:col-6">
                        <label for="lastname">Apellidos</label>
                        <InputText
                            id="lastname"
                            type="text"
                            v-model="user.lastname"
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
                            v-model="user.phone"
                            placeholder="Teléfono"
                            class="w-full"
                            style="background-color: #fdfbdf; border: none;" />
                    </div>
                    <div class="field col-12 md:col-6">
                        <label for="email">Correo electrónico</label>
                        <InputText
                            id="email"
                            type="text"
                            v-model="user.email"
                            placeholder="Correo electrónico"
                            class="w-full"
                            style="background-color: #fdfbdf; border: none;" />
                    </div>
                </div>

                <div class="grid formgrid mb-3">
                    <div class="field col-12 md:col-6">
                        <label for="birth_date">Fecha de nacimiento</label>
                        <Calendar 
                            v-model="user.birth_date"
                            dateFormat="dd/mm/yy"
                            showIcon 
                            iconDisplay="input"
                            inputId="birth_date"
                            class="w-full"
                            inputStyle="background-color: #fdfbdf; border: none;" />

                    </div>
                    <div class="field col-12 md:col-6">
                        <label for="start_date">Fecha de ingreso</label>
                        <Calendar 
                            v-model="user.start_date"
                            dateFormat="dd/mm/yy"
                            showIcon 
                            iconDisplay="input"
                            inputId="birth_date"
                            class="w-full"
                            inputStyle="background-color: #fdfbdf; border: none;" />
                    </div>
                </div>
            </div>
            <div class="flex justify-content-end">
                <Button label="Cancelar" icon="pi pi-times-circle" class="p-button-text"
                    severity="danger" @click="onCancel"/>
                <Button label="Guardar" icon="pi pi-save" 
                    style="background-color: #EFE627; color: white; border-radius: 5px; font-weight: 700 !important;"
                    class="p-button-text ml-4 mr-2 font-bold" @click="createUser"/>
            </div>
        </div>
    </div>
</template>
  
<script lang="ts" setup>
    import ConfirmDialog from '../../components/dialogs/ConfirmDialog.vue';
    import { ref, reactive } from 'vue';
    import InputText from 'primevue/inputtext';
    import Dropdown from 'primevue/dropdown';
    import Calendar from 'primevue/calendar';
    import Button from 'primevue/button';
    import { useRouter } from 'vue-router';

    const router = useRouter();
  
    const user = ref({
        ci: '',
        name: '',
        lastname: '',
        email: '',
        role: '',
        roleId: '',
        phone: '',
        birth_date: new Date(),
        start_date: new Date()
    });

    const roles = ref([
        { name: 'Administrador', roleId: '1' },
        { name: 'Auditor', roleId: '2' },
        { name: 'Empleado', roleId: '3' }
    ]);

    function onCancel() {
        router.push({ path: `/users` });
    }

    function createUser() {
    }
</script>
  
<style>
    label {
        font-size: 13px;
        font-weight: 600;
        color: #9594A4;
    }
</style>