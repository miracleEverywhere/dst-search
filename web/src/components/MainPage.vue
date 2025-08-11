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
            <v-dialog @after-leave="handleCloseDetailDialog">
              <template #activator="{props: activatorProps}">
                <v-btn v-bind="activatorProps" color="info" variant="plain" @click="handleDetail(item)">
                  查看
                </v-btn>
              </template>

              <template #default="{isActive}">
                <v-card title="详细信息">
                  <v-card-text>
                    <v-row>
                      <v-col class="d-flex justify-center">
                        <span class="text-h5">{{detailData.name}}</span>
                      </v-col>
                    </v-row>
                    <v-row>
                      <v-col class="d-flex justify-center">
                        <span class="text-subtitle-1">{{detailData.desc}}</span>
                      </v-col>
                    </v-row>
                    <v-divider class="mt-8 mb-8"></v-divider>
                    <v-row>
                      <v-col class="d-flex justify-end align-center">
                        天数：
                      </v-col>
                      <v-col class="d-flex justify-start align-center">
                        <v-chip>
                          {{parsedDetailData.days}}
                        </v-chip>
                      </v-col>
                      <v-col class="d-flex justify-end align-center">
                        季节：
                      </v-col>
                      <v-col class="d-flex justify-start align-center">
                        <v-chip v-if="detailData.season==='autumn'" color="amber">
                          <v-icon icon="mdi-leaf-maple"></v-icon>秋
                        </v-chip>
                        <v-chip v-else-if="detailData.season==='winter'" color="light-blue">
                          <v-icon icon="mdi-snowflake"></v-icon>冬
                        </v-chip>
                        <v-chip v-else-if="detailData.season==='spring'" color="success">
                          <v-icon icon="mdi-leaf"></v-icon>春
                        </v-chip>
                        <v-chip v-else-if="detailData.season==='summer'" color="red-lighten-2">
                          <v-icon icon="mdi-fire"></v-icon>夏
                        </v-chip>
                        <v-chip v-else>
                          {{detailData.season}}
                        </v-chip>
                      </v-col>
                      <v-col class="d-flex justify-end align-center">
                        类型：
                      </v-col>
                      <v-col class="d-flex justify-start align-center">
                        <v-chip v-if="detailData.intent==='survival'" color="orange">
                          生存
                        </v-chip>
                        <v-chip v-else-if="detailData.intent==='endless'" color="success">
                          无尽
                        </v-chip>
                        <v-chip v-else-if="detailData.intent==='relaxed'" color="info">
                          轻松
                        </v-chip>
                        <v-chip v-else>
                          {{detailData.intent}}
                        </v-chip>
                      </v-col>
                    </v-row>
                    <v-row>
                      <v-col class="d-flex justify-end align-center">
                        密码：
                      </v-col>
                      <v-col class="d-flex justify-start align-center">
                        <v-chip v-if="detailData.password">
                          有
                        </v-chip>
                        <v-chip v-else>
                          无
                        </v-chip>
                      </v-col>
                      <v-col class="d-flex justify-end align-center">
                        版本：
                      </v-col>
                      <v-col class="d-flex justify-start align-center">
                        <v-chip>
                          {{detailData.v}}
                        </v-chip>
                      </v-col>
                      <v-col class="d-flex justify-end align-center">
                        Tick：
                      </v-col>
                      <v-col class="d-flex justify-start align-center">
                        <v-chip>
                          {{detailData.tick}}
                        </v-chip>
                      </v-col>
                    </v-row>
                    <v-row>
                      <v-col cols="2" class="d-flex justify-end align-center">
                        直连代码：
                      </v-col>
                      <v-col class="d-flex justify-start align-center">
                        <v-btn variant="plain" color="info"
                               @click="copyText(parsedDetailData.connectionCode)">
                          {{parsedDetailData.connectionCode}}
                        </v-btn>
                      </v-col>
                    </v-row>
                    <v-divider class="mt-8 mb-8"></v-divider>
                    <v-row class="d-flex justify-center">
                      <v-data-table :headers="playerHeaders" :items="parsedDetailData.players" class="w-75">
                        <template #top>
                          <div class="d-flex justify-center text-h6">在线玩家</div>
                        </template>
                        <template #no-data>
                          没有发现在线玩家
                        </template>
                        <template #item.prefab="{value}">
                          {{getValueOrKey(frefabMap, value)}}
                        </template>
                        <template #item.color="{value}">
                          <v-icon icon="mdi-checkbox-blank" :color="'#'+value"></v-icon>
                        </template>
                      </v-data-table>
                    </v-row>
                    <v-divider class="mt-8 mb-8"></v-divider>
                    <v-row>
                      <v-col class="d-flex justify-center">
                        <v-data-table :headers="modHeaders" :items="modInfo" class="w-75">
                          <template #top>
                            <div class="d-flex justify-center text-h6">模组信息</div>
                          </template>
                          <template #no-data>
                            没有发现模组
                          </template>
                          <template #item.name="{value}">
                            <v-chip v-if="value" color="success">
                              {{value}}
                            </v-chip>
                            <v-chip v-else color="error">
                              获取失败
                            </v-chip>
                          </template>
                          <template #item.image="{value}">
                            <v-img v-if="value" :width="100" :src="value" class="ma-4"></v-img>
                            <v-chip v-else color="error">
                              获取失败
                            </v-chip>
                          </template>
                          <template #item.size="{value}">
                            <v-chip v-if="value" color="info">
                              {{formatBytes(value)}}
                            </v-chip>
                            <v-chip v-else color="error">
                              获取失败
                            </v-chip>
                          </template>
                        </v-data-table>
                      </v-col>
                    </v-row>
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
  tableData.value = []
  detailData.value = {}
  parsedDetailData.value = {
    dats: 0,
    connectionCode: '',
    players: [],
    mods: [],
  }
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

