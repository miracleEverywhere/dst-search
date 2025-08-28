import {createApp} from "vue";
import './style.css'
import App from './App.vue'

import 'vuetify/styles'
import {createVuetify} from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import {aliases, mdi} from "vuetify/iconsets/mdi";
import '@mdi/font/css/materialdesignicons.css'

import router from "./router/index.js";

import {createPinia} from "pinia";


const vuetify = createVuetify({
    components,
    directives,
    icons: {
        defaultSet: 'mdi',
        aliases,
        sets: {
            mdi,
        },
    },
})

const pinia = createPinia()

const app = createApp(App)

app.use(vuetify)
app.use(router)
app.use(pinia)

app.mount('#app')
