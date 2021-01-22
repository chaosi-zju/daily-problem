import Vue from 'vue'
import Element from 'element-ui'
import VueParticles from 'vue-particles'
import VueClipboard from 'vue-clipboard2'
import mavonEditor from 'mavon-editor'

import App from '@/App.vue'
import router from '@/router/router'
import store from '@/store/store'

import i18n from '@/assets/lang'
import {messages} from '@/assets/js/common.js'
import '@/assets/icon/iconfont.js'
import '@/assets/icon/iconfont.css'
import '@/assets/css/public.css'
import '@/assets/scss/element-variables.scss'
import 'mavon-editor/dist/css/index.css'

Vue.use(VueParticles)
Vue.use(VueClipboard)
Vue.use(Element)
Vue.use(mavonEditor)
Vue.prototype.$messages = messages
Vue.config.productionTip = false

new Vue({
    i18n,
    router,
    store,
    render: h => h(App)
}).$mount('#app')