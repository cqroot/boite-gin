<template>
  <div id="app">
    <Layout>
      <router-view />
    </Layout>
  </div>
</template>

<script>
import Layout from "@/components/layouts/Layout.vue";

export default {
  name: "App",
  components: {
    Layout,
  },
};
</script>

<style lang="scss" scoped>
#app {
  height: 100%;
}
</style>
