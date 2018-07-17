<template>
    <div class="modal-card" style="width: auto">
        <header class="modal-card-head">
            <span class="title is-6">目標実行数変更</span>
        </header>
        <section class="modal-card-body">
            <b-field label="回数:">
                <b-input type="number" v-model="data.message_call" min="1"></b-input>
            </b-field>
        </section>
        <footer class="modal-card-foot">
            <button class="button" type="button" @click="$parent.close()">キャンセル</button>
            <button class="button is-primary" @click="change">変更</button>
        </footer>
    </div>
</template>

<script>
export default {
  data() {
    return {
      data: {
        goal_id: "",
        message: 0
      },
    };
  },
  methods: {
    change() {
      this.$emit("change", this.data);
    }
  }
};
</script>