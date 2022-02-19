<template>
  <a-layout id="components-layout">
    <a-layout-sider id="sider" v-model="collapsed" collapsible>
      <div class="logo">
        {{ collapsed ? "B" : "Boite" }}
      </div>
      <a-menu theme="dark" mode="inline" :selectedKeys="[routeName]">
        <a-menu-item v-for="item in routes" :key="item.name">
          <router-link :to="item.path" tag="div">
            <a-icon :type="item.meta.icon" />
            <span>{{ item.name }}</span>
          </router-link>
        </a-menu-item>
      </a-menu>
    </a-layout-sider>

    <a-layout :style="collapsed ? 'margin-left: 80px' : 'margin-left: 200px'">
      <a-layout-content class="main-content">
        <slot />
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>
<script>
import routes from "@/router/routes";

export default {
  computed: {
    routeName() {
      const name = this.$route.name;
      if (typeof name === "string") {
        return name;
      } else {
        return "";
      }
    },
  },
  data() {
    return {
      routes,
      collapsed: false,
    };
  },
};
</script>

<style lang="scss" scoped>
#components-layout {
  min-height: 100%;
}

#sider {
  overflow: auto;
  height: 100vh;
  position: fixed;
  left: 0;
}

.logo {
  height: 32px;
  line-height: 32px;
  margin: 16px;
  color: white;
  font-weight: bold;
  font-size: 30px;
  text-align: center;
  margin-top: 30px;

  cursor: pointer;

  -moz-user-select: none;
  -webkit-user-select: none;
  -ms-user-select: none;
  -khtml-user-select: none;
  user-select: none;
}

.main-content {
  margin: 24px 16px;
  padding: 40px;
  background: #fff;
  minheight: 280px;
}
</style>
