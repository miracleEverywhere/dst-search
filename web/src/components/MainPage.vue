<template>
  <v-container class="fill-height">
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

  <v-dialog v-model="dialogVisible" fullscreen @after-leave="closeDialog">
    <v-card>
      <v-toolbar>
        <v-toolbar-title>
          共找到 <span style="color: #2196F3">{{tableData?.length || 0}}</span> 个服务器
        </v-toolbar-title>
        <v-toolbar-items>
          <v-btn @click="closeDialog" icon="mdi-close"></v-btn>
        </v-toolbar-items>
      </v-toolbar>

      <v-card-text>
        <v-data-table :headers="headers" :items="tableData">
          <template #item.dedicated="{value}">
            <v-chip v-if="value" color="green" label>
              是
            </v-chip>
            <v-chip v-else color="red" label>
              否
            </v-chip>
          </template>

          <template #item.intent="{value}">
            <v-chip v-if="value==='survival'" color="orange">
              生存
            </v-chip>
            <v-chip v-else-if="value==='endless'" color="success">
              无尽
            </v-chip>
            <v-chip v-else-if="value==='relaxed'" color="info">
              轻松
            </v-chip>
            <v-chip v-else>
              {{value}}
            </v-chip>
          </template>

          <template #item.mods="{value}">
            <v-chip v-if="value">
              是
            </v-chip>
            <v-chip v-else>
              否
            </v-chip>
          </template>

          <template #item.season="{value}">
            <v-chip v-if="value==='autumn'" color="amber">
              <v-icon icon="mdi-leaf-maple"></v-icon>秋
            </v-chip>
            <v-chip v-else-if="value==='winter'" color="light-blue">
              <v-icon icon="mdi-snowflake"></v-icon>冬
            </v-chip>
            <v-chip v-else-if="value==='spring'" color="success">
              <v-icon icon="mdi-leaf"></v-icon>春
            </v-chip>
            <v-chip v-else-if="value==='summer'" color="red-lighten-2">
              <v-icon icon="mdi-fire"></v-icon>夏
            </v-chip>
            <v-chip v-else>
              {{value}}
            </v-chip>
          </template>

          <template #item.pvp="{value}">
            <v-chip v-if="value">
              是
            </v-chip>
            <v-chip v-else>
              否
            </v-chip>
          </template>

          <template #item.__rowId="{item}">
            <v-dialog>
              <template #activator="{props: activatorProps}">
                <v-btn v-bind="activatorProps" color="info" variant="plain" @click="handleDetail(item)">
                  查看
                </v-btn>
              </template>

              <template #default="{isActive}">
                <v-card title="详细信息">
                  <v-card-text>
                    {{detailData}}
                  </v-card-text>
                </v-card>
              </template>
            </v-dialog>
          </template>
        </v-data-table>
      </v-card-text>
    </v-card>

  </v-dialog>
</template>

<script setup>
import logo from '@/assets/logo.png'
import {ref} from "vue";
import Api from '@/api'
import { showSnackbar } from '@/utils/snackbar'


const regions = ref({
  'ap-southeast-1': true,
  'ap-east-1': true,
  'us-east-1': false,
  'eu-central-1': false,
})
const searchText = ref('')

const searchButtonLoading = ref(false)

const dialogVisible = ref(false)

const closeDialog = () => {
  searchText.value = ''
  dialogVisible.value = false
}

const tableData = ref([])

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
    tableData.value = response.data
    dialogVisible.value = true
    searchButtonLoading.value = false
  } catch {
    searchButtonLoading.value = false
  }
}

const detailData = ref({

})

const handleDetail = (row) => {
  const reqForm = {
    rowId: row.__rowId,
    region: row.region
  }

  Api.detail(reqForm).then(response => {
    detailData.value = response.data
  })
}

const headers = [
  { title: '服务器名', value: 'name' },
  { title: '在线玩家', value: 'connected' },
  { title: '最大玩家', value: 'maxconnections' },
  { title: '专服', value: 'dedicated' },
  { title: '类型', value: 'intent' },
  { title: '含有模组', value: 'mods' },
  { title: '季节', value: 'season' },
  { title: '玩家对战', value: 'pvp' },
  { title: '详细信息', value: '__rowId' }
]

</script>

<style>
</style>
