<template>
    <div class="modal-card" style="width: auto">
        <app-header :title='title'></app-header>
        <div class="contents">
            <b-loading :is-full-page="false" :active.sync="isLoading" :can-cancel="true"></b-loading>
            <div v-for='goal in goals'>
            <card v-for='(child_goal, index) in goal.child_goals'
                :key="index"
                :goalid="child_goal.goal_id"
                :nickname="goal.nickname"
                :content="child_goal.content" 
                :run="child_goal.run"
                :selected="selected"
                @remove="removeGoal"
                @select="select"></card>
            </div>
        </div>
        <fab :icon="fabIcon" @click="isComponentModalActive = true"></fab>
        <app-footer></app-footer>
        <under-tab :index='1'></under-tab>
        <b-modal :active.sync="isComponentModalActive" has-modal-card>
            <modal-form @add="addGoal"></modal-form>
        </b-modal>
        
    </div>
</template>
<script>
import moment from "moment";
import http from "../../service/service";
import UnderTab from "../modules/underTab.vue";
import AppHeader from "../modules/header.vue";
import AppFooter from "../modules/footer.vue";
import Card from "../modules/goalCard.vue";
import ModalForm from "../modules/addGoalModal.vue";
import Fab from "../modules/fab.vue";

export default {
  name: "goals",
  components: {
    UnderTab,
    AppHeader,
    AppFooter,
    Card,
    ModalForm,
    Fab
  },
  data() {
    return {
      title: "目標一覧",
      fabIcon: "plus",
      goals: [],
      //child_goals: [],
      selected: "",
      isComponentModalActive: false,
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
    },
    addGoal(data) {
      this.isComponentModalActive = false;
      var deadline = moment(data.deadline);
      var child_id = localStorage.getItem("child_id");
      console.log(data);
      http
        .addGoal(
          Number(child_id),
          data.content,
          Number(data.criteria),
          deadline.format("YYYY-MM-DD")
        )
        .then(response => {
          this.getGoal();
        })
        .catch(err => {
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
    },
    removeGoal(id, content) {
      this.$dialog.confirm({
        title: "目標削除",
        message: "『" + content + "』を削除しますか？",
        confirmText: "削除",
        type: "is-danger",
        hasIcon: true,
        onConfirm: () =>
          http
            .removeGoal(id)
            .then(response => {
              var goal_id = localStorage.getItem("goal_id");
              if ((goal_id = id)) {
                localStorage.removeItem("goal_id");
              }
              this.getGoal();
            })
            .catch(err => {
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
            })
      });
    },
    select(id) {
      localStorage.setItem("goal_id", id);
      this.selected = id;
    }
  },
  created() {
    this.selected = localStorage.getItem("goal_id");
    this.getGoal();
  }
};
</script>

<style>
</style>