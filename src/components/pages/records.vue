<template>
    <div class="modal-card" style="width: auto">
        <app-header :title='title'></app-header>
        <div class="contents">
            <b-field grouped group-multiline>
                <div class="is-left"><span class="button" @click="prev">＜</span></div>
                <div class="is-center"><b-input type="text" v-model="values.week_range" readonly></b-input></div>
                <div class="is-right"><span class="button" @click="next">＞</span></div>
            </b-field>
            <div style="position:relative;">
                <b-loading :is-full-page="false" :active.sync="isLoading" :can-cancel="true"></b-loading>
                <graph :chartData='chartData' :options='options' :width="900" :height="750"></graph>
            </div>
        </div>
        <fab :icon="fabIcon" @click="getRecords"></fab>
        <app-footer></app-footer>
        <under-tab :index='1'></under-tab>
    </div>
</template>
<script>
import moment from "moment";
import http from '../../service/service';
import UnderTab from '../modules/underTab.vue'
import AppHeader from '../modules/header.vue'
import AppFooter from '../modules/footer.vue'
import Graph from '../modules/graph.vue'
import Fab from '../modules/fab.vue'

export default {
    name :"records",
    components:{
        UnderTab,
        AppHeader,
        AppFooter,
        Graph,
        Fab
    },
    data() {
        return {
            title: "目標",
            child_id: "",
            filter: this.$route.params.filter,
            chartData:{},
            options:{},
            records:[],
            values:{
                week:[],
                week_range:"",
                solved:[],
                correct:[]
            },
            isLoading: false,
            fabIcon: "sync",
        }
    },
    methods:{
        fillData () {
            var solved_data = this.values.solved
            var correct_data = this.values.correct
            var rate_data = this.values.rate
            
            var datasets = [
                {
                    label: '回答数',
                    type: 'bar',
                    data: solved_data,
                    borderColor: "rgba(254,97,132,0.8)",
                    backgroundColor: "rgba(254,97,132,0.5)",
                    yAxisID: "count-axis"
                },
                {
                    label: '正答数',
                    type: 'line',
                    data: correct_data,
                    borderColor: "rgba(54,164,235,1.0)",
                    pointBackgroundColor: "rgba(54,164,235,1.0)",
                    fill: false,
                    yAxisID: "count-axis"
                }
            ]

            this.chartData = {
                labels: this.values.week,
                datasets: datasets                 
            }

            this.options = {
                scales:{
                    yAxes:[
                        {
                            id: "count-axis",
                            position: "left",
                            ticks:{
                                min: 0,
                                max: 30,
                            }
                        }
                    ],
                },
                onClick: 
                    (e, el)=>{
                        this.showSolvedList(e, el)
                    } 
            }
            this.isLoading = false
        },
        showSolvedList(e,el){
            if (! el || el.length == 0) return;
            var date = moment(el[0]._model.label,"MM/DD").format("MMDD")
            this.$router.push({path: '/records/date/' + date})
        },
        getRecords(){
            this.isLoading = true
            http.getRecords(this.child_id,this.filter)
                .then((response)=>{
                    console.log(response)
                    var records = response.data.records
                    this.records = records
                    this.aggregate()  
                })
                .catch((err)=>{
                    this.isLoading = true
                    if(err){
                        this.$dialog.alert({
                            title: 'Error',
                            message: err.response.data.error,
                            type: 'is-danger',
                            hasIcon: true,
                            icon: 'times-circle',
                            iconPack: 'fa'
                        })
                        switch(err.response.status){
                            case 401:
                                http.RemoveToken()
                                this.$router.push({path:'/'})
                                break;
                            default:
                                break;
                        }
                    }
                });
        },
        aggregate(date){
            var m = moment(date).day(1)
            var records = this.records
            for( var i = 0; i < 7; i++, m.add(1,"day") ){
                this.values.week[i] = m.format('MM/DD')
                this.values.solved[i] = 0
                this.values.correct[i] = 0
                records.forEach((record)=>{
                    if( moment(record.date).isSame(m,'day') ){
                        this.values.solved[i] = Number(record.solved)
                        this.values.correct[i] = Number(record.correct)
                    }
                })                
            }
            this.values.week_range = this.values.week[0] + " ~ " + this.values.week[i-1]
            this.fillData()
        },
        prev(){
            this.aggregate(moment(this.values.week[0],'MM/DD').add(-7,'day').toDate())
        },
        next(){
            var date = moment(this.values.week[0],'MM/DD').add(7,'day')
            if(date.isBefore(moment())){
                this.aggregate(date.toDate())
            }
            
        }
    },
    created() {
        this.child_id = localStorage.getItem('child_id')
        this.getRecords()
    }
}
</script>

<style>
</style>
