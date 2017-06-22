<template>
    <div>
        <el-button @click="createTask" type="primary" icon="plus">Create</el-button>
    <el-col :span="16" :offset="3" >

    <el-table :data="tasks" border style="width: 100%">
        <el-table-column
                prop="id"
                label="Id"
                width="50">
        </el-table-column>
        <el-table-column
                prop="name"
                label="Name"
                width="170">
        </el-table-column>
        <el-table-column
                prop="commandConsole"
                label="Command Console"
                width="200">
        </el-table-column>
        <el-table-column
                prop="timeStart"
                label="Time Start"
                width="200">
        </el-table-column>
        <el-table-column
                prop="timeOut"
                label="Time Out"
                width="200">
        </el-table-column>
        <el-table-column
                label="Operations"
                width="150">
            <template scope="scope">
                <el-button type="text" @click="openEdit(scope.row)" size="small">Edit</el-button>
            </template>
        </el-table-column>
    </el-table>
    </el-col>
    <el-dialog  :title=" buttonModal ? 'Create Task' : 'Edit Task'" :visible.sync="modalEdit">
        <el-form :model="form">
            <el-form-item label="Name" :label-width="formLabelWidth">
                <el-input v-model="form.name" auto-complete="off"></el-input>
            </el-form-item>
            <el-form-item label="Command console" :label-width="formLabelWidth">
            <el-input v-model="form.console" auto-complete="off"></el-input>
            </el-form-item>
            <el-form-item label="Time start" :label-width="formLabelWidth">
                <el-date-picker v-model="form.timeStart" type="datetime" placeholder="Select date and time"></el-date-picker>
            </el-form-item>

            <el-form-item label="Time Out" :label-width="formLabelWidth">
                <el-date-picker v-model="form.timeOut" type="datetime" placeholder="Select date and time"></el-date-picker>
            </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
    <el-button @click="modalEdit = false">Cancel</el-button>
    <el-button type="primary" @click="createTask((buttonModal ? 'save' : 'update'))">{{ buttonModal ? 'save' : 'update'}}</el-button>
    </span>
    </el-dialog>
    </div>
</template>
<style>

</style>
<script>

    export default{
        data(){
            return{
                tasks:[],
                modalEdit:false,
                buttonModal:true,
                form: {
                    name: '',
                    console:'',
                    timeStart: '',
                    timeOut: '',
                },
                formLabelWidth: '100px'
            }
        },
        methods : {
            createTask(typeAction){
                this.form.name      = ''
                this.form.console   = ''
                this.form.timeStart = ''
                this.form.timeOut   = ''
                this.buttonModal    = true
                this.modalEdit      = true
                switch (typeAction) {
                    case 'save' :
                    case 'update':
                }
            },
            openEdit(row){
                this.form.name      = row.name
                this.form.console   = row.commandConsole
                this.form.timeStart = row.timeStart
                this.form.timeOut   = row.timeOut

               this.buttonModal    = false
               this.modalEdit = true
            }
        },
        mounted(){
            this.$socket.emit('list:queues', function(data){
                console.log('list task: ', JSON.parse(data))

                this.tasks = JSON.parse(data)
            }.bind(this))
        }
    }
</script>