const detailData = ref({})
const parsedDetailData = ref({
  dats: 0,
  connectionCode: '',
  players: [],
  mods: [],
})
const detailLoading = ref(false)

const handleDetail = (row) => {
  const reqForm = {
    rowId: row.__rowId,
    region: row.region
  }

  detailLoading.value = true

  Api.detail(reqForm).then(response => {
    detailData.value = response.data
    parsedDetailData.value['days'] = detailData.value.data.match(/day=(\d+)/)[1]
    parsedDetailData.value['connectionCode'] = `c_connect("${detailData.value.__addr}", ${detailData.value.port})`
    const playerMatches = detailData.value.players.matchAll(/colour="([^"]+)".*?name="([^"]+)".*?prefab="([^"]+)"/gs);
    const players = []
    for (const match of playerMatches) {
      players.push({
        color: match[1],
        name: match[2],
        prefab: match[3]
      })
    }
    parsedDetailData.value['players'] = players
    if (detailData.value.mods_info === null) {
      detailData.value.mods_info = []
    }
    parsedDetailData.value['mods'] = []
    for (let i = 0; i < detailData.value.mods_info.length; i += 5) {
      parsedDetailData.value['mods'].push(parseInt(detailData.value.mods_info[i].split('-')[1]))
    }
    handleGetModInfo()
  }).finally(() => {
    detailLoading.value = false
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

const copyText = async (text) => {
  try {
    await navigator.clipboard.writeText(text); // 写入剪贴板
    showSnackbar('已复制直连代码')
  } catch (err) {
    showSnackbar('复制失败，请手动复制', 'error');
  }
}

const frefabMap = {
  "wilson": "威尔逊",
  "willow": "薇洛",
  "wolfgang": "沃尔夫冈",
  "wendy": "温蒂",
  "wx78": "WX-78",
  "wickerbottom": "薇克巴顿",
  "woodie": "伍迪",
  "wes": "韦斯",
  "waxwell": "麦斯威尔",
  "wathgrithr": "薇格弗德",
  "webber": "韦伯",
  "winona": "薇诺娜",
  "warly": "沃利",
  "walter": "沃尔特",
  "wortox": "沃拓克斯",
  "wormwood": "沃姆伍德",
  "wurt": "沃特",
  "wanda": "旺达",
  "wonkey": "芜猴"
}

const getValueOrKey = (obj, key) => obj?.hasOwnProperty(key) ? obj[key] : key

const modInfo = ref([])

const handleGetModInfo = () => {
  const reqForm = {
    mod: parsedDetailData.value.mods
  }

  Api.modInfo(reqForm).then(response => {
    modInfo.value = response.data
  })
}

const formatBytes = (bytes, num=2) => {
  if (bytes === 0) return '0 B';

  const k = 1024;
  const sizes = ['B', 'KiB', 'MiB', 'GiB', 'TiB', 'PiB', 'EiB', 'ZiB', 'YiB'];
  const i = Math.floor(Math.log(bytes) / Math.log(k));

  return parseFloat((bytes / Math.pow(k, i)).toFixed(num)) + ' ' + sizes[i];
}

const modHeaders = [
  { title: '名称', value: 'name' },
  { title: '预览', value: 'image' },
  { title: 'ID', value: 'id' },
  { title: '大小', value: 'size' },
]

const playerHeaders = [
  { title: '玩家昵称', value: 'name' },
  { title: '玩家角色', value: 'prefab' },
  { title: '游戏中颜色', value: 'color' },
]

const handleCloseDetailDialog = () => {
  modInfo.value = []
  const parsedDetailData = {
    dats: 0,
    connectionCode: '',
    players: [],
    mods: [],
  }
}

</script>

<style>
</style>
