<template>
    <el-table :data="tasks" border style="width: 100%">
        <el-table-column
                fixed
                prop="id"
                label="Id"
                width="50">
        </el-table-column>
        <!--<el-table-column-->
                <!--prop="name"-->
                <!--label="Name"-->
                <!--width="170">-->
        <!--</el-table-column>-->
        <!--<el-table-column-->
                <!--prop="commandConsole"-->
                <!--label="Command Console"-->
                <!--width="170">-->
        <!--</el-table-column>-->
        <!--<el-table-column-->
                <!--prop="timeStart"-->
                <!--label="Time Start"-->
                <!--width="120">-->
        <!--</el-table-column>-->
        <!--<el-table-column-->
                <!--prop="timeOut"-->
                <!--label="Time Out"-->
                <!--width="120">-->
        <!--</el-table-column>-->
        <!--<el-table-column-->
                <!--fixed="right"-->
                <!--label="Operations"-->
                <!--width="120">-->
            <!--<template scope="scope">-->
                <!--<el-button @click="handleClick" type="text" size="small">Detail</el-button>-->
                <!--<el-button type="text" size="small">Edit</el-button>-->
            <!--</template>-->
        <!--</el-table-column>-->
    </el-table>
</template>
<style>

</style>
<script>

    export default{
        data(){
            return{
                tasks:[]
            }
        },
        mounted(){
            this.$socket.emit('list:queues', function(data){
                console.log('list task: ', JSON.parse(data))

                this.tasks = JSON.parse(data)
            })
        }
    }
</script>
