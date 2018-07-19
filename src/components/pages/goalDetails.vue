<template>
    <div class="modal-card" style="width: auto">
        <app-header :title='title' id="titlegoaldeta"></app-header>
        <div class="contents">
            <b-field>
                <b-select placeholder="Select a filter" v-model="filter" @input="getRecords"> 
                    <option v-for="option in options.filter" :key="option.value" :value="option.value">{{option.name}}</option>
                </b-select>
            </b-field>
            <by-date ref="date" v-if="filter=='date'" :isLoading="isLoading" @isLoading="isLoading=false"></by-date>
            <by-month ref="month" v-if="filter=='month'" :isLoading="isLoading" @isLoading="isLoading=false"></by-month>
        </div>
        <fab icon="sync" @click="getRecords"></fab>    
        <app-footer id="footergoaldeta"></app-footer>
        <under-tab :index='1'></under-tab>
    </div>
</template>
<script>
import http from "../../service/service";
import UnderTab from "../modules/underTab.vue";
import AppHeader from "../modules/header.vue";
import AppFooter from "../modules/footer.vue";
import Fab from "../modules/fab.vue";
import ByDate from "../modules/goaldetailsByDate.vue";

export default {
  name: "records",
  components: {
    UnderTab,
    AppHeader,
    AppFooter,
    Fab,
    ByDate
  },
  data() {
    return {
      title: "目標総数",
      child_id: "",
      filter: "date",
      goals: [],
      options: {
        filter: [
          { name: "日別", value: "date" },
          { name: "月別", value: "month" }
        ]
      },
      isLoading: false
    };
  },
  methods: {
    getGoal() {
      this.isLoading = true;
      http
        .getGoal()
        .then(response => {
          console.log(response);
          this.isLoading = false;
          this.goals = response.data.goals;
        })
        .catch(err => {
          this.isLoading = false;
          if (err) {
            this.$dialog.alert({
              title: "Error",
              message: err.response.data.error,
              type: "is-danger",
              hasIcon: true,
              icon: "times-circle",
              iconPack: "fa"
            });
            switch (err.response.status) {
              case 401:
                http.RemoveToken();
                this.$router.push({ path: "/" });
                break;
              default:
                break;
            }
          }
        });
    }
  },
  created() {
    this.child_id = localStorage.getItem("child_id");
    this.getGoals();
  }
};
</script>

<style>
#titlegoaldeta,#footergoaldeta{
background: coral
}
</style>
