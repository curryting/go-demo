<template>
 <div class="wx_message">
   <div class="head_con">
      消息
      <el-dropdown class="add_icon">
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
      </el-dropdown>
   </div>
   <div class="head_search">

   </div>
   <el-row class="friends_list">
      <el-col :span="24">
         <img src="../assets/logo.png" alt="" class="head_img">
         <div class="content">
            <p class="remark_name">备注</p>
            <p>消息</p>
            <span class="time">12:00</span>
         </div>
      </el-col>
      <el-col :span="24">
         <img src="../assets/logo.png" alt="" class="head_img">
         <div class="content">
            <p class="remark_name">备注</p>
            <p>消息</p>
            <span class="time">12:00</span>
         </div>
      </el-col>
      <el-col :span="24">
         <img src="../assets/logo.png" alt="" class="head_img">
         <div class="content">
            <p class="remark_name">备注</p>
            <p>消息</p>
            <span class="time">12:00</span>
         </div>
      </el-col>
   </el-row>
 </div>
</template>
<style lang="scss">
.wx_message{
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
    .head_search{
      width: 100%;
      margin: 10px auto;
      border-bottom: 1px solid #ccc;
    }
    .friends_list{
      .el-col{
         padding: 10px 0 0px 10px;
         .head_img{
            width: 40px;
            height: 40px;
            margin-right: 6px;
            vertical-align: top;
         }
         .content{
            width: calc(100% - 50px);
            display: inline-block;
            position: relative;
            color: #ccc;
            font-size: 12px;
            border-bottom: 1px solid #ccc;
            padding-bottom: 10px;
            .remark_name{
               color: #333;
               font-size: 14px;
            }
            .time{
               position: absolute;
               top: 0px;
               right: 10px;
               color: #ccc;
               font-size: 12px;
               transform: scale(0.8);
            }
         }
      }
    }
}
</style>
<script>
export default {
 name: 'message',
 data() {
    return {
       
    }
 },
 methods: {
    
 }
}
</script>