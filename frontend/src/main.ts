import {createApp} from 'vue'
import App from './App.vue'
import './style.css';
import { PrimeVue } from '@primevue/core';
import Aura from '@primeuix/themes/aura';
import router from './routes/wtgRouter';
import 'primeflex/primeflex.css';
import 'primeicons/primeicons.css';

const app = createApp(App);
app.use(router);
app.use(PrimeVue, {
  theme: {
      preset: Aura,
      options: {
        darkModeSelector: false || 'none',
      }
  }
});
app.mount('#app')
