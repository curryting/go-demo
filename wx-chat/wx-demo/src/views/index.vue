<template>
 <div class="wx_home">
    <!-- 消息 -->
    <MessageView v-if="checkedTab=='message'"/>
    <!-- 通讯录 -->
    <MailList v-if="checkedTab=='mailList'"/>
    <!-- 我的 -->
    <My v-if="checkedTab=='my'"/>
    <div class="wx_footer_tab">
        <el-row>
            <el-col :span="8" @click="checked('message')" :class="checkedTab=='message'?'checked_tab':''">
                <el-icon :size="20"><ChatRound /></el-icon>
                <p>消息</p>
            </el-col>
            <el-col :span="8" @click="checked('mailList')" :class="checkedTab=='mailList'?'checked_tab':''">
                <el-icon :size="20"><Memo /></el-icon>
                <p>通讯录</p>
            </el-col>
            <el-col :span="8" @click="checked('my')" :class="checkedTab=='my'?'checked_tab':''">
                <el-icon :size="20"><User /></el-icon>
                <p>我的</p>
            </el-col>
        </el-row>
    </div>
 </div>
</template>

<script>
import MessageView from "./message-view.vue"
import MailList from "./mail-list.vue"
import My from "./my.vue"
export default {
    name: 'Index',
    components: {
        MessageView,
        MailList,
        My
    },
    data() {
        return {
            checkedTab: 'message'
        }
    },
    methods: {
        checked(type) {
            this.checkedTab = type
        }
    }
}
</script>
<style lang="scss">
.wx_home{
    width: 100%;
    height: 100%;
    background: rgb(228, 227, 227);
    .wx_footer_tab{
        width: 100%;
        position: fixed;
        bottom: 0;
        left: 0;
        background: rgb(228, 227, 227);
        border-top: 1px solid #ccc;
        font-size: 12px;
        padding: 5px 0;
        text-align: center;
        .checked_tab{
            color: #48c264;
        }
    }
}
</style>