<template>
    <div class="card">
        <div class="card-content" @click="select">
            <div class="media">
                <div class="media-content">
                    <p class="title is-4">{{content}}</p>
                    <p class="subtitle is-6">{{nickname}}:{{run}}回 目標:{{criteria}}回</p>
                </div>
                <b-icon v-if="selected==goalid" icon="check"></b-icon>
                <div class="media-right">
                    <button class="button is-info" @click="info">
                        <b-icon icon="pencil"></b-icon>
                    </button>
                    <button class="button is-danger" @click="remove">
                        <b-icon icon="delete"></b-icon>
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
  data() {
    return {
      icon: {}
    };
  },
  methods: {
    info() {
      this.$emit("info", this.goalid, this.content);
    },
    remove() {
      this.$emit("remove", this.goalid, this.content);
    },
    select() {
      this.$emit("select", this.goalid);
    }
  },
  created() {},
  props: ["goalid", "nickname", "content", "run", "criteria", "selected"]
};
</script>