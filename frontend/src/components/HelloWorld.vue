<script lang="ts" setup>
import { reactive } from 'vue'
import { CreateEmployee } from '../../wailsjs/go/desktop/EmployeeHandler'
import { employee } from '../../wailsjs/go/models'

const {Employee} = employee

const form = reactive({
  name: "Juan Pérez",
  ci: "12345678",
  birth_date: new Date().toISOString(),
  start_date: new Date().toISOString(),
  photo_url: "https://placehold.co/200x200",
  auth: 1,
  category_id: 1,
  email: "juan@example.com",
})

const message = reactive({
  status: "",
})

async function create() {
  try {
    const emp = new Employee({
      id: 0, 
      name: form.name,
      ci: form.ci,
      birth_date: form.birth_date,
      start_date: form.start_date,
      photo_url: form.photo_url,
      auth: form.auth,
      category_id: form.category_id,
      email: form.email,
    })

    await CreateEmployee(emp)
    message.status = "✅ Empleado creado exitosamente"
  } catch (error) {
    message.status = "❌ Error al crear: " + error
  }
}
</script>

<template>
  <main>
    <div id="form" class="input-box">
      <input v-model="form.name" placeholder="Nombre" class="input" />
      <input v-model="form.ci" placeholder="CI" class="input" />
      <input v-model="form.email" placeholder="Email" class="input" />
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