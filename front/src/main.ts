import { createApp, VueElement } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import ElementPlus from 'element-plus'
import { Mgr } from './ocid/oidcConfig'
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import { ITokenService, TokenService } from './service/TokenService'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(ElementPlus)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

router.beforeEach((to, from, next) => {
  const tokenService: ITokenService = new TokenService()
  if (to.name == "signin-oidc" || tokenService.get()) {
    next()
  } else {
    Mgr.signinRedirect();
  }
})

app.mount('#app')
