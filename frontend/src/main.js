import { createApp } from 'vue'
// import StudyApp from './views/Study.vue'
import App from './App.vue'




import './assets/css/tailwind.css';
import router from './router'
createApp(App).use(router)
.mount('#app')
