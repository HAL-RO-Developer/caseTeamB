<template>
    <footer id="app-footer" class="modal-card-foot">
        <div class="buttons is-left">
            <span class="button" @click="$router.go(-1)">戻る</span>
        </div>
    </footer>
</template>

<script>
export default {
    name :"app-footer",
    props:[],
    data() {
        return {
        }
    }
}
</script>

<style>
    #app-footer{
        display: block;
        position: fixed;
        width: 100%;
        bottom: 40px;
        left:0px;
    }
</style>