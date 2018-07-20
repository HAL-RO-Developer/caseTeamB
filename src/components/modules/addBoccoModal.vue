<template>
    <div class="modal-card" style="width: auto">
        <header class="modal-card-head">
            <span class="title is-6">BOCCO設定</span>
        </header>
        <section class="modal-card-body">
            <b-field label="メールアドレス:">
                <b-input type="email" v-model="data.email" icon="email"></b-input>
            </b-field>
            <b-field label="パスワード:">
                <b-input type="password" v-model="data.pass" password-reveal icon="key"></b-input>
            </b-field>
        </section>
        <footer class="modal-card-foot">
            <button class="button" type="button" @click="$parent.close()">キャンセル</button>
            <button class="button is-primary" @click="add">設定</button>
        </footer>
    </div>
</template>

<script>
export default {
  data() {
    return {
      data: {
        email: "",
        pass:""
      }
    };
  },
  methods: {
    add() {
      this.$emit("add", this.data);
    }
  }
};
</script>