<template>
    <div class="card relative">
        <img src="https://placehold.co/200x200" class="ml-5 w-12rem h-12rem border-circle absolute" style="top: -6rem;"/>
        <div class="p-6 bg-white border-round shadow-2" style="margin-top: 7rem; padding-top: 7.5rem !important;">
            <div class="flex">
                <h2 class="text-xl m-0" style="color: #9594A4;">Información del usuario</h2>
            </div>
            <div class="mt-2" style="margin-right: 20%;">
                <div style="display: flex; gap: 5%; margin-bottom: 15px;">
                    <div style="flex: 1;">
                        <label>Nombre</label><br/>
                        <input type="text" v-model="user.name" class="card-yellow" readonly/>
                    </div>
                    <div style="flex: 1;">
                        <label>Apellidos</label><br/>
                        <input type="text" v-model="user.lastname" class="card-yellow" readonly/>
                    </div>
                </div>
                <div style="display: flex; gap: 5%; margin-bottom: 15px;">
                    <div style="flex: 1;">
                        <label>Carnet de identidad</label><br/>
                        <input type="text" v-model="user.ci" class="card-yellow" readonly/>
                    </div>
                    <div style="flex: 1;">
                        <label>Teléfono</label><br/>
                        <input type="text" v-model="user.phone" class="card-yellow" readonly/>
                    </div>
                </div>
                <div style="display: flex; gap: 5%; margin-bottom: 15px;">
                    <div style="flex: 1;">
                        <label>Correo electrónico</label><br/>
                        <input type="text" v-model="user.email" class="card-yellow" readonly/>
                    </div>
                    <div style="flex: 1;">
                        <label>Fecha de nacimiento</label><br/>
                        <input type="text" v-model="user.birth_date" class="card-yellow" readonly/>
                    </div>
                </div>
                <div style="display: flex; gap: 5%; margin-bottom: 15px;">
                    <div style="flex: 1;">
                        <label>Fecha de ingreso</label><br/>
                        <input type="text" v-model="user.start_date" class="card-yellow" readonly/>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
  
<script lang="ts" setup>
    import { ref, reactive } from 'vue';        
    import InputGroup from 'primevue/inputgroup';
    import InputGroupAddon from 'primevue/inputgroupaddon';
  
    const user = ref({
        id: 1,
        ci: '12345678',
        name: 'Juan Luis',
        lastname: 'Pérez Paredes',
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