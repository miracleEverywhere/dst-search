import http from "../utils/axios.js";


export default {
  search: async function (data) {
    return await http.post('/search', data);
  },
  detail: async function (data) {
    return await http.post('/detail', data);
  },
  modInfo: async function (data) {
    return await http.post('/mod_info', data);
  },
}