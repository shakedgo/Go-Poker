import { defineStore } from 'pinia'
import { useFetch } from '@/composables/api'

interface AuthState {
  logged: boolean
  loginForm: boolean
}

export const useAuthStore = defineStore('auth', {
  state: (): AuthState => ({
    loginForm: true,
    logged: false
  }),
  getters: {
    isLoggedIn(): boolean {
      return this.logged
    }
  },
  actions: {
    async login(form: { username: string; password: string }) {
      try {
        const res = await useFetch('login', form, 'POST')
        this.logged = res.status == 200;
      } catch (err) {
        console.error(err)
      }
    },
    async signup(form: { username: string; password: string }) {
      try {
        const res = await useFetch('signup', form, 'POST')
        this.logged = res.status == 200;
      } catch (err) {
        console.error(err)
      }
    },
    async logout() {
      try {
        const res = await useFetch('logout', {}, 'POST')
        if (res.status === 200) {
          this.logged = false
          // Optionally, you can clear any user-related data or perform additional cleanup
        }
      } catch (err) {
        console.error(err)
      }
    },
    toggleLoginForm() {
      this.loginForm = !this.loginForm
    },
    initializeStore() {
      const jwtToken = this.getCookie('jwt_token')
      this.logged = !!jwtToken // Set logged state based on the presence of the JWT token cookie
    },
    getCookie(name: string): string | null {
      const cookies = document.cookie.split(';')
      for (let i = 0; i < cookies.length; i++) {
        const cookie = cookies[i].trim()
        if (cookie.startsWith(name + '=')) {
          return cookie.substring(name.length + 1)
        }
      }
      return null
    }
  },
})

const authStore = useAuthStore()
authStore.initializeStore()