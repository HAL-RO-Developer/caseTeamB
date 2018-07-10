<template>
    <div class="modal-card" style="width: auto">
        <header class="modal-card-head">
            <span class="title is-6">メッセージ</span>
        </header>
        <section class="modal-card-body">
            <b-field label="回数:">
                <b-input type="number" v-model="data.message_call" min="1"></b-input>
            </b-field>
            <b-field label="メッセージ:">
                <b-input maxlength="140" type="textarea" v-model="data.message" required></b-input>
            </b-field>
        </section>
        <footer class="modal-card-foot">
            <button class="button" type="button" @click="$parent.close()">キャンセル</button>
            <button class="button is-primary" @click="add">追加</button>
        </footer>
    </div>
</template>

<script>
export default {
  data() {
    return {
      data: {
        goal_id: "",
        message_call: null,
        message: ""
      },
    };
  },
  methods: {
    add() {
      this.$emit("add", this.data);
    }
  }
};
</script>