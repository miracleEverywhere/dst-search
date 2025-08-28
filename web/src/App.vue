<template>
  <v-snackbar v-model="snackbar" :timeout="timeout" transition="slide-y-transition"
              :location="location" variant="flat" color="#FFFFFF00">
    <v-card>
      <v-alert :type="TYPE" :color="COLOR" variant="text" density="compact">
        {{ text }}
      </v-alert>
    </v-card>
  </v-snackbar>

  <router-view />
</template>

<script setup>
import {RouterView} from "vue-router";
import { useSnackbar } from './utils/snackbar'

const { snackbar, TYPE, COLOR, text, location, icon, timeout } = useSnackbar()
</script>
