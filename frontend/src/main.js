import { createApp } from 'vue'
// import StudyApp from './views/Study.vue'
import App from './App.vue'
import { library } from "@fortawesome/fontawesome-svg-core";
import { 
    faMagnifyingGlass,
    faArrowsToCircle,
    faHighlighter,
    faHouse,
} from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

library.add(
    faMagnifyingGlass, 
    faArrowsToCircle, 
    faHighlighter,
    faHouse,
);

import './assets/css/tailwind.css';
import router from './router'
createApp(App).use(router)
.component("font-awesome-icon", FontAwesomeIcon)
.mount('#app')
