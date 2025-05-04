<template>
    <div class="card relative">
        <img :src="user.profileImage" class="ml-5 w-12rem h-12rem border-circle absolute" style="top: -6rem;" />
        <div class="p-6 bg-white border-round shadow-2" style="margin-top: 7rem; padding-top: 7.5rem !important;">
            <div class="flex flex-wrap w-full">
                <h2 class="text-xl m-0 mr-3" style="color: #9594A4;">
                    {{ user.name }}
                </h2>
                <div class="w-full flex justify-content-end sm:w-auto mt-2 sm:mt-0 border-round p-2 pt-1 pb-1 text-sm font-medium"
                    style="color: white; background-color: #EFE627;">
                    {{ user.role }}
                </div>
            </div>
            <div class="col-12 md:col-10 mt-2">
                <div class="grid formgrid mb-3">
                    <div class="field col-12 md:col-6 text-left">
                        <label>Carnet de identidad</label>
                        <div class="card-yellow">{{ user.ci }}</div>
                    </div>
                    <div class="field col-12 md:col-6 text-left">
                        <label>Teléfono</label>
                        <div class="card-yellow">{{ user.phone }}</div>
                    </div>
                </div>
                <div class="grid formgrid mb-3">
                    <div class="field col-12 md:col-6 text-left">
                        <label>Correo electrónico</label>
                        <div class="card-yellow">{{ user.email }}</div>
                    </div>
                    <div class="field col-12 md:col-6 text-left">
                        <label>Fecha de nacimiento</label>
                        <div class="card-yellow">{{ user.birth_date }}</div>
                    </div>
                </div>
                <div class="grid formgrid mb-3">
                    <div class="field col-12 md:col-6 text-left">
                        <label>Fecha de ingreso</label>
                        <div class="card-yellow">{{ user.start_date }}</div>
                    </div>
                    <div class="field col-12 md:col-6 flex">
                        <button @click="visible = true" class="btn-password text-gray-400 hover:text-gray-600 mt-auto">
                            <i class="pi pi-lock mr-1"></i>
                            <span class="underline">Cambiar contraseña</span>
                        </button>
                    </div>
                </div>
            </div>
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
                inputStyle="background-color: #fdfbdf; border: none; padding: 0.75rem;"
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
                inputStyle="background-color: #fdfbdf; border: none; padding: 0.75rem;"
            />
        </div>
        
        <template #footer>
            <div class="flex justify-content-between w-full">
                <Button icon="pi pi-times-circle" class="p-button-text"
                    label="Cancelar" text severity="danger" @click="visible = false" autofocus />
                <Button label="Cambiar" icon="pi pi-angle-right" iconPos="right" class="p-button-primary" 
                    @click="visible = false" autofocus/>
            </div>
        </template>
    </Dialog>
</template>
  
<script lang="ts" setup>
    import { ref, reactive } from 'vue';
    import { computed } from 'vue';
    import Dialog from 'primevue/dialog';
    import Button from 'primevue/button';
    import Password from 'primevue/password';

    const visible = ref(false);
  
    const user = ref({
        id: 1,
        profileImage: 'https://image.tmdb.org/t/p/w235_and_h235_face/xKs4zD0ze9aw3KtLZdzFxLYmVAu.jpg',
        ci: '12345678',
        name: 'Juan Luis Pérez Paredes',
        email: 'juan.perez@example.com',
        role: 'Administrador',
        phone: '555-1234',
        birth_date: '1990-01-01',
        start_date: '2020-01-01',
    });
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