<template>
 <div class="wx_mail_list">
   <div class="head_con">
      通讯录
      <!-- <el-dropdown class="add_icon">
         <span class="el-dropdown-link">
            <el-icon :size="18" color="#333"><CirclePlus /></el-icon>
         </span>
         <template #dropdown>
            <el-dropdown-menu>
               <el-dropdown-item>Action 1</el-dropdown-item>
               <el-dropdown-item>Action 2</el-dropdown-item>
               <el-dropdown-item>Action 3</el-dropdown-item>
               <el-dropdown-item disabled>Action 4</el-dropdown-item>
               <el-dropdown-item divided>Action 5</el-dropdown-item>
            </el-dropdown-menu>
         </template>
      </el-dropdown> -->
   </div>
   <el-row class="friends_list">
      <el-col :span="24">
         <img src="../assets/logo.png" alt="" class="head_img">
         <div class="content">
            备注
         </div>
      </el-col>
      <el-col :span="24">
         <img src="../assets/logo.png" alt="" class="head_img">
         <div class="content">
            备注
         </div>
      </el-col>
      <el-col :span="24">
         <img src="../assets/logo.png" alt="" class="head_img">
         <div class="content">
            备注
         </div>
      </el-col>
   </el-row>
 </div>
</template>
<style lang="scss">
.wx_mail_list{
   width: 100%;
   .head_con{
      padding: 10px 0;
      position: relative;
      text-align: center;
      font-weight: 900;
      .add_icon{
         position: absolute;
         top: 15px;
         right: 10px;
      }
   }
   .friends_list{
      background: #fff;
      .el-col{
         padding: 10px 0 0px 10px;
         .head_img{
            width: 30px;
            height: 30px;
            margin-right: 10px;
            vertical-align: top;
         }
         .content{
            width: calc(100% - 50px);
            display: inline-block;
            position: relative;
            color: #333;
            font-size: 14px;
            border-bottom: 1px solid rgb(238, 237, 237);
            padding-bottom: 2px;
            line-height: 30px;
         }
      }
    }
}
</style>
<script>
export default {
 name: 'MailList',
 data() {
    return {
       
    }
 },
 methods: {
    
 }
}
</script>