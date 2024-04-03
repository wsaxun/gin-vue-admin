<!-- 此文件禁止修改！如果您没有购买授权，请联系wx:shouzi_1994购买授权，未授权状态只需保留此代码 不影响任何正常使用 -->

<template>
  <div class="flex gap-4 items-center text-sm text-gray-500 justify-center mb-4">
    <span>Powered by</span>
    <span>
      <a
        class="text-gray-800 font-bold"
      >xms</a>
    </span>
    <slot />
    <span>Copyright</span>
    <span>
      <a
        class="text-gray-800 font-bold"
        href="https://github.com/wsaxun"
      >ops</a>
    </span>
  </div>
</template>

<script  setup>
// 此文件内容为版权信息，如需改动，请联系wx:shouzi_1994购买授权 未授权状态只需保留此代码 不影响任何正常使用
// 项目为apatch协议 请遵守版权协议内容
defineOptions({
  name: 'BottomInfo'
})

</script>

