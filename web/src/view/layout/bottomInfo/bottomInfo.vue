<template>
  <div class="bottom-info">
    <div>
      <span>版权所有</span>
      <span>
        <!-- <a href="https://github.com/flipped-aurora/gin-vue-admin">{{ $GIN_VUE_ADMIN.appName }}</a> -->
        {{ $GIN_VUE_ADMIN.appName }}
      </span>
      <!-- <el-divider direction="vertical" />
      <span>版权所有</span>
      <span>
        <a href="https://github.com/flipped-aurora">flipped-aurora团队</a>
        Abelit
      </span> -->
    </div>
  </div>
</template>

<script>
// 此文件内容为版权信息，如需改动，请联系wx:shouzi_1994购买授权 未授权状态只需保留此代码 不影响任何正常使用
// 项目为apatch协议 请遵守版权协议内容
export default {
  name: 'BottomInfo'
}
</script>

<style lang="scss">
.bottom-info {
  color: #888;
  height: 30px;
  line-height: 12px;
  a {
    color: #888;
  }
  div {
    display: flex;
    justify-content: center;
    span{
      margin: 0 3px;
    }
  }
}
</style>
