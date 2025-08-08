<template>
  <v-app>
    <v-snackbar v-model="snackbar" :timeout="timeout" transition="slide-y-transition"
                :location="location" variant="flat" color="#FFFFFF00">
      <v-card>
        <v-alert :type="color" variant="text" density="compact">
          {{ text }}
        </v-alert>
      </v-card>
    </v-snackbar>

    <v-main>
      <main-page />
    </v-main>

    <app-footer/>
  </v-app>
</template>

<script setup>
import MainPage from '@/components/MainPage.vue'
import AppFooter from '@/components/AppFooter.vue'
import { useSnackbar } from './utils/snackbar'

const { snackbar, color, text, location, icon, timeout } = useSnackbar()
</script>
