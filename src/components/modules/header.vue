<template>
    <header id="app-header" class="modal-card-head">
        <p class="modal-card-title">{{title}}</p>
    </header>
</template>

<script>
export default {
    name :"app-header",
    props:['title'],
    data() {
        return {
        }
    }
}
</script>

<style>
    #app-header{
        display: block;
        position: fixed;
        width: 100%;
        top: 0px;
        left:0px;
        margin-bottom: 5%;
    }
</style>