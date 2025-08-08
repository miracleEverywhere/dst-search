<template>
  <v-container class="fill-height">
    <v-row class="justify-center">
      <v-col cols="12" class="d-flex justify-center align-center">
        <div style="height: 100px">
          <v-img :src="logo" :width="100"></v-img>
        </div>
        <div class="text-h2">饥荒服务器查询</div>
      </v-col>
      <v-col cols="12" class="d-flex mt-16">
        <div class="d-flex">
          <v-checkbox v-model="regions['ap-east-1']" label="东亚区" class="mr-8"></v-checkbox>
          <v-checkbox v-model="regions['ap-southeast-1']" label="东南亚区" class="mr-8"></v-checkbox>
          <v-checkbox v-model="regions['us-east-1']" label="美洲区" class="mr-8"></v-checkbox>
          <v-checkbox v-model="regions['eu-central-1']" label="欧洲区" class="mr-8"></v-checkbox>
        </div>
      </v-col>
      <v-col cols="12" class="d-flex justify-center">
        <v-text-field clearable label="房间名" variant="outlined">
          <template #append>
            <v-btn size="x-large" variant="outlined" :loading="searchButtonLoading" @click="openDialog">
              查询
            </v-btn>
          </template>
        </v-text-field>
      </v-col>
    </v-row>
  </v-container>

  <v-dialog v-model="dialogVisible" fullscreen>
    <v-card>
      <v-toolbar>
        <v-btn @click="closeDialog" icon="mdi-close"></v-btn>
      </v-toolbar>
    </v-card>

  </v-dialog>
</template>

<script setup>
import logo from '@/assets/logo.png'
import {ref} from "vue";

const regions = ref({
  'ap-southeast-1': true,
  'ap-east-1': true,
  'us-east-1': false,
  'eu-central-1': false,
})

const searchButtonLoading = ref(false)

const dialogVisible = ref(false)
const openDialog = () => {
  searchButtonLoading.value = true
  searchButtonLoading.value = false
  dialogVisible.value = true
}
const closeDialog = () => {
  dialogVisible.value = false
}

</script>

<style>
</style>
