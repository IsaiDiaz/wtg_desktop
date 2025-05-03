<template>
    <div class="card relative">
        <img src="https://placehold.co/200x200" class="ml-5 w-12rem h-12rem border-circle absolute" style="top: -6rem;"/>
        <div class="p-6 bg-white border-round shadow-2" style="margin-top: 7rem; padding-top: 7.5rem !important;">
            <div class="flex">
                <h2 class="text-xl m-0" style="color: #9594A4;">{{ user.name }}</h2>
                <div class="ml-3 border-round p-2 pt-1 pb-1 text-sm font-medium" style="color: white; background-color: #EFE627;">
                    {{ user.role }}
                </div>
            </div>
            <div class="col-12 md:col-8 mt-2" style="margin-right: 20%;">
                <div class="grid formgrid" style="display: flex; gap: 5%; margin-bottom: 15px;">
                    <div class="field col-9 md:col-4" style="flex: 1;">
                        <label>Carnet de identidad</label>
                        <div class="card-yellow">{{ user.ci }}</div>
                    </div>
                    <div class="field col-9 md:col-4" style="flex: 1;">
                        <label>Teléfono</label>
                        <div class="card-yellow">{{ user.phone }}</div>
                    </div>
                </div>
                <div style="display: flex; gap: 5%; margin-bottom: 15px;">
                    <div style="flex: 1;">
                        <label>Correo electrónico</label>
                        <div class="card-yellow">{{ user.email }}</div>
                    </div>
                    <div style="flex: 1;">
                        <label>Fecha de nacimiento</label>
                        <div class="card-yellow">{{ user.birth_date }}</div>
                    </div>
                </div>
                <div style="display: flex; gap: 5%; margin-bottom: 15px;">
                    <div style="flex: 1;">
                        <label>Fecha de ingreso</label>
                        <div class="card-yellow">{{ user.start_date }}</div>
                    </div>
                    <div class="flex flex-col flex-1">
                        <button class="btn-password flex flex-col items-center text-gray-400 hover:text-gray-600 mt-auto">
                            <i class="pi pi-lock mr-1"></i>
                            <span class="underline">Cambiar contraseña</span>
                        </button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
  
<script lang="ts" setup>
    import { ref, reactive } from 'vue'
  
    const user = ref({
        id: 1,
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