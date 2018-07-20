<template>
  <div>
    <header class="modal-card-head" style=" background: royalblue">
      <p class="modal-card-title">{{title}}</p>
    </header>
    <div class="contents"></div>
    <app-footer style=" background: royalblue"></app-footer>
    <under-tab :index='2'></under-tab>
  </div>
</template>

<script>
import http from "../../service/service";
import UnderTab from "../modules/underTab.vue";
import AppHeader from "../modules/header.vue";
import AppFooter from "../modules/footer.vue";

export default {
  name: "settings",
  components: {
    UnderTab,
    AppHeader,
    AppFooter
  },
  data() {
    return {
      title: "取扱説明書"
    };
  },
  methods: {},
  created() {}
};
</script>

