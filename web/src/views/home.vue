<template>
  <v-container class="fill-height" height="10">
    <v-row class="justify-center" style="margin-top: -25vh">
      <v-col cols="12" class="d-flex justify-center align-center">
        <div style="height: 100px">
          <v-img :src="logo" :width="100"></v-img>
        </div>
        <div class="text-h2">饥荒服务器查询</div>
      </v-col>

      <!-- 修改：使用 v-row + v-col 控制复选框布局 -->
      <v-col cols="12" class="mt-16">
        <v-row>
          <!-- 第一行复选框 -->
          <v-col cols="6" sm="3">
            <v-checkbox v-model="regions['ap-east-1']" label="东亚区"></v-checkbox>
          </v-col>
          <v-col cols="6" sm="3">
            <v-checkbox v-model="regions['ap-southeast-1']" label="东南亚区"></v-checkbox>
          </v-col>
          <!-- 第二行复选框 -->
          <v-col cols="6" sm="3">
            <v-checkbox v-model="regions['us-east-1']" label="美洲区"></v-checkbox>
          </v-col>
          <v-col cols="6" sm="3">
            <v-checkbox v-model="regions['eu-central-1']" label="欧洲区"></v-checkbox>
          </v-col>
        </v-row>
      </v-col>

      <v-col cols="12" class="d-flex justify-center">
        <v-text-field v-model="searchText" clearable label="房间名" variant="outlined" @keyup.enter="handleSearch">
          <template #append>
            <v-btn size="x-large" variant="outlined" :loading="searchButtonLoading" @click="handleSearch">
              查询
            </v-btn>
          </template>
        </v-text-field>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import logo from '../assets/logo.png'
import {ref} from "vue";
import Api from '../api'
import { showSnackbar } from '../utils/snackbar'
import router from "../router/index.js";
import {useStore} from '../store/index.js'


const store = useStore()

const regions = ref({
  'ap-southeast-1': true,
  'ap-east-1': true,
  'us-east-1': false,
  'eu-central-1': false,
})
const searchText = ref('')

const searchButtonLoading = ref(false)

const handleSearch = async () => {
  if (!searchText.value) {
    showSnackbar("请输入要查询的房间名", "error")
    return
  }
  searchButtonLoading.value = true
  const reqForm = {
    text: searchText.value,
    regions: []
  }
  Object.entries(regions.value).forEach(([key, value]) => {
    if (value) {
      reqForm.regions.push(key)
    }
  })
  try {
    const response = await Api.search(reqForm)
    if (response.data === null) {
      searchButtonLoading.value = false
      showSnackbar("没有找到对应的结果", "warning")
      return
    }
    store.data = response.data
    searchButtonLoading.value = false
    await router.push('/result')
  } catch {
    searchButtonLoading.value = false
  }
}

</script>

<style scoped>
</style>