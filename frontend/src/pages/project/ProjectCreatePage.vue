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
                            v-model="project.name"
                            placeholder="Nombre del proyecto"
                            class="w-full"
                            style="background-color: #fdfbdf; border: none;" />

                        <label for="description" class="mt-3">Descripcion</label>
                        <Textarea 
                            v-model="project.description"
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
                            v-model="project.start_date"
                            dateFormat="dd/mm/yy"
                            showIcon 
                            iconDisplay="input"
                            inputId="start_date"
                            class="w-full"
                            inputStyle="background-color: #fdfbdf; border: none;" />

                        <label for="end_date" class="mt-3">Fecha de finalización</label>
                        <Calendar 
                            v-model="project.end_date"
                            dateFormat="dd/mm/yy"
                            showIcon 
                            iconDisplay="input"
                            inputId="end_date"
                            class="w-full"
                            inputStyle="background-color: #fdfbdf; border: none;" />

                        <label for="status" class="mt-3">Estado</label>
                        <Dropdown 
                            v-model="project.statusId"
                            :options="status" 
                            optionLabel="name" 
                            optionValue="statusId"
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
                optionLabel="name"
                placeholder="Agregar nuevo integrante"
                @complete="searchEmployees"
                @item-select="addEmployee"
                class="col-12 md:col-5 m-0"
                inputStyle="background-color: #fdfbdf; border: none; width: 100% !important;"/>
            <div class="col-12 md:col-10">
                <div class="bg-gray-100 flex flex-wrap gap-2 mb-2 p-3 border-round" style="min-height: 62px;">
                    <div v-for="employee in selectedEmployees" :key="employee.id" 
                        class="text-base bg-gray-400 text-white flex border-round align-items-center justify-content-between p-1 pl-2">
                        {{ employee.name }}
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
</template>
  
<script lang="ts" setup>
    import ConfirmDialog from '../../components/dialogs/ConfirmDialog.vue';
    import { ref, reactive } from 'vue';
    import InputText from 'primevue/inputtext';
    import Textarea from 'primevue/textarea';
    import Dropdown from 'primevue/dropdown';
    import Calendar from 'primevue/calendar';
    import Button from 'primevue/button';
    import { useRouter } from 'vue-router';
    import MultiSelect from 'primevue/multiselect';
    import AutoComplete from 'primevue/autocomplete';

    const router = useRouter();
  
    const project = ref({
        name: '',
        description: '',
        start_date: new Date(),
        end_date: new Date(),
        status: '',
        statusId: '',
        members: [],
    });

    const status = ref([
        { name: 'Activo', statusId: '1' },
        { name: 'Inactivo', statusId: '2' },
        { name: 'Finalizado', statusId: '3' },
        { name: 'Cancelado', statusId: '4' }
    ]);

    const employees = ref([
        { id: 1, name: 'Juan Perez Lopez' },
        { id: 2, name: 'Micaela Gordillo Alcocer' },
        { id: 3, name: 'Luis Fernández' },
        { id: 4, name: 'María Chávez' },
        { id: 5, name: 'Carlos López' },
        { id: 6, name: 'Ana María' },
        { id: 7, name: 'Pedro González' },
        { id: 8, name: 'Laura Martínez' },
        { id: 9, name: 'Javier Torres' },
        { id: 10, name: 'Sofía Ramírez' },
        { id: 11, name: 'Diego Morales' },
        { id: 12, name: 'Valeria Castro' },
        { id: 13, name: 'Andrés Herrera' },
        { id: 14, name: 'Camila Ortega' },
        { id: 15, name: 'Fernando Ruiz' }
    ]);

    const selectedEmployees = ref([]);
    const selectedEmployee = ref(null);
    const filteredEmployees = ref([]);

    function searchEmployees(event: any) {
        const query = event.query.toLowerCase();
        filteredEmployees.value = employees.value.filter(
            (e) =>
            e.name.toLowerCase().includes(query) &&
            !selectedEmployees.value.some((sel) => sel.id === e.id)
        );
    }

    function addEmployee(event: any) {
        const employee = event.value;
        if (!selectedEmployees.value.some((e) => e.id === employee.id)) {
            selectedEmployees.value.push(employee);
        }
        selectedEmployee.value = null;
    }

    function removeEmployee(employeeToRemove: any) {
        selectedEmployees.value = selectedEmployees.value.filter(
            (e) => e.id !== employeeToRemove.id
        );
    }

    function onCancel() {
        router.push({ path: `/projects` });
    }
</script>
  
<style>
    label {
        font-size: 13px;
        font-weight: 600;
        color: #9594A4;
    }
</style>