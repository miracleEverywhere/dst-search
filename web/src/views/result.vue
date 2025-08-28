<template>
  <v-card>
    <v-layout style="width: 100%">
      <v-app-bar>
        <template #prepend>
          <v-btn icon="mdi-arrow-left" @click="gotoHome"></v-btn>
        </template>
        <v-app-bar-title>
          查询结果
        </v-app-bar-title>
      </v-app-bar>

      <v-data-table :headers="headers" :items="store.data" hover style="width: 100%">
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
            <v-icon icon="mdi-sprout"></v-icon>春
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
          <v-btn color="info" variant="plain" @click="handleDetail(item)">
            查看
          </v-btn>
        </template>
      </v-data-table>
    </v-layout>
  </v-card>

  <v-dialog v-model="detailDialog" fullscreen @after-leave="handleCloseDetailDialog">
    <v-card>
      <v-layout>
        <v-app-bar>
          <template #append>
            <div class="d-flex justify-end">
              <v-btn icon="mdi-window-close" @click="detailDialog=false"></v-btn>
            </div>
          </template>
          <v-app-bar-title>
            {{detailData.name}}
          </v-app-bar-title>
        </v-app-bar>
        <v-card-text class="mt-16">
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
                     v-copy="parsedDetailData.connectionCode">
                {{parsedDetailData.connectionCode}}
              </v-btn>
            </v-col>
          </v-row>
          <v-divider class="mt-8 mb-8"></v-divider>
          <v-row class="d-flex justify-center">
            <v-data-table :headers="playerHeaders" :loading="detailLoading"
                          :items="parsedDetailData.players" class="w-75">
              <template v-slot:loading>
                <v-skeleton-loader type="table-row@4"></v-skeleton-loader>
              </template>
              <template #top>
                <div class="d-flex justify-center text-h6">在线玩家</div>
              </template>
              <template #no-data>
                没有发现在线玩家
              </template>
              <template #item.prefab="{value}">
                {{getValueOrKey(prefabMap, value)}}
              </template>
              <template #item.color="{value}">
                <v-icon icon="mdi-checkbox-blank" :color="'#'+value"></v-icon>
              </template>
            </v-data-table>
          </v-row>
          <v-divider class="mt-8 mb-8"></v-divider>
          <v-row>
            <v-col class="d-flex justify-center">
              <v-data-table :headers="modHeaders" :items="modInfo" :loading="modLoading" class="w-75">
                <template v-slot:loading>
                  <v-skeleton-loader type="table-row@4"></v-skeleton-loader>
                </template>
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

      </v-layout>
    </v-card>
  </v-dialog>

</template>

<script setup>
import router from "../router/index.js";
import {useStore} from '../store/index.js'
import {ref} from "vue";
import Api from "../api/index.js";


const store = useStore()

const gotoHome = () => {
  store.clean()
  router.push('/')
}

const detailData = ref({})
const parsedDetailData = ref({
  dats: 0,
  connectionCode: '',
  players: [],
  mods: [],
})
const detailLoading = ref(false)

const detailDialog = ref(false)

const handleDetail = (row) => {
  const reqForm = {
    rowId: row.__rowId,
    region: row.region
  }

  detailLoading.value = true
  detailDialog.value = true
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

const prefabMap = {
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
const modLoading = ref(false)

const handleGetModInfo = () => {
  modLoading.value = true
  const reqForm = {
    mod: parsedDetailData.value.mods
  }

  Api.modInfo(reqForm).then(response => {
    modInfo.value = response.data
  }).finally(() => {
    modLoading.value = false
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

<style scoped>
</style>