<template>
  <div>
    <header class="modal-card-head" style="background: lime">
      <p class="modal-card-title">{{title}}</p>
    </header>

    <div class="contents">
        <b-loading :is-full-page="false" :active.sync="isLoading" :can-cancel="true"></b-loading>
        <card
          :email="email"
          @remove="removeBOCCO"></card>
    </div>

    <fab :icon="fabIcon" @click="isComponentModalActive = true"></fab>
    <app-footer style="background: lime"></app-footer>
    <under-tab :index='2'></under-tab>
    <b-modal :active.sync="isComponentModalActive" has-modal-card>
      <modal-form @add="addBOCCO"></modal-form>
    </b-modal>
  </div>
</template>

<script>
import axios from "axios";
import http from "../../service/service";
import UnderTab from "../modules/underTab.vue";
import AppHeader from "../modules/header.vue";
import AppFooter from "../modules/footer.vue";
import Card from "../modules/boccoCard.vue";
import ModalForm from "../modules/addBoccoModal.vue";
import Fab from "../modules/fab.vue";

export default {
  name: "bocco",
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
      title: "BOCCO設定",
      email: "",
      pass: "",
      fabIcon: "plus",
      isComponentModalActive: false,
      isLoading: false
    };
  },
  methods: {
    addBOCCO(data) {
      this.isComponentModalActive = false;
      console.log(data);
      http
        .addBocco(data.email, data.pass)
        .then(response => {
          console.log(response);
          this.getBOCCO();
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
    getBOCCO() {
      this.isLoading = true;
      http
        .getBocco()
        .then(response => {
          this.isLoading = false;
          this.email = response.data.email;
          console.log("email");
          console.log(this.email);
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
    removeBOCCO(email) {
      this.$dialog.confirm({
        title: "登録情報削除",
        message: "『" + email + "』を削除しますか？",
        confirmText: "削除",
        type: "is-danger",
        hasIcon: true,
        onConfirm: () =>
          http
            .removeBocco()
            .then(response => {
              this.isLoading = false;
              console.log("delete");
              this.getBOCCO();
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
            })
      });
    }
  },
  created() {
    this.email = localStorage.getItem("email");
    this.getBOCCO()
  }
};
</script>