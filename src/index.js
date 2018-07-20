import Vue from 'vue'
import VueRouter from 'vue-router'
import dashboad from './components/pages/dashboad.vue'
import Login from './components/pages/login.vue'
import Device from './components/pages/device.vue'
import Goals from './components/pages/goalsettings.vue'
import GoalList from './components/pages/goallist.vue'
import Graph from './components/pages/goalgraph.vue'
import Children from './components/pages/children.vue'
import Messages from './components/pages/messages.vue'
import BOCCO from './components/pages/bocco.vue'
import Settings from './components/pages/settings.vue'
import NotFound from './components/pages/notFound.vue'
import Buefy from 'buefy'
import 'buefy/lib/buefy.css'

Vue.use(Buefy)
Vue.use(VueRouter)

const routes = [
    { path: "/", component: dashboad },
    { path: "/login", component: Login },
    { path: "/device", component: Device },
    { path: "/goals", component: Goals },
    { path: "/goals/list", component: GoalList },
    { path: "/goals/graph", component: Graph },
    { path: "/children", component: Children },
    { path: "/messages", component: Messages },
    { path: "/bocco", component: BOCCO },
    { path: "/settings", component: Settings },
    { path: "*", component: NotFound },
]
const router = new VueRouter({ mode: 'history', routes })

new Vue({
    router
}).$mount("#app")
