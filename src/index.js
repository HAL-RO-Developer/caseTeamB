import Vue from 'vue'
import VueRouter from 'vue-router'
import dashboad from './components/pages/dashboad.vue'
import Login from './components/pages/login.vue'
import Device from './components/pages/device.vue'
import Records from './components/pages/records.vue'
import SolvedList from './components/pages/solvedList.vue'
import Settings from './components/pages/settings.vue'
import Children from './components/pages/children.vue'
import Goals from './components/pages/goals.vue'
import Messages from './components/pages/messages.vue'
import Manual from './components/pages/Manual.vue'
import NotFound from './components/pages/notFound.vue'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'

Vue.use(Buefy)
Vue.use(VueRouter)

const routes = [
    { path: "/", component: dashboad },
    { path: "/login", component: Login },
    { path: "/device", component: Device },
    { path: "/records/:filter", component: Records },
    { path: "/records/:filter/:date", component: SolvedList },
    { path: "/settings", component: Settings },
    { path: "/children", component: Children },
    { path: "/goals", component: Goals },
    { path: "/messages", component: Messages },
    { path: "/manual", component: Manual },
    { path: "*", component: NotFound },
]
const router = new VueRouter({ mode: 'history', routes })

new Vue({
    router
}).$mount("#app")
