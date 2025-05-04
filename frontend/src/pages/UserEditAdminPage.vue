<template>
    <div class="card relative">
        <Button class="absolute" icon="pi pi-trash" severity="danger" 
            style="right: 0; top: -4rem; background-color: var(--error-color);" 
            @click="showDialogDelete" label="Eliminar" />
        <img :src="user.profileImage" class="ml-5 w-12rem h-12rem border-circle absolute" style="top: -6rem;"/>
        <div class="p-6 bg-white border-round shadow-2" style="margin-top: 7rem; padding-top: 7.5rem !important;">
            <div class="flex">
                <i class="pi pi-info-circle" style="font-size: 1.5rem; color: #9594A4;"></i>
                <h2 class="text-xl m-0 ml-2" style="color: #9594A4;">Información del usuario</h2>
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
                            inputStyle="background-color: #fdfbdf; border: none; padding: 0.75rem;" />

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
                            inputStyle="background-color: #fdfbdf; border: none; padding: 0.75rem;" />
                    </div>
                </div>
            </div>
            <div class="flex justify-content-end">
                <Button label="Guardar" icon="pi pi-save" 
                    style="background-color: #EFE627; color: white; border-radius: 5px; font-weight: 700 !important;"
                    class="p-button-text mr-2 font-bold" @click="showDialogEdit"/>
            </div>
        </div>
    </div>

    <!-- Dialog, eliminación -->
    <ConfirmDialog :visible="isDialogDeleteVisible" :title="'Eliminación de usuario'"
      :content="'Se enviará una solicitud al administrador del sistema para eliminar este usuario. El administrador revisará la solicitud y podrá aprobarla o rechazarla. Hasta que esto suceda, el usuario seguirá activo en el sistema. Te notificaremos cuando haya una respuesta.'"
      :severity="'warn'" :cancelText="'Cancelar'" :acceptText="'Enviar solicitud'" :onAcceptAction="handleAcceptDelete"
      acceptIcon="pi pi-send"
      @update:visible="isDialogDeleteVisible = $event" />
    <!-- Dialog, actualización -->
    <ConfirmDialog :visible="isDialogEditVisible" :title="'Actualización de información'"
      :content="'Se enviará una solicitud al administrador del sistema para actualizar la información. Una vez revisada, el administrador podrá aprobar o rechazar la modificación. Hasta que esto suceda, los cambios no serán visibles. Te notificaremos cuando haya una respuesta.'"
      :severity="'warn'" :cancelText="'Cancelar'" :acceptText="'Enviar solicitud'" :onAcceptAction="handleAcceptEdit"
      acceptIcon="pi pi-send"
      @update:visible="isDialogEditVisible = $event" />
</template>
  
<script lang="ts" setup>
    import ConfirmDialog from '../components/dialogs/ConfirmDialog.vue';
    import { ref, reactive } from 'vue';
    import InputText from 'primevue/inputtext';
    import Dropdown from 'primevue/dropdown';
    import Calendar from 'primevue/calendar';
    import Button from 'primevue/button';
  
    const user = ref({
        id: 1,
        profileImage: 'https://image.tmdb.org/t/p/w235_and_h235_face/xKs4zD0ze9aw3KtLZdzFxLYmVAu.jpg',
        ci: '12345678',
        name: 'Juan Luis',
        lastname: 'Pérez Paredes',
        email: 'juan.perez@example.com',
        role: 'Administrador',
        roleId: '1',
        phone: '555-1234',
        birth_date: new Date('1990-01-01T00:00:00'),
        start_date: new Date('2020-01-01T00:00:00')
    });

    const roles = ref([
        { name: 'Administrador', roleId: '1' },
        { name: 'Auditor', roleId: '2' },
        { name: 'Empleado', roleId: '3' }
    ]);

    const isDialogDeleteVisible = ref(false);
    const showDialogDelete = () => {
        isDialogDeleteVisible.value = true;
    };
    const handleAcceptDelete = () => {
        console.log('Accepted!');
    };

    const isDialogEditVisible = ref(false);
    const showDialogEdit = () => {
        isDialogEditVisible.value = true;
    };
    const handleAcceptEdit = () => {
        console.log('Accepted!');
    };
</script>
  
<style>
    label {
        font-size: 13px;
        font-weight: 600;
        color: #9594A4;
    }
</style>