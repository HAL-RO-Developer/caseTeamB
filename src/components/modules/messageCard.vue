<template>
    <div class="card">
        <div class="card-content" >
            <div class="media">
                <div class="media-left">
                    
                </div>
                <div class="media-content">
                    <p class="title is-4">{{content}} </p>
                    <p class="subtitle is-6">メッセージ:{{message}} {{call}}回</p>
                </div>
                <div class="media-right">
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
                options: {
                    condition:[

                    ]
                }
            }
        },
        methods:{
            edit(){},
            remove(){
                this.$emit('remove', this.id,this.call,this.content )
            }
        },
        props:['content','call','message','id']
    }
</script>