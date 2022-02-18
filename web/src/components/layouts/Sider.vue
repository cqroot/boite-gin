<template>
  <div>
    <div class="logo">{{ collapsed ? "B" : "Boite" }}</div>
    <a-menu theme="dark" mode="inline" :selectedKeys="[routeName]">
      <a-menu-item v-for="item in routes" :key="item.name">
        <router-link :to="item.path" tag="div">
          <a-icon :type="item.meta.icon" />
          <span>{{ item.name }}</span>
        </router-link>
      </a-menu-item>
    </a-menu>
  </div>
</template>

<script>
import routes from "@/router/routes";

export default {
  name: "Sider",
  props: {
    collapsed: Boolean,
  },
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
    };
  },
};
</script>

<style lang="scss" scoped>
.logo {
  height: 32px;
  line-height: 32px;
  margin: 16px;
  color: white;
  font-weight: bold;
  font-size: 30px;
  text-align: center;
  margin-top: 30px;
}
</style>
