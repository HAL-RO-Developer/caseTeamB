<template>
    <div>
      <header class="modal-card-head" style="background: lightcoral">
                <p class="modal-card-title">{{title}}</p>
            </header>

        <div class="contents">
            <div class="buttons" v-for="item in menu_items" :key="item.id">
                <span class="button full-width" @click="click(item.id)">
                  <b-icon :icon = item.icon id="icon">
                        </b-icon>{{item.title}}</span>
            </div>
        </div>
        <under-tab :index='1'></under-tab>
    </div>
</template>
<script>
import http from "../../service/service";
import UnderTab from "../modules/underTab.vue";
import AppHeader from "../modules/header.vue";
import AppFooter from "../modules/footer.vue";

export default {
  name: "Goalsettings",
  components: {
    UnderTab,
    AppHeader,
    AppFooter
  },
  data() {
    return {
      title: "目標",
      menu_items: [
<<<<<<< HEAD
        { id: 1, title: "目標一覧" , icon: "briefcase" },
        { id: 2, title: "目標実行数変更", icon: "redo" },
        { id: 3, title: "子ども別グラフ", icon: "chart-bar" }
=======
        { id: 1, title: "目標一覧" },
        { id: 2, title: "子ども別グラフ" }
>>>>>>> c69722da1625f61c3897ce9fd63b8f0962bdf7c6
      ]
    };
  },
  methods: {
    click(id) {
      switch (id) {
        case 1:
          this.$router.push({ path: "/goals/list" });
          break;
        case 2:
          this.$router.push({ path: "/goals/graph" });
          break;
        default:
          break;
      }
    }
  }
};
</script>

<style>
#icon{
  margin-right: 2%
}
</style>
