<template>
    <div >
                    <header class="modal-card-head" style="background: coral">
                <p class="modal-card-title">{{title}}</p>
            </header>
        <div class="contents">
            <b-field>
              <b-select placeholder="Select a child" v-model="child_id" @input="aggregate"> 
                    <option v-for="option in options.children" :key="option.child_id" :value="option.child_id">{{option.nickname}}</option>
                </b-select>
            </b-field>
            <!--<by-date ref="date" v-if="filter=='date'" :isLoading="isLoading" @isLoading="isLoading=false"></by-date>-->
            <!--<by-month ref="month" v-if="filter=='month'" :isLoading="isLoading" @isLoading="isLoading=false"></by-month>-->
            <graph :chartData='chartData' :options='settings' :width="900" :height="750"></graph>
        </div> 
        <app-footer style="background: coral"></app-footer>
        <under-tab :index='1'></under-tab>
    </div>
</template>
<script>
import http from "../../service/service";
import UnderTab from "../modules/underTab.vue";
import AppHeader from "../modules/header.vue";
import AppFooter from "../modules/footer.vue";
import Fab from "../modules/fab.vue";
import graph from "../modules/graph.vue";

export default {
  name: "records",
  components: {
    UnderTab,
    AppHeader,
    AppFooter,
    Fab,
    graph
  },
  data() {
    return {
      title: "子ども別グラフ",
      child_id: "",
      filter: "date",
      goals: [],
      options: {
        children: []
      },
      isLoading: false,

      chartData: {},
      settings: {}
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
          this.aggregate();
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
    getChild() {
      http
        .getChild()
        .then(response => {
          this.options.children = response.data.children;
        })
        .catch(err => {
          this.isLoading = false;
          if (err) {
            this.$dialog.alert({
              title: "Error",
              message: err.response.data.error,
              type: "is-danger",
              hasIcon: true,
              icon: "alert-circle"
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
    aggregate() {
      var labels = [];
      var data = [13, 20, 16];

      var child = this.goals.find(item => {
        if (item.child_id == this.child_id) return true;
      });
      child.child_goals.forEach(item => {
        labels.push(item.content);
        //data.push(item.run)
      });

      console.log(labels);
      console.log(data);
      var datasets = [
        {
          label: "回数",
          type: "bar",
          data: data
        }
      ];

      this.chartData = {
        labels: labels,
        datasets: datasets
      };
      this.settings = {
        scales: {
          yAxes: [
            {
              ticks: {
                min: 0
              }
            }
          ]
        }
      };
    }
  },
  created() {
    this.child_id = localStorage.getItem("child_id");
    this.getGoal();
    this.getChild();
  }
};
</script>


