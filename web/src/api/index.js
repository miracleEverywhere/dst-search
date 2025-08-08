import http from "@/utils/axios.js";


export default {
  search: async function (data) {
    return await http.post('/search', data);
  },
}