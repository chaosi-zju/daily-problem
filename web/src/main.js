import Vue from 'vue'
import Element from 'element-ui'
import VueParticles from 'vue-particles'
import VueClipboard from 'vue-clipboard2'

import App from '@/App.vue'
import router from '@/router/router'
import store from '@/store/store'

import i18n from '@/assets/lang'
import {messages} from '@/assets/js/common.js'
import '@/assets/icon/iconfont.js'
import '@/assets/icon/iconfont.css'
import '@/assets/css/public.css'
import '@/assets/scss/element-variables.scss'

Vue.use(VueParticles)
Vue.use(VueClipboard)
Vue.use(Element)
Vue.prototype.$messages = messages
Vue.config.productionTip = false

new Vue({
    i18n,
    router,
    store,
    render: h => h(App)
}).$mount('#app')