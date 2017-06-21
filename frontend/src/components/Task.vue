<template>
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
                <el-button type="text" @click="openEdit" size="small">Edit</el-button>
                <el-button  type="text" size="small">Detail</el-button>
            </template>
        </el-table-column>
    </el-table>

        <el-dialog title="Edit" :visible.sync="modalEdit">
            <el-form :model="form">
                <el-form-item label="Name" :label-width="formLabelWidth">
                    <el-input v-model="form.name" auto-complete="off"></el-input>
                </el-form-item>
                <el-form-item label="Command console" :label-width="formLabelWidth">
                    <el-input v-model="form.console" auto-complete="off"></el-input>
                </el-form-item>
                <el-form-item label="Command console" :label-width="formLabelWidth">
                    <el-input v-model="form.console" auto-complete="off"></el-input>
                </el-form-item>

                <el-form-item label="Zones" :label-width="formLabelWidth">
                    <el-select v-model="form.region" placeholder="Please select a zone">
                        <el-option label="Zone No.1" value="shanghai"></el-option>
                        <el-option label="Zone No.2" value="beijing"></el-option>
                    </el-select>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
    <el-button @click="modalEdit = false">Cancel</el-button>
    <el-button type="primary" @click="modalEdit = false">Confirm</el-button>
    </span>
    </el-dialog>
    </el-col>
</template>
<style>

</style>
<script>

    export default{
        data(){
            return{
                tasks:[],
                modalEdit:true,
                form: {
                    name: '',
                    console:'',
                    region: '',
                    date1: '',
                    date2: '',
                    delivery: false,
                    type: [],
                    resource: '',
                    desc: ''
                },
                formLabelWidth: '120px'
            }
        },
        methods : {
        },
        mounted(){
            this.$socket.emit('list:queues', function(data){
                console.log('list task: ', JSON.parse(data))

                this.tasks = JSON.parse(data)
            }.bind(this))
        }
    }
</script>
