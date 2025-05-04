<template>
  <Dialog
    modal
    :visible="visible"
    :closable="false"
    @hide="onHide"
    :style="{ width: '40vw' }"
    :breakpoints="{ '1199px': '75vw', '575px': '90vw' }"
  >
    <template #header>
      <span class="dialog-header flex align-items-center gap-2">
        <i :class="[
          {
            'pi pi-info-circle': severity === 'info',
            'pi pi-check-circle': severity === 'success',
            'pi pi-exclamation-triangle': severity === 'warn',
            'pi pi-times-circle': severity === 'error'
          },
          'text-xl'
        ]" style="color: var(--primary-color)"></i>
        <span class="font-bold text-lg" style="color: #666666">{{ title }}</span>
      </span>
    </template>
    <div class="dialog-content">
      <p>{{ content }}</p>
    </div>
    <template #footer>
      <div class="flex justify-content-between w-full">
        <Button
          v-if="cancelText"
          icon="pi pi-times-circle"
          :label="cancelText ? cancelText : 'Cancelar'"
          class="p-button-text"
          severity="danger"
          @click="onCancel"
        />
        <Button
          :label="acceptText"
          :icon="acceptIcon ? acceptIcon : 'pi pi-check'"
          class="p-button-primary"
          @click="onAccept"
        />
      </div>
    </template>
  </Dialog>
</template>

<script lang="ts" setup>
import { computed } from 'vue';
import Dialog from 'primevue/dialog';
import Button from 'primevue/button';

interface Props {
  visible: boolean;
  title: string;
  content: string;
  severity?: 'info' | 'success' | 'warn' | 'error';
  acceptText?: string;
  acceptIcon?: string;
  cancelText?: string | null;
  onAcceptAction?: () => void;
}

const props = defineProps<Props>();
const emit = defineEmits(['update:visible']);

const onHide = () => {
  emit('update:visible', false);
};

const onAccept = () => {
  if (props.onAcceptAction) {
    props.onAcceptAction();
  }
  emit('update:visible', false);
};

const onCancel = () => {
  emit('update:visible', false);
};
</script>

<style scoped>
.dialog-content p{
  margin:  0;
}
</style>
