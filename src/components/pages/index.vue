<template>
  <div>
    <app-header ></app-header>
    <under-tab></under-tab>
    <app-footer></app-footer>
  </div>
</template>

<script>
import http from "../../service/service";
import UnderTab from "../modules/underTab.vue";
import AppHeader from "../modules/header.vue";
import AppFooter from "../modules/footer.vue";

export default {
  name: "index",
  components: {
    AppHeader,
    AppFooter,
    UnderTab
  },
  data() {
    return {};
  },
  methods: {},
  created() {
    if (!http.GetToken()) {
      this.$router.push({ path: "/login" });
    } else if (!localStorage.getItem("child_id")) {
      this.$router.push({ path: "/children" });
    } else {
      this.$router.push({ path: "/goals" });
    }
  }
};
</script>