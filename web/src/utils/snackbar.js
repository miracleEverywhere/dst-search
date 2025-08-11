// src/utils/snackbar.js
import { ref } from 'vue';

// 定义 Snackbar 的状态
const snackbar = ref(false);
const TYPE = ref('success')
const COLOR = ref('success');
const text = ref('');
const location = ref('top');
const timeout = ref(3000);
const icon = ref('ri-checkbox-circle-fill')

/**
 * 显示 Snackbar 的函数
 * @param message 要显示的内容
 * @param type 默认值为success，一共有4种：[success, info, warning, error]
 */
export function showSnackbar(message, type = 'success') {
    snackbar.value = true;
    TYPE.value = type
    text.value = message;
    if (type === 'success') {
        COLOR.value = type
        icon.value = 'mdi-check-circle'
    }
    if (type === 'warning') {
        COLOR.value = type
        icon.value = 'mdi-alert-circle'
    }
    if (type === 'error') {
        COLOR.value = '#F56C6C'
        icon.value = 'mdi-close-circle'
    }
    if (type === 'info') {
        COLOR.value = type
        icon.value = 'mdi-information'
    }
}

// 提供 Snackbar 的状态和方法
export function useSnackbar() {
    return {
        snackbar,
        TYPE,
        COLOR,
        text,
        location,
        icon,
        timeout,
        showSnackbar,
    };
}