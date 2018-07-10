<template>
    <div class="modal-card" style="width: auto">
        <header class="modal-card-head">
            <p class="modal-card-title">目標追加</p>
        </header>
        <section class="modal-card-body">
            <b-field label="目標名">
                <b-input
                    type="text"
                    v-model="data.content"
                    placeholder="目標名"
                    required>
                </b-input>
           </b-field>
           <b-field label="達成目標数">
             <b-input
              type="number"
              v-model="data.criteria"
              placeholder="達成目標数"
              required>
              </b-input>
            </b-field>
          <b-field label="期限">
            <b-datepicker
              placeholder="Click to select..."
              icon="calendar-today"
              :readonly="false"
              v-model="data.deadline"
              >
            </b-datepicker>
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
        child_id: null,
        content: "",
        criteria: null,
        deadline: new Date()
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