<template>
  <form action="">
    <div>
      <header class="modal-card-head" style="background: mistyrose">
        <p class="modal-card-title">ログイン</p>
      </header>
      <section class="modal-card-body">
        <b-field label="User">
          <b-input
            type="text"
            icon="account"
            v-model="name"
            placeholder="ユーザー名"
            required>
          </b-input>
        </b-field>
        <b-field label="Password">
          <b-input
            type="password"
            icon="key"
            v-model="password"
            password-reveal
            placeholder="パスワード"
            required>
          </b-input>
        </b-field>
      </section>
      <footer class="modal-card-foot" style="background: mistyrose">
        <button class="button" type="button" @click="signup"><b-icon icon = "pen" id="icon"></b-icon>新規登録</button>
        <button class="button" type="button" @click="signin"><b-icon icon = "login" id="icon"></b-icon>ログイン</button>
      </footer>
    </div>
  </form>
</template>

<script>
import axios from "axios";
import http from "../../service/service";
export default {
  data() {
    return {
      name: "",
      password: "",
      flag: 0
    };
  },
  methods: {
    signup() {
      http
        .signup(this.name, this.password)
        .then(response => {
          console.log(response);
          this.$dialog.alert({
            title: "ユーザー作成",
            message: response.data.success,
            type: "is-info",
            hasIcon: true,
            icon: "times-circle",
            iconPack: "fa"
          });
          this.flag = 1;
          this.signin();
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
    signin() {
      http
        .signin(this.name, this.password)
        .then(response => {
          http.SetToken(response.data.token);
          if (this.flag == 1) {
            this.$router.push({ path: "/children" });
          } else {
            this.$router.push({ path: "/" });
          }
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
                this.$router.push({ path: "/" });
                break;
              default:
                break;
            }
          }
        });
    }
  }
};
</script>

<style>
#icon {
  margin-right: 2%;
}
</style>