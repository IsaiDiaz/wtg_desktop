<script lang="ts" setup>
import { reactive } from 'vue'
import { CreateEmployee, GetAllEmployees } from '../../wailsjs/go/desktop/EmployeeHandler'
import { employee } from '../../wailsjs/go/models'

const {Employee} = employee

const form = reactive({
  Name: "Juan Pérez",
  CI: "12345678",
  BirthDate: new Date().toISOString(),
  StartDate: new Date().toISOString(),
  PhotoURL: "https://placehold.co/200x200",
  Auth: 1,
  CategoryID: 1,
  Email: "juan@example.com",
})

const message = reactive({
  status: "",
})

async function create() {
  try {
    const emp = new Employee({
      ID: 0,
      Name: form.Name,
      CI: form.CI,
      BirthDate: form.BirthDate,
      StartDate: form.StartDate,
      PhotoURL: form.PhotoURL,
      Auth: form.Auth,
      CategoryID: form.CategoryID,
      Email: form.Email,
    })

    await CreateEmployee(emp)
    console.log(await GetAllEmployees())
    message.status = "✅ Empleado creado exitosamente"
  } catch (error) {
    message.status = "❌ Error al crear: " + error
  }
}
</script>

<template>
  <main>
    <div id="form" class="input-box">
      <input v-model="form.Name" placeholder="Nombre" class="input" />
      <input v-model="form.CI" placeholder="CI" class="input" />
      <input v-model="form.Email" placeholder="Email" class="input" />
      <button class="btn" @click="create">Crear empleado</button>
    </div>
    <div class="result">{{ message.status }}</div>
  </main>
</template>

<style scoped>
.result {
  margin-top: 1rem;
  color: green;
}

.input-box {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.input-box .btn {
  width: 160px;
  height: 40px;
  line-height: 40px;
  border-radius: 6px;
  border: none;
  padding: 0 8px;
  cursor: pointer;
  background-color: #007bff;
  color: white;
}

.input-box .input {
  border: 1px solid #ccc;
  border-radius: 4px;
  height: 36px;
  padding: 0 10px;
}
</style>