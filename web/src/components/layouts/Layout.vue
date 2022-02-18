<template>
  <a-layout id="components-layout">
    <a-layout-sider v-model="collapsed" :trigger="null" collapsible>
      <Sider v-bind:collapsed="collapsed" />
    </a-layout-sider>

    <a-layout>
      <a-layout-header style="background: #fff; padding: 0">
        <a-icon
          class="trigger"
          :type="collapsed ? 'menu-unfold' : 'menu-fold'"
          @click="() => (collapsed = !collapsed)"
        />
      </a-layout-header>
      <a-layout-content
        :style="{
          margin: '24px 16px',
          padding: '24px',
          background: '#fff',
          minHeight: '280px',
        }"
      >
        <slot />
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>
<script>
import Sider from "./Sider.vue";

export default {
  components: {
    Sider,
  },
  data() {
    return {
      collapsed: false,
    };
  },
};
</script>

<style lang="scss" scoped>
#components-layout {
  min-height: 100%;

  .trigger {
    font-size: 18px;
    line-height: 64px;
    padding: 0 24px;
    cursor: pointer;
    transition: color 0.3s;
  }

  .trigger:hover {
    color: #1890ff;
  }
}
</style>
