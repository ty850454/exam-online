import {computed, ref} from 'vue'
import {defineStore} from 'pinia'
import {post} from '@/utils/http'

export const useUserStore = defineStore('user', () => {

    const token = ref(localStorage.getItem('token'))
    const isLogin = computed(() => token.value !== null)

    function login(username, password) {
        return post('/api/login',
            {
                "username": username,
                "password": password
            })
            .then(response => {
                token.value = response['token']
                localStorage.setItem('token', token.value)
            })
    }

    function logout(username, password) {
        token.value = null
        localStorage.removeItem('token')
        location.reload()
    }

    return {isLogin, login, logout}
})
